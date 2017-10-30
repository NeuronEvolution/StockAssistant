package handler

import (
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/services"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
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
		return restful.Responder(err)
	}

	return operations.NewUserIndexListOK().WithPayload(fromIndexList(indexList))
}

func (h *StockAssistantHandler) UserIndexGet(p operations.UserIndexGetParams) middleware.Responder {
	index, err := h.service.UserIndexGet(p.UserID, p.IndexID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexGetOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserIndexSave(p operations.UserIndexSaveParams) middleware.Responder {
	index, err := h.service.UserIndexSave(p.UserID, toIndex(p.Index))
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexSaveOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserIndexDelete(p operations.UserIndexDeleteParams) middleware.Responder {
	err := h.service.UserIndexDelete(p.UserID, p.IndexID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexDeleteOK()
}

func (h *StockAssistantHandler) UserIndexRename(p operations.UserIndexRenameParams) middleware.Responder {
	indexRenamed, err := h.service.UserIndexRename(p.UserID, p.OldName, p.NewName)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexRenameOK().WithPayload(fromIndex(indexRenamed))
}

func (h *StockAssistantHandler) UserStockEvaluateList(p operations.UserStockEvaluateListParams) middleware.Responder {
	query := &models.UserStockEvaluateListQuery{}
	query.UserId = p.UserID
	if p.NotEvaluated == nil {
		query.Evaluated = true
	} else {
		query.Evaluated = false
	}
	if p.Sort != nil {
		query.Sort = *p.Sort
	}
	if p.PageToken != nil {
		query.PageToken = *p.PageToken
	}
	if p.PageSize != nil {
		query.PageSize = *p.PageSize
	}

	list, err := h.service.UserStockEvaluateList(query)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockEvaluateListOK().WithPayload(fromStockEvaluateList(list))
}

func (h *StockAssistantHandler) UserStockEvaluateGet(p operations.UserStockEvaluateGetParams) middleware.Responder {
	se, err := h.service.UserStockEvaluateGet(p.UserID, p.StockID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockEvaluateGetOK().WithPayload(fromStockEvaluate(se))
}

func (h *StockAssistantHandler) UserIndexEvaluateList(p operations.UserIndexEvaluateListParams) middleware.Responder {
	list, err := h.service.UserIndexEvaluateList(p.UserID, p.StockID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexEvaluateListOK().WithPayload(fromIndexEvaluateList(list))
}

func (h *StockAssistantHandler) UserIndexEvaluateGet(p operations.UserIndexEvaluateGetParams) middleware.Responder {
	ie, err := h.service.UserIndexEvaluateGet(p.UserID, p.StockID, p.IndexName)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexEvaluateGetOK().WithPayload(fromIndexEvaluate(ie))
}

func (h *StockAssistantHandler) UserIndexEvaluateSave(p operations.UserIndexEvaluateSaveParams) middleware.Responder {
	ie, err := h.service.UserIndexEvaluateSave(p.UserID, p.StockID, toIndexEvaluate(p.IndexEvaluate))
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserIndexEvaluateSaveOK().WithPayload(fromIndexEvaluate(ie))
}

func (h *StockAssistantHandler) UserSettingsList(p operations.UserSettingListParams) middleware.Responder {
	list, err := h.service.UserSettingList(p.UserID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserSettingListOK().WithPayload(fromUserSettingList(list))
}

func (h *StockAssistantHandler) UserSettingsGet(p operations.UserSettingGetParams) middleware.Responder {
	setting, err := h.service.UserSettingGet(p.UserID, p.ConfigKey)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserSettingGetOK().WithPayload(fromUserSetting(setting))
}

func (h *StockAssistantHandler) UserSettingsSave(p operations.UserSettingSaveParams) middleware.Responder {
	setting, err := h.service.UserSettingSave(p.UserID, toUserSetting(p.Setting))
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserSettingSaveOK().WithPayload(fromUserSetting(setting))
}

func (h *StockAssistantHandler) UserSettingsDelete(p operations.UserSettingDeleteParams) middleware.Responder {
	err := h.service.UserSettingDelete(p.UserID, p.ConfigKey)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserSettingDeleteOK()
}
