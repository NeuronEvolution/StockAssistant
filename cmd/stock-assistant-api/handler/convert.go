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
	r.NiWeight = p.NIWeight

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
	r.NIWeight = p.NiWeight

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
	r.EvalRemark = p.EvalRemark

	r.ExchangeID = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCN
	r.LaunchDate = strfmt.DateTime(p.LaunchDate)
	r.WebsiteURL = p.WebsiteUrl
	r.IndustryName = p.IndustryName
	r.CityNameCN = p.CityNameCN
	r.ProvinceNameCN = p.ProvinceNameCN

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
