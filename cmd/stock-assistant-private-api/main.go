package main

import (
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi"
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/cmd/stock-assistant-private-api/handler"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	log.Init(true)

	middleware.Debug = false

	logger := zap.L().Named("main")

	var bindAddr string
	var storageConnectionString= ""

	cmd := cobra.Command{}
	cmd.PersistentFlags().StringVar(&bindAddr, "bind-addr", ":8082", "api server bind addr")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return errors.WithStack(err)
		}
		api := operations.NewStockAssistantAPI(swaggerSpec)

		h, err := handler.NewStockAssistantHandler(&handler.StockAssistantHandlerOptions{FinStockAssistantConnectionString: storageConnectionString})
		if err != nil {
			return err
		}

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
		api.UserSettingListHandler = operations.UserSettingListHandlerFunc(h.UserSettingsList)
		api.UserSettingGetHandler = operations.UserSettingGetHandlerFunc(h.UserSettingsGet)
		api.UserSettingSaveHandler = operations.UserSettingSaveHandlerFunc(h.UserSettingsSave)
		api.UserSettingDeleteHandler = operations.UserSettingDeleteHandlerFunc(h.UserSettingsDelete)
		api.StockIndexAdviceListHandler = operations.StockIndexAdviceListHandlerFunc(h.StockIndexAdviceList)
		api.StockGetHandler = operations.StockGetHandlerFunc(h.StockGet)

		logger.Info("Start server", zap.String("addr", bindAddr))
		err = http.ListenAndServe(bindAddr,
			restful.Recovery(cors.AllowAll().Handler(api.Serve(nil))))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
	cmd.Execute()
}
