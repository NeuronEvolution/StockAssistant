package handler

import (
	api "github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/go-openapi/strfmt"
	"time"
)

func fromIndex(p *models.UserStockIndex) (r *api.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = &api.UserStockIndex{}
	r.Name = p.IndexName
	r.Desc = p.IndexDesc
	r.EvalWeight = p.EvalWeight
	r.AiWeight = p.AIWeight

	return r
}

func toIndex(p *api.UserStockIndex) (r *models.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = &models.UserStockIndex{}
	r.IndexName = p.Name
	r.IndexDesc = p.Desc
	r.EvalWeight = p.EvalWeight
	r.AIWeight = p.AiWeight

	return r
}

func fromIndexList(p []*models.UserStockIndex) (r []*api.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = make([]*api.UserStockIndex, len(p))
	for i, v := range p {
		r[i] = fromIndex(v)
	}

	return r
}

func fromIndexEvaluate(p *models.UserIndexEvaluate) (r *api.UserIndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &api.UserIndexEvaluate{}
	r.IndexName = p.IndexName
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = strfmt.DateTime(p.UpdateTime)

	return r
}

func toIndexEvaluate(p *api.UserIndexEvaluate) (r *models.UserIndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &models.UserIndexEvaluate{}
	r.IndexName = p.IndexName
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = time.Time(p.UpdateTime)

	return r
}

func fromIndexEvaluateList(p []*models.UserIndexEvaluate) (r []*api.UserIndexEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*api.UserIndexEvaluate, len(p))
	for i, v := range p {
		r[i] = fromIndexEvaluate(v)
	}

	return r
}

func fromStockEvaluate(p *models.UserStockEvaluate) (r *api.UserStockEvaluate) {
	if p == nil {
		return nil
	}

	r = &api.UserStockEvaluate{}
	r.StockID = p.StockId
	r.TotalScore = p.TotalScore
	r.IndexCount = p.IndexCount
	r.EvalRemark = p.EvalRemark

	r.ExchangeID = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCN
	r.LaunchDate = strfmt.DateTime(p.LaunchDate)
	r.IndustryName = p.IndustryName

	return r
}

func fromStockEvaluateList(p []*models.UserStockEvaluate) (r []*api.UserStockEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*api.UserStockEvaluate, len(p))
	for i, v := range p {
		r[i] = fromStockEvaluate(v)
	}

	return r
}

func fromUserSetting(p *models.UserSetting) (r *api.UserSetting) {
	if p == nil {
		return nil
	}

	r = &api.UserSetting{}
	r.Key = p.ConfigKey
	r.Value = p.ConfigValue

	return r
}

func toUserSetting(p *api.UserSetting) (r *models.UserSetting) {
	if p == nil {
		return nil
	}

	r = &models.UserSetting{}
	r.ConfigKey = p.Key
	r.ConfigValue = p.Value

	return r
}

func fromUserSettingList(p []*models.UserSetting) (r []*api.UserSetting) {
	if p == nil {
		return nil
	}

	r = make([]*api.UserSetting, len(p))
	for i, v := range p {
		r[i] = fromUserSetting(v)
	}

	return r
}

func fromStockIndexAdvice(p *models.StockIndexAdvice) (r *api.StockIndexAdvice) {
	if p == nil {
		return nil
	}

	r = &api.StockIndexAdvice{}
	r.IndexName = p.IndexName
	r.UsedCount = p.UsedCount
	r.HaveUsed = p.HaveUsed

	return r
}

func fromStockIndexAdviceList(p []*models.StockIndexAdvice) (r []*api.StockIndexAdvice) {
	if p == nil {
		return nil
	}

	r = make([]*api.StockIndexAdvice, len(p))
	for i, v := range p {
		r[i] = fromStockIndexAdvice(v)
	}

	return r
}

func fromStockUrl(p *models.StockUrl) (r *api.StockURL) {
	if p == nil {
		return nil
	}

	r = &api.StockURL{}
	r.Name = p.Name
	r.Icon = p.Icon
	r.URL = p.Url

	return r
}

func fromStockUrlList(p []*models.StockUrl) (r []*api.StockURL) {
	if p == nil {
		return nil
	}

	r = make([]*api.StockURL, len(p))
	for i, v := range p {
		r[i] = fromStockUrl(v)
	}

	return r
}

func fromStock(p *models.Stock) (r *api.Stock) {
	if p == nil {
		return nil
	}

	r = &api.Stock{}
	r.StockID = p.StockID
	r.ExchangeID = p.ExchangeID
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCN
	r.LaunchDate = strfmt.DateTime(p.LaunchDate)
	r.IndustryName = p.IndustryName
	r.CityNameCN = p.CityNameCN
	r.ProvinceNameCN = p.ProvinceNameCN
	r.WebsiteURL = p.WebsiteURL
	r.StockURLList = fromStockUrlList(p.StockUrlList)

	return r
}
