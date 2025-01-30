package dto

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type PaymentDetail struct {
	PaymentDetailData
	TotalPrice int `json:"total_price"`
}

type PaymentDetailData struct {
	BankAccountName   string `json:"bank_account_name"`
	BankAccountHolder string `json:"bank_account_holder"`
	BankAccountNumber string `json:"bank_account_number"`
}
