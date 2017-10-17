package storages

import (
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronEvolution/log"
	"go.uber.org/zap"
)

type StorageOptions struct {
	FinStockAssistantConnectionString string
}

type Storage struct {
	logger         *zap.Logger
	StockAssistant *fin_stock_assistant.DB
}

func NewStorage(options *StorageOptions) (s *Storage, err error) {
	s = &Storage{}
	s.logger = log.TypedLogger(s)
	fin_stock_assistant_db, err := fin_stock_assistant.NewDB(options.FinStockAssistantConnectionString)
	if err != nil {
		return nil, err
	}
	s.StockAssistant = fin_stock_assistant_db

	return s, nil
}
