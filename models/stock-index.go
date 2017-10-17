package models

type StockIndex struct {
	IndexId    string
	IndexName  string
	IndexDesc  string
	EvalWeight int32
	AIWeight   int32
	NIWeight   int32
}
