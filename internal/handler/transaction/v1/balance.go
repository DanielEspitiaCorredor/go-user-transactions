package transactionv1

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"gopkg.in/gomail.v2"

	"github.com/DanielEspitiaCorredor/go-user-transactions/tools/dataframeops"
)

type AccountBalance struct {
	Account        string
	TotalBalance   float64
	BalacePerMonth *dataframe.DataFrame
	DebitTx        *TransactionData
	CreditTx       *TransactionData
}

//
// Getters
//

// GetAccount returns the account field
func (a *AccountBalance) GetAccount() string {
	if a == nil {
		return ""
	}
	return a.Account
}

// GetTotalBalance returns the total balance field
func (a *AccountBalance) GetTotalBalance() float64 {
	if a == nil {
		return 0.0
	}
	return a.TotalBalance
}

// GetBalacePerMonth returns the dataframe with total balance per month
func (a *AccountBalance) GetBalacePerMonth() *dataframe.DataFrame {
	if a == nil {
		return nil
	}
	return a.BalacePerMonth
}

// GetDebitTx returns the debit transactions data
func (a *AccountBalance) GetDebitTx() *TransactionData {
	if a == nil {
		return nil
	}
	return a.DebitTx
}

// GetCreditTx returns the credit transactions data
func (a *AccountBalance) GetCreditTx() *TransactionData {
	if a == nil {
		return nil
	}
	return a.CreditTx
}

//
// Functions
//

func (a *AccountBalance) SendReport(receiverEmail string) error {

	// Load the HTML content from a file
	bytes, err := os.ReadFile("./assets/email/template.html")

	if err != nil {
		fmt.Println("[AccountBalance][SendReport] err opening file", err)
		return err
	}

	templateStr := string(bytes)

	// Replace transactions in month
	templateStr = strings.Replace(templateStr, "[account_number]", a.GetAccount(), 1)
	templateStr = strings.Replace(templateStr, "[total_balance]", fmt.Sprintf("%.2f", a.GetTotalBalance()), 1)

	monthTxStr := ""
	for _, v := range a.GetBalacePerMonth().Maps() {

		month, _ := v["month"].(int)
		txCount, _ := v["transactions_count"].(float64)
		templateMonthTx := fmt.Sprintf(`<p><span class="label">Transactions in %s:</span> %.0f</p><br>`, time.Month(month).String(), txCount)

		if monthTxStr == "" {
			monthTxStr = templateMonthTx
			continue
		}

		monthTxStr = fmt.Sprintf("%s%s", monthTxStr, templateMonthTx)
	}

	templateStr = strings.Replace(templateStr, "[month_transactions]", monthTxStr, 1)

	// Replace debit and credit average

	templateStr = strings.Replace(templateStr, "[average_debit_amount]", fmt.Sprintf("%.2f", a.GetDebitTx().GetAverageTxValue()), 1)
	templateStr = strings.Replace(templateStr, "[average_credit_amount]", fmt.Sprintf("%.2f", a.GetCreditTx().GetAverageTxValue()), 1)

	// Replace names of top debit and credit transactions
	templateStr = strings.Replace(templateStr, "[top_debit_transactions]", a.GetTopTransactionStr(TransactionType_DEBIT, 3), 1)
	templateStr = strings.Replace(templateStr, "[top_credit_transactions]", a.GetTopTransactionStr(TransactionType_CREDIT, 3), 1)

	senderEmail := os.Getenv("SMTP_SENDER_EMAIL")

	mail := gomail.NewMessage()
	mail.SetHeader("From", senderEmail)
	mail.SetHeader("To", receiverEmail)
	mail.SetHeader("Subject", "Your account report is now avalaible")
	mail.Embed("./assets/email/logo.png")
	mail.SetBody("text/html", templateStr)

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, senderEmail, os.Getenv("SMTP_PASSWD"))
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println("[AccountBalance][SendReport] err sending email", err)
	}

	fmt.Println("[AccountBalance][SendReport] email sent to", senderEmail)
	return nil

}

func (a *AccountBalance) GetTopTransactionStr(txType TransactionType, topNum int) (topTxStr string) {

	if txType == TransactionType_UNKNOWN {
		return
	}

	multiplier := 1.0
	topDf := a.GetCreditTx().GetTopTransactions()

	if txType == TransactionType_DEBIT {

		multiplier = -1.0
		topDf = a.GetDebitTx().GetTopTransactions()
	}

	topDf = dataframeops.GetTop(topDf, topNum)

	for _, v := range topDf.Maps() {

		txName, _ := v["name"].(string)
		txTotal, _ := v["total"].(float64)
		if topTxStr == "" {

			topTxStr = fmt.Sprintf("%s:%.2f", txName, txTotal*multiplier)
			continue
		}
		topTxStr = fmt.Sprintf("%s, %s:%.2f", topTxStr, txName, txTotal*multiplier)
	}

	return
}
