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
UserIndexDelete user index delete API
*/
func (a *Client) UserIndexDelete(params *UserIndexDeleteParams) (*UserIndexDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexDelete",
		Method:             "DELETE",
		PathPattern:        "/{userId}/indices/{indexId}",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexDeleteOK), nil

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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
UserIndexGet gets user index
*/
func (a *Client) UserIndexGet(params *UserIndexGetParams) (*UserIndexGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexGet",
		Method:             "GET",
		PathPattern:        "/{userId}/indices/{indexId}",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexGetOK), nil

}

/*
UserIndexList gets user indices
*/
func (a *Client) UserIndexList(params *UserIndexListParams) (*UserIndexListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexList",
		Method:             "GET",
		PathPattern:        "/{userId}/indices",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexListOK), nil

}

/*
UserIndexRename user index rename API
*/
func (a *Client) UserIndexRename(params *UserIndexRenameParams) (*UserIndexRenameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexRenameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexRename",
		Method:             "POST",
		PathPattern:        "/{userId}/indices/rename",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexRenameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexRenameOK), nil

}

/*
UserIndexSave saves
*/
func (a *Client) UserIndexSave(params *UserIndexSaveParams) (*UserIndexSaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIndexSaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserIndexSave",
		Method:             "POST",
		PathPattern:        "/{userId}/indices",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserIndexSaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserIndexSaveOK), nil

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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
