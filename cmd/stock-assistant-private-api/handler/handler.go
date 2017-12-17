package handler

import (
	"context"
	api "github.com/NeuronEvolution/StockAssistant/api/private/gen/models"
	"github.com/NeuronEvolution/StockAssistant/api/private/gen/restapi/operations"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/StockAssistant/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type StockAssistantHandler struct {
	logger  *zap.Logger
	service *services.StockAssistantService
}

func NewStockAssistantHandler() (h *StockAssistantHandler, err error) {
	h = &StockAssistantHandler{}
	h.logger = log.TypedLogger(h)
	s, err := services.NewStockAssistantService()
	if err != nil {
		return nil, err
	}
	h.service = s

	return h, nil
}

func (h *StockAssistantHandler) UserStockIndexList(p operations.UserStockIndexListParams) middleware.Responder {
	indexList, err := h.service.UserStockIndexList(context.Background(), p.UserID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockIndexListOK().WithPayload(fromIndexList(indexList))
}

func (h *StockAssistantHandler) UserStockIndexGet(p operations.UserStockIndexGetParams) middleware.Responder {
	index, err := h.service.UserStockIndexGet(context.Background(), p.UserID, p.IndexName)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockIndexGetOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexAdd(p operations.UserStockIndexAddParams) middleware.Responder {
	index, err := h.service.UserStockIndexAdd(context.Background(), p.UserID, toIndex(p.Index))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockIndexAddOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexUpdate(p operations.UserStockIndexUpdateParams) middleware.Responder {
	index, err := h.service.UserStockIndexUpdate(context.Background(), p.UserID, toIndex(p.Index))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockIndexUpdateOK().WithPayload(fromIndex(index))
}

func (h *StockAssistantHandler) UserStockIndexDelete(p operations.UserStockIndexDeleteParams) middleware.Responder {
	err := h.service.UserStockIndexDelete(context.Background(), p.UserID, p.IndexName)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockIndexDeleteOK()
}

func (h *StockAssistantHandler) UserStockIndexRename(p operations.UserStockIndexRenameParams) middleware.Responder {
	indexRenamed, err := h.service.UserStockIndexRename(context.Background(), p.UserID, p.NameOld, p.NameNew)
	if err != nil {
		return errors.Wrap(err)
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

	result, nextPageToken, err := h.service.UserStockEvaluateList(context.Background(), query)
	if err != nil {
		return errors.Wrap(err)
	}

	response := &api.UserStockEvaluateListResponse{
		Items:         fromStockEvaluateList(result),
		NextPageToken: nextPageToken,
	}

	return operations.NewUserStockEvaluateListOK().
		WithPayload(response)
}

func (h *StockAssistantHandler) UserStockEvaluateGet(p operations.UserStockEvaluateGetParams) middleware.Responder {
	se, err := h.service.UserStockEvaluateGet(context.Background(), p.UserID, p.StockID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserStockEvaluateGetOK().WithPayload(fromStockEvaluate(se))
}

func (h *StockAssistantHandler) UserIndexEvaluateList(p operations.UserIndexEvaluateListParams) middleware.Responder {
	list, err := h.service.UserIndexEvaluateList(context.Background(), p.UserID, p.StockID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserIndexEvaluateListOK().WithPayload(fromIndexEvaluateList(list))
}

func (h *StockAssistantHandler) UserIndexEvaluateGet(p operations.UserIndexEvaluateGetParams) middleware.Responder {
	ie, err := h.service.UserIndexEvaluateGet(context.Background(), p.UserID, p.StockID, p.IndexName)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserIndexEvaluateGetOK().WithPayload(fromIndexEvaluate(ie))
}

func (h *StockAssistantHandler) UserIndexEvaluateSave(p operations.UserIndexEvaluateSaveParams) middleware.Responder {
	ie, err := h.service.UserIndexEvaluateSave(context.Background(), p.UserID, p.StockID, toIndexEvaluate(p.IndexEvaluate))
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewUserIndexEvaluateSaveOK().WithPayload(fromIndexEvaluate(ie))
}

func (h *StockAssistantHandler) StockIndexAdviceList(p operations.StockIndexAdviceListParams) middleware.Responder {
	query := &models.StockIndexAdviceQuery{}

	if p.UserID != nil {
		query.UserId = *p.UserID
	}

	if p.PageToken != nil {
		query.PageToken = *p.PageToken
	}

	if p.PageSize != nil {
		query.PageSize = int64(*p.PageSize)
	}

	result, nextPageToken, err := h.service.StockIndexAdviceList(context.Background(), query)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewStockIndexAdviceListOK().
		WithNeuronXNextPageToken(nextPageToken).
		WithPayload(fromStockIndexAdviceList(result))
}

func (h *StockAssistantHandler) StockGet(p operations.StockGetParams) middleware.Responder {
	stock, err := h.service.StockGet(context.Background(), p.StockID)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewStockGetOK().WithPayload(fromStock(stock))
}
