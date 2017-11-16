package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"time"
)

func (s *StockAssistantService) UserIndexEvaluateList(userId string, stockId string) (result []*models.UserIndexEvaluate, err error) {
	dbEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).QueryList(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	dbIndexList, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).OrderBy(fin_stock_assistant.USER_STOCK_INDEX_FIELD_UI_ORDER, true).
		QueryList(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	result = make([]*models.UserIndexEvaluate, 0)
	for _, dbIndex := range dbIndexList {
		has := false
		for _, dbEvaluate := range dbEvaluateList {
			if dbEvaluate.IndexName == dbIndex.IndexName {
				result = append(result, fin_stock_assistant.FromIndexEvaluate(dbEvaluate))
				has = true
				break
			}
		}
		if !has {
			notEvaluatedIndex := &models.UserIndexEvaluate{}
			notEvaluatedIndex.IndexName = dbIndex.IndexName
			notEvaluatedIndex.EvalStars = -1
			notEvaluatedIndex.EvalRemark = ""
			notEvaluatedIndex.UpdateTime = time.Unix(0, 0)

			result = append(result, notEvaluatedIndex)
		}
	}

	return result, nil
}

func (s *StockAssistantService) UserIndexEvaluateGet(userId string, stockId string, indexName string) (indexEvaluate *models.UserIndexEvaluate, err error) {
	dbIndexEvaluate, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).And().IndexName_Equal(indexName).
		QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}

func (s *StockAssistantService) UserIndexEvaluateSave(userId string, stockId string, indexEvaluate *models.UserIndexEvaluate) (indexEvaluateSaved *models.UserIndexEvaluate, err error) {
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	//save index evaluate
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

	adding := false

	if dbIndexEvaluate == nil {
		adding = true

		dbIndexEvaluate = &fin_stock_assistant.UserIndexEvaluate{}
		dbIndexEvaluate.UserId = userId
		dbIndexEvaluate.StockId = stockId
		dbIndexEvaluate.IndexName = indexEvaluate.IndexName
		dbIndexEvaluate.EvalStars = indexEvaluate.EvalStars
		dbIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
		dbIndexEvaluate.UpdateTime = time.Now()
		dbIndexEvaluate.CreateTime = time.Now()
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

	//update stock evaluate
	dbIndexEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().ForShare().
		UserId_Equal(userId).And().StockId_Equal(stockId).
		QueryList(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexEvaluateList != nil {
		totalScore := int32(0)
		for _, v := range dbIndexEvaluateList {
			totalScore += v.EvalStars
		}

		se, err := s.db.UserStockEvaluate.GetQuery().ForUpdate().
			UserId_Equal(userId).And().StockId_Equal(stockId).
			QueryOne(context.Background(), tx)
		if err != nil {
			return nil, err
		}

		dbStock, err := s.db.Stock.GetQuery().StockId_Equal(stockId).QueryOne(context.Background(), nil)
		if err != nil {
			return nil, err
		}

		if se == nil {
			se = &fin_stock_assistant.UserStockEvaluate{}
			se.UserId = userId
			se.StockId = stockId
			se.TotalScore = float64(totalScore)
			se.IndexCount = 1
			se.EvalRemark = ""
			se.CreateTime = time.Now()
			se.UpdateTime = time.Now()
			se.ExchangeId = dbStock.ExchangeId
			se.StockCode = dbStock.StockCode
			se.StockNameCn = dbStock.StockNameCn
			se.LaunchDate = dbStock.LaunchDate
			se.IndustryName = dbStock.IndustryName
			_, err = s.db.UserStockEvaluate.Insert(context.Background(), tx, se)
			if err != nil {
				return nil, err
			}
		} else {
			se.TotalScore = float64(totalScore)
			if adding {
				se.IndexCount++
			}
			se.UpdateTime = time.Now()
			err = s.db.UserStockEvaluate.Update(context.Background(), tx, se)
			if err != nil {
				return nil, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}
