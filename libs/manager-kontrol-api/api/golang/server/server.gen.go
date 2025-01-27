// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	. "github.com/kurtosis-tech/kardinal/libs/manager-kontrol-api/api/golang/types"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Cluster resource definition
	// (GET /tenant/{uuid}/cluster-resources)
	GetTenantUuidClusterResources(ctx echo.Context, uuid Uuid) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTenantUuidClusterResources converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenantUuidClusterResources(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid Uuid

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", ctx.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenantUuidClusterResources(ctx, uuid)
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

	router.GET(baseURL+"/tenant/:uuid/cluster-resources", wrapper.GetTenantUuidClusterResources)

}

type NotOkJSONResponse ResponseInfo

type GetTenantUuidClusterResourcesRequestObject struct {
	Uuid Uuid `json:"uuid"`
}

type GetTenantUuidClusterResourcesResponseObject interface {
	VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error
}

type GetTenantUuidClusterResources200JSONResponse ClusterResources

func (response GetTenantUuidClusterResources200JSONResponse) VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTenantUuidClusterResourcesdefaultJSONResponse struct {
	Body       ResponseInfo
	StatusCode int
}

func (response GetTenantUuidClusterResourcesdefaultJSONResponse) VisitGetTenantUuidClusterResourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Cluster resource definition
	// (GET /tenant/{uuid}/cluster-resources)
	GetTenantUuidClusterResources(ctx context.Context, request GetTenantUuidClusterResourcesRequestObject) (GetTenantUuidClusterResourcesResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTenantUuidClusterResources operation middleware
func (sh *strictHandler) GetTenantUuidClusterResources(ctx echo.Context, uuid Uuid) error {
	var request GetTenantUuidClusterResourcesRequestObject

	request.Uuid = uuid

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTenantUuidClusterResources(ctx.Request().Context(), request.(GetTenantUuidClusterResourcesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTenantUuidClusterResources")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTenantUuidClusterResourcesResponseObject); ok {
		return validResponse.VisitGetTenantUuidClusterResourcesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RWTXPbNhD9K5ptj5RgxZcOb5468Wg6tTOK3R4yHg8CraiNSQBdLJSwHv73DgjqI7Ek",
	"+9A2J4Lct4u3D9A+PYFxjXcWrQQon8Br1g0Kcv8WIy3Sc4HBMHkhZ6GEu7vZ5cgtR7LCEWNwkQ1CAZRi",
	"XssKCrC6QShzfgGMf0ViXEApHLGAYFbY6FRYWp9wQZhsBV3XJXDwzgbsCVw7uXlMC+OsoJW01N7XZHQi",
	"oz6HxOhpr+LPjEso4Se160vlaFDzofTMLl3e7LvGLH71aAQXI2R2DAkyJKfav9YxCPJ86DkLxs4jC+U3",
	"HWXlmP7u2T14V5MZIiTY9Iuv48qNh77X008oejq52E97n7JaKHbIMTXecd/8IOyQCEUWvAQKQm5CTpma",
	"0Mq4cso/Vkp7CiqgiUzSqk1WamtgoJl1C70SvnZts7kHB+lq78N6OrncQk+TzPAdx8dfQmKoPakUUutj",
	"TIKQzQpyrE/Kp2u/0ueTy13KPNb4knY563XiWZQvjh/JVmqbeIg12rVrH5ZUb347pxm/TfB3PfpHsK20",
	"4BfdHqV3NcT/B2oBeU3mxCkbx7ieTj5k3GlKGXvwyqXQsSu3Jpao64cXuWwF+iNnvIrUf3GEuw/u02c0",
	"krr4Zrw9G03GLTA9l44bLWk4k5XzN7AtRFawQk6VGgxBV3hgQm/Qrxu0twmbh/rGAT7mArs9iszs/kRD",
	"t8OWaGOTKrydz2/mUMDs+t0NFPDnxfx6dn21V2LfT2hQQ0jqFPtdW10hq9+cFXb16OL9DApYI4dsAtPJ",
	"2eQs7e48Wu0JSjjvP+XT67VUglZbUU/J4Dplsi+Med8YKuzvQDqCfjDNFlDCFcptn3oXafHMTopv7Pfj",
	"YY13ENXba3f/nWe+OTv71xzzGcUDrvkhGoMhLGM92vDIY3ypYy3HdthSVtnj+1EQm0ZzC+XGabd/L0YL",
	"XJKlfscCRFdJH3iu+33Xdd0/AQAA//8hn+431AgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
