package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/errors"
	"strconv"
)

func (s *StockAssistantService) StockIndexAdviceList(query *models.StockIndexAdviceQuery) (result []*models.StockIndexAdvice, nextPageToken string, err error) {
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

	rows, err := s.db.UserStockIndex.GetQuery().
		GroupBy(fin_stock_assistant.USER_STOCK_INDEX_FIELD_INDEX_NAME).
		OrderByGroupCount(
			nil,
			[]fin_stock_assistant.USER_STOCK_INDEX_FIELD{fin_stock_assistant.USER_STOCK_INDEX_FIELD_INDEX_NAME},
			false).
		Limit(start, count).
		QueryGroupBy(context.Background(), nil)
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

		result = append(result, e)
	}
	if rows.Err() != nil {
		return nil, "", rows.Err()
	}

	return result, fmt.Sprint(start + count), nil
}
