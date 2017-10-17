package handler

import (
	api "github.com/NeuronEvolution/StockAssistant/api/models"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/go-openapi/strfmt"
)

func fromIndex(p *models.StockIndex) (r *api.StockIndex) {
	if p == nil {
		return nil
	}

	r = &api.StockIndex{}
	r.ID = p.IndexId
	r.Name = p.IndexName
	r.Desc = p.IndexDesc
	r.EvalWeight = p.EvalWeight
	r.AiWeight = p.AIWeight
	r.NiWeight = p.NIWeight

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
	r.IndexID = p.IndexId
	r.EvalStars = p.EvalStars
	r.EvalRemark = p.EvalRemark
	r.UpdateTime = strfmt.DateTime(p.UpdateTime)

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
	r.IndexEvaluates = fromIndexEvaluateList(p.IndexEvaluates)

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
