package dto

type PurchasedItemReq struct {
	ProductID  string `json:"productId" validate:"required"`
	Qty        int    `json:"qty" validate:"required,min=2"`
	PurchaseID string `json:"purchaseID,omitempty" validate:"min=2,omitempty"`
}

type PurchasedItemFilter struct {
}

type PurchasedItemData struct {
	ProductID        int    `json:"productId"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Qty              int    `json:"qty"`
	Price            int    `json:"price"`
	SKU              string `json:"sku"`
	FileID           int    `json:"fileId"`
	FileURI          string `json:"fileUri"`
	FileThumbnailURI string `json:"fileThumbnailUri"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
