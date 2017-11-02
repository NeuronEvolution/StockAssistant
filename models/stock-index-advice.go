package models

type StockIndexAdvice struct {
	IndexName string
	UsedCount int64
}

type StockIndexAdviceQuery struct {
	PageToken string
	PageSize  int32
}
