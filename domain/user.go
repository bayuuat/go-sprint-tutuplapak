package domain

import "time"

type User struct {
	Id                string    `db:"id"`
	Email             *string   `db:"email"`
	Password          string    `db:"password"`
	Phone             *string   `db:"phone"`
	BankAccountName   *string   `db:"bank_account_name"`
	BankAccountHolder *string   `db:"bank_account_holder"`
	BankAccountNumber *string   `db:"bank_account_number"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
