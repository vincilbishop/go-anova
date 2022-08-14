package goanova

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetOneWayAnovaInput() AnovaInput {
	// Values taken from here: https://www.statology.org/one-way-anova-by-hand/
	oneWayAnoveInput := AnovaInput{
		DataPoints: []AnovaDataPoint{
			{
				Value: 85,
				Group: "Group1",
			},
			{
				Value: 86,
				Group: "Group1",
			}, {
				Value: 88,
				Group: "Group1",
			}, {
				Value: 75,
				Group: "Group1",
			}, {
				Value: 78,
				Group: "Group1",
			}, {
				Value: 94,
				Group: "Group1",
			}, {
				Value: 98,
				Group: "Group1",
			}, {
				Value: 79,
				Group: "Group1",
			}, {
				Value: 71,
				Group: "Group1",
			}, {
				Value: 80,
				Group: "Group1",
			},
			{
				Value: 91,
				Group: "Group2",
			}, {
				Value: 92,
				Group: "Group2",
			}, {
				Value: 93,
				Group: "Group2",
			}, {
				Value: 85,
				Group: "Group2",
			}, {
				Value: 87,
				Group: "Group2",
			}, {
				Value: 84,
				Group: "Group2",
			}, {
				Value: 82,
				Group: "Group2",
			}, {
				Value: 88,
				Group: "Group2",
			}, {
				Value: 95,
				Group: "Group2",
			}, {
				Value: 96,
				Group: "Group2",
			}, {
				Value: 79,
				Group: "Group3",
			}, {
				Value: 78,
				Group: "Group3",
			}, {
				Value: 88,
				Group: "Group3",
			}, {
				Value: 94,
				Group: "Group3",
			}, {
				Value: 92,
				Group: "Group3",
			}, {
				Value: 85,
				Group: "Group3",
			}, {
				Value: 83,
				Group: "Group3",
			}, {
				Value: 85,
				Group: "Group3",
			}, {
				Value: 82,
				Group: "Group3",
			}, {
				Value: 81,
				Group: "Group3",
			},
		},
	}
	return oneWayAnoveInput
}

func TestGroupedDataPoints(t *testing.T) {
	input := GetOneWayAnovaInput()
	actual := input.groupedDataPoints()
	assert.NotNil(t, actual)

	for _, value := range actual {
		assert.NotNil(t, value)
		assert.Equal(t, 10, len(value))
	}
}

func TestGetGroupedDataPointValues(t *testing.T) {
	input := GetOneWayAnovaInput()
	groupedDataPoints := input.groupedDataPoints()
	actual := getGroupedDataPointValues(&groupedDataPoints)
	assert.NotNil(t, actual)

	for _, value := range actual {
		assert.NotNil(t, value)
		assert.Equal(t, 10, len(value))
	}
}

func TestAnovaTableMeansResult(t *testing.T) {
	input := GetOneWayAnovaInput()
	actual := input.anovaTableMeansResult()
	assert.NotNil(t, actual)
	assert.Equal(t, 83.4, actual.GroupMeans["Group1"])
	assert.Equal(t, 89.3, actual.GroupMeans["Group2"])
	assert.Equal(t, 84.7, actual.GroupMeans["Group3"])
	assert.Equal(t, 85.8, actual.OverallMean)
}

func TestCalculateSSR(t *testing.T) {
	input := GetOneWayAnovaInput()
	meansResult := input.anovaTableMeansResult()
	actual := calculateSSR(&meansResult)
	assert.NotNil(t, actual)
	assert.Equal(t, 192.2, actual)
}

func TestCalculateSSE(t *testing.T) {
	input := GetOneWayAnovaInput()
	meansResult := input.anovaTableMeansResult()
	groupedDataPoints := input.groupedDataPoints()
	groupedDataPointValues := getGroupedDataPointValues(&groupedDataPoints)
	actual := calculateSSE(&groupedDataPointValues, &meansResult)
	assert.NotNil(t, actual)
	assert.Equal(t, 1100.6, actual)
}

func TestCalculateSST(t *testing.T) {
	input := GetOneWayAnovaInput()
	meansResult := input.anovaTableMeansResult()
	groupedDataPoints := input.groupedDataPoints()
	groupedDataPointValues := getGroupedDataPointValues(&groupedDataPoints)
	ssr := calculateSSR(&meansResult)
	sse := calculateSSE(&groupedDataPointValues, &meansResult)
	actual := calculateSST(ssr, sse)
	assert.NotNil(t, actual)
	assert.Equal(t, 1292.8, actual)

}

func TestCalculateAnovaResult(t *testing.T) {
	input := GetOneWayAnovaInput()
	actual := input.CalculateAnovaTable()
	assert.NotNil(t, actual)
	assert.Equal(t, float64(2), actual.DfTreatment)
	assert.Equal(t, float64(27), actual.DfError)
	assert.Equal(t, float64(29), actual.DfTotal)
	assert.Equal(t, float64(96.1), actual.MSTreatment)
	assert.Equal(t, float64(40.76), actual.MSError)
	assert.Equal(t, float64(192.2), actual.SSTreatment)
	assert.Equal(t, float64(1100.6), actual.SSError)
	assert.Equal(t, float64(1292.8), actual.SSTotal)

	assert.Equal(t, float64(2.358), actual.F)
}
