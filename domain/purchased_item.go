package domain

import "time"

type PurchasedItem struct {
	ItemID           int       `db:"item_id"`
	PurchaseID       int       `db:"purchase_id"`
	ProductID        int       `db:"product_id"`
	Name             string    `db:"name"`
	Category         string    `db:"category"`
	Qty              int       `db:"qty"`
	Price            float64   `db:"price"`
	SKU              string    `db:"sku"`
	FileID           int       `db:"file_id"`
	FileURI          string    `db:"file_uri"`
	FileThumbnailURI string    `db:"file_thumbnail_uri"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
