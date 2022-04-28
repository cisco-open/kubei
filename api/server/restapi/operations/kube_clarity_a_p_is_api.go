// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewKubeClarityAPIsAPI creates a new KubeClarityAPIs instance
func NewKubeClarityAPIsAPI(spec *loads.Document) *KubeClarityAPIsAPI {
	return &KubeClarityAPIsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		DeleteApplicationsIDHandler: DeleteApplicationsIDHandlerFunc(func(params DeleteApplicationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteApplicationsID has not yet been implemented")
		}),
		GetApplicationResourcesHandler: GetApplicationResourcesHandlerFunc(func(params GetApplicationResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetApplicationResources has not yet been implemented")
		}),
		GetApplicationResourcesIDHandler: GetApplicationResourcesIDHandlerFunc(func(params GetApplicationResourcesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetApplicationResourcesID has not yet been implemented")
		}),
		GetApplicationsHandler: GetApplicationsHandlerFunc(func(params GetApplicationsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetApplications has not yet been implemented")
		}),
		GetApplicationsIDHandler: GetApplicationsIDHandlerFunc(func(params GetApplicationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetApplicationsID has not yet been implemented")
		}),
		GetDashboardCountersHandler: GetDashboardCountersHandlerFunc(func(params GetDashboardCountersParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardCounters has not yet been implemented")
		}),
		GetDashboardMostVulnerableHandler: GetDashboardMostVulnerableHandlerFunc(func(params GetDashboardMostVulnerableParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardMostVulnerable has not yet been implemented")
		}),
		GetDashboardPackagesPerLanguageHandler: GetDashboardPackagesPerLanguageHandlerFunc(func(params GetDashboardPackagesPerLanguageParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardPackagesPerLanguage has not yet been implemented")
		}),
		GetDashboardPackagesPerLicenseHandler: GetDashboardPackagesPerLicenseHandlerFunc(func(params GetDashboardPackagesPerLicenseParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardPackagesPerLicense has not yet been implemented")
		}),
		GetDashboardTrendsVulnerabilitiesHandler: GetDashboardTrendsVulnerabilitiesHandlerFunc(func(params GetDashboardTrendsVulnerabilitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardTrendsVulnerabilities has not yet been implemented")
		}),
		GetDashboardVulnerabilitiesWithFixHandler: GetDashboardVulnerabilitiesWithFixHandlerFunc(func(params GetDashboardVulnerabilitiesWithFixParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboardVulnerabilitiesWithFix has not yet been implemented")
		}),
		GetNamespacesHandler: GetNamespacesHandlerFunc(func(params GetNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetNamespaces has not yet been implemented")
		}),
		GetPackagesHandler: GetPackagesHandlerFunc(func(params GetPackagesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPackages has not yet been implemented")
		}),
		GetPackagesIDHandler: GetPackagesIDHandlerFunc(func(params GetPackagesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPackagesID has not yet been implemented")
		}),
		GetPackagesIDApplicationResourcesHandler: GetPackagesIDApplicationResourcesHandlerFunc(func(params GetPackagesIDApplicationResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPackagesIDApplicationResources has not yet been implemented")
		}),
		GetRuntimeQuickscanConfigHandler: GetRuntimeQuickscanConfigHandlerFunc(func(params GetRuntimeQuickscanConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRuntimeQuickscanConfig has not yet been implemented")
		}),
		GetRuntimeScanProgressHandler: GetRuntimeScanProgressHandlerFunc(func(params GetRuntimeScanProgressParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRuntimeScanProgress has not yet been implemented")
		}),
		GetRuntimeScanResultsHandler: GetRuntimeScanResultsHandlerFunc(func(params GetRuntimeScanResultsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRuntimeScanResults has not yet been implemented")
		}),
		GetVulnerabilitiesHandler: GetVulnerabilitiesHandlerFunc(func(params GetVulnerabilitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetVulnerabilities has not yet been implemented")
		}),
		GetVulnerabilitiesVulIDPkgIDHandler: GetVulnerabilitiesVulIDPkgIDHandlerFunc(func(params GetVulnerabilitiesVulIDPkgIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetVulnerabilitiesVulIDPkgID has not yet been implemented")
		}),
		PostApplicationsHandler: PostApplicationsHandlerFunc(func(params PostApplicationsParams) middleware.Responder {
			return middleware.NotImplemented("operation PostApplications has not yet been implemented")
		}),
		PostApplicationsContentAnalysisIDHandler: PostApplicationsContentAnalysisIDHandlerFunc(func(params PostApplicationsContentAnalysisIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PostApplicationsContentAnalysisID has not yet been implemented")
		}),
		PostApplicationsVulnerabilityScanIDHandler: PostApplicationsVulnerabilityScanIDHandlerFunc(func(params PostApplicationsVulnerabilityScanIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PostApplicationsVulnerabilityScanID has not yet been implemented")
		}),
		PutApplicationsIDHandler: PutApplicationsIDHandlerFunc(func(params PutApplicationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PutApplicationsID has not yet been implemented")
		}),
		PutRuntimeQuickscanConfigHandler: PutRuntimeQuickscanConfigHandlerFunc(func(params PutRuntimeQuickscanConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation PutRuntimeQuickscanConfig has not yet been implemented")
		}),
		PutRuntimeScanStartHandler: PutRuntimeScanStartHandlerFunc(func(params PutRuntimeScanStartParams) middleware.Responder {
			return middleware.NotImplemented("operation PutRuntimeScanStart has not yet been implemented")
		}),
		PutRuntimeScanStopHandler: PutRuntimeScanStopHandlerFunc(func(params PutRuntimeScanStopParams) middleware.Responder {
			return middleware.NotImplemented("operation PutRuntimeScanStop has not yet been implemented")
		}),
		PutRuntimeScheduleScanStartHandler: PutRuntimeScheduleScanStartHandlerFunc(func(params PutRuntimeScheduleScanStartParams) middleware.Responder {
			return middleware.NotImplemented("operation PutRuntimeScheduleScanStart has not yet been implemented")
		}),
	}
}

