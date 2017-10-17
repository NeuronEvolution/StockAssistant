package fin_stock_assistant

import (
	"github.com/NeuronEvolution/StockAssistant/models"
)

func FromStockIndex(p *StockIndex) (r *models.StockIndex) {
	if p == nil {
		return nil
	}

	r = &models.StockIndex{}
	r.IndexId = p.IndexId
	r.IndexName = p.IndexName
	r.IndexDesc = p.IndexDesc
	r.EvalWeight = p.EvalWeight
	r.AIWeight = p.AiWeight
	r.NIWeight = p.NiWeight

	return r
}

func FromStockIndexList(p []*StockIndex) (r []*models.StockIndex) {
	if p == nil {
		return nil
	}

	r = make([]*models.StockIndex, 0)
	for _, v := range p {
		r = append(r, FromStockIndex(v))
	}

	return r
}

func FromIndexEvaluate(p *IndexEvaluate) (r *models.IndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &models.IndexEvaluate{}
	r.IndexId = p.IndexId
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = p.UpdateTime

	return r
}

func FromIndexEvaluateList(p []*IndexEvaluate) (r []*models.IndexEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*models.IndexEvaluate, 0)
	for _, v := range p {
		r = append(r, FromIndexEvaluate(v))
	}

	return r
}

func FromSetting(p *UserSetting) (r *models.Setting) {
	if p == nil {
		return nil
	}

	r = &models.Setting{}
	r.Key = p.ConfigKey
	r.Value = p.ConfigValue

	return r
}

func FromSettingList(p []*UserSetting) (r []*models.Setting) {
	if p == nil {
		return nil
	}

	r = make([]*models.Setting, 0)
	for _, v := range p {
		r = append(r, FromSetting(v))
	}

	return r
}
