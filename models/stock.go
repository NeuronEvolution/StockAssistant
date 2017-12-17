package models

import (
	"time"
)

type StockUrl struct {
	Name string
	Icon string
	Url  string
}

type Stock struct {
	StockID        string
	ExchangeID     string
	StockCode      string
	StockNameCN    string
	LaunchDate     time.Time
	IndustryName   string
	CityNameCN     string
	ProvinceNameCN string
	WebsiteURL     string
	StockUrlList   []*StockUrl
}
