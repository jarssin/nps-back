package order

import "time"

type DTO struct {
	ID            int64     `json:"id"`
	Status        string    `json:"status"`
	OrderType     string    `json:"order_type"`
	OderTiming    string    `json:"order_timing"`
	SalesChannel  string    `json:"sales_channel"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	Items         []ItemDTO `json:"items"`
}

type ItemDTO struct {
	ItemID       int64   `json:"item_id"`
	OrderItemID  int64   `json:"order_item_id"`
	ExternalCode string  `json:"external_code"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	TotalPrice   float64 `json:"total_price"`
	Kind         string  `json:"kind"`
	Status       string  `json:"status"`
	Observation  string  `json:"observation"`
	Options      []any   `json:"options"`
	Payments     []any   `json:"payments"`
}
