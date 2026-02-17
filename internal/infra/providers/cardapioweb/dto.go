package cardapioweb

type OrderDTO struct {
	ID           int64  `json:"id"`
	Status       string `json:"status"`
	OrderType    string `json:"order_type"`
	OderTiming   string `json:"order_timing"`
	SalesChannel string `json:"sales_channel"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type OrderDetailsDTO struct {
	ID           int64          `json:"id"`
	Status       string         `json:"status"`
	OrderType    string         `json:"order_type"`
	OderTiming   string         `json:"order_timing"`
	SalesChannel string         `json:"sales_channel"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	Customer     CustomerDTO    `json:"customer"`
	Items        []OrderItemDTO `json:"items"`
}

type OrderItemDTO struct {
	ItemID       int64               `json:"item_id"`
	OrderItemID  int64               `json:"order_item_id"`
	ExternalCode string              `json:"external_code"`
	Name         string              `json:"name"`
	Quantity     int                 `json:"quantity"`
	UnitPrice    float64             `json:"unit_price"`
	TotalPrice   float64             `json:"total_price"`
	Kind         string              `json:"kind"`
	Status       string              `json:"status"`
	Observation  string              `json:"observation"`
	Options      []map[string]string `json:"options"`
	Payments     []map[string]string `json:"payments"`
}

type CustomerDTO struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
