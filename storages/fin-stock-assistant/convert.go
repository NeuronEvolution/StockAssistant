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
	r.NIWeight = p.NiWeight

	return r
}

func FromStockIndexList(p []*UserStockIndex) (r []*models.UserStockIndex) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserStockIndex, 0)
	for _, v := range p {
		r = append(r, FromStockIndex(v))
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

	return r
}

func FromStockEvaluateList(p []*UserStockEvaluate) (r []*models.UserStockEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*models.UserStockEvaluate, 0)
	for _, v := range p {
		r = append(r, FromStockEvaluate(v))
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

	r = make([]*models.UserIndexEvaluate, 0)
	for _, v := range p {
		r = append(r, FromIndexEvaluate(v))
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

	r = make([]*models.UserSetting, 0)
	for _, v := range p {
		r = append(r, FromSetting(v))
	}

	return r
}
