package fin_stock_assistant

import (
	"github.com/NeuronEvolution/StockAssistant/models"
)

func FromStockIndex(p *UserStockIndex) (r *models.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = &models.UserStockIndex{}
	r.IndexName = p.IndexName
	r.IndexDesc = p.IndexDesc
	r.EvalWeight = p.EvalWeight
	r.AIWeight = p.AiWeight

	return r
}

func FromStockIndexList(p []*UserStockIndex) (r []*models.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserStockIndex, len(p))
	for i, v := range p {
		r[i] = FromStockIndex(v)
	}

	return r
}

func FromStockEvaluate(p *UserStockEvaluate) (r *models.UserStockEvaluate) {
	if p == nil {
		return r
	}

	r = &models.UserStockEvaluate{}
	r.StockId = p.StockId
	r.TotalScore = p.TotalScore
	r.EvalRemark = p.EvalRemark
	r.ExchangeId = p.ExchangeId
	r.StockCode = p.StockCode
	r.StockNameCN = p.StockNameCn
	r.LaunchDate = p.LaunchDate
	r.IndustryName = p.IndustryName

	return r
}

func FromStockEvaluateList(p []*UserStockEvaluate) (r []*models.UserStockEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserStockEvaluate, len(p))
	for i, v := range p {
		r[i] = FromStockEvaluate(v)
	}

	return r
}

func FromIndexEvaluate(p *UserIndexEvaluate) (r *models.UserIndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &models.UserIndexEvaluate{}
	r.IndexName = p.IndexName
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = p.UpdateTime

	return r
}

func FromIndexEvaluateList(p []*UserIndexEvaluate) (r []*models.UserIndexEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserIndexEvaluate, len(p))
	for i, v := range p {
		r[i] = FromIndexEvaluate(v)
	}

	return r
}

func FromSetting(p *UserSetting) (r *models.UserSetting) {
	if p == nil {
		return nil
	}

	r = &models.UserSetting{}
	r.ConfigKey = p.ConfigKey
	r.ConfigValue = p.ConfigValue

	return r
}

func FromSettingList(p []*UserSetting) (r []*models.UserSetting) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserSetting, len(p))
	for i, v := range p {
		r[i] = FromSetting(v)
	}

	return r
}

func FromStockIndexAdvice(p *StockIndexAdvice) (r *models.StockIndexAdvice) {
	if p == nil {
		return nil
	}

	r = &models.StockIndexAdvice{}
	r.IndexName = p.IndexName
	r.UsedCount = p.UsedCount

	return r
}

func FromStockIndexAdviceList(p []*StockIndexAdvice) (r []*models.StockIndexAdvice) {
	if p == nil {
		return nil
	}

	r = make([]*models.StockIndexAdvice, len(p))
	for i, v := range p {
		r[i] = FromStockIndexAdvice(v)
	}

	return r
}
