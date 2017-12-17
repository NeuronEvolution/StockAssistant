package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/errors"
	"strconv"
)

func (s *StockAssistantService) StockIndexAdviceList(ctx context.Context, query *models.StockIndexAdviceQuery) (result []*models.StockIndexAdvice, nextPageToken string, err error) {
	start := int64(0)
	count := int64(0)

	if query.PageToken == "" {
		start = 0
	} else {
		i, err := strconv.Atoi(query.PageToken)
		if err != nil {
			return nil, "", errors.InvalidParams(&errors.ParamError{Field: "pageToken"})
		}
		start = int64(i)
	}

	if query.PageSize == 0 {
		count = 40
	} else {
		count = query.PageSize
	}

	indexMap := make(map[string]*fin_stock_assistant.UserStockIndex)
	if query.UserId != "" {
		dbIndexList, err := s.db.UserStockIndex.GetQuery().UserId_Equal(query.UserId).QueryList(ctx, nil)
		if err != nil {
			return nil, "", err
		}

		if dbIndexList != nil {
			for _, dbIndex := range dbIndexList {
				indexMap[dbIndex.IndexName] = dbIndex
			}
		}
	}

	rows, err := s.db.UserStockIndex.GetQuery().
		GroupBy(fin_stock_assistant.USER_STOCK_INDEX_FIELD_INDEX_NAME).
		OrderByGroupCount(false).
		OrderBy(fin_stock_assistant.USER_STOCK_INDEX_FIELD_INDEX_NAME, true).
		Limit(start, count).
		QueryGroupBy(ctx, nil)
	if err != nil {
		return nil, "", err
	}

	result = make([]*models.StockIndexAdvice, 0)
	for rows.Next() {
		e := &models.StockIndexAdvice{}
		err = rows.Scan(&e.IndexName, &e.UsedCount)
		if err != nil {
			return nil, "", err
		}

		_, has := indexMap[e.IndexName]
		if has {
			e.HaveUsed = true
		}

		result = append(result, e)
	}
	if rows.Err() != nil {
		return nil, "", rows.Err()
	}

	return result, fmt.Sprint(start + count), nil
}
