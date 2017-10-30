package services

import (
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronEvolution/StockAssistant/models"
	"context"
	"database/sql"
	"time"
	"fmt"
)

func (s *StockAssistantService) UserSettingList(userId string) (settingsList []*models.UserSetting, err error) {
	dbList, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).SelectList(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSettingList(dbList), nil
}

func (s *StockAssistantService) UserSettingGet(userId string, configKey string) (setting *models.UserSetting, err error) {
	dbSetting, err := s.db.UserSetting.GetQuery().
		UserId_Equal(userId).
		ConfigKey_Equal(configKey).Select(context.Background())
	if err != nil {
		return nil, err
	}

	return fin_stock_assistant.FromSetting(dbSetting), nil
}

func (s *StockAssistantService) UserSettingSave(userId string, setting *models.UserSetting) (settingSaved *models.UserSetting, err error) {
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
