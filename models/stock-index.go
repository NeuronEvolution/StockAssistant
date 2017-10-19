package models

type StockIndex struct {
	IndexName  string
	IndexDesc  string
	EvalWeight int32
	AIWeight   int32
	NIWeight   int32
}
