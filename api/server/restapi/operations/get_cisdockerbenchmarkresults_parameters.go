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

// NewGetCisdockerbenchmarkresultsParams creates a new GetCisdockerbenchmarkresultsParams object
// with the default values initialized.
func NewGetCisdockerbenchmarkresultsParams() GetCisdockerbenchmarkresultsParams {

	var (
		// initialize parameters with default values

		sortDirDefault = string("ASC")
	)

	return GetCisdockerbenchmarkresultsParams{
		SortDir: &sortDirDefault,
	}
}

// GetCisdockerbenchmarkresultsParams contains all the bound params for the get cisdockerbenchmarkresults operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetCisdockerbenchmarkresults
type GetCisdockerbenchmarkresultsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*application ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	ApplicationID *string
	/*application resource ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	ApplicationResourceID *string
	/*greater than or equal
	  In: query
	*/
	ApplicationResourcesGte *int64
	/*
	  In: query
	*/
	ApplicationResourcesIsNot []int64
	/*
	  In: query
	*/
	ApplicationResourcesIs []int64
	/*less than or equal
	  In: query
	*/
	ApplicationResourcesLte *int64
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
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCisdockerbenchmarkresultsParams() beforehand.
func (o *GetCisdockerbenchmarkresultsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplicationID, qhkApplicationID, _ := qs.GetOK("applicationID")
	if err := o.bindApplicationID(qApplicationID, qhkApplicationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourceID, qhkApplicationResourceID, _ := qs.GetOK("applicationResourceID")
	if err := o.bindApplicationResourceID(qApplicationResourceID, qhkApplicationResourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesGte, qhkApplicationResourcesGte, _ := qs.GetOK("applicationResources[gte]")
	if err := o.bindApplicationResourcesGte(qApplicationResourcesGte, qhkApplicationResourcesGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesIsNot, qhkApplicationResourcesIsNot, _ := qs.GetOK("applicationResources[isNot]")
	if err := o.bindApplicationResourcesIsNot(qApplicationResourcesIsNot, qhkApplicationResourcesIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesIs, qhkApplicationResourcesIs, _ := qs.GetOK("applicationResources[is]")
	if err := o.bindApplicationResourcesIs(qApplicationResourcesIs, qhkApplicationResourcesIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesLte, qhkApplicationResourcesLte, _ := qs.GetOK("applicationResources[lte]")
	if err := o.bindApplicationResourcesLte(qApplicationResourcesLte, qhkApplicationResourcesLte, route.Formats); err != nil {
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

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
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
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplicationID binds and validates parameter ApplicationID from query.
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindApplicationResourceID binds and validates parameter ApplicationResourceID from query.
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationResourceID = &raw

	return nil
}

// bindApplicationResourcesGte binds and validates parameter ApplicationResourcesGte from query.
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationResourcesGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("applicationResources[gte]", "query", "int64", raw)
	}
	o.ApplicationResourcesGte = &value

	return nil
}

// bindApplicationResourcesIsNot binds and validates array parameter ApplicationResourcesIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationResourcesIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationResourcesIsNot string
	if len(rawData) > 0 {
		qvApplicationResourcesIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationResourcesIsNotIC := swag.SplitByFormat(qvApplicationResourcesIsNot, "")
	if len(applicationResourcesIsNotIC) == 0 {
		return nil
	}

	var applicationResourcesIsNotIR []int64
	for i, applicationResourcesIsNotIV := range applicationResourcesIsNotIC {
		applicationResourcesIsNotI, err := swag.ConvertInt64(applicationResourcesIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applicationResources[isNot]", i), "query", "int64", applicationResourcesIsNotI)
		}

		applicationResourcesIsNotIR = append(applicationResourcesIsNotIR, applicationResourcesIsNotI)
	}

	o.ApplicationResourcesIsNot = applicationResourcesIsNotIR

	return nil
}

// bindApplicationResourcesIs binds and validates array parameter ApplicationResourcesIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationResourcesIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationResourcesIs string
	if len(rawData) > 0 {
		qvApplicationResourcesIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationResourcesIsIC := swag.SplitByFormat(qvApplicationResourcesIs, "")
	if len(applicationResourcesIsIC) == 0 {
		return nil
	}

	var applicationResourcesIsIR []int64
	for i, applicationResourcesIsIV := range applicationResourcesIsIC {
		applicationResourcesIsI, err := swag.ConvertInt64(applicationResourcesIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applicationResources[is]", i), "query", "int64", applicationResourcesIsI)
		}

		applicationResourcesIsIR = append(applicationResourcesIsIR, applicationResourcesIsI)
	}

	o.ApplicationResourcesIs = applicationResourcesIsIR

	return nil
}

// bindApplicationResourcesLte binds and validates parameter ApplicationResourcesLte from query.
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationResourcesLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("applicationResources[lte]", "query", "int64", raw)
	}
	o.ApplicationResourcesLte = &value

	return nil
}

// bindApplicationsGte binds and validates parameter ApplicationsGte from query.
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationsGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationsIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationsIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) bindApplicationsLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) bindCisDockerBenchmarkLevelGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) validateCisDockerBenchmarkLevelGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[gte]", "query", *o.CisDockerBenchmarkLevelGte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCisDockerBenchmarkLevelLte binds and validates parameter CisDockerBenchmarkLevelLte from query.
func (o *GetCisdockerbenchmarkresultsParams) bindCisDockerBenchmarkLevelLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) validateCisDockerBenchmarkLevelLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[lte]", "query", *o.CisDockerBenchmarkLevelLte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCurrentRuntimeScan binds and validates parameter CurrentRuntimeScan from query.
func (o *GetCisdockerbenchmarkresultsParams) bindCurrentRuntimeScan(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindPage binds and validates parameter Page from query.
func (o *GetCisdockerbenchmarkresultsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", o.PageSize, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", o.PageSize, 50, false); err != nil {
		return err
	}

	return nil
}

// bindSortDir binds and validates parameter SortDir from query.
func (o *GetCisdockerbenchmarkresultsParams) bindSortDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetCisdockerbenchmarkresultsParams()
		return nil
	}
	o.SortDir = &raw

	if err := o.validateSortDir(formats); err != nil {
		return err
	}

	return nil
}

// validateSortDir carries on validations for parameter SortDir
func (o *GetCisdockerbenchmarkresultsParams) validateSortDir(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortDir", "query", *o.SortDir, []interface{}{"ASC", "DESC"}, true); err != nil {
		return err
	}

	return nil
}

// bindSortKey binds and validates parameter SortKey from query.
func (o *GetCisdockerbenchmarkresultsParams) bindSortKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCisdockerbenchmarkresultsParams) validateSortKey(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortKey", "query", o.SortKey, []interface{}{"code", "title", "level"}, true); err != nil {
		return err
	}

	return nil
}
