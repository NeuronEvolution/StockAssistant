package handler

import (
	api "github.com/NeuronEvolution/StockAssistant/api/gen/models"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/go-openapi/strfmt"
	"time"
)

func fromIndex(p *models.StockIndex) (r *api.StockIndex) {
	if p == nil {
		return nil
	}

	r = &api.StockIndex{}
	r.Name = p.IndexName
	r.Desc = p.IndexDesc
	r.EvalWeight = p.EvalWeight
	r.AiWeight = p.AIWeight
	r.NiWeight = p.NIWeight

	return r
}

func toIndex(p *api.StockIndex) (r *models.StockIndex) {
	if p == nil {
		return nil
	}

	r = &models.StockIndex{}
	r.IndexName = p.Name
	r.IndexDesc = p.Desc
	r.EvalWeight = p.EvalWeight
	r.AIWeight = p.AiWeight
	r.NIWeight = p.NiWeight

	return r
}

func fromIndexList(p []*models.StockIndex) (r []*api.StockIndex) {
	if p == nil {
		return nil
	}

	r = make([]*api.StockIndex, 0)
	for _, v := range p {
		r = append(r, fromIndex(v))
	}

	return r
}

func fromIndexEvaluate(p *models.IndexEvaluate) (r *api.IndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &api.IndexEvaluate{}
	r.IndexName = p.IndexName
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = strfmt.DateTime(p.UpdateTime)

	return r
}

func toIndexEvaluate(p *api.IndexEvaluate) (r *models.IndexEvaluate) {
	if p == nil {
		return nil
	}

	r = &models.IndexEvaluate{}
	r.IndexName = p.IndexName
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = time.Time(p.UpdateTime)

	return r
}

func fromIndexEvaluateList(p []*models.IndexEvaluate) (r []*api.IndexEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*api.IndexEvaluate, 0)
	for _, v := range p {
		r = append(r, fromIndexEvaluate(v))
	}

	return r
}

func fromStockEvaluate(p *models.StockEvaluate) (r *api.StockEvaluate) {
	if p == nil {
		return nil
	}

	r = &api.StockEvaluate{}
	r.StockID = p.StockId
	r.TotalScore = p.TotalScore

	return r
}

func toStockEvaluate(p *api.StockEvaluate) (r *models.StockEvaluate) {
	if p == nil {
		return nil
	}

	r = &models.StockEvaluate{}
	r.StockId = p.StockID
	r.TotalScore = p.TotalScore

	return r
}

func fromStockEvaluateList(p []*models.StockEvaluate) (r []*api.StockEvaluate) {
	if p == nil {
		return nil
	}

	r = make([]*api.StockEvaluate, 0)
	for _, v := range p {
		r = append(r, fromStockEvaluate(v))
	}

	return r
}

func fromUserSetting(p *models.Setting) (r *api.Setting) {
	if p == nil {
		return nil
	}

	r = &api.Setting{}
	r.Key = p.ConfigKey
	r.Value = p.ConfigValue

	return r
}

func toUserSetting(p *api.Setting) (r *models.Setting) {
	if p == nil {
		return nil
	}

	r = &models.Setting{}
	r.ConfigKey = p.Key
	r.ConfigValue = p.Value

	return r
}

func fromUserSettingList(p []*models.Setting) (r []*api.Setting) {
	if p == nil {
		return nil
	}

	r = make([]*api.Setting, 0)
	for _, v := range p {
		r = append(r, fromUserSetting(v))
	}

	return r
}
