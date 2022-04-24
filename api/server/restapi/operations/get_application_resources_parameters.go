// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetApplicationResourcesParams creates a new GetApplicationResourcesParams object
// with the default values initialized.
func NewGetApplicationResourcesParams() GetApplicationResourcesParams {

	var (
		// initialize parameters with default values

		sortDirDefault = string("ASC")
	)

	return GetApplicationResourcesParams{
		SortDir: &sortDirDefault,
	}
}

// GetApplicationResourcesParams contains all the bound params for the get application resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetApplicationResources
type GetApplicationResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*application ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	ApplicationID *string
	/*greater than or equal
	  In: query
	*/
	ApplicationsGte *int64
	/*
	  In: query
	*/
	ApplicationsIsNot []int64
	/*
	  In: query
	*/
	ApplicationsIs []int64
	/*less than or equal
	  In: query
	*/
	ApplicationsLte *int64
	/*
	  In: query
	*/
	CisDockerBenchmarkLevelGte *string
	/*
	  In: query
	*/
	CisDockerBenchmarkLevelLte *string
	/*current runtime scan system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	CurrentRuntimeScan *bool
	/*package ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	PackageID *string
	/*greater than or equal
	  In: query
	*/
	PackagesGte *int64
	/*
	  In: query
	*/
	PackagesIsNot []int64
	/*
	  In: query
	*/
	PackagesIs []int64
	/*less than or equal
	  In: query
	*/
	PackagesLte *int64
	/*Page number of the query
	  Required: true
	  In: query
	*/
	Page int64
	/*Maximum items to return
	  Required: true
	  Maximum: 50
	  Minimum: 1
	  In: query
	*/
	PageSize int64
	/*
	  In: query
	*/
	ReportingSBOMAnalyzersContainElements []string
	/*
	  In: query
	*/
	ReportingSBOMAnalyzersDoesntContainElements []string
	/*
	  In: query
	*/
	ResourceHashContains []string
	/*
	  In: query
	*/
	ResourceHashEnd *string
	/*
	  In: query
	*/
	ResourceHashIsNot []string
	/*
	  In: query
	*/
	ResourceHashIs []string
	/*
	  In: query
	*/
	ResourceHashStart *string
	/*
	  In: query
	*/
	ResourceNameContains []string
	/*
	  In: query
	*/
	ResourceNameEnd *string
	/*
	  In: query
	*/
	ResourceNameIsNot []string
	/*
	  In: query
	*/
	ResourceNameIs []string
	/*
	  In: query
	*/
	ResourceNameStart *string
	/*
	  In: query
	*/
	ResourceTypeIs []string
	/*Sorting direction
	  In: query
	  Default: "ASC"
	*/
	SortDir *string
	/*Sort key
	  Required: true
	  In: query
	*/
	SortKey string
	/*
	  In: query
	*/
	VulnerabilitySeverityGte *string
	/*
	  In: query
	*/
	VulnerabilitySeverityLte *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetApplicationResourcesParams() beforehand.
