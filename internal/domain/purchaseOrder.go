package domain

type PurchaseOrder struct {
	ID            int    `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerId       int    `json:"buyer_id"`
	OrderStatusId int    `json:"order_status_id"`
	CarrierId     int    `json:"carrier_id"`
	WarehouseId   int    `json:"warehouse_id"`
}
