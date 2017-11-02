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
	}

	dbList, err := s.db.StockIndexAdvice.GetQuery().Limit(start, count).QueryList(context.Background(),nil)
	if err != nil {
		return nil, "", err
	}

	return fin_stock_assistant.FromStockIndexAdviceList(dbList), fmt.Sprint(start + count), nil
}
