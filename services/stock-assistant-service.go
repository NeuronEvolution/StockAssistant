package services

import (
	"context"
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

	result = make([]*models.StockEvaluate, 0)

	for _, u := range dbStockEvaluateList {
		se := &models.StockEvaluate{}
		se.StockId = u.StockId
		se.TotalScore = u.TotalScore
		se.EvalRemark = u.EvalRemark

		dbIndexEvaluateList, err := s.storage.StockAssistant.IndexEvaluate.
			SelectListByUserIdAndStockId(context.Background(), nil, userId, u.StockId)
		if err != nil {
			return nil, err
		}

		se.IndexEvaluates = make([]*models.IndexEvaluate, 0)
		for _, v := range dbIndexEvaluateList {
			ie := &models.IndexEvaluate{}
			ie.IndexId = v.IndexId
			ie.EvalStars = v.EvalStars
			ie.EvalRemark = v.EvalRemark
			ie.UpdateTime = v.UpdateTime

			se.IndexEvaluates = append(se.IndexEvaluates, ie)
		}

		result = append(result, se)
	}

	return result, nil
}

func (s *StockAssistantService) UserStockEvaluateGet(userId string, stockId string) (eval *models.StockEvaluate, err error) {
	return nil, nil
}

func (s *StockAssistantService) UserStockEvaluateAddOrUpdate(userId string, eval *models.StockEvaluate) (evalUpdated *models.StockEvaluate, err error) {
	return nil, nil
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
