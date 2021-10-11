package transaction

import (
	"banking_example_v1/account"
	"banking_example_v1/rounding"
	"time"
)

type Transaction struct {
	PayeeUserID *account.Account
	PayerUserID *account.Account
	Title       string
	Amount      float64
	Description string
	Date        time.Time
}

func Transfer(trans Transaction) {
	trans.PayeeUserID.Balance += trans.Amount
	trans.PayerUserID.Balance -= trans.Amount

	trans.PayeeUserID.Balance = rounding.ToFixed(trans.PayeeUserID.Balance, 2)
	trans.PayerUserID.Balance = rounding.ToFixed(trans.PayerUserID.Balance, 2)
}
