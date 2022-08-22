package goanova

import "sync"

type AnovaDataPoint struct {
	Value float64
	Group string
}

type AnovaInput struct {
	mu         sync.Mutex
	DataPoints []AnovaDataPoint
}

func (anoveInput *AnovaInput) Append(dataPoint *AnovaDataPoint) {
	anoveInput.mu.Lock()
	anoveInput.DataPoints = append(anoveInput.DataPoints, *dataPoint)
	anoveInput.mu.Unlock()
}

func (anoveInput *AnovaInput) Reset() {
	anoveInput.mu.Lock()
	anoveInput.DataPoints = nil
	anoveInput.mu.Unlock()
}

func (anoveInput *AnovaInput) Read() *[]AnovaDataPoint {
	anoveInput.mu.Lock()
	// Lock so only one goroutine at a time can access the sc.DataPoints.
	defer anoveInput.mu.Unlock()
	return &anoveInput.DataPoints
}

type AnovaTable struct {
	SSRegression, SSTotal, SSError float64
	DfRegression, DfTotal, DfError float64
	MSRegression, MSTotal, MSError float64
	F                              float64
	P                              float64
	Significant                    bool
}
