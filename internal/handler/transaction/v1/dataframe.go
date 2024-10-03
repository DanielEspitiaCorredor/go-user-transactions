package transactionv1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"

	"github.com/DanielEspitiaCorredor/go-user-transactions/tools/dataframeops"
)

type TransactionType int

const (
	TransactionType_UNKNOWN TransactionType = iota
	TransactionType_DEBIT
	TransactionType_CREDIT
)

type TransactionDataframe struct {
	account string
	df      dataframe.DataFrame
	year    int
}

func NewTransactionDF(account, csvPath string, year int) (tx *TransactionDataframe, err error) {

	file, err := os.Open(csvPath)

	if err != nil {
		fmt.Println("[GetDataframeFromCsv] error open file", err)
		return
	}

	df := dataframe.ReadCSV(file,
		dataframe.HasHeader(true),
	)

	defer file.Close()

	tx = &TransactionDataframe{
		account: account,
		df:      df,
		year:    year,
	}

	return
}

func (t *TransactionDataframe) PreProcessData() {

	dateCol := t.df.Col("date").Records()
	months := []int{}
	// dates := []time.Time{}

	for _, value := range dateCol {

		dateValues := strings.Split(value, "/")

		if len(dateValues) < 2 {

			continue
		}

		month, _ := strconv.Atoi(dateValues[0])
		months = append(months, month)
	}

	// Add new column with time.Time values
	t.df = t.df.Mutate(
		series.New(months, series.Int, "month"),
	)

}

func (t *TransactionDataframe) NewAccountBalance() *AccountBalance {

	txPerMonth := (t.df.
		GroupBy("month").
		Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"value"}).
		Rename("transactions_count", "value_COUNT"))

	return &AccountBalance{
		Account:        t.account,
		TotalBalance:   dataframeops.GetAggregatedValue(t.df.Col("value"), dataframeops.AggregationType_SUM),
		BalacePerMonth: &txPerMonth,
		DebitTx:        t.getTransactionData(TransactionType_DEBIT),  // Get debit data
		CreditTx:       t.getTransactionData(TransactionType_CREDIT), // Get credit data
	}
}

func (t *TransactionDataframe) getTransactionData(txType TransactionType) (txData *TransactionData) {

	var comparator series.Comparator
	switch txType {
	case TransactionType_CREDIT:

		comparator = series.Greater

	case TransactionType_DEBIT:

		comparator = series.Less
	default:
		return nil
	}

	filteredTx := t.df.Filter(dataframe.F{Colname: "value", Comparator: comparator, Comparando: 0.0})

	// For debit transactions apply absolute
	if comparator == series.Less {

		dataframeops.ApplyDf(&filteredTx, "value", dataframeops.Operation_ABSOULTE)
	}

	topTx := (filteredTx.
		GroupBy("name").
		Aggregation([]dataframe.AggregationType{dataframe.Aggregation_SUM}, []string{"value"}).
		Rename("total", "value_SUM").
		Arrange(dataframe.RevSort("total")))

	txData = &TransactionData{
		AverageTxValue:  dataframeops.GetAggregatedValue(filteredTx.Col("value"), dataframeops.AggregationType_MEAN),
		TopTransactions: &topTx,
	}

	return
}
