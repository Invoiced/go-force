package sobjects

type OrderItem struct {
	BaseSObject
	OrderID     string  `force:",omitempty"`
	Product2ID  string  `force:",omitempty"`
	Quantity    float64 `force:",omitempty"`
	ServiceDate string  `force:",omitempty"`
	UnitPrice   float64 `force:",omitempty"`
}

func (t *OrderItem) ApiName() string {
	return "OrderItem"
}

type OrderItemQueryResponse struct {
	BaseQuery
	Records []OrderItem `json:"Records" force:"records"`
}