/*KubeClarityAPIsAPI the kube clarity a p is API */
type KubeClarityAPIsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// DeleteApplicationsIDHandler sets the operation handler for the delete applications ID operation
	DeleteApplicationsIDHandler DeleteApplicationsIDHandler
	// GetApplicationResourcesHandler sets the operation handler for the get application resources operation
	GetApplicationResourcesHandler GetApplicationResourcesHandler
	// GetApplicationResourcesIDHandler sets the operation handler for the get application resources ID operation
	GetApplicationResourcesIDHandler GetApplicationResourcesIDHandler
	// GetApplicationsHandler sets the operation handler for the get applications operation
	GetApplicationsHandler GetApplicationsHandler
	// GetApplicationsIDHandler sets the operation handler for the get applications ID operation
	GetApplicationsIDHandler GetApplicationsIDHandler
	// GetDashboardCountersHandler sets the operation handler for the get dashboard counters operation
	GetDashboardCountersHandler GetDashboardCountersHandler
	// GetDashboardMostVulnerableHandler sets the operation handler for the get dashboard most vulnerable operation
	GetDashboardMostVulnerableHandler GetDashboardMostVulnerableHandler
	// GetDashboardPackagesPerLanguageHandler sets the operation handler for the get dashboard packages per language operation
	GetDashboardPackagesPerLanguageHandler GetDashboardPackagesPerLanguageHandler
	// GetDashboardPackagesPerLicenseHandler sets the operation handler for the get dashboard packages per license operation
	GetDashboardPackagesPerLicenseHandler GetDashboardPackagesPerLicenseHandler
	// GetDashboardTrendsVulnerabilitiesHandler sets the operation handler for the get dashboard trends vulnerabilities operation
	GetDashboardTrendsVulnerabilitiesHandler GetDashboardTrendsVulnerabilitiesHandler
	// GetDashboardVulnerabilitiesWithFixHandler sets the operation handler for the get dashboard vulnerabilities with fix operation
	GetDashboardVulnerabilitiesWithFixHandler GetDashboardVulnerabilitiesWithFixHandler
	// GetNamespacesHandler sets the operation handler for the get namespaces operation
	GetNamespacesHandler GetNamespacesHandler
	// GetPackagesHandler sets the operation handler for the get packages operation
	GetPackagesHandler GetPackagesHandler
	// GetPackagesIDHandler sets the operation handler for the get packages ID operation
	GetPackagesIDHandler GetPackagesIDHandler
	// GetPackagesIDApplicationResourcesHandler sets the operation handler for the get packages ID application resources operation
	GetPackagesIDApplicationResourcesHandler GetPackagesIDApplicationResourcesHandler
	// GetRuntimeQuickscanConfigHandler sets the operation handler for the get runtime quickscan config operation
	GetRuntimeQuickscanConfigHandler GetRuntimeQuickscanConfigHandler
	// GetRuntimeScanProgressHandler sets the operation handler for the get runtime scan progress operation
	GetRuntimeScanProgressHandler GetRuntimeScanProgressHandler
	// GetRuntimeScanResultsHandler sets the operation handler for the get runtime scan results operation
	GetRuntimeScanResultsHandler GetRuntimeScanResultsHandler
	// GetVulnerabilitiesHandler sets the operation handler for the get vulnerabilities operation
	GetVulnerabilitiesHandler GetVulnerabilitiesHandler
	// GetVulnerabilitiesVulIDPkgIDHandler sets the operation handler for the get vulnerabilities vul ID pkg ID operation
	GetVulnerabilitiesVulIDPkgIDHandler GetVulnerabilitiesVulIDPkgIDHandler
	// PostApplicationsHandler sets the operation handler for the post applications operation
	PostApplicationsHandler PostApplicationsHandler
	// PostApplicationsContentAnalysisIDHandler sets the operation handler for the post applications content analysis ID operation
	PostApplicationsContentAnalysisIDHandler PostApplicationsContentAnalysisIDHandler
	// PostApplicationsVulnerabilityScanIDHandler sets the operation handler for the post applications vulnerability scan ID operation
	PostApplicationsVulnerabilityScanIDHandler PostApplicationsVulnerabilityScanIDHandler
	// PutApplicationsIDHandler sets the operation handler for the put applications ID operation
	PutApplicationsIDHandler PutApplicationsIDHandler
	// PutRuntimeQuickscanConfigHandler sets the operation handler for the put runtime quickscan config operation
	PutRuntimeQuickscanConfigHandler PutRuntimeQuickscanConfigHandler
	// PutRuntimeScanStartHandler sets the operation handler for the put runtime scan start operation
	PutRuntimeScanStartHandler PutRuntimeScanStartHandler
	// PutRuntimeScanStopHandler sets the operation handler for the put runtime scan stop operation
	PutRuntimeScanStopHandler PutRuntimeScanStopHandler
	// PutRuntimeScheduleScanStartHandler sets the operation handler for the put runtime schedule scan start operation
	PutRuntimeScheduleScanStartHandler PutRuntimeScheduleScanStartHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *KubeClarityAPIsAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *KubeClarityAPIsAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *KubeClarityAPIsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KubeClarityAPIsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KubeClarityAPIsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KubeClarityAPIsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KubeClarityAPIsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KubeClarityAPIsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KubeClarityAPIsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KubeClarityAPIsAPI
