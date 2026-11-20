package services

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountGenerate = errors.New("failed to generate unique account number")
)

// ກຳນົດ Column ທີ່ຈະດຶງອອກມາ (ຫຼີກລ້ຽງ SELECT *)
const accountColumns = `id, user_id, account_number, account_type, status, opened_at, closed_at, balance, available_balance, hold_amount, currency`

func CreateAccount(userID int64, accountType string, currency string) (*models.Account, error) {
	
	accountNumber, err := utils.GenerateAccountNumber()
	if err != nil {
		return nil, err
	}

	// SQL ສຳລັບສ້າງ Account ໃໝ່ (ຍອດເງິນຕ່າງໆ ຈະຖືກ Default ເປັນ 0 ຈາກ Database)
	query := `
		INSERT INTO account (user_id, account_number, account_type, status, currency) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING ` + accountColumns

	var acc models.Account
	err = config.DB.Get(&acc, query, userID, accountNumber, accountType, models.AccountStatusActive, currency)
	
	if err != nil {
		// (Optional) ກວດສອບ Error ຖ້າ account_number ຊ້ຳກັນ (Postgres Error Code 23505)
		// ໃນຄວາມເປັນຈິງ 15 ຫຼັກມີໂອກາດຊ້ຳກັນໜ້ອຍຫຼາຍ ແຕ່ກໍ່ຄວນມີການກັນໄວ້
		return nil, err
	}

	return &acc, nil
}

func GetAccountByID(id int64) (*models.Account, error) {
	var acc models.Account
	query := `SELECT ` + accountColumns + ` FROM account WHERE id = $1`
	
	err := config.DB.Get(&acc, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &acc, nil
}

func GetAccountsByUserID(userID int64) ([]models.Account, error) {
	list := []models.Account{} 
	query := `SELECT ` + accountColumns + ` FROM account WHERE user_id = $1 ORDER BY id DESC`
	
	err := config.DB.Select(&list, query, userID)
	return list, err
}

func UpdateAccountStatus(id int64, newStatus string) (*models.Account, error) {
	
	query := `
		UPDATE account 
		SET status = $1, 
		    closed_at = CASE WHEN $1 = 'closed' THEN NOW() ELSE NULL END
		WHERE id = $2
		RETURNING ` + accountColumns

	var acc models.Account
	err := config.DB.Get(&acc, query, newStatus, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &acc, nil
}

// ສຳລັບການທົດສອບ (Mock) - ບໍ່ຄວນມີໃນ Production ຖ້າບໍ່ຈຳເປັນ
func PrintAccountDetails(acc *models.Account) string {
	return fmt.Sprintf("Account: %s | Balance: %.2f %s", acc.AccountNumber, acc.Balance, acc.Currency)
}

// UpdateAccount ອັບເດດຂໍ້ມູນ Account (Type ແລະ Status)
func UpdateAccount(id int64, accountType *string, status *string) (*models.Account, error) {
	// ໃຊ້ COALESCE ເພື່ອຖ້າຄ່າໃດເປັນ nil ໃຫ້ໃຊ້ຄ່າເກົ່າໃນ Database
	query := `
		UPDATE account 
		SET account_type = COALESCE($1, account_type), 
		    status = COALESCE($2, status),
		    closed_at = CASE WHEN $2 = 'closed' THEN NOW() ELSE closed_at END
		WHERE id = $3
		RETURNING ` + accountColumns

	var acc models.Account
	err := config.DB.Get(&acc, query, accountType, status, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &acc, nil
}

// CloseAccount ສຳລັບປິດ Account (Soft Delete)
func CloseAccount(id int64) error {
	query := `UPDATE account SET status = 'closed', closed_at = NOW() WHERE id = $1`
	res, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrAccountNotFound
	}
	return nil
}

// GetAllAccounts ດຶງ Account ທັງໝົດໃນລະບົບ (ສຳລັບ Admin ເບິ່ງພາບລວມ)
func GetAllAccounts() ([]models.Account, error) {
	list := []models.Account{}
	query := `SELECT ` + accountColumns + ` FROM account ORDER BY id DESC`
	
	err := config.DB.Select(&list, query)
	return list, err
}