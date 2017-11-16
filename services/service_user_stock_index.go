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
		QueryList(context.Background(), tx)
	if err != nil {
		return err
	}

	if dbIndexEvaluateList != nil {
		for _, dbIndexEvaluate := range dbIndexEvaluateList {
			dbStockEvaluate, err := s.db.UserStockEvaluate.GetQuery().ForUpdate().
				UserId_Equal(userId).And().StockId_Equal(dbIndexEvaluate.StockId).
				QueryOne(context.Background(), tx)
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

			err = s.db.UserStockEvaluate.Update(context.Background(), tx, dbStockEvaluate)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *StockAssistantService) UserStockIndexList(userId string) (indexList []*models.UserStockIndex, err error) {
	dbIndexList, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).
		QueryList(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndexList(dbIndexList), nil
}

func (s *StockAssistantService) UserStockIndexGet(userId string, indexName string) (index *models.UserStockIndex, err error) {
	dbIndex, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).And().IndexName_Equal(indexName).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserStockIndexAdd(userId string, index *models.UserStockIndex) (indexAdded *models.UserStockIndex, err error) {
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
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

	//new index last ui
	indexCount, err := s.db.UserStockIndex.GetQuery().ForShare().
		UserId_Equal(userId).
		QueryCount(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	dbIndex = &fin_stock_assistant.UserStockIndex{}
	dbIndex.UserId = userId
	dbIndex.IndexName = index.IndexName
	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
	dbIndex.AiWeight = index.AIWeight
	dbIndex.UiOrder = int32(indexCount + 1)
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
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
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

	if index.EvalWeight != dbIndex.EvalWeight || index.AIWeight != dbIndex.AiWeight {
		err = s.reCalcStockEvaluates(context.Background(), tx,
			userId, index.IndexName, dbIndex.EvalWeight, index.EvalWeight, dbIndex.AiWeight, index.AIWeight, false)
		if err != nil {
			return nil, err
		}
	}

	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
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
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.UserStockIndex.GetQuery().ForUpdate().
		UserId_Equal(userId).And().IndexName_Equal(indexName).
		QueryOne(context.Background(), tx)
	if err != nil {
		return err
	}

	if dbIndex == nil {
		return errors.NotFound("指标不存在")
	}

	err = s.reCalcStockEvaluates(context.Background(), tx,
		userId, indexName, dbIndex.EvalWeight, 0, dbIndex.AiWeight, 0, true)
	if err != nil {
		return err
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
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
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
		UserId_Equal(userId).And().IndexName_Equal(indexNameNew).
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
		UserId_Equal(userId).And().IndexName_Equal(indexNameOld).
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

	return fin_stock_assistant.FromStockIndex(dbIndexOld), err
}
