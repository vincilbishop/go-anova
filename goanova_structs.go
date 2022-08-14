package goanova

type AnovaDataPoint struct {
	Value float64
	Group string
}

type AnovaInput struct {
	DataPoints []AnovaDataPoint
}

type AnovaTable struct {
	SSRegression, SSTotal, SSError float64
	DfRegression, DfTotal, DfError float64
	MSRegression, MSTotal, MSError float64
	F                              float64
	P                              float64
	Significant                    bool
}
