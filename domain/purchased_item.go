package domain

type PurchasedItem struct {
	ItemID     int `db:"item_id"`
	PurchaseID int `db:"purchase_id"`
	ProductID  int `db:"product_id"`
	Qty        int `db:"qty"`
	Price      int `db:"price"`
}

type PurchasedItemReq struct {
	ItemID     string `db:"item_id,omitempty"`
	PurchaseID string `db:"purchase_id,omitempty"`
	ProductID  string `db:"product_id,omitempty"`
	Qty        int    `db:"qty,omitempty"`
	Price      int    `db:"price,omitempty"`
}
