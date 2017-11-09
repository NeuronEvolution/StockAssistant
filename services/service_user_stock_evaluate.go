package services

import (
	"context"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
)

func (s *StockAssistantService) UserStockEvaluateList(query *models.UserStockEvaluateListQuery) (result []*models.UserStockEvaluate, nextPageToken string, err error) {
	if query.PageSize == 0 {
		query.PageSize = 40
	}

	if query.Evaluated {
		dbStockEvaluateList, err := s.db.UserStockEvaluate.GetQuery().
			UserId_Equal(query.UserId).
			OrderBy(fin_stock_assistant.USER_STOCK_EVALUATE_FIELD_TOTAL_SCORE, false).
			QueryList(context.Background(), nil)
		if err != nil {
			return nil, "", err
		}

		return fin_stock_assistant.FromStockEvaluateList(dbStockEvaluateList), "", nil
	} else {
		dbStockList, err := s.db.Stock.GetQuery().QueryList(context.Background(), nil)
		if err != nil {
			return nil, "", err
		}
		if dbStockList == nil {
			return nil, "", nil
		}

		//filter todo
		//sort todo

		//filter evaluated
		dbStockEvaluateList, err := s.db.UserStockEvaluate.GetQuery().
			UserId_Equal(query.UserId).
			QueryList(context.Background(), nil)
		if err != nil {
			return nil, "", err
		}
		evaluatedMap := make(map[string]bool)
		if dbStockEvaluateList != nil {
			for _, v := range dbStockEvaluateList {
				evaluatedMap[v.StockId] = true
			}
		}

		//limit select
		result = make([]*models.UserStockEvaluate, 0)
		for i := 0; i < len(dbStockList); i++ {
			v := dbStockList[i]
			_, has := evaluatedMap[v.StockId]
			if has {
				continue
			}

			e := &models.UserStockEvaluate{}
			e.StockId = v.StockId
			e.TotalScore = 0
			e.EvalRemark = ""
			e.ExchangeId = v.ExchangeId
			e.StockCode = v.StockCode
			e.StockNameCN = v.StockNameCn
			e.LaunchDate = v.LaunchDate
			e.IndustryName = v.IndustryName

			result = append(result, e)

			if len(result) >= int(query.PageSize) {
				break
			}
		}

		return result, "", nil
	}
}

func (s *StockAssistantService) UserStockEvaluateGet(userId string, stockId string) (eval *models.UserStockEvaluate, err error) {
	dbStockEvaluate, err := s.db.UserStockEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).
		QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockEvaluate(dbStockEvaluate), nil
}
