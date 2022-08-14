package goanova

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
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

func groupedDataPointValues(groupedDataPoints *map[string][]AnovaDataPoint) map[string][]float64 {
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
	groupedValues := groupedDataPointValues(&groupedDataPoints)
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
