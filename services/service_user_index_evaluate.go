package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"time"
)

func (s *StockAssistantService) UserIndexEvaluateList(userId string, stockId string) (result []*models.UserIndexEvaluate, err error) {
	dbList, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).QueryList(context.Background(),nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluateList(dbList), nil
}

func (s *StockAssistantService) UserIndexEvaluateGet(userId string, stockId string, indexName string) (indexEvaluate *models.UserIndexEvaluate, err error) {
	dbIndexEvaluate, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).And().IndexName_Equal(indexName).
		QueryOne(context.Background(),nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}

func (s *StockAssistantService) UserIndexEvaluateSave(userId string, stockId string, indexEvaluate *models.UserIndexEvaluate) (indexEvaluateSaved *models.UserIndexEvaluate, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.UserStockIndex.GetQuery().ForShare().
		UserId_Equal(userId).And().IndexName_Equal(indexEvaluate.IndexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndex == nil {
		return nil, fmt.Errorf("user have no this index")
	}

	dbIndexEvaluate, err := s.db.UserIndexEvaluate.GetQuery().ForUpdate().
		UserId_Equal(userId).And().StockId_Equal(stockId).And().IndexName_Equal(indexEvaluate.IndexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexEvaluate == nil {
		dbIndexEvaluate = &fin_stock_assistant.UserIndexEvaluate{}
		dbIndexEvaluate.UserId = userId
		dbIndexEvaluate.StockId = stockId
		dbIndexEvaluate.IndexName = indexEvaluate.IndexName
		dbIndexEvaluate.EvalStars = indexEvaluate.EvalStars
		dbIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
		dbIndexEvaluate.UpdateTime = time.Now()
		_, err = s.db.UserIndexEvaluate.Insert(context.Background(), tx, dbIndexEvaluate)
		if err != nil {
			return nil, err
		}
	} else {
		dbIndexEvaluate.EvalStars = indexEvaluate.EvalStars
		dbIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
		dbIndexEvaluate.UpdateTime = time.Now()
		err = s.db.UserIndexEvaluate.Update(context.Background(), tx, dbIndexEvaluate)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}
