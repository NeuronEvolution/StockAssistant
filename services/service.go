package services

import (
	"github.com/NeuronEvolution/StockAssistant/remotes/stock/gen/client"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
	"github.com/NeuronFramework/log"
	"go.uber.org/zap"
)

type StockAssistantServiceOptions struct {
	FinStockAssistantConnectionString string
}

type StockAssistantService struct {
	logger      *zap.Logger
	options     *StockAssistantServiceOptions
	db          *fin_stock_assistant.DB
	stockClient *client.Stock
}

func NewStockAssistantService(options *StockAssistantServiceOptions) (s *StockAssistantService, err error) {
	s = &StockAssistantService{}
	s.logger = log.TypedLogger(s)
	s.options = options

	s.db, err = fin_stock_assistant.NewDB(options.FinStockAssistantConnectionString)
	if err != nil {
		return nil, err
	}

	s.stockClient = client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("127.0.0.1:8081"))

	return s, nil
}
