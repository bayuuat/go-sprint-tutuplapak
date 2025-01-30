package domain

import "time"

type Product struct {
	ProductID int       `db:"product_id"`
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	Qty       int       `db:"qty"`
	Price     int       `db:"price"`
	SKU       string    `db:"sku"`
	FileID    int       `db:"file_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
