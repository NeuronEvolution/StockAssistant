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

func (h *StockAssistantHandler) UserStockIndexList(p operations.UserStockIndexListParams) middleware.Responder {
	indexList, err := h.service.UserStockIndexList(p.UserID)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexListOK().WithPayload(fromIndexList(indexList))
}

func (h *StockAssistantHandler) UserStockIndexGet(p operations.UserStockIndexGetParams) middleware.Responder {
	index, err := h.service.UserStockIndexGet(p.UserID, p.IndexName)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexGetOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexAdd(p operations.UserStockIndexAddParams) middleware.Responder {
	index, err := h.service.UserStockIndexAdd(p.UserID, toIndex(p.Index))
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexAddOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexUpdate(p operations.UserStockIndexUpdateParams) middleware.Responder {
	index, err := h.service.UserStockIndexUpdate(p.UserID, toIndex(p.Index))
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexUpdateOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexDelete(p operations.UserStockIndexDeleteParams) middleware.Responder {
	err := h.service.UserStockIndexDelete(p.UserID, p.IndexName)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexDeleteOK()
}

func (h *StockAssistantHandler) UserStockIndexRename(p operations.UserStockIndexRenameParams) middleware.Responder {
	indexRenamed, err := h.service.UserStockIndexRename(p.UserID, p.OldName, p.NewName)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockIndexRenameOK().WithPayload(fromIndex(indexRenamed))
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

	result, nextPageToken, err := h.service.UserStockEvaluateList(query)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewUserStockEvaluateListOK().WithNeuronXNextPageToken(nextPageToken).
		WithPayload(fromStockEvaluateList(result))
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

func (h *StockAssistantHandler) StockIndexAdviceList(p operations.StockIndexAdviceListParams) middleware.Responder {
	query := &models.StockIndexAdviceQuery{}

	if p.PageToken != nil {
		query.PageToken = *p.PageToken
	}

	if p.PageSize != nil {
		query.PageSize = *p.PageSize
	}

	result, nextPageToken, err := h.service.StockIndexAdviceList(query)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewStockIndexAdviceListOK().
		WithNeuronXNextPageToken(nextPageToken).
		WithPayload(fromStockIndexAdviceList(result))
}
