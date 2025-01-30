package dto

type ProductFilter struct {
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	ProductID     string `json:"product_id"`
	SKU           string `json:"sku"`
	Category      string `json:"category"`
	SortBy        string `json:"sort_by"`
	SoldInSeconds int    `json:"-"`
}

type UpdateProductReq struct {
}

type Product struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`
	Category string `json:"category" validate:"required,oneof=Food Beverage Clothes Furniture Tools"`
	Qty      int    `json:"qty" validate:"required,numeric,min=1"`
	Price    int    `json:"price" validate:"required,numeric,min=100"`
	Sku      string `json:"sku" validate:"required,max=32"`
	FileId   string `json:"fileId" validate:"required,min=1"`
}

type ProductReq struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`
	Category string `json:"category" validate:"required,oneof=Food Beverage Clothes Furniture Tools"`
	Qty      int    `json:"qty" validate:"required,numeric,min=1"`
	Price    int    `json:"price" validate:"required,numeric,min=100"`
	Sku      string `json:"sku" validate:"required,max=32"`
	FileId   string `json:"fileId" validate:"required,min=1"`
}

type ProductData struct {
	ProductId        string `json:"productId"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Qty              int    `json:"qty"`
	Price            int    `json:"price"`
	Sku              string `json:"sku"`
	FileId           string `json:"fileId"`
	FileUri          string `json:"fileUri"`
	FileThumbnailUri string `json:"fileThumbnailUri"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
