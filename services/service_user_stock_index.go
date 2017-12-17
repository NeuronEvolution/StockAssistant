package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/sql/wrap"
	"time"
)

func (s *StockAssistantService) reCalcStockEvaluates(ctx context.Context, tx *wrap.Tx,
	userId string, indexName string,
	evalWeightOld int32, evalWeightNew int32,
	aiWeightOld int32, aiWeightNew int32, deleting bool) (err error) {
	dbIndexEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().ForShare().
		UserId_Equal(userId).And().IndexName_Equal(indexName).
		QueryList(ctx, tx)
	if err != nil {
		return err
	}

	if dbIndexEvaluateList != nil {
		for _, dbIndexEvaluate := range dbIndexEvaluateList {
			dbStockEvaluate, err := s.db.UserStockEvaluate.GetQuery().ForUpdate().
				UserId_Equal(userId).And().StockId_Equal(dbIndexEvaluate.StockId).
				QueryOne(ctx, tx)
			if err != nil {
				return err
			}
			if dbStockEvaluate == nil {
				return errors.InternalServerError("指标评估存在，但股票评估不存在 userId=" +
					userId + ",stockId=" + dbIndexEvaluate.StockId + ",indexName=" + dbIndexEvaluate.IndexName)
			}

			dbStockEvaluate.UpdateTime = time.Now()
			dbStockEvaluate.TotalScore += float64(evalWeightNew - evalWeightOld + aiWeightNew - aiWeightOld)
			if deleting {
				dbStockEvaluate.IndexCount--
			}

			err = s.db.UserStockEvaluate.Update(ctx, tx, dbStockEvaluate)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *StockAssistantService) UserStockIndexList(ctx context.Context, userId string) (indexList []*models.UserStockIndex, err error) {
	dbIndexList, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).
		QueryList(ctx, nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndexList(dbIndexList), nil
}

func (s *StockAssistantService) UserStockIndexGet(ctx context.Context, userId string, indexName string) (index *models.UserStockIndex, err error) {
	dbIndex, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).And().IndexName_Equal(indexName).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserStockIndexAdd(ctx context.Context, userId string, index *models.UserStockIndex) (indexAdded *models.UserStockIndex, err error) {
	var dbUserStockIndex *fin_stock_assistant.UserStockIndex
	err = s.db.TransactionReadCommitted(ctx, func(tx *wrap.Tx) (err error) {
		dbUserStockIndex, err = s.db.UserStockIndex.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(index.IndexName).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbUserStockIndex != nil {
			return errors.AlreadyExists("指标已经存在")
		}

		//new index last ui
		indexCount, err := s.db.UserStockIndex.GetQuery().ForShare().
			UserId_Equal(userId).
			QueryCount(ctx, tx)
		if err != nil {
			return err
		}

		dbUserStockIndex = &fin_stock_assistant.UserStockIndex{}
		dbUserStockIndex.UserId = userId
		dbUserStockIndex.IndexName = index.IndexName
		dbUserStockIndex.IndexDesc = index.IndexDesc
		dbUserStockIndex.EvalWeight = index.EvalWeight
		dbUserStockIndex.AiWeight = index.AIWeight
		dbUserStockIndex.UiOrder = int32(indexCount + 1)
		dbUserStockIndex.UpdateTime = time.Now()
		dbUserStockIndex.CreateTime = time.Now()
		_, err = s.db.UserStockIndex.Insert(ctx, tx, dbUserStockIndex)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbUserStockIndex), nil
}

func (s *StockAssistantService) UserStockIndexUpdate(ctx context.Context, userId string, index *models.UserStockIndex) (indexUpdated *models.UserStockIndex, err error) {
	var dbUserStockIndex *fin_stock_assistant.UserStockIndex
	err = s.db.TransactionReadCommitted(ctx, func(tx *wrap.Tx) (err error) {
		dbUserStockIndex, err = s.db.UserStockIndex.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(index.IndexName).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbUserStockIndex == nil {
			return errors.NotFound("指标不存在")
		}

		if index.EvalWeight != dbUserStockIndex.EvalWeight || index.AIWeight != dbUserStockIndex.AiWeight {
			err = s.reCalcStockEvaluates(ctx, tx,
				userId, index.IndexName, dbUserStockIndex.EvalWeight, index.EvalWeight,
				dbUserStockIndex.AiWeight, index.AIWeight, false)
			if err != nil {
				return err
			}
		}

		dbUserStockIndex.IndexDesc = index.IndexDesc
		dbUserStockIndex.EvalWeight = index.EvalWeight
		dbUserStockIndex.AiWeight = index.AIWeight
		dbUserStockIndex.UpdateTime = time.Now()
		err = s.db.UserStockIndex.Update(ctx, tx, dbUserStockIndex)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbUserStockIndex), nil
}

func (s *StockAssistantService) UserStockIndexDelete(ctx context.Context, userId string, indexName string) (err error) {
	err = s.db.TransactionReadCommitted(ctx, func(tx *wrap.Tx) (err error) {
		dbIndex, err := s.db.UserStockIndex.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(indexName).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbIndex == nil {
			return errors.NotFound("指标不存在")
		}

		err = s.reCalcStockEvaluates(ctx, tx,
			userId, indexName, dbIndex.EvalWeight, 0, dbIndex.AiWeight, 0, true)
		if err != nil {
			return err
		}

		err = s.db.UserStockIndex.Delete(ctx, tx, dbIndex.Id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *StockAssistantService) UserStockIndexRename(ctx context.Context, userId string, indexNameOld string, indexNameNew string) (indexNew *models.UserStockIndex, err error) {
	var dbIndexOld *fin_stock_assistant.UserStockIndex
	err = s.db.TransactionReadCommitted(ctx, func(tx *wrap.Tx) (err error) {
		if indexNameNew == indexNameOld {
			return fmt.Errorf("name same")
		}

		dbIndexOld, err = s.db.UserStockIndex.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(indexNameOld).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbIndexOld == nil {
			return fmt.Errorf("index old not exist")
		}

		dbIndexNew, err := s.db.UserStockIndex.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(indexNameNew).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbIndexNew != nil {
			return fmt.Errorf("index new exist")
		}

		dbIndexOld.IndexName = indexNameNew
		dbIndexOld.UpdateTime = time.Now()
		err = s.db.UserStockIndex.Update(ctx, tx, dbIndexOld)
		if err != nil {
			return err
		}

		//update index evaluates
		dbIndexEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().ForUpdate().
			UserId_Equal(userId).And().IndexName_Equal(indexNameOld).
			QueryList(ctx, tx)
		if err != nil {
			return err
		}
		if dbIndexEvaluateList != nil {
			for _, v := range dbIndexEvaluateList {
				v.IndexName = indexNameNew
				err := s.db.UserIndexEvaluate.Update(ctx, tx, v)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndexOld), err
}