func (o *GetApplicationResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplicationID, qhkApplicationID, _ := qs.GetOK("applicationID")
	if err := o.bindApplicationID(qApplicationID, qhkApplicationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsGte, qhkApplicationsGte, _ := qs.GetOK("applications[gte]")
	if err := o.bindApplicationsGte(qApplicationsGte, qhkApplicationsGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsIsNot, qhkApplicationsIsNot, _ := qs.GetOK("applications[isNot]")
	if err := o.bindApplicationsIsNot(qApplicationsIsNot, qhkApplicationsIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsIs, qhkApplicationsIs, _ := qs.GetOK("applications[is]")
	if err := o.bindApplicationsIs(qApplicationsIs, qhkApplicationsIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsLte, qhkApplicationsLte, _ := qs.GetOK("applications[lte]")
	if err := o.bindApplicationsLte(qApplicationsLte, qhkApplicationsLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCisDockerBenchmarkLevelGte, qhkCisDockerBenchmarkLevelGte, _ := qs.GetOK("cisDockerBenchmarkLevel[gte]")
	if err := o.bindCisDockerBenchmarkLevelGte(qCisDockerBenchmarkLevelGte, qhkCisDockerBenchmarkLevelGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCisDockerBenchmarkLevelLte, qhkCisDockerBenchmarkLevelLte, _ := qs.GetOK("cisDockerBenchmarkLevel[lte]")
	if err := o.bindCisDockerBenchmarkLevelLte(qCisDockerBenchmarkLevelLte, qhkCisDockerBenchmarkLevelLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCurrentRuntimeScan, qhkCurrentRuntimeScan, _ := qs.GetOK("currentRuntimeScan")
	if err := o.bindCurrentRuntimeScan(qCurrentRuntimeScan, qhkCurrentRuntimeScan, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageID, qhkPackageID, _ := qs.GetOK("packageID")
	if err := o.bindPackageID(qPackageID, qhkPackageID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesGte, qhkPackagesGte, _ := qs.GetOK("packages[gte]")
	if err := o.bindPackagesGte(qPackagesGte, qhkPackagesGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesIsNot, qhkPackagesIsNot, _ := qs.GetOK("packages[isNot]")
	if err := o.bindPackagesIsNot(qPackagesIsNot, qhkPackagesIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesIs, qhkPackagesIs, _ := qs.GetOK("packages[is]")
	if err := o.bindPackagesIs(qPackagesIs, qhkPackagesIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesLte, qhkPackagesLte, _ := qs.GetOK("packages[lte]")
	if err := o.bindPackagesLte(qPackagesLte, qhkPackagesLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qReportingSBOMAnalyzersContainElements, qhkReportingSBOMAnalyzersContainElements, _ := qs.GetOK("reportingSBOMAnalyzers[containElements]")
	if err := o.bindReportingSBOMAnalyzersContainElements(qReportingSBOMAnalyzersContainElements, qhkReportingSBOMAnalyzersContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qReportingSBOMAnalyzersDoesntContainElements, qhkReportingSBOMAnalyzersDoesntContainElements, _ := qs.GetOK("reportingSBOMAnalyzers[doesntContainElements]")
	if err := o.bindReportingSBOMAnalyzersDoesntContainElements(qReportingSBOMAnalyzersDoesntContainElements, qhkReportingSBOMAnalyzersDoesntContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceHashContains, qhkResourceHashContains, _ := qs.GetOK("resourceHash[contains]")
	if err := o.bindResourceHashContains(qResourceHashContains, qhkResourceHashContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceHashEnd, qhkResourceHashEnd, _ := qs.GetOK("resourceHash[end]")
	if err := o.bindResourceHashEnd(qResourceHashEnd, qhkResourceHashEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceHashIsNot, qhkResourceHashIsNot, _ := qs.GetOK("resourceHash[isNot]")
	if err := o.bindResourceHashIsNot(qResourceHashIsNot, qhkResourceHashIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceHashIs, qhkResourceHashIs, _ := qs.GetOK("resourceHash[is]")
	if err := o.bindResourceHashIs(qResourceHashIs, qhkResourceHashIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceHashStart, qhkResourceHashStart, _ := qs.GetOK("resourceHash[start]")
	if err := o.bindResourceHashStart(qResourceHashStart, qhkResourceHashStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceNameContains, qhkResourceNameContains, _ := qs.GetOK("resourceName[contains]")
	if err := o.bindResourceNameContains(qResourceNameContains, qhkResourceNameContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceNameEnd, qhkResourceNameEnd, _ := qs.GetOK("resourceName[end]")
	if err := o.bindResourceNameEnd(qResourceNameEnd, qhkResourceNameEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceNameIsNot, qhkResourceNameIsNot, _ := qs.GetOK("resourceName[isNot]")
	if err := o.bindResourceNameIsNot(qResourceNameIsNot, qhkResourceNameIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceNameIs, qhkResourceNameIs, _ := qs.GetOK("resourceName[is]")
	if err := o.bindResourceNameIs(qResourceNameIs, qhkResourceNameIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceNameStart, qhkResourceNameStart, _ := qs.GetOK("resourceName[start]")
	if err := o.bindResourceNameStart(qResourceNameStart, qhkResourceNameStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceTypeIs, qhkResourceTypeIs, _ := qs.GetOK("resourceType[is]")
	if err := o.bindResourceTypeIs(qResourceTypeIs, qhkResourceTypeIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortDir, qhkSortDir, _ := qs.GetOK("sortDir")
	if err := o.bindSortDir(qSortDir, qhkSortDir, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortKey, qhkSortKey, _ := qs.GetOK("sortKey")
	if err := o.bindSortKey(qSortKey, qhkSortKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qVulnerabilitySeverityGte, qhkVulnerabilitySeverityGte, _ := qs.GetOK("vulnerabilitySeverity[gte]")
	if err := o.bindVulnerabilitySeverityGte(qVulnerabilitySeverityGte, qhkVulnerabilitySeverityGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qVulnerabilitySeverityLte, qhkVulnerabilitySeverityLte, _ := qs.GetOK("vulnerabilitySeverity[lte]")
	if err := o.bindVulnerabilitySeverityLte(qVulnerabilitySeverityLte, qhkVulnerabilitySeverityLte, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplicationID binds and validates parameter ApplicationID from query.
func (o *GetApplicationResourcesParams) bindApplicationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationID = &raw

	return nil
}

// bindApplicationsGte binds and validates parameter ApplicationsGte from query.
func (o *GetApplicationResourcesParams) bindApplicationsGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applications[gte]", "query", "int64", raw)
	}
	o.ApplicationsGte = &value

	return nil
}

// bindApplicationsIsNot binds and validates array parameter ApplicationsIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindApplicationsIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationsIsNot string
	if len(rawData) > 0 {
		qvApplicationsIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationsIsNotIC := swag.SplitByFormat(qvApplicationsIsNot, "")
	if len(applicationsIsNotIC) == 0 {
		return nil
	}

	var applicationsIsNotIR []int64
	for i, applicationsIsNotIV := range applicationsIsNotIC {
		applicationsIsNotI, err := swag.ConvertInt64(applicationsIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applications[isNot]", i), "query", "int64", applicationsIsNotI)
		}

		applicationsIsNotIR = append(applicationsIsNotIR, applicationsIsNotI)
	}

	o.ApplicationsIsNot = applicationsIsNotIR

	return nil
}

// bindApplicationsIs binds and validates array parameter ApplicationsIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindApplicationsIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationsIs string
	if len(rawData) > 0 {
		qvApplicationsIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationsIsIC := swag.SplitByFormat(qvApplicationsIs, "")
	if len(applicationsIsIC) == 0 {
		return nil
	}

	var applicationsIsIR []int64
	for i, applicationsIsIV := range applicationsIsIC {
		applicationsIsI, err := swag.ConvertInt64(applicationsIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applications[is]", i), "query", "int64", applicationsIsI)
		}

		applicationsIsIR = append(applicationsIsIR, applicationsIsI)
	}

	o.ApplicationsIs = applicationsIsIR

	return nil
}

// bindApplicationsLte binds and validates parameter ApplicationsLte from query.
func (o *GetApplicationResourcesParams) bindApplicationsLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applications[lte]", "query", "int64", raw)
	}
	o.ApplicationsLte = &value

	return nil
}

// bindCisDockerBenchmarkLevelGte binds and validates parameter CisDockerBenchmarkLevelGte from query.
func (o *GetApplicationResourcesParams) bindCisDockerBenchmarkLevelGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CisDockerBenchmarkLevelGte = &raw

	if err := o.validateCisDockerBenchmarkLevelGte(formats); err != nil {
		return err
	}

	return nil
}

// validateCisDockerBenchmarkLevelGte carries on validations for parameter CisDockerBenchmarkLevelGte
func (o *GetApplicationResourcesParams) validateCisDockerBenchmarkLevelGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[gte]", "query", *o.CisDockerBenchmarkLevelGte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCisDockerBenchmarkLevelLte binds and validates parameter CisDockerBenchmarkLevelLte from query.
func (o *GetApplicationResourcesParams) bindCisDockerBenchmarkLevelLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CisDockerBenchmarkLevelLte = &raw

	if err := o.validateCisDockerBenchmarkLevelLte(formats); err != nil {
		return err
	}

	return nil
}

// validateCisDockerBenchmarkLevelLte carries on validations for parameter CisDockerBenchmarkLevelLte
func (o *GetApplicationResourcesParams) validateCisDockerBenchmarkLevelLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[lte]", "query", *o.CisDockerBenchmarkLevelLte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCurrentRuntimeScan binds and validates parameter CurrentRuntimeScan from query.
func (o *GetApplicationResourcesParams) bindCurrentRuntimeScan(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("currentRuntimeScan", "query", "bool", raw)
	}
	o.CurrentRuntimeScan = &value

	return nil
}

// bindPackageID binds and validates parameter PackageID from query.
func (o *GetApplicationResourcesParams) bindPackageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageID = &raw

	return nil
}

// bindPackagesGte binds and validates parameter PackagesGte from query.
func (o *GetApplicationResourcesParams) bindPackagesGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("packages[gte]", "query", "int64", raw)
	}
	o.PackagesGte = &value

	return nil
}

// bindPackagesIsNot binds and validates array parameter PackagesIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindPackagesIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackagesIsNot string
	if len(rawData) > 0 {
		qvPackagesIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packagesIsNotIC := swag.SplitByFormat(qvPackagesIsNot, "")
	if len(packagesIsNotIC) == 0 {
		return nil
	}

	var packagesIsNotIR []int64
	for i, packagesIsNotIV := range packagesIsNotIC {
		packagesIsNotI, err := swag.ConvertInt64(packagesIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "packages[isNot]", i), "query", "int64", packagesIsNotI)
		}

		packagesIsNotIR = append(packagesIsNotIR, packagesIsNotI)
	}

	o.PackagesIsNot = packagesIsNotIR

	return nil
}

// bindPackagesIs binds and validates array parameter PackagesIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindPackagesIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackagesIs string
	if len(rawData) > 0 {
		qvPackagesIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packagesIsIC := swag.SplitByFormat(qvPackagesIs, "")
	if len(packagesIsIC) == 0 {
		return nil
	}

	var packagesIsIR []int64
	for i, packagesIsIV := range packagesIsIC {
		packagesIsI, err := swag.ConvertInt64(packagesIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "packages[is]", i), "query", "int64", packagesIsI)
		}

		packagesIsIR = append(packagesIsIR, packagesIsI)
	}

	o.PackagesIs = packagesIsIR

	return nil
}

// bindPackagesLte binds and validates parameter PackagesLte from query.
func (o *GetApplicationResourcesParams) bindPackagesLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("packages[lte]", "query", "int64", raw)
	}
	o.PackagesLte = &value

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetApplicationResourcesParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("page", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("page", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetApplicationResourcesParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("pageSize", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("pageSize", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *GetApplicationResourcesParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", o.PageSize, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", o.PageSize, 50, false); err != nil {
		return err
	}

	return nil
}

// bindReportingSBOMAnalyzersContainElements binds and validates array parameter ReportingSBOMAnalyzersContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindReportingSBOMAnalyzersContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvReportingSBOMAnalyzersContainElements string
	if len(rawData) > 0 {
		qvReportingSBOMAnalyzersContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	reportingSBOMAnalyzersContainElementsIC := swag.SplitByFormat(qvReportingSBOMAnalyzersContainElements, "")
	if len(reportingSBOMAnalyzersContainElementsIC) == 0 {
		return nil
	}

	var reportingSBOMAnalyzersContainElementsIR []string
	for _, reportingSBOMAnalyzersContainElementsIV := range reportingSBOMAnalyzersContainElementsIC {
		reportingSBOMAnalyzersContainElementsI := reportingSBOMAnalyzersContainElementsIV

		reportingSBOMAnalyzersContainElementsIR = append(reportingSBOMAnalyzersContainElementsIR, reportingSBOMAnalyzersContainElementsI)
	}

	o.ReportingSBOMAnalyzersContainElements = reportingSBOMAnalyzersContainElementsIR

	return nil
}

// bindReportingSBOMAnalyzersDoesntContainElements binds and validates array parameter ReportingSBOMAnalyzersDoesntContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindReportingSBOMAnalyzersDoesntContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvReportingSBOMAnalyzersDoesntContainElements string
	if len(rawData) > 0 {
		qvReportingSBOMAnalyzersDoesntContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	reportingSBOMAnalyzersDoesntContainElementsIC := swag.SplitByFormat(qvReportingSBOMAnalyzersDoesntContainElements, "")
	if len(reportingSBOMAnalyzersDoesntContainElementsIC) == 0 {
		return nil
	}

	var reportingSBOMAnalyzersDoesntContainElementsIR []string
	for _, reportingSBOMAnalyzersDoesntContainElementsIV := range reportingSBOMAnalyzersDoesntContainElementsIC {
		reportingSBOMAnalyzersDoesntContainElementsI := reportingSBOMAnalyzersDoesntContainElementsIV

		reportingSBOMAnalyzersDoesntContainElementsIR = append(reportingSBOMAnalyzersDoesntContainElementsIR, reportingSBOMAnalyzersDoesntContainElementsI)
	}

	o.ReportingSBOMAnalyzersDoesntContainElements = reportingSBOMAnalyzersDoesntContainElementsIR

	return nil
}

// bindResourceHashContains binds and validates array parameter ResourceHashContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceHashContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceHashContains string
	if len(rawData) > 0 {
		qvResourceHashContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceHashContainsIC := swag.SplitByFormat(qvResourceHashContains, "")
	if len(resourceHashContainsIC) == 0 {
		return nil
	}

	var resourceHashContainsIR []string
	for _, resourceHashContainsIV := range resourceHashContainsIC {
		resourceHashContainsI := resourceHashContainsIV

		resourceHashContainsIR = append(resourceHashContainsIR, resourceHashContainsI)
	}

	o.ResourceHashContains = resourceHashContainsIR

	return nil
}

// bindResourceHashEnd binds and validates parameter ResourceHashEnd from query.
func (o *GetApplicationResourcesParams) bindResourceHashEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ResourceHashEnd = &raw

	return nil
}

// bindResourceHashIsNot binds and validates array parameter ResourceHashIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceHashIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceHashIsNot string
	if len(rawData) > 0 {
		qvResourceHashIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceHashIsNotIC := swag.SplitByFormat(qvResourceHashIsNot, "")
	if len(resourceHashIsNotIC) == 0 {
		return nil
	}

	var resourceHashIsNotIR []string
	for _, resourceHashIsNotIV := range resourceHashIsNotIC {
		resourceHashIsNotI := resourceHashIsNotIV

		resourceHashIsNotIR = append(resourceHashIsNotIR, resourceHashIsNotI)
	}

	o.ResourceHashIsNot = resourceHashIsNotIR

	return nil
}

// bindResourceHashIs binds and validates array parameter ResourceHashIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceHashIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceHashIs string
	if len(rawData) > 0 {
		qvResourceHashIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceHashIsIC := swag.SplitByFormat(qvResourceHashIs, "")
	if len(resourceHashIsIC) == 0 {
		return nil
	}

	var resourceHashIsIR []string
	for _, resourceHashIsIV := range resourceHashIsIC {
		resourceHashIsI := resourceHashIsIV

		resourceHashIsIR = append(resourceHashIsIR, resourceHashIsI)
	}

	o.ResourceHashIs = resourceHashIsIR

	return nil
}

// bindResourceHashStart binds and validates parameter ResourceHashStart from query.
func (o *GetApplicationResourcesParams) bindResourceHashStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ResourceHashStart = &raw

	return nil
}

// bindResourceNameContains binds and validates array parameter ResourceNameContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceNameContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceNameContains string
	if len(rawData) > 0 {
		qvResourceNameContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceNameContainsIC := swag.SplitByFormat(qvResourceNameContains, "")
	if len(resourceNameContainsIC) == 0 {
		return nil
	}

	var resourceNameContainsIR []string
	for _, resourceNameContainsIV := range resourceNameContainsIC {
		resourceNameContainsI := resourceNameContainsIV

		resourceNameContainsIR = append(resourceNameContainsIR, resourceNameContainsI)
	}

	o.ResourceNameContains = resourceNameContainsIR

	return nil
}

// bindResourceNameEnd binds and validates parameter ResourceNameEnd from query.
func (o *GetApplicationResourcesParams) bindResourceNameEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ResourceNameEnd = &raw

	return nil
}

// bindResourceNameIsNot binds and validates array parameter ResourceNameIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceNameIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceNameIsNot string
	if len(rawData) > 0 {
		qvResourceNameIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceNameIsNotIC := swag.SplitByFormat(qvResourceNameIsNot, "")
	if len(resourceNameIsNotIC) == 0 {
		return nil
	}

	var resourceNameIsNotIR []string
	for _, resourceNameIsNotIV := range resourceNameIsNotIC {
		resourceNameIsNotI := resourceNameIsNotIV

		resourceNameIsNotIR = append(resourceNameIsNotIR, resourceNameIsNotI)
	}

	o.ResourceNameIsNot = resourceNameIsNotIR

	return nil
}

// bindResourceNameIs binds and validates array parameter ResourceNameIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceNameIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceNameIs string
	if len(rawData) > 0 {
		qvResourceNameIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceNameIsIC := swag.SplitByFormat(qvResourceNameIs, "")
	if len(resourceNameIsIC) == 0 {
		return nil
	}

	var resourceNameIsIR []string
	for _, resourceNameIsIV := range resourceNameIsIC {
		resourceNameIsI := resourceNameIsIV

		resourceNameIsIR = append(resourceNameIsIR, resourceNameIsI)
	}

	o.ResourceNameIs = resourceNameIsIR

	return nil
}

// bindResourceNameStart binds and validates parameter ResourceNameStart from query.
func (o *GetApplicationResourcesParams) bindResourceNameStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ResourceNameStart = &raw

	return nil
}

// bindResourceTypeIs binds and validates array parameter ResourceTypeIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationResourcesParams) bindResourceTypeIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvResourceTypeIs string
	if len(rawData) > 0 {
		qvResourceTypeIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	resourceTypeIsIC := swag.SplitByFormat(qvResourceTypeIs, "")
	if len(resourceTypeIsIC) == 0 {
		return nil
	}

	var resourceTypeIsIR []string
	for i, resourceTypeIsIV := range resourceTypeIsIC {
		resourceTypeIsI := resourceTypeIsIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "resourceType[is]", i), "query", resourceTypeIsI, []interface{}{"IMAGE", "DIRECTORY", "FILE"}, true); err != nil {
			return err
		}

		resourceTypeIsIR = append(resourceTypeIsIR, resourceTypeIsI)
	}

	o.ResourceTypeIs = resourceTypeIsIR

	return nil
}

// bindSortDir binds and validates parameter SortDir from query.
func (o *GetApplicationResourcesParams) bindSortDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetApplicationResourcesParams()
		return nil
	}
	o.SortDir = &raw

	if err := o.validateSortDir(formats); err != nil {
		return err
	}

	return nil
}

// validateSortDir carries on validations for parameter SortDir
func (o *GetApplicationResourcesParams) validateSortDir(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortDir", "query", *o.SortDir, []interface{}{"ASC", "DESC"}, true); err != nil {
		return err
	}

	return nil
}

// bindSortKey binds and validates parameter SortKey from query.
func (o *GetApplicationResourcesParams) bindSortKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sortKey", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("sortKey", "query", raw); err != nil {
		return err
	}
	o.SortKey = raw

	if err := o.validateSortKey(formats); err != nil {
		return err
	}

	return nil
}

// validateSortKey carries on validations for parameter SortKey
func (o *GetApplicationResourcesParams) validateSortKey(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortKey", "query", o.SortKey, []interface{}{"resourceName", "resourceHash", "resourceType", "vulnerabilities", "applications", "packages"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityGte binds and validates parameter VulnerabilitySeverityGte from query.
func (o *GetApplicationResourcesParams) bindVulnerabilitySeverityGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.VulnerabilitySeverityGte = &raw

	if err := o.validateVulnerabilitySeverityGte(formats); err != nil {
		return err
	}

	return nil
}

// validateVulnerabilitySeverityGte carries on validations for parameter VulnerabilitySeverityGte
func (o *GetApplicationResourcesParams) validateVulnerabilitySeverityGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[gte]", "query", *o.VulnerabilitySeverityGte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityLte binds and validates parameter VulnerabilitySeverityLte from query.
func (o *GetApplicationResourcesParams) bindVulnerabilitySeverityLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.VulnerabilitySeverityLte = &raw

	if err := o.validateVulnerabilitySeverityLte(formats); err != nil {
		return err
	}

	return nil
}

// validateVulnerabilitySeverityLte carries on validations for parameter VulnerabilitySeverityLte
func (o *GetApplicationResourcesParams) validateVulnerabilitySeverityLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[lte]", "query", *o.VulnerabilitySeverityLte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}
