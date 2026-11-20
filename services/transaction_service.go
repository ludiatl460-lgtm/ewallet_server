package services

import (
	"database/sql"
	"errors"

	"backend/config"
	"backend/models"

	"github.com/shopspring/decimal"
)

// error ທາງທຸລະກິດ: controller ຈະແປງເປັນ HTTP status ທີ່ເໝາະສົມ
var (
	ErrSourceNotFound    = errors.New("source account not found")
	ErrDestNotFound      = errors.New("destination account not found")
	ErrSameAccount       = errors.New("cannot transfer to the same account")
	ErrAccountInactive   = errors.New("account is not active")
	ErrCurrencyMismatch  = errors.New("currency of the two accounts does not match")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// ຂໍ້ມູນບັນຊີທີ່ຖືກລັອກ ຕອນໂອນ (ອ່ານຈາກຕາຕະລາງ account ຂອງໝູ່)
type lockedAccount struct {
	Id               int64           `db:"id"`
	Status           string          `db:"status"`
	Balance          decimal.Decimal `db:"balance"`
	AvailableBalance decimal.Decimal `db:"available_balance"`
	Currency         string          `db:"currency"`
}

func GetTransactionByID(id int64) (*models.Transaction, error) {
	var t models.Transaction
	err := config.DB.Get(&t, `SELECT * FROM transactions WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	t.Ledgers, err = GetLedgersByTransactionID(id)
	return &t, err
}

func getByIdempotencyKey(key string) (*models.Transaction, error) {
	var t models.Transaction
	err := config.DB.Get(&t, `SELECT * FROM transactions WHERE idempotency_key = $1`, key)
	if err != nil {
		return nil, err
	}
	t.Ledgers, err = GetLedgersByTransactionID(t.Id)
	return &t, err
}

// ໂອນເງິນພາຍໃນລະບົບ. ຄືນ (ຜົນ, duplicate, error)
// duplicate = true ແປວ່າ idempotency_key ນີ້ເຄີຍເຮັດແລ້ວ ຄືນຜົນເກົ່າ ບໍ່ຫັກເງິນຊ້ຳ
func CreateTransfer(in *models.Transaction, ip string) (*models.Transaction, bool, error) {
	// 1. ເຄີຍເຮັດແລ້ວບໍ?
	old, err := getByIdempotencyKey(in.IdempotencyKey)
	if err == nil {
		return old, true, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, false, err
	}

	// 2. ເລີ່ມ DB transaction: ຕັ້ງແຕ່ນີ້ຈົນ Commit ເປັນກ້ອນດຽວ (ພັງ = ຍົກເລີກທັງໝົດ)
	tx, err := config.DB.Beginx()
	if err != nil {
		return nil, false, err
	}
	defer tx.Rollback() // ຫຼັງ Commit ແລ້ວ ບໍ່ເຮັດຫຍັງ

	// 3. ຫາບັນຊີປາຍທາງຈາກເລກບັນຊີ
	var toID int64
	err = tx.Get(&toID, `SELECT id FROM account WHERE account_number = $1`, in.ToAccountNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, ErrDestNotFound
	}
	if err != nil {
		return nil, false, err
	}
	if toID == in.AccountId {
		return nil, false, ErrSameAccount
	}

	// 4. ລັອກທັງ 2 ບັນຊີ ຕາມລຳດັບ id (ກັນ deadlock) ແລະ ກັນໂອນພ້ອມກັນຈົນຍອດຕິດລົບ
	var accs []lockedAccount
	err = tx.Select(&accs,
		`SELECT id, status, balance, available_balance, currency
		 FROM account WHERE id IN ($1, $2) ORDER BY id FOR UPDATE`,
		in.AccountId, toID)
	if err != nil {
		return nil, false, err
	}
	var from, to *lockedAccount
	for i := range accs {
		switch accs[i].Id {
		case in.AccountId:
			from = &accs[i]
		case toID:
			to = &accs[i]
		}
	}
	if from == nil {
		return nil, false, ErrSourceNotFound
	}
	if to == nil {
		return nil, false, ErrDestNotFound
	}
	if from.Status != "active" || to.Status != "active" {
		return nil, false, ErrAccountInactive
	}
	if from.Currency != to.Currency {
		return nil, false, ErrCurrencyMismatch
	}

	fee := decimal.Zero // ຍັງບໍ່ຄິດຄ່າທຳນຽມ (ເພີ່ມຕາມກົດ limit/fee ທີຫຼັງ)
	total := in.Amount.Add(fee)

	// 5. ບັນທຶກ transaction ກ່ອນກວດຍອດ. ຖ້າ key ຊ້ຳ (ຄົນສົ່ງພ້ອມກັນ ແລະ ສຳເລັດກ່ອນ)
	//    ຈະບໍ່ມີແຖວໃໝ່ → ຄືນຜົນເກົ່າ ແທນທີ່ຈະຕອບ "ເງິນບໍ່ພໍ" ຜິດໆ
	var created models.Transaction
	err = tx.Get(&created,
		`INSERT INTO transactions
		   (idempotency_key, type, amount, currency, transaction_status, account_id,
		    to_account_number, ip_address, description, complete_at, fee, total_amount)
		 VALUES ($1, 'transfer', $2, $3, 'success', $4, $5, $6, $7, NOW(), $8, $9)
		 ON CONFLICT (idempotency_key) DO NOTHING
		 RETURNING *`,
		in.IdempotencyKey, in.Amount, from.Currency, from.Id,
		in.ToAccountNumber, ip, in.Description, fee, total)
	if errors.Is(err, sql.ErrNoRows) {
		existing, e := getByIdempotencyKey(in.IdempotencyKey)
		return existing, true, e
	}
	if err != nil {
		return nil, false, err
	}

	// 6. ກວດຍອດ (ລົ້ມ = Rollback ແຖວ transaction ຂ້າງເທິງຫາຍໄປນຳ)
	if from.AvailableBalance.LessThan(total) {
		return nil, false, ErrInsufficientFunds
	}

	// 7. ຫັກຕົ້ນທາງ ແລະ ເພີ່ມປາຍທາງ
	if _, err = tx.Exec(
		`UPDATE account SET balance = balance - $1, available_balance = available_balance - $1 WHERE id = $2`,
		total, from.Id); err != nil {
		return nil, false, err
	}
	if _, err = tx.Exec(
		`UPDATE account SET balance = balance + $1, available_balance = available_balance + $1 WHERE id = $2`,
		in.Amount, to.Id); err != nil {
		return nil, false, err
	}

	// 8. ບັນທຶກ ledger 2 ແຖວ (debit ຕົ້ນທາງ, credit ປາຍທາງ)
	if err = insertLedger(tx, from.Id, created.Id, "debit", total, from.Balance, from.Balance.Sub(total)); err != nil {
		return nil, false, err
	}
	if err = insertLedger(tx, to.Id, created.Id, "credit", in.Amount, to.Balance, to.Balance.Add(in.Amount)); err != nil {
		return nil, false, err
	}

	// 9. ສຳເລັດທັງໝົດ
	if err = tx.Commit(); err != nil {
		return nil, false, err
	}
	res, err := GetTransactionByID(created.Id)
	return res, false, err
}