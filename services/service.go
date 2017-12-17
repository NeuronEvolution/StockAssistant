package services

import (
	"github.com/NeuronEvolution/StockAssistant/remotes/stock/gen/client"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/log"
	"go.uber.org/zap"
)

type StockAssistantService struct {
	logger      *zap.Logger
	db          *fin_stock_assistant.DB
	stockClient *client.Stock
}

func NewStockAssistantService() (s *StockAssistantService, err error) {
	s = &StockAssistantService{}
	s.logger = log.TypedLogger(s)

	s.db, err = fin_stock_assistant.NewDB("root:123456@tcp(127.0.0.1:3307)/fin-stock-assistant?parseTime=true")
	if err != nil {
		return nil, err
	}

	s.stockClient = client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("127.0.0.1:8081"))

	return s, nil
}
