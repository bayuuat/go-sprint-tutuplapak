package domain

import (
	"time"
)

type Purchase struct {
	PurchaseID        int       `db:"purchase_id"`
	TotalPrice        int       `db:"total_price"`
	BankAccountName   string    `db:"bank_account_name"`
	BankAccountHolder string    `db:"bank_account_holder"`
	BankAccountNumber string    `db:"bank_account_number"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type PurchaseReq struct {
	TotalPrice          int    `db:"total_price" goqu:"omitempty"`
	SenderName          string `db:"sender_name" goqu:"omitempty"`
	SenderContactType   string `db:"sender_contact_type" goqu:"omitempty"`
	SenderContactDetail string `db:"sender_contact_detail" goqu:"omitempty"`
	UserIds             []int  `db:"user_ids" goqu:"omitempty"`
}
