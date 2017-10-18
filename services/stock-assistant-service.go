package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/runtime"
	"go.uber.org/zap"
)

type StockAssistantServiceOptions struct {
	FinStockAssistantConnectionString string
}

type StockAssistantService struct {
	logger  *zap.Logger
	options *StockAssistantServiceOptions
	storage *storages.Storage
}

func NewStockAssistantService(options *StockAssistantServiceOptions) (s *StockAssistantService, err error) {
	s = &StockAssistantService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	storage, err := storages.NewStorage(&storages.StorageOptions{FinStockAssistantConnectionString: options.FinStockAssistantConnectionString})
	if err != nil {
		return nil, err
	}
	s.storage = storage

	return s, nil
}

func (s *StockAssistantService) UserIndexList(userId string) (indexList []*models.StockIndex, err error) {
	dbIndexList, err := s.storage.StockAssistant.StockIndex.SelectList(context.Background(), nil, "")
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndexList(dbIndexList), nil
}

func (s *StockAssistantService) UserIndexAdd(userId string, index *models.StockIndex) (indexAdded *models.StockIndex, err error) {
	return s.storage.StockAssistant.StockIndex.InsertStockIndex(userId, index)
}

func (s *StockAssistantService) UserIndexGet(userId string, indexId string) (index *models.StockIndex, err error) {
	dbIndex, err := s.storage.StockAssistant.StockIndex.SelectByIndexId(context.Background(), nil, indexId)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserIndexUpdate(userId string, index *models.StockIndex) (indexUpdated *models.StockIndex, err error) {
	return s.storage.StockAssistant.StockIndex.UpdateStockIndex(userId, index)
}

func (s *StockAssistantService) UserIndexDelete(userId string, indexId string) (err error) {
	return s.storage.StockAssistant.StockIndex.DeleteStockIndex(userId, indexId)
}

func (s *StockAssistantService) UserStockEvaluateList(userId string) (result []*models.StockEvaluate, err error) {
	dbStockEvaluateList, err := s.storage.StockAssistant.StockEvaluate.GetQuery().UserId_Column(runtime.RELATION_EQUAL, userId).
		Limit(0, 10).
		Sort(fin_stock_assistant.STOCK_EVALUATE_FIELD_TOTAL_SCORE, false).
		SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	if dbStockEvaluateList == nil {
		return nil, nil
	}

	result = fin_stock_assistant.FromStockEvaluateList(dbStockEvaluateList)
	for _, u := range result {
		dbIndexEvaluateList, err := s.storage.StockAssistant.IndexEvaluate.
			SelectListByUserIdAndStockId(context.Background(), nil, userId, u.StockId)
		if err != nil {
			return nil, err
		}

		u.IndexEvaluates = fin_stock_assistant.FromIndexEvaluateList(dbIndexEvaluateList)
	}

	return result, nil
}

func (s *StockAssistantService) UserStockEvaluateGet(userId string, stockId string) (eval *models.StockEvaluate, err error) {
	dbStockEvaluate, err := s.storage.StockAssistant.StockEvaluate.SelectByUserIdAndStockId(context.Background(), nil, userId, stockId)
	if err != nil {
		return nil, err
	}

	if dbStockEvaluate == nil {
		return nil, nil
	}

	dbIndexEvaluateList, err := s.storage.StockAssistant.IndexEvaluate.
		SelectListByUserIdAndStockId(context.Background(), nil, userId, stockId)
	if err != nil {
		return nil, err
	}

	eval = fin_stock_assistant.FromStockEvaluate(dbStockEvaluate)
	eval.IndexEvaluates = fin_stock_assistant.FromIndexEvaluateList(dbIndexEvaluateList)

	return eval, nil
}

func (s *StockAssistantService) UserStockEvaluateAddOrUpdate(userId string, eval *models.StockEvaluate) (evalUpdated *models.StockEvaluate, err error) {
	tx, err := s.storage.StockAssistant.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	//valid indices
	dbIndexList, err := s.storage.StockAssistant.StockIndex.GetQuery().
		UserId_Column(runtime.RELATION_EQUAL, userId).
		SelectListForShare(context.Background(), tx)
	if err != nil {
		return nil, err
	}
	if dbIndexList == nil || len(dbIndexList) == 0 {
		return nil, fmt.Errorf("user have no index")
	}
	for _, u := range eval.IndexEvaluates {
		has := false
		for _, v := range dbIndexList {
			if v.IndexId == u.IndexId {
				has = true
				break
			}
		}
		if !has {
			return nil, fmt.Errorf("user have no index:%s", u.IndexId)
		}
	}

	//result
	evalUpdated = &models.StockEvaluate{}
	evalUpdated.IndexEvaluates = make([]*models.IndexEvaluate, 0)

	//update index evaluates
	dbIndexEvaluateListOld, err := s.storage.StockAssistant.IndexEvaluate.GetQuery().
		UserId_Column(runtime.RELATION_EQUAL, userId).
		StockId_Column(runtime.RELATION_EQUAL, eval.StockId).SelectListForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}
	if eval.IndexEvaluates != nil {
		for _, u := range eval.IndexEvaluates {
			var ie *fin_stock_assistant.IndexEvaluate
			if dbIndexEvaluateListOld != nil {
				for _, v := range dbIndexEvaluateListOld {
					if v.IndexId == u.IndexId {
						ie = v
						break
					}
				}
			}
			if ie == nil {
				ie = &fin_stock_assistant.IndexEvaluate{}
				ie.UserId = userId
				ie.StockId = eval.StockId
				ie.IndexId = u.IndexId
				ie.EvalStars = u.EvalStars
				ie.EvalRemark = u.EvalRemark
				_, err := s.storage.StockAssistant.IndexEvaluate.Insert(context.Background(), tx, ie)
				if err != nil {
					return nil, err
				}
			} else {
				ie.EvalStars = u.EvalStars
				ie.EvalRemark = u.EvalRemark
				err = s.storage.StockAssistant.IndexEvaluate.Update(context.Background(), tx, ie)
				if err != nil {
					return nil, err
				}
			}
			evalUpdated.IndexEvaluates = append(evalUpdated.IndexEvaluates, &models.IndexEvaluate{
				IndexId:    ie.IndexId,
				EvalStars:  ie.EvalStars,
				EvalRemark: ie.EvalRemark,
				UpdateTime: ie.UpdateTime,
			})
		}
	}

	//update stock evaluates
	se, err := s.storage.StockAssistant.StockEvaluate.GetQuery().
		UserId_Column(runtime.RELATION_EQUAL, userId).
		StockId_Column(runtime.RELATION_EQUAL, eval.StockId).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if se == nil {
		se = &fin_stock_assistant.StockEvaluate{}
		se.UserId = userId
		se.StockId = eval.StockId
		se.TotalScore = 10000

		_, err := s.storage.StockAssistant.StockEvaluate.Insert(context.Background(), tx, se)
		if err != nil {
			return nil, err
		}
	} else {
		se.EvalRemark = eval.EvalRemark
		se.TotalScore = 10000
		err = s.storage.StockAssistant.StockEvaluate.Update(context.Background(), tx, se)
		if err != nil {
			return nil, err
		}
	}

	evalUpdated.StockId = se.StockId
	evalUpdated.EvalRemark = se.EvalRemark
	evalUpdated.TotalScore = se.TotalScore

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return evalUpdated, nil
}

func (s *StockAssistantService) UserSettingsList(userId string) (settingsList []*models.Setting, err error) {
	return nil, nil
}

func (s *StockAssistantService) UserSettingsAdd(userId string, setting *models.Setting) (settingAdded *models.Setting, err error) {
	return nil, nil
}

func (s *StockAssistantService) UserSettingsGet(userId string, key string) (setting *models.Setting, err error) {
	return nil, nil
}

func (s *StockAssistantService) UserSettingsUpdate(userId string, setting *models.Setting) (settingUpdated *models.Setting, err error) {
	return nil, nil
}

func (s *StockAssistantService) UserSettingsDelete(userId string, key string) (err error) {
	return nil
}
