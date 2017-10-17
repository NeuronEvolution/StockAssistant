package handler

import (
	"github.com/NeuronEvolution/StockAssistant/api/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/services"
	"github.com/NeuronEvolution/log"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type StockAssistantHandlerOptions struct {
	FinStockAssistantConnectionString string
}

type StockAssistantHandler struct {
	logger  *zap.Logger
	options *StockAssistantHandlerOptions
	service *services.StockAssistantService
}

func NewStockAssistantHandler(options *StockAssistantHandlerOptions) (h *StockAssistantHandler, err error) {
	h = &StockAssistantHandler{}
	h.logger = log.TypedLogger(h)
	h.options = options
	s, err := services.NewStockAssistantService(&services.StockAssistantServiceOptions{
		FinStockAssistantConnectionString: options.FinStockAssistantConnectionString})
	if err != nil {
		return nil, err
	}
	h.service = s

	return h, nil
}

func (h *StockAssistantHandler) UserIndexList(params operations.UserIndexListParams) middleware.Responder {
	indexList, err := h.service.UserIndexList(params.UserID)
	if err != nil {
		return operations.NewUserIndexListInternalServerError().WithPayload(err.Error())
	}

	return operations.NewUserIndexListOK().WithPayload(fromIndexList(indexList))
}

func (h *StockAssistantHandler) UserIndexAdd(params operations.UserIndexAddParams) middleware.Responder {
	return operations.NewUserIndexAddBadRequest().WithPayload("UserIndexAdd")
}

func (h *StockAssistantHandler) UserIndexGet(params operations.UserIndexGetParams) middleware.Responder {
	return operations.NewUserIndexGetBadRequest().WithPayload("UserIndexGet")
}

func (h *StockAssistantHandler) UserIndexUpdate(params operations.UserIndexUpdateParams) middleware.Responder {
	return operations.NewUserIndexUpdateBadRequest().WithPayload("UserIndexUpdate")
}

func (h *StockAssistantHandler) UserIndexDelete(params operations.UserIndexDeleteParams) middleware.Responder {
	return operations.NewUserIndexDeleteBadRequest().WithPayload("UserIndexDelete")
}

func (h *StockAssistantHandler) UserStockEvaluateList(params operations.UserStockEvaluateListParams) middleware.Responder {
	return operations.NewUserStockEvaluateListBadRequest().WithPayload("UserStockEvaluateList")
}

func (h *StockAssistantHandler) UserStockEvaluateGet(params operations.UserStockEvaluateGetParams) middleware.Responder {
	return operations.NewUserStockEvaluateGetBadRequest().WithPayload("UserStockEvaluateGet")
}

func (h *StockAssistantHandler) UserStockEvaluateAddOrUpdate(params operations.UserStockEvaluateAddOrUpdateParams) middleware.Responder {
	return operations.NewUserStockEvaluateAddOrUpdateBadRequest().WithPayload("UserStockEvaluateAddOrUpdate")
}
