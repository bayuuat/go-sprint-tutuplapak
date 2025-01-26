package dto

type AuthResponse struct {
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	Token string  `json:"token"`
}

type AuthEmailReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type AuthPhoneReq struct {
	Phone    string `json:"phone" validate:"required,phonenumber"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserData struct {
	Email             *string `json:"email"`
	Phone             *string `json:"phone"`
	FileID            int     `json:"fileId"`
	FileURI           string  `json:"fileUri"`
	FileThumbnailURI  string  `json:"fileThumbnailUri"`
	BankAccountName   *string `json:"bankAccountName"`
	BankAccountHolder *string `json:"bankAccountHolder"`
	BankAccountNumber *string `json:"bankAccountNumber"`
}

type UpdateUser struct {
	FileID            int    `json:"fileId" validate:"required"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=8,max=32"`
	BankAccountHolder string `json:"bankAccountHolder" validate:"required,min=8,max=32"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=8,max=32"`
}
