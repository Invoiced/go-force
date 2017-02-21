package sobjects

type Product struct {
	BaseSObject
	Description string `force:",omitempty"`
	ProductCode string `force:",omitempty"`
	isActive    bool   `force:",omitempty"`
}

func (t *Product) ApiName() string {
	return "Product2"
}

type ProductQueryResponse struct {
	BaseQuery
	Records []Product `json:"Records" force:"records"`
}
