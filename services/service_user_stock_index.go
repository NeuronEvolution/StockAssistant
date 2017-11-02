package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/errors"
	"time"
)

func (s *StockAssistantService) UserStockIndexList(userId string) (indexList []*models.UserStockIndex, err error) {
	dbIndexList, err := s.db.UserStockIndex.GetQuery().QueryList(context.Background(),nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndexList(dbIndexList), nil
}

func (s *StockAssistantService) UserStockIndexGet(userId string, indexName string) (index *models.UserStockIndex, err error) {
	dbIndex, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).And().IndexName_Equal(indexName).QueryOne(context.Background(),nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserStockIndexAdd(userId string, index *models.UserStockIndex) (indexAdded *models.UserStockIndex, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).And().IndexName_Equal(index.IndexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndex != nil {
		return nil, errors.AlreadyExists("指标已经存在")
	}

	dbIndex = &fin_stock_assistant.UserStockIndex{}
	dbIndex.UserId = userId
	dbIndex.IndexName = index.IndexName
	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
	dbIndex.NiWeight = index.NIWeight
	dbIndex.AiWeight = index.AIWeight
	dbIndex.UpdateTime = time.Now()
	dbIndex.CreateTime = time.Now()
	_, err = s.db.UserStockIndex.Insert(context.Background(), tx, dbIndex)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserStockIndexUpdate(userId string, index *models.UserStockIndex) (indexUpdated *models.UserStockIndex, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).And().IndexName_Equal(index.IndexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndex == nil {
		return nil, errors.NotFound("指标不存在")
	}

	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
	dbIndex.NiWeight = index.NIWeight
	dbIndex.AiWeight = index.AIWeight
	dbIndex.UpdateTime = time.Now()
	err = s.db.UserStockIndex.Update(context.Background(), tx, dbIndex)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserStockIndexDelete(userId string, indexName string) (err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).And().IndexName_Equal(indexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil
	}

	if dbIndex == nil {
		return fmt.Errorf("not exist")
	}

	err = s.db.UserStockIndex.Delete(context.Background(), tx, dbIndex.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockAssistantService) UserStockIndexRename(userId string, indexNameOld string, indexNameNew string) (indexNew *models.UserStockIndex, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if indexNameNew == indexNameOld {
		return nil, fmt.Errorf("name same")
	}

	dbIndexOld, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).And().IndexName_Equal(indexNameOld).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexOld == nil {
		return nil, fmt.Errorf("index old not exist")
	}

	dbIndexNew, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).
		IndexName_Equal(indexNameNew).
		QueryOne(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexNew != nil {
		return nil, fmt.Errorf("index new exist")
	}

	dbIndexOld.IndexName = indexNameNew
	dbIndexOld.UpdateTime = time.Now()
	err = s.db.UserStockIndex.Update(context.Background(), tx, dbIndexOld)
	if err != nil {
		return nil, err
	}

	//update index evaluates
	dbIndexEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().ForUpdate().
		UserId_Equal(userId).
		QueryList(context.Background(), tx)
	if err != nil {
		return nil, err
	}
	if dbIndexEvaluateList != nil {
		for _, v := range dbIndexEvaluateList {
			v.IndexName = indexNameNew
			err := s.db.UserIndexEvaluate.Update(context.Background(), tx, v)
			if err != nil {
				return nil, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return nil, err
}
