package validators

import (
	"fmt"
	"strings"

	"backend/models"
)

type CreateAccountRequest struct {
	UserID      int64  `json:"user_id"`
	AccountType string `json:"account_type"`
	Currency    string `json:"currency"`
}

func ValidateCreateAccount(r *CreateAccountRequest) map[string]string {
	errs := map[string]string{}

	if r.UserID <= 0 {
		errs["user_id"] = "user_id must be a positive number"
	}

	// 2. ກວດສອບ Account Type (ຕ້ອງເປັນ personal ຫຼື business ເທົ່ານັ້ນ)
	r.AccountType = strings.ToLower(strings.TrimSpace(r.AccountType))
	if r.AccountType != models.AccountTypePersonal && r.AccountType != models.AccountTypeBusiness {
		errs["account_type"] = fmt.Sprintf("account_type must be '%s' or '%s'", models.AccountTypePersonal, models.AccountTypeBusiness)
	}

	// 3. ກວດສອບ Currency (ຕ້ອງມີ 3 ຕົວອັກສອນ ແລະ ເປັນຕົວພິມໃຫຍ່)
	r.Currency = strings.ToUpper(strings.TrimSpace(r.Currency))
	if len(r.Currency) != 3 {
		errs["currency"] = "currency must be exactly 3 characters (e.g., LAK, USD, THB)"
	}

	return errs
}