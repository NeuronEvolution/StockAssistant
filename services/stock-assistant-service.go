package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronEvolution/log"
	"go.uber.org/zap"
	"time"
)

type StockAssistantServiceOptions struct {
	FinStockAssistantConnectionString string
}

type StockAssistantService struct {
	logger  *zap.Logger
	options *StockAssistantServiceOptions
	db      *fin_stock_assistant.DB
}

func NewStockAssistantService(options *StockAssistantServiceOptions) (s *StockAssistantService, err error) {
	s = &StockAssistantService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.db, err = fin_stock_assistant.NewDB(options.FinStockAssistantConnectionString)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StockAssistantService) UserIndexList(userId string) (indexList []*models.StockIndex, err error) {
	dbIndexList, err := s.db.StockIndex.GetQuery().SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndexList(dbIndexList), nil
}

func (s *StockAssistantService) UserIndexGet(userId string, indexName string) (index *models.StockIndex, err error) {
	dbIndex, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).IndexName_Equal(indexName).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserIndexSave(userId string, index *models.StockIndex) (indexSaved *models.StockIndex, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).
		IndexName_Equal(index.IndexName).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndex == nil {
		dbIndex = &fin_stock_assistant.StockIndex{}
		dbIndex.UserId = userId
		dbIndex.IndexName = index.IndexName
		dbIndex.IndexDesc = index.IndexDesc
		dbIndex.EvalWeight = index.EvalWeight
		dbIndex.NiWeight = index.NIWeight
		dbIndex.AiWeight = index.AIWeight
		dbIndex.UpdateTime = time.Now()
		_, err = s.db.StockIndex.Insert(context.Background(), tx, dbIndex)
		if err != nil {
			return nil, err
		}
	} else {
		dbIndex.IndexDesc = index.IndexDesc
		dbIndex.EvalWeight = index.EvalWeight
		dbIndex.NiWeight = index.NIWeight
		dbIndex.AiWeight = index.AIWeight
		dbIndex.UpdateTime = time.Now()
		err = s.db.StockIndex.Update(context.Background(), tx, dbIndex)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockIndex(dbIndex), nil
}

