package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/sql/wrap"
	"time"
)

func (s *StockAssistantService) UserIndexEvaluateList(ctx context.Context, userId string, stockId string) (result []*models.UserIndexEvaluate, err error) {
	dbEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).QueryList(ctx, nil)
	if err != nil {
		return nil, err
	}

	dbIndexList, err := s.db.UserStockIndex.GetQuery().
		UserId_Equal(userId).OrderBy(fin_stock_assistant.USER_STOCK_INDEX_FIELD_UI_ORDER, true).
		QueryList(ctx, nil)
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

func (s *StockAssistantService) UserIndexEvaluateGet(ctx context.Context, userId string, stockId string, indexName string) (indexEvaluate *models.UserIndexEvaluate, err error) {
	dbIndexEvaluate, err := s.db.UserIndexEvaluate.GetQuery().
		UserId_Equal(userId).And().StockId_Equal(stockId).And().IndexName_Equal(indexName).
		QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}

func (s *StockAssistantService) UserIndexEvaluateSave(ctx context.Context, userId string, stockId string, indexEvaluate *models.UserIndexEvaluate) (indexEvaluateSaved *models.UserIndexEvaluate, err error) {
	var dbUserIndexEvaluate *fin_stock_assistant.UserIndexEvaluate
	err = s.db.TransactionReadCommitted(ctx, func(tx *wrap.Tx) (err error) {
		//save index evaluate
		dbIndex, err := s.db.UserStockIndex.GetQuery().ForShare().
			UserId_Equal(userId).And().IndexName_Equal(indexEvaluate.IndexName).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		if dbIndex == nil {
			return fmt.Errorf("user have no this index")
		}

		dbUserIndexEvaluate, err = s.db.UserIndexEvaluate.GetQuery().ForUpdate().
			UserId_Equal(userId).And().StockId_Equal(stockId).And().IndexName_Equal(indexEvaluate.IndexName).
			QueryOne(ctx, tx)
		if err != nil {
			return err
		}

		adding := false

		if dbUserIndexEvaluate == nil {
			adding = true

			dbUserIndexEvaluate = &fin_stock_assistant.UserIndexEvaluate{}
			dbUserIndexEvaluate.UserId = userId
			dbUserIndexEvaluate.StockId = stockId
			dbUserIndexEvaluate.IndexName = indexEvaluate.IndexName
			dbUserIndexEvaluate.EvalStars = indexEvaluate.EvalStars
			dbUserIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
			dbUserIndexEvaluate.UpdateTime = time.Now()
			dbUserIndexEvaluate.CreateTime = time.Now()
			_, err = s.db.UserIndexEvaluate.Insert(ctx, tx, dbUserIndexEvaluate)
			if err != nil {
				return err
			}
		} else {
			dbUserIndexEvaluate.EvalStars = indexEvaluate.EvalStars
			dbUserIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
			dbUserIndexEvaluate.UpdateTime = time.Now()
			err = s.db.UserIndexEvaluate.Update(ctx, tx, dbUserIndexEvaluate)
			if err != nil {
				return err
			}
		}

		//update stock evaluate
		dbIndexEvaluateList, err := s.db.UserIndexEvaluate.GetQuery().ForShare().
			UserId_Equal(userId).And().StockId_Equal(stockId).
			QueryList(ctx, tx)
		if err != nil {
			return err
		}

		if dbIndexEvaluateList != nil {
			totalScore := int32(0)
			for _, v := range dbIndexEvaluateList {
				totalScore += v.EvalStars
			}

			se, err := s.db.UserStockEvaluate.GetQuery().ForUpdate().
				UserId_Equal(userId).And().StockId_Equal(stockId).
				QueryOne(ctx, tx)
			if err != nil {
				return err
			}

			dbStock, err := s.db.Stock.GetQuery().StockId_Equal(stockId).QueryOne(ctx, nil)
			if err != nil {
				return err
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
				_, err = s.db.UserStockEvaluate.Insert(ctx, tx, se)
				if err != nil {
					return err
				}
			} else {
				se.TotalScore = float64(totalScore)
				if adding {
					se.IndexCount++
				}
				se.UpdateTime = time.Now()
				err = s.db.UserStockEvaluate.Update(ctx, tx, se)
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

	return fin_stock_assistant.FromIndexEvaluate(dbUserIndexEvaluate), nil
}
