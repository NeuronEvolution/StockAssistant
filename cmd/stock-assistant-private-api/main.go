package main

import (
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi"
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/cmd/stock-assistant-private-api/handler"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/loads"
	"net/http"
	"os"
)

func main() {
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "8082")

	restful.Run(func() (http.Handler, error) {
		h, err := handler.NewStockAssistantHandler()
		if err != nil {
			return nil, err
		}

		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return nil, err
		}

		api := operations.NewStockAssistantPrivateAPI(swaggerSpec)
		api.UserStockIndexListHandler = operations.UserStockIndexListHandlerFunc(h.UserStockIndexList)
		api.UserStockIndexGetHandler = operations.UserStockIndexGetHandlerFunc(h.UserStockIndexGet)
		api.UserStockIndexAddHandler = operations.UserStockIndexAddHandlerFunc(h.UserStockIndexAdd)
		api.UserStockIndexUpdateHandler = operations.UserStockIndexUpdateHandlerFunc(h.UserStockIndexUpdate)
		api.UserStockIndexDeleteHandler = operations.UserStockIndexDeleteHandlerFunc(h.UserStockIndexDelete)
		api.UserStockIndexRenameHandler = operations.UserStockIndexRenameHandlerFunc(h.UserStockIndexRename)
		api.UserStockEvaluateListHandler = operations.UserStockEvaluateListHandlerFunc(h.UserStockEvaluateList)
		api.UserStockEvaluateGetHandler = operations.UserStockEvaluateGetHandlerFunc(h.UserStockEvaluateGet)
		api.UserIndexEvaluateListHandler = operations.UserIndexEvaluateListHandlerFunc(h.UserIndexEvaluateList)
		api.UserIndexEvaluateGetHandler = operations.UserIndexEvaluateGetHandlerFunc(h.UserIndexEvaluateGet)
		api.UserIndexEvaluateSaveHandler = operations.UserIndexEvaluateSaveHandlerFunc(h.UserIndexEvaluateSave)
		api.StockIndexAdviceListHandler = operations.StockIndexAdviceListHandlerFunc(h.StockIndexAdviceList)
		api.StockGetHandler = operations.StockGetHandlerFunc(h.StockGet)

		return api.Serve(nil), nil
	})
}
