package domain

import "time"

type Purchase struct {
	PurchaseID        int       `db:"purchase_id"`
	TotalPrice        float64   `db:"total_price"`
	BankAccountName   string    `db:"bank_account_name"`
	BankAccountHolder string    `db:"bank_account_holder"`
	BankAccountNumber string    `db:"bank_account_number"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
