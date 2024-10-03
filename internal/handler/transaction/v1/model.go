package transactionv1

import "github.com/go-gota/gota/dataframe"

type ExtractRequest struct {
	Account       string `json:"account" binding:"required"`
	Year          int    `json:"year" binding:"required"`
	ReceiverEmail string `json:"receiver_email" binding:"required"`
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

type TransactionData struct {
	AverageTxValue  float64
	TopTransactions *dataframe.DataFrame
}

// GetAverageTxValue returns the average transaction value
func (t *TransactionData) GetAverageTxValue() float64 {
	if t == nil {
		return 0.0
	}
	return t.AverageTxValue
}

// GetTopTransactions returns the dataframe with top transactions field
func (t *TransactionData) GetTopTransactions() *dataframe.DataFrame {
	if t == nil {
		return nil
	}
	return t.TopTransactions
}
