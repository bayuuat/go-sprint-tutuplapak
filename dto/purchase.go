package dto

type PurchaseReq struct {
	PurchasedItemsReq   []PurchasedItemReq `json:"purchasedItems" validate:"required,min=1"`
	SenderName          string             `json:"senderName"`
	SenderContactType   string             `json:"senderContactType" validate:"required"`
	SenderContactDetail string             `json:"senderContactDetail" validate:"required"`
}

type PurchaseData struct {
	PurchaseID     string          `json:"purchaseId"`
	PurchasedItems []ProductData   `json:"purchasedItems"`
	PaymentDetails []PaymentDetail `json:"paymentDetails"`
}
