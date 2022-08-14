package goanova

type AnovaDataPoint struct {
	Value float64
	Group string
}

type AnovaInput struct {
	DataPoints []AnovaDataPoint
}

type AnovaTable struct {
	SSTreatment, SSTotal, SSError float64
	DfTreatment, DfTotal, DfError float64
	MSTreatment, MSTotal, MSError float64
	F                             float64
	P                             float64
}
