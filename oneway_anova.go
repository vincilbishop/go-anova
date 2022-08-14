package goanova

import (
	"math"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"gonum.org/v1/gonum/stat/distuv"
)

type anovaTableMeansResult struct {
	GroupMeans  map[string]float64
	OverallMean float64
}

func OneWayAnova() {}

// Gets a map of each group in the input
func (anovaInput *AnovaInput) groupedDataPoints() map[string][]AnovaDataPoint {
	return lop.GroupBy(anovaInput.DataPoints, func(dataPoint AnovaDataPoint) string {
		return dataPoint.Group
	})
}

func getGroupedDataPointValues(groupedDataPoints *map[string][]AnovaDataPoint) map[string][]float64 {
	var result = make(map[string][]float64)
	for key, value := range *groupedDataPoints {
		result[key] = lop.Map(value, func(dataPoint AnovaDataPoint, _ int) float64 {
			return dataPoint.Value
		})
	}
	return result
}

// Calculates the mean for groups along with the overall mean
func (anovaInput *AnovaInput) anovaTableMeansResult() anovaTableMeansResult {
	groupedDataPoints := anovaInput.groupedDataPoints()
	groupedValues := getGroupedDataPointValues(&groupedDataPoints)
	var anovaTableMeansResult = anovaTableMeansResult{
		GroupMeans: make(map[string]float64),
	}
	var overallTotal = float64(0.0)
	var overallCount = float64(0.0)
	for key, value := range groupedValues {
		sum := lo.SumBy(value, func(item float64) float64 {
			return item
		})
		count := float64(len(value))
		avg := sum / count
		anovaTableMeansResult.GroupMeans[key] = avg
		overallTotal += sum
		overallCount += count
	}
	anovaTableMeansResult.OverallMean = overallTotal / overallCount
	return anovaTableMeansResult
}

func calculateSSR(meansResult *anovaTableMeansResult) float64 {
	ssrValue := float64(0.0)
	for _, value := range meansResult.GroupMeans {
		ssrValue += 10 * math.Pow(value-meansResult.OverallMean, 2)
	}
	return math.Round(ssrValue*100) / 100
}

func calculateSSE(groupedValues *map[string][]float64, meansResult *anovaTableMeansResult) float64 {
	sseValue := float64(0.0)
	for key, values := range *groupedValues {
		groupSSEValue := float64(0.0)

		for _, value := range values {
			groupSSEValue += math.Pow(value-meansResult.GroupMeans[key], 2)
		}

		sseValue += groupSSEValue
	}
	return math.Round(sseValue*100) / 100
}

func calculateSST(ssr float64, sse float64) float64 {
	return ssr + sse
}

// Calculates an Anova table on an instance of AnoveInput
func (anovaInput *AnovaInput) CalculateAnovaTable(alpha float64) AnovaTable {
	var anovaTable = AnovaTable{}
	meansResult := anovaInput.anovaTableMeansResult()
	groupedDataPoints := anovaInput.groupedDataPoints()
	groupedDataPointValues := getGroupedDataPointValues(&groupedDataPoints)
	ssr := calculateSSR(&meansResult)
	sse := calculateSSE(&groupedDataPointValues, &meansResult)
	sst := calculateSST(ssr, sse)
	// See: https://www.statology.org/one-way-anova-by-hand/
	n := float64(len(anovaInput.DataPoints))  // # of observations
	k := float64(len(meansResult.GroupMeans)) // # of results
	anovaTable.DfRegression = k - 1
	anovaTable.DfError = n - k
	anovaTable.DfTotal = n - 1
	anovaTable.MSRegression = ssr / anovaTable.DfRegression
	anovaTable.MSError = math.Round((sse/anovaTable.DfError)*100) / 100
	anovaTable.SSRegression = ssr
	anovaTable.SSError = sse
	anovaTable.SSTotal = sst

	// Round to 3x decimal places
	anovaTable.F = math.Round((anovaTable.MSRegression/anovaTable.MSError)*1000) / 1000

	// See: https://pkg.go.dev/gonum.org/v1/gonum/stat/distuv#F
	var f = distuv.F{
		D1: anovaTable.DfRegression,
		D2: anovaTable.DfError,
	}
	// Round to 6x decimal places
	anovaTable.P = math.Round((1-f.CDF(anovaTable.F))*1000000) / 1000000
	anovaTable.Significant = anovaTable.P <= alpha
	return anovaTable
}
