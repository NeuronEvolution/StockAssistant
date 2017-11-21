package services

import (
	"context"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/errors"
	"strconv"
)

func (s *StockAssistantService) UserStockEvaluateList(query *models.UserStockEvaluateListQuery) (result []*models.UserStockEvaluate, nextPageToken string, err error) {
	start := int64(0)
	if query.PageToken != "" {
		iStart, err := strconv.Atoi(query.PageToken)
		if err != nil {
			return nil, "", errors.InvalidParams(&errors.ParamError{Field: "PageToken", Code: "InvalidFormat", Message: "PageToken错误"})
		}
		start = int64(iStart)
	}

	count := int64(0)
	if query.PageSize > 0 {
		count = int64(query.PageSize)
	}

	if query.Evaluated {
		dbStockEvaluateList, err := s.db.UserStockEvaluate.GetQuery().
			UserId_Equal(query.UserId).
			OrderBy(fin_stock_assistant.USER_STOCK_EVALUATE_FIELD_TOTAL_SCORE, false).Limit(start, count).
			QueryList(context.Background(), nil)
		if err != nil {
			return nil, "", err
		}

		result = fin_stock_assistant.FromStockEvaluateList(dbStockEvaluateList)

		if len(result) < int(count) {
			nextPageToken = ""
		} else {
			nextPageToken = strconv.Itoa(int(start + count))
		}

		return result, nextPageToken, nil
	} else {
		dbStockList, err := s.db.Stock.GetQuery().OrderBy(fin_stock_assistant.STOCK_FIELD_STOCK_CODE, true).QueryList(context.Background(), nil)
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

		skipped := int64(0)

		//limit select
		result = make([]*models.UserStockEvaluate, 0)
		for i := 0; i < len(dbStockList); i++ {
			v := dbStockList[i]
			_, has := evaluatedMap[v.StockId]
			if has {
				continue
			}

			if skipped < start {
				skipped++
				continue
			}

			e := &models.UserStockEvaluate{}
			e.StockId = v.StockId
			e.TotalScore = 0
			e.IndexCount = 0
			e.EvalRemark = ""
			e.ExchangeId = v.ExchangeId
			e.StockCode = v.StockCode
			e.StockNameCN = v.StockNameCn
			e.LaunchDate = v.LaunchDate
			e.IndustryName = v.IndustryName

			result = append(result, e)

			if len(result) >= int(count) {
				break
			}
		}

		if len(result) < int(count) {
			nextPageToken = ""
		} else {
			nextPageToken = strconv.Itoa(int(start + count))
		}

		return result, nextPageToken, nil
	}
}

func (s *StockAssistantService) UserStockEvaluateGet(userId string, stockId string) (eval *models.UserStockEvaluate, err error) {
	dbStockEvaluate, err := s.db.UserStockEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).
		QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	stockEvaluate := fin_stock_assistant.FromStockEvaluate(dbStockEvaluate)

	dbStock, err := s.db.Stock.GetQuery().StockId_Equal(stockId).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if dbStock == nil {
		return nil, errors.NotFound("Stock not exists")
	}

	return stockEvaluate, nil
}
