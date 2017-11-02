// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
StockIndexAdviceList lists
*/
func (a *Client) StockIndexAdviceList(params *StockIndexAdviceListParams) (*StockIndexAdviceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStockIndexAdviceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StockIndexAdviceList",
		Method:             "GET",
		PathPattern:        "/stockIndexAdvices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StockIndexAdviceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StockIndexAdviceListOK), nil

}

/*
UserIndexEvaluateGet user index evaluate get API
*/
func (a *Client) UserIndexEvaluateGet(params *UserIndexEvaluateGetParams) (*UserIndexEvaluateGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexEvaluateGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexEvaluateGet",
		Method:             "GET",
		PathPattern:        "/{userId}/stockEvaluates/{stockId}/indexEvaluates/{indexName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexEvaluateGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexEvaluateGetOK), nil

}

/*
UserIndexEvaluateList user index evaluate list API
*/
func (a *Client) UserIndexEvaluateList(params *UserIndexEvaluateListParams) (*UserIndexEvaluateListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexEvaluateListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexEvaluateList",
		Method:             "GET",
		PathPattern:        "/{userId}/stockEvaluates/{stockId}/indexEvaluates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexEvaluateListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexEvaluateListOK), nil

}

/*
UserIndexEvaluateSave user index evaluate save API
*/
func (a *Client) UserIndexEvaluateSave(params *UserIndexEvaluateSaveParams) (*UserIndexEvaluateSaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexEvaluateSaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexEvaluateSave",
		Method:             "POST",
		PathPattern:        "/{userId}/stockEvaluates/{stockId}/indexEvaluates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexEvaluateSaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexEvaluateSaveOK), nil

}

/*
UserSettingDelete user setting delete API
*/
func (a *Client) UserSettingDelete(params *UserSettingDeleteParams) (*UserSettingDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserSettingDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserSettingDelete",
		Method:             "DELETE",
		PathPattern:        "/{userId}/settings/{configKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserSettingDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserSettingDeleteOK), nil

}

/*
UserSettingGet user setting get API
*/
func (a *Client) UserSettingGet(params *UserSettingGetParams) (*UserSettingGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserSettingGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserSettingGet",
		Method:             "GET",
		PathPattern:        "/{userId}/settings/{configKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserSettingGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserSettingGetOK), nil

}

/*
UserSettingList lists
*/
func (a *Client) UserSettingList(params *UserSettingListParams) (*UserSettingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserSettingListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserSettingList",
		Method:             "GET",
		PathPattern:        "/{userId}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserSettingListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserSettingListOK), nil

}

/*
UserSettingSave saves
*/
func (a *Client) UserSettingSave(params *UserSettingSaveParams) (*UserSettingSaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserSettingSaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserSettingSave",
		Method:             "POST",
		PathPattern:        "/{userId}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserSettingSaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserSettingSaveOK), nil

}

/*
UserStockEvaluateGet user stock evaluate get API
*/
func (a *Client) UserStockEvaluateGet(params *UserStockEvaluateGetParams) (*UserStockEvaluateGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockEvaluateGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockEvaluateGet",
		Method:             "GET",
		PathPattern:        "/{userId}/stockEvaluates/{stockId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockEvaluateGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockEvaluateGetOK), nil

}

/*
UserStockEvaluateList user stock evaluate list API
*/
func (a *Client) UserStockEvaluateList(params *UserStockEvaluateListParams) (*UserStockEvaluateListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockEvaluateListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockEvaluateList",
		Method:             "GET",
		PathPattern:        "/{userId}/stockEvaluates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockEvaluateListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockEvaluateListOK), nil

}

/*
UserStockIndexAdd adds
*/
func (a *Client) UserStockIndexAdd(params *UserStockIndexAddParams) (*UserStockIndexAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexAdd",
		Method:             "POST",
		PathPattern:        "/{userId}/stockIndices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexAddReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexAddOK), nil

}

/*
UserStockIndexDelete user stock index delete API
*/
func (a *Client) UserStockIndexDelete(params *UserStockIndexDeleteParams) (*UserStockIndexDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexDelete",
		Method:             "DELETE",
		PathPattern:        "/{userId}/stockIndices/{indexId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexDeleteOK), nil

}

/*
UserStockIndexGet gets user index
*/
func (a *Client) UserStockIndexGet(params *UserStockIndexGetParams) (*UserStockIndexGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexGet",
		Method:             "GET",
		PathPattern:        "/{userId}/stockIndices/{indexId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexGetOK), nil

}

/*
UserStockIndexList gets user indices
*/
func (a *Client) UserStockIndexList(params *UserStockIndexListParams) (*UserStockIndexListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexList",
		Method:             "GET",
		PathPattern:        "/{userId}/stockIndices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexListOK), nil

}

/*
UserStockIndexRename user stock index rename API
*/
func (a *Client) UserStockIndexRename(params *UserStockIndexRenameParams) (*UserStockIndexRenameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexRenameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexRename",
		Method:             "POST",
		PathPattern:        "/{userId}/stockIndices/rename",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexRenameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexRenameOK), nil

}

/*
UserStockIndexUpdate updates
*/
func (a *Client) UserStockIndexUpdate(params *UserStockIndexUpdateParams) (*UserStockIndexUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStockIndexUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserStockIndexUpdate",
		Method:             "POST",
		PathPattern:        "/{userId}/stockIndices/{indexId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserStockIndexUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserStockIndexUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
