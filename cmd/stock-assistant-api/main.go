package main

import (
	"github.com/NeuronEvolution/StockAssistant/api/restapi"
	"github.com/NeuronEvolution/StockAssistant/api/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/cmd/stock-assistant-api/handler"
	"github.com/NeuronEvolution/log"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/i-pkg/httphelper"
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

	var bind_addr string
	var storageConnectionString = ""

	cmd := cobra.Command{}
	cmd.PersistentFlags().StringVar(&bind_addr, "bind-addr", ":8082", "api server bind addr")
	cmd.PersistentFlags().StringVar(&storageConnectionString, "mysql-connection-string",
		"root:123456@tcp(127.0.0.1:3307)/fin-stock-assistant?parseTime=true", "mysql connection string")
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

		api.UserIndexListHandler = operations.UserIndexListHandlerFunc(h.UserIndexList)
		api.UserIndexAddHandler = operations.UserIndexAddHandlerFunc(h.UserIndexAdd)
		api.UserIndexGetHandler = operations.UserIndexGetHandlerFunc(h.UserIndexGet)
		api.UserIndexUpdateHandler = operations.UserIndexUpdateHandlerFunc(h.UserIndexUpdate)
		api.UserIndexDeleteHandler = operations.UserIndexDeleteHandlerFunc(h.UserIndexDelete)
		api.UserStockEvaluateListHandler = operations.UserStockEvaluateListHandlerFunc(h.UserStockEvaluateList)
		api.UserStockEvaluateGetHandler = operations.UserStockEvaluateGetHandlerFunc(h.UserStockEvaluateGet)
		api.UserStockEvaluateAddOrUpdateHandler = operations.UserStockEvaluateAddOrUpdateHandlerFunc(h.UserStockEvaluateAddOrUpdate)

		logger.Info("Start server", zap.String("addr", bind_addr))
		err = http.ListenAndServe(bind_addr,
			httphelper.Recovery(cors.Default().Handler(api.Serve(nil))))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
	cmd.Execute()
}
