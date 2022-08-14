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

func TestGroupedDataPointValues(t *testing.T) {
	input := GetOneWayAnovaInput()
	groupedDataPoints := input.groupedDataPoints()
	actual := groupedDataPointValues(&groupedDataPoints)
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

func TestOneWayAnova(t *testing.T) {

	// input := GetOneWayAnovaInput()
	// OneWayAnova()
}
