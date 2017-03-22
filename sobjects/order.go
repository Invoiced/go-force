package sobjects

type Order struct {
	BaseSObject
	AccountID     string  `force:",omitempty"`
	EffectiveDate string  `force:",omitempty"`
	AccountNumber string  `force:",omitempty"`
	ContractID    string  `force:",omitempty"`
	CreatedDate   string  `force:",omitempty"`
	OrderNumber   string  `force:",omitempty"`
	IsDeleted     bool    `force:",omitempty"`
	Status        string  `force:",omitempty"`
	TotalAmount   float64 `force:",omitempty"`
	OpportunityId string  `force:",omitempty"`
}

func (t *Order) ApiName() string {
	return "Order"
}

type OrderQueryResponse struct {
	BaseQuery
	Records []Order `json:"Records" force:"records"`
}
