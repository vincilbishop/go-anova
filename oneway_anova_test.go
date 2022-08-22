package goanova

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetGroupOneValues() []float64 {
	return []float64{
		85, 86, 88, 75, 78, 94, 98, 79, 71, 80,
	}
}

func GetGroupTwoValues() []float64 {
	return []float64{
		91, 92, 93, 85, 87, 84, 82, 88, 95, 96,
	}
}

func GetGroupThreeValues() []float64 {
	return []float64{
		79, 78, 88, 94, 92, 85, 83, 85, 82, 81,
	}
}

func GetOneWayAnovaInput() *AnovaInput {
	// Values taken from here: https://www.statology.org/one-way-anova-by-hand/

	oneWayAnoveInput := AnovaInput{}

	oneWayAnoveInput.AddDataPointGroup("Group1", GetGroupOneValues())
	oneWayAnoveInput.AddDataPointGroup("Group2", GetGroupTwoValues())
	oneWayAnoveInput.AddDataPointGroup("Group3", GetGroupThreeValues())

	return &oneWayAnoveInput
}

func TestAddDataPointGroup(t *testing.T) {
	input := GetOneWayAnovaInput()
	actual := input.groupedDataPoints()
	assert.NotNil(t, actual)

	for _, value := range actual {
		assert.NotNil(t, value)
		assert.Equal(t, 10, len(value))
	}
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
	alpha := 0.05 // Alpha level (designated threshold of significance)
	actual := input.CalculateAnovaTable(alpha)
	assert.NotNil(t, actual)
	assert.Equal(t, float64(2), actual.DfRegression)
	assert.Equal(t, float64(27), actual.DfError)
	assert.Equal(t, float64(29), actual.DfTotal)
	assert.Equal(t, float64(96.1), actual.MSRegression)
	assert.Equal(t, float64(40.76), actual.MSError)
	assert.Equal(t, float64(192.2), actual.SSRegression)
	assert.Equal(t, float64(1100.6), actual.SSError)
	assert.Equal(t, float64(1292.8), actual.SSTotal)

	assert.Equal(t, float64(2.358), actual.F)
	assert.Equal(t, float64(0.113803), actual.P)
	// Since this p-value is not less than 0.05
	// we fail to reject the null hypothesis
	assert.Equal(t, false, actual.Significant)
}
