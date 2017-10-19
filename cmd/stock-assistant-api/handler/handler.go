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

func (h *StockAssistantHandler) UserIndexList(p operations.UserIndexListParams) middleware.Responder {
	indexList, err := h.service.UserIndexList(p.UserID)
	if err != nil {
		return operations.NewUserIndexListBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserIndexListOK().WithPayload(fromIndexList(indexList))
}

func (h *StockAssistantHandler) UserIndexGet(p operations.UserIndexGetParams) middleware.Responder {
	index, err := h.service.UserIndexGet(p.UserID, p.IndexID)
	if err != nil {
		return operations.NewUserIndexGetBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserIndexGetOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserIndexSave(p operations.UserIndexSaveParams) middleware.Responder {
	index, err := h.service.UserIndexSave(p.UserID, toIndex(p.Index))
	if err != nil {
		return operations.NewUserIndexSaveBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserIndexSaveOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserIndexDelete(p operations.UserIndexDeleteParams) middleware.Responder {
	err := h.service.UserIndexDelete(p.UserID, p.IndexID)
	if err != nil {
		return operations.NewUserIndexDeleteBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserIndexDeleteOK()
}

func (h *StockAssistantHandler) UserIndexRename(p operations.UserIndexRenameParams) middleware.Responder {
	indexRenamed, err := h.service.UserIndexRename(p.UserID, p.OldName, p.NewName)
	if err != nil {
		return operations.NewUserIndexRenameBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserIndexRenameOK().WithPayload(fromIndex(indexRenamed))
}

func (h *StockAssistantHandler) UserStockEvaluateList(p operations.UserStockEvaluateListParams) middleware.Responder {
	list, err := h.service.UserStockEvaluateList(p.UserID)
	if err != nil {
		return operations.NewUserStockEvaluateListBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserStockEvaluateListOK().WithPayload(fromStockEvaluateList(list))
}

func (h *StockAssistantHandler) UserStockEvaluateGet(p operations.UserStockEvaluateGetParams) middleware.Responder {
	se, err := h.service.UserStockEvaluateGet(p.UserID, p.StockID)
	if err != nil {
		return operations.NewUserStockEvaluateGetBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserStockEvaluateGetOK().WithPayload(fromStockEvaluate(se))
}

func (h *StockAssistantHandler) UserStockEvaluateSave(p operations.UserStockEvaluateSaveParams) middleware.Responder {
	se, err := h.service.UserStockEvaluateSave(p.UserID, toStockEvaluate(p.StockEvaluate))
	if err != nil {
		return operations.NewUserStockEvaluateSaveBadRequest().WithPayload(err.Error())
	}

	return operations.NewUserStockEvaluateSaveOK().WithPayload(fromStockEvaluate(se))
}

func (h *StockAssistantHandler) UserIndexEvaluateList(p operations.UserIndexEvaluateListParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserIndexEvaluateGet(p operations.UserIndexEvaluateGetParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserIndexEvaluateSave(p operations.UserIndexEvaluateSaveParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserSettingsList(p operations.UserSettingsListParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserSettingsGet(p operations.UserSettingsGetParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserSettingsSave(p operations.UserSettingsSaveParams) middleware.Responder {
	return nil
}

func (h *StockAssistantHandler) UserSettingsDelete(p operations.UserSettingsDeleteParams) middleware.Responder {
	return nil
}
