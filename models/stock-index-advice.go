package models

type StockIndexAdvice struct {
	IndexName string
	UsedCount int64
	HaveUsed  bool
}

type StockIndexAdviceQuery struct {
	UserId    string
	PageToken string
	PageSize  int64
}
