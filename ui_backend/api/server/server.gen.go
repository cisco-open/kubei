// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	. "github.com/openclarity/vmclarity/ui_backend/api/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of findings impact for the dashboard.
	// (GET /dashboard/findingsImpact)
	GetDashboardFindingsImpact(ctx echo.Context) error
	// Get a list of finding trends for all finding types.
	// (GET /dashboard/findingsTrends)
	GetDashboardFindingsTrends(ctx echo.Context, params GetDashboardFindingsTrendsParams) error
	// Get a list of riskiest assets for the dashboard.
	// (GET /dashboard/riskiestAssets)
	GetDashboardRiskiestAssets(ctx echo.Context) error
	// Get a list of riskiest regions for the dashboard.
	// (GET /dashboard/riskiestRegions)
	GetDashboardRiskiestRegions(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDashboardFindingsImpact converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardFindingsImpact(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDashboardFindingsImpact(ctx)
	return err
}

// GetDashboardFindingsTrends converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardFindingsTrends(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDashboardFindingsTrendsParams
	// ------------- Required query parameter "startTime" -------------

	err = runtime.BindQueryParameter("form", true, true, "startTime", ctx.QueryParams(), &params.StartTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter startTime: %s", err))
	}

	// ------------- Required query parameter "endTime" -------------

	err = runtime.BindQueryParameter("form", true, true, "endTime", ctx.QueryParams(), &params.EndTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter endTime: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDashboardFindingsTrends(ctx, params)
	return err
}

// GetDashboardRiskiestAssets converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardRiskiestAssets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDashboardRiskiestAssets(ctx)
	return err
}

// GetDashboardRiskiestRegions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardRiskiestRegions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDashboardRiskiestRegions(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/dashboard/findingsImpact", wrapper.GetDashboardFindingsImpact)
	router.GET(baseURL+"/dashboard/findingsTrends", wrapper.GetDashboardFindingsTrends)
	router.GET(baseURL+"/dashboard/riskiestAssets", wrapper.GetDashboardRiskiestAssets)
	router.GET(baseURL+"/dashboard/riskiestRegions", wrapper.GetDashboardRiskiestRegions)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
