package goanova

type AnovaDataPoint struct {
	Value float64
	Group string
}

type AnovaInput struct {
	DataPoints []AnovaDataPoint
}

type AnovaResult struct {
	SSM, SST, SSE float64
	Dfm, Dft, Dfe float64
	MSM, MST, MSE float64
	F             float64
	P             float64
}
