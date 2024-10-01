package transactionv1

type ExtractRequest struct {
	Account       string `json:"account,required"`
	Year          int    `json:"year,required"`
	ReceiverEmail string `json:"receiver_email,required"`
}

// GetAccount returns the Account field
func (e *ExtractRequest) GetAccount() string {
	if e == nil {
		return ""
	}
	return e.Account
}

// GetYear returns the Year field
func (e *ExtractRequest) GetYear() int {
	if e == nil {
		return 0
	}
	return e.Year
}

// GetReceiverEmail returns the ReceiverEmail field
func (e *ExtractRequest) GetReceiverEmail() string {
	if e == nil {
		return ""
	}
	return e.ReceiverEmail
}
