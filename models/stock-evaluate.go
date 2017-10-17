package models

type StockEvaluate struct {
	StockId        string
	TotalScore     float64
	EvalRemark     string
	IndexEvaluates []*IndexEvaluate
}