func (o *KubeClarityAPIsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DeleteApplicationsIDHandler == nil {
		unregistered = append(unregistered, "DeleteApplicationsIDHandler")
	}
	if o.GetApplicationResourcesHandler == nil {
		unregistered = append(unregistered, "GetApplicationResourcesHandler")
	}
	if o.GetApplicationResourcesIDHandler == nil {
		unregistered = append(unregistered, "GetApplicationResourcesIDHandler")
	}
	if o.GetApplicationsHandler == nil {
		unregistered = append(unregistered, "GetApplicationsHandler")
	}
	if o.GetApplicationsIDHandler == nil {
		unregistered = append(unregistered, "GetApplicationsIDHandler")
	}
	if o.GetDashboardCountersHandler == nil {
		unregistered = append(unregistered, "GetDashboardCountersHandler")
	}
	if o.GetDashboardMostVulnerableHandler == nil {
		unregistered = append(unregistered, "GetDashboardMostVulnerableHandler")
	}
	if o.GetDashboardPackagesPerLanguageHandler == nil {
		unregistered = append(unregistered, "GetDashboardPackagesPerLanguageHandler")
	}
	if o.GetDashboardPackagesPerLicenseHandler == nil {
		unregistered = append(unregistered, "GetDashboardPackagesPerLicenseHandler")
	}
	if o.GetDashboardTrendsVulnerabilitiesHandler == nil {
		unregistered = append(unregistered, "GetDashboardTrendsVulnerabilitiesHandler")
	}
	if o.GetDashboardVulnerabilitiesWithFixHandler == nil {
		unregistered = append(unregistered, "GetDashboardVulnerabilitiesWithFixHandler")
	}
	if o.GetNamespacesHandler == nil {
		unregistered = append(unregistered, "GetNamespacesHandler")
	}
	if o.GetPackagesHandler == nil {
		unregistered = append(unregistered, "GetPackagesHandler")
	}
	if o.GetPackagesIDHandler == nil {
		unregistered = append(unregistered, "GetPackagesIDHandler")
	}
	if o.GetPackagesIDApplicationResourcesHandler == nil {
		unregistered = append(unregistered, "GetPackagesIDApplicationResourcesHandler")
	}
	if o.GetRuntimeQuickscanConfigHandler == nil {
		unregistered = append(unregistered, "GetRuntimeQuickscanConfigHandler")
	}
	if o.GetRuntimeScanProgressHandler == nil {
		unregistered = append(unregistered, "GetRuntimeScanProgressHandler")
	}
	if o.GetRuntimeScanResultsHandler == nil {
		unregistered = append(unregistered, "GetRuntimeScanResultsHandler")
	}
	if o.GetVulnerabilitiesHandler == nil {
		unregistered = append(unregistered, "GetVulnerabilitiesHandler")
	}
	if o.GetVulnerabilitiesVulIDPkgIDHandler == nil {
		unregistered = append(unregistered, "GetVulnerabilitiesVulIDPkgIDHandler")
	}
	if o.PostApplicationsHandler == nil {
		unregistered = append(unregistered, "PostApplicationsHandler")
	}
	if o.PostApplicationsContentAnalysisIDHandler == nil {
		unregistered = append(unregistered, "PostApplicationsContentAnalysisIDHandler")
	}
	if o.PostApplicationsVulnerabilityScanIDHandler == nil {
		unregistered = append(unregistered, "PostApplicationsVulnerabilityScanIDHandler")
	}
	if o.PutApplicationsIDHandler == nil {
		unregistered = append(unregistered, "PutApplicationsIDHandler")
	}
	if o.PutRuntimeQuickscanConfigHandler == nil {
		unregistered = append(unregistered, "PutRuntimeQuickscanConfigHandler")
	}
	if o.PutRuntimeScanStartHandler == nil {
		unregistered = append(unregistered, "PutRuntimeScanStartHandler")
	}
	if o.PutRuntimeScanStopHandler == nil {
		unregistered = append(unregistered, "PutRuntimeScanStopHandler")
	}
	if o.PutRuntimeScheduleScanStartHandler == nil {
		unregistered = append(unregistered, "PutRuntimeScheduleScanStartHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KubeClarityAPIsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KubeClarityAPIsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *KubeClarityAPIsAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *KubeClarityAPIsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *KubeClarityAPIsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KubeClarityAPIsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the kube clarity a p is API
func (o *KubeClarityAPIsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KubeClarityAPIsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/applications/{id}"] = NewDeleteApplicationsID(o.context, o.DeleteApplicationsIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/applicationResources"] = NewGetApplicationResources(o.context, o.GetApplicationResourcesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/applicationResources/{id}"] = NewGetApplicationResourcesID(o.context, o.GetApplicationResourcesIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/applications"] = NewGetApplications(o.context, o.GetApplicationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/applications/{id}"] = NewGetApplicationsID(o.context, o.GetApplicationsIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/counters"] = NewGetDashboardCounters(o.context, o.GetDashboardCountersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/mostVulnerable"] = NewGetDashboardMostVulnerable(o.context, o.GetDashboardMostVulnerableHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/packagesPerLanguage"] = NewGetDashboardPackagesPerLanguage(o.context, o.GetDashboardPackagesPerLanguageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/packagesPerLicense"] = NewGetDashboardPackagesPerLicense(o.context, o.GetDashboardPackagesPerLicenseHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/trends/vulnerabilities"] = NewGetDashboardTrendsVulnerabilities(o.context, o.GetDashboardTrendsVulnerabilitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard/vulnerabilitiesWithFix"] = NewGetDashboardVulnerabilitiesWithFix(o.context, o.GetDashboardVulnerabilitiesWithFixHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/namespaces"] = NewGetNamespaces(o.context, o.GetNamespacesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/packages"] = NewGetPackages(o.context, o.GetPackagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/packages/{id}"] = NewGetPackagesID(o.context, o.GetPackagesIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/packages/{id}/applicationResources"] = NewGetPackagesIDApplicationResources(o.context, o.GetPackagesIDApplicationResourcesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/quickscan/config"] = NewGetRuntimeQuickscanConfig(o.context, o.GetRuntimeQuickscanConfigHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/scan/progress"] = NewGetRuntimeScanProgress(o.context, o.GetRuntimeScanProgressHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/scan/results"] = NewGetRuntimeScanResults(o.context, o.GetRuntimeScanResultsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/vulnerabilities"] = NewGetVulnerabilities(o.context, o.GetVulnerabilitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/vulnerabilities/{vul_id}/{pkg_id}"] = NewGetVulnerabilitiesVulIDPkgID(o.context, o.GetVulnerabilitiesVulIDPkgIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/applications"] = NewPostApplications(o.context, o.PostApplicationsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/applications/contentAnalysis/{id}"] = NewPostApplicationsContentAnalysisID(o.context, o.PostApplicationsContentAnalysisIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/applications/vulnerabilityScan/{id}"] = NewPostApplicationsVulnerabilityScanID(o.context, o.PostApplicationsVulnerabilityScanIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/applications/{id}"] = NewPutApplicationsID(o.context, o.PutApplicationsIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/runtime/quickscan/config"] = NewPutRuntimeQuickscanConfig(o.context, o.PutRuntimeQuickscanConfigHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/runtime/scan/start"] = NewPutRuntimeScanStart(o.context, o.PutRuntimeScanStartHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/runtime/scan/stop"] = NewPutRuntimeScanStop(o.context, o.PutRuntimeScanStopHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/runtime/scheduleScan/start"] = NewPutRuntimeScheduleScanStart(o.context, o.PutRuntimeScheduleScanStartHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KubeClarityAPIsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *KubeClarityAPIsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *KubeClarityAPIsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *KubeClarityAPIsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *KubeClarityAPIsAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
