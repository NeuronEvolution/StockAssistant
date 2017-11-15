package services

import (
	"context"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/storages/fin-stock-assistant"
)

func (s *StockAssistantService) StockGet(stockId string) (stock *models.Stock, err error) {
	dbStock, err := s.db.Stock.GetQuery().StockId_Equal(stockId).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	stock = fin_stock_assistant.FromStock(dbStock)

	stock.StockUrlList = make([]*models.StockUrl, 0)
	if stock.WebsiteURL != "" {
		stock.StockUrlList = append(stock.StockUrlList, &models.StockUrl{Name: "官网", Icon: "", Url: stock.WebsiteURL})
	}

	stock.StockUrlList = append(stock.StockUrlList, &models.StockUrl{
		Name: "交易所公告",
		Icon: "",
		Url:  "http://disclosure.szse.cn/m/szmb/drgg_search.htm?secode=" + stock.StockCode,
	})

	stock.StockUrlList = append(stock.StockUrlList, &models.StockUrl{
		Name: "同花顺", Icon: "",
		Url: "http://stockpage.10jqka.com.cn/" + stock.StockCode + "/"})

	return stock, nil
}
