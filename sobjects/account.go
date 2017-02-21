package sobjects

type Account struct {
	AccountNumber string `force:",omitempty"`
	BaseSObject
	BillingCity       string `force:",omitempty"`
	BillingCountry    string `force:",omitempty"`
	BillingPostalCode string `force:",omitempty"`
	BillingState      string `force:",omitempty"`
	BillingStreet     string `force:",omitempty"`
	Phone             string `force:",omitempty"`
}

func (a Account) ApiName() string {
	return "Account"
}
