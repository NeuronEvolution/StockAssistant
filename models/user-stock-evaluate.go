package models

import "time"

type UserStockEvaluate struct {
	StockId    string
	TotalScore float64
	EvalRemark string

	ExchangeId     string
	StockCode      string
	StockNameCN    string
	LaunchDate     time.Time
	WebsiteUrl     string
	IndustryName   string
	CityNameCN     string
	ProvinceNameCN string
}

type UserStockEvaluateListQuery struct {
	UserId    string
	Evaluated bool
	Sort      string
	PageToken string
	PageSize  int32
}
