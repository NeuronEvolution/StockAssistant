package services

import (
	"context"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"time"
)

func (s *StockAssistantService) UserSettingList(userId string) (settingsList []*models.UserSetting, err error) {
	dbList, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).QueryList(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSettingList(dbList), nil
}

func (s *StockAssistantService) UserSettingGet(userId string, configKey string) (setting *models.UserSetting, err error) {
	dbSetting, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).And().ConfigKey_Equal(configKey).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSetting(dbSetting), nil
}

func (s *StockAssistantService) UserSettingSave(userId string, setting *models.UserSetting) (settingSaved *models.UserSetting, err error) {
	tx, err := s.db.BeginReadCommittedTx(context.Background(), false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbSetting, err := s.db.UserSetting.GetQuery().ForUpdate().
		UserId_Equal(userId).And().ConfigKey_Equal(setting.ConfigKey).QueryOne(context.Background(), tx)
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
	tx, err := s.db.BeginReadCommittedTx(context.Background(),false)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbSetting, err := s.db.UserSetting.GetQuery().ForUpdate().
		UserId_Equal(userId).And().ConfigKey_Equal(configKey).QueryOne(context.Background(), tx)
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