func (s *StockAssistantService) UserIndexDelete(userId string, indexName string) (err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).
		IndexName_Equal(indexName).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil
	}

	if dbIndex == nil {
		return fmt.Errorf("not exist")
	}

	err = s.db.StockIndex.Delete(context.Background(), tx, dbIndex.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *StockAssistantService) UserIndexRename(userId string, indexNameOld string, indexNameNew string) (indexNew *models.StockIndex, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if indexNameNew == indexNameOld {
		return nil, fmt.Errorf("name same")
	}

	dbIndexOld, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).
		IndexName_Equal(indexNameOld).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexOld == nil {
		return nil, fmt.Errorf("index old not exist")
	}

	dbIndexNew, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).
		IndexName_Equal(indexNameNew).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexNew != nil {
		return nil, fmt.Errorf("index new exist")
	}

	dbIndexOld.IndexName = indexNameNew
	dbIndexOld.UpdateTime = time.Now()
	err = s.db.StockIndex.Update(context.Background(), tx, dbIndexOld)
	if err != nil {
		return nil, err
	}

	//update index evaluates
	dbIndexEvaluateList, err := s.db.IndexEvaluate.GetQuery().
		UserId_Equal(userId).
		SelectListForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}
	if dbIndexEvaluateList != nil {
		for _, v := range dbIndexEvaluateList {
			v.IndexName = indexNameNew
			err := s.db.IndexEvaluate.Update(context.Background(), tx, v)
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

func (s *StockAssistantService) UserStockEvaluateList(userId string) (result []*models.StockEvaluate, err error) {
	dbStockEvaluateList, err := s.db.StockEvaluate.GetQuery().
		UserId_Equal(userId).
		Limit(0, 10).
		Sort(fin_stock_assistant.STOCK_EVALUATE_FIELD_TOTAL_SCORE, false).
		SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockEvaluateList(dbStockEvaluateList), nil
}

func (s *StockAssistantService) UserStockEvaluateGet(userId string, stockId string) (eval *models.StockEvaluate, err error) {
	dbStockEvaluate, err := s.db.StockEvaluate.GetQuery().
		UserId_Equal(userId).
		StockId_Equal(stockId).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromStockEvaluate(dbStockEvaluate), nil
}

func (s *StockAssistantService) UserStockEvaluateSave(userId string, eval *models.StockEvaluate) (evalSaved *models.StockEvaluate, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	//update stock evaluates
	se, err := s.db.StockEvaluate.GetQuery().
		UserId_Equal(userId).
		StockId_Equal(eval.StockId).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if se == nil {
		se = &fin_stock_assistant.StockEvaluate{}
		se.UserId = userId
		se.StockId = eval.StockId
		se.TotalScore = 0

		_, err := s.db.StockEvaluate.Insert(context.Background(), tx, se)
		if err != nil {
			return nil, err
		}
	} else {
		se.EvalRemark = eval.EvalRemark
		err = s.db.StockEvaluate.Update(context.Background(), tx, se)
		if err != nil {
			return nil, err
		}
	}

	evalSaved.StockId = se.StockId
	evalSaved.EvalRemark = se.EvalRemark
	evalSaved.TotalScore = se.TotalScore

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return evalSaved, nil
}

func (s *StockAssistantService) UserIndexEvaluateList(userId string, stockId string) (result []*models.IndexEvaluate, err error) {
	dbList, err := s.db.IndexEvaluate.GetQuery().
		UserId_Equal(userId).
		StockId_Equal(stockId).SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluateList(dbList), nil
}

func (s *StockAssistantService) UserIndexEvaluateGet(userId string, stockId string, indexName string) (indexEvaluate *models.IndexEvaluate, err error) {
	dbIndexEvaluate, err := s.db.IndexEvaluate.GetQuery().
		UserId_Equal(userId).
		StockId_Equal(stockId).
		IndexName_Equal(indexName).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromIndexEvaluate(dbIndexEvaluate), nil
}

func (s *StockAssistantService) UserIndexEvaluateSave(userId string, stockId string, indexEvaluate *models.IndexEvaluate) (indexEvaluateSaved *models.IndexEvaluate, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := s.db.StockIndex.GetQuery().
		UserId_Equal(userId).
		IndexName_Equal(indexEvaluate.IndexName).SelectForShare(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndex == nil {
		return nil, fmt.Errorf("user have no this index")
	}

	dbIndexEvaluate, err := s.db.IndexEvaluate.GetQuery().
		UserId_Equal(userId).
		StockId_Equal(stockId).
		IndexName_Equal(indexEvaluate.IndexName).SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexEvaluate == nil {
		dbIndexEvaluate = &fin_stock_assistant.IndexEvaluate{}
		dbIndexEvaluate.UserId = userId
		dbIndexEvaluate.StockId = stockId
		dbIndexEvaluate.IndexName = indexEvaluate.IndexName
		dbIndexEvaluate.EvalStars = indexEvaluate.EvalStars
		dbIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
		dbIndexEvaluate.UpdateTime = time.Now()
		_, err = s.db.IndexEvaluate.Insert(context.Background(), tx, dbIndexEvaluate)
		if err != nil {
			return nil, err
		}
	} else {
		dbIndexEvaluate.EvalStars = indexEvaluate.EvalStars
		dbIndexEvaluate.EvalRemark = indexEvaluate.EvalRemark
		dbIndexEvaluate.UpdateTime = time.Now()
		err = s.db.IndexEvaluate.Update(context.Background(), tx, dbIndexEvaluate)
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

func (s *StockAssistantService) UserSettingList(userId string) (settingsList []*models.Setting, err error) {
	dbList, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSettingList(dbList), nil
}

func (s *StockAssistantService) UserSettingGet(userId string, configKey string) (setting *models.Setting, err error) {
	dbSetting, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).
		ConfigKey_Equal(configKey).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSetting(dbSetting), nil
}

func (s *StockAssistantService) UserSettingSave(userId string, setting *models.Setting) (settingSaved *models.Setting, err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbSetting, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).
		ConfigKey_Equal(setting.ConfigKey).SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbSetting == nil {
		dbSetting = &fin_stock_assistant.UserSetting{}
		dbSetting.UserId = userId
		dbSetting.ConfigKey = setting.ConfigKey
		dbSetting.ConfigValue = setting.ConfigValue
		dbSetting.UpdateTime = time.Now()
		_, err := s.db.UserSetting.Insert(context.Background(), tx, dbSetting)
		if err != nil {
			return nil, err
		}
	} else {
		dbSetting.ConfigValue = setting.ConfigValue
		dbSetting.UpdateTime = time.Now()
		err := s.db.UserSetting.Update(context.Background(), tx, dbSetting)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSetting(dbSetting), nil
}

func (s *StockAssistantService) UserSettingDelete(userId string, configKey string) (err error) {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbSetting, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).
		ConfigKey_Equal(configKey).SelectForUpdate(context.Background(), tx)
	if err != nil {
		return err
	}

	if dbSetting == nil {
		return fmt.Errorf("not exist")
	}

	err = s.db.UserSetting.Delete(context.Background(), tx, dbSetting.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
