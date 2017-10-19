// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// UserIndexEvaluateListURL generates an URL for the user index evaluate list operation
type UserIndexEvaluateListURL struct {
	StockID string
	UserID  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UserIndexEvaluateListURL) WithBasePath(bp string) *UserIndexEvaluateListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UserIndexEvaluateListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UserIndexEvaluateListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/{userId}/stockEvaluates/{stockId}/indexEvaluates"

	stockID := o.StockID
	if stockID != "" {
		_path = strings.Replace(_path, "{stockId}", stockID, -1)
	} else {
		return nil, errors.New("StockID is required on UserIndexEvaluateListURL")
	}
	userID := o.UserID
	if userID != "" {
		_path = strings.Replace(_path, "{userId}", userID, -1)
	} else {
		return nil, errors.New("UserID is required on UserIndexEvaluateListURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/stock-assistant/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UserIndexEvaluateListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UserIndexEvaluateListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UserIndexEvaluateListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UserIndexEvaluateListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UserIndexEvaluateListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UserIndexEvaluateListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
