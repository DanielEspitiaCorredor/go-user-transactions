package dataframeops

import (
	"math"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type DataframeOperation int
type AggregationType int

const (
	Operation_UNKNOWN DataframeOperation = iota
	Operation_ABSOULTE
)
const (
	AggregationType_UNKNOWN AggregationType = iota
	AggregationType_MEAN
	AggregationType_SUM
)

func GetAggregatedValue(s series.Series, aggType AggregationType) (v float64) {

	values := s.Float()

	if len(values) == 0 {
		return
	}

	var sum float64
	for _, f := range values {
		sum += f
	}

	switch aggType {
	case AggregationType_MEAN:

		v = sum / float64(len(values))

	case AggregationType_SUM:

		v = sum
	}

	return
}

func ApplyDf(df *dataframe.DataFrame, columnName string, operation DataframeOperation) {

	if df == nil {
		return
	}

	switch operation {

	case Operation_ABSOULTE:

		values := df.Col(columnName).Float()

		if len(values) == 0 {
			return
		}

		for idx, v := range values {

			values[idx] = math.Abs(v)
		}

		*df = df.Mutate(series.New(values, series.Float, columnName))
	}

}
