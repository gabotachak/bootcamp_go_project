package domain

type ProductRecord struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float32 `json:"purchase_price"`
	SalePrice      float32 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}
