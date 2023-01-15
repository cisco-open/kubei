// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Defines values for CloudProvider.
const (
	AWS CloudProvider = "AWS"
)

// Defines values for MalwareType.
const (
	ADWARE     MalwareType = "ADWARE"
	RANSOMWARE MalwareType = "RANSOMWARE"
	SPYWARE    MalwareType = "SPYWARE"
	TROJAN     MalwareType = "TROJAN"
	VIRUS      MalwareType = "VIRUS"
	WORM       MalwareType = "WORM"
)

// Defines values for RootkitType.
const (
	APPLICATION RootkitType = "APPLICATION"
	FIRMWARE    RootkitType = "FIRMWARE"
	KERNEL      RootkitType = "KERNEL"
	MEMORY      RootkitType = "MEMORY"
)

// Defines values for ScanType.
const (
	EXPLOIT          ScanType = "EXPLOIT"
	MALWARE          ScanType = "MALWARE"
	MISCONFIGURATION ScanType = "MISCONFIGURATION"
	ROOTKIT          ScanType = "ROOTKIT"
	SBOM             ScanType = "SBOM"
	SECRET           ScanType = "SECRET"
	VULNERABILITY    ScanType = "VULNERABILITY"
)

// Defines values for TargetScanStateState.
const (
	DONE       TargetScanStateState = "DONE"
	INIT       TargetScanStateState = "INIT"
	INPROGRESS TargetScanStateState = "IN_PROGRESS"
	NOTSCANNED TargetScanStateState = "NOT_SCANNED"
)

// ApiResponse An object that is returned in all cases of failures.
type ApiResponse struct {
	Message *string `json:"message,omitempty"`
}

// AwsRegion AWS region
type AwsRegion struct {
	Id   *string   `json:"id,omitempty"`
	Name *string   `json:"name,omitempty"`
	Vpcs *[]AwsVPC `json:"vpcs,omitempty"`
}

// AwsScanScope The scope of a configured scan.
type AwsScanScope struct {
	All *bool `json:"all,omitempty"`

	// InstanceTagExclusion VM instances will not be scanned if they contain all of these tags (even if they match instanceTagSelector). If empty, not taken into account.
	InstanceTagExclusion *[]Tag `json:"instanceTagExclusion,omitempty"`

	// InstanceTagSelector VM instances will be scanned if they contain all of these tags. If empty, not taken into account.
	InstanceTagSelector        *[]Tag       `json:"instanceTagSelector,omitempty"`
	ObjectType                 string       `json:"objectType"`
	Regions                    *[]AwsRegion `json:"regions,omitempty"`
	ShouldScanStoppedInstances *bool        `json:"shouldScanStoppedInstances,omitempty"`
}

// AwsSecurityGroup AWS security group
type AwsSecurityGroup struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// AwsVPC AWS VPC
type AwsVPC struct {
	Id             *string             `json:"id,omitempty"`
	Name           *string             `json:"name,omitempty"`
	SecurityGroups *[]AwsSecurityGroup `json:"securityGroups,omitempty"`
}

// ByDaysScheduleScanConfig defines model for ByDaysScheduleScanConfig.
type ByDaysScheduleScanConfig struct {
	DaysInterval *int       `json:"daysInterval,omitempty"`
	ObjectType   string     `json:"objectType"`
	TimeOfDay    *TimeOfDay `json:"timeOfDay,omitempty"`
}

// ByHoursScheduleScanConfig defines model for ByHoursScheduleScanConfig.
type ByHoursScheduleScanConfig struct {
	HoursInterval *int   `json:"hoursInterval,omitempty"`
	ObjectType    string `json:"objectType"`
}

// CloudProvider defines model for CloudProvider.
type CloudProvider string

// DirInfo defines model for DirInfo.
type DirInfo struct {
	DirName    *string `json:"dirName,omitempty"`
	Location   *string `json:"location,omitempty"`
	ObjectType string  `json:"objectType"`
}

// ExploitInfo defines model for ExploitInfo.
type ExploitInfo struct {
	Description     *string   `json:"description,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Vulnerabilities *[]string `json:"vulnerabilities,omitempty"`
}

// ExploitScan defines model for ExploitScan.
type ExploitScan struct {
	Exploits *[]ExploitInfo `json:"exploits,omitempty"`
}

// ExploitsConfig defines model for ExploitsConfig.
type ExploitsConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// MalwareConfig defines model for MalwareConfig.
type MalwareConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// MalwareInfo defines model for MalwareInfo.
type MalwareInfo struct {
	Id          *string      `json:"id,omitempty"`
	MalwareName *string      `json:"malwareName,omitempty"`
	MalwareType *MalwareType `json:"malwareType,omitempty"`

	// Path Path of the file that contains malware
	Path *string `json:"path,omitempty"`
}

// MalwareScan defines model for MalwareScan.
type MalwareScan struct {
	Malware *[]MalwareInfo `json:"malware,omitempty"`
}

// MalwareType defines model for MalwareType.
type MalwareType string

// MisconfigurationInfo defines model for MisconfigurationInfo.
type MisconfigurationInfo struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`

	// Path Path of the file that contains misconfigurations
	Path *string `json:"path,omitempty"`
}

// MisconfigurationScan defines model for MisconfigurationScan.
type MisconfigurationScan struct {
	Misconfigurations *[]MisconfigurationInfo `json:"misconfigurations,omitempty"`
}

// MisconfigurationsConfig defines model for MisconfigurationsConfig.
type MisconfigurationsConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Id          *string      `json:"id,omitempty"`
	PackageInfo *PackageInfo `json:"packageInfo,omitempty"`
}

// PackageInfo defines model for PackageInfo.
type PackageInfo struct {
	Id             *string `json:"id,omitempty"`
	PackageName    *string `json:"packageName,omitempty"`
	PackageVersion *string `json:"packageVersion,omitempty"`
}

// PodInfo defines model for PodInfo.
type PodInfo struct {
	Location   *string `json:"location,omitempty"`
	ObjectType string  `json:"objectType"`
	PodName    *string `json:"podName,omitempty"`
}

// RootkitInfo defines model for RootkitInfo.
type RootkitInfo struct {
	Id *string `json:"id,omitempty"`

	// Path Path of the file that contains rootkit
	Path        *string      `json:"path,omitempty"`
	RootkitName *string      `json:"rootkitName,omitempty"`
	RootkitType *RootkitType `json:"rootkitType,omitempty"`
}

// RootkitScan defines model for RootkitScan.
type RootkitScan struct {
	Rootkits *[]RootkitInfo `json:"rootkits,omitempty"`
}

// RootkitType defines model for RootkitType.
type RootkitType string

// RootkitsConfig defines model for RootkitsConfig.
type RootkitsConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// RuntimeScheduleScanConfigType defines model for RuntimeScheduleScanConfigType.
type RuntimeScheduleScanConfigType struct {
	union json.RawMessage
}

// SBOMConfig defines model for SBOMConfig.
type SBOMConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// SbomScan defines model for SbomScan.
type SbomScan struct {
	Packages *[]Package `json:"packages,omitempty"`
}

// Scan Describes a multi-target scheduled scan.
type Scan struct {
	EndTime *time.Time `json:"endTime,omitempty"`
	Id      *string    `json:"id,omitempty"`

	// ScanConfigId The ID of the config that this scan was initiated from (optionanl)
	ScanConfigId *string `json:"scanConfigId,omitempty"`

	// ScanFamiliesConfig The configuration of the scanner families within a scan config
	ScanFamiliesConfig *ScanFamiliesConfig `json:"scanFamiliesConfig,omitempty"`
	StartTime          *time.Time          `json:"startTime,omitempty"`

	// TargetIDs List of target IDs that are targeted for scanning as part of this scan
	TargetIDs *[]string `json:"targetIDs,omitempty"`
}

// ScanConfig Describes a multi-target scheduled scan config.
type ScanConfig struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	// ScanFamiliesConfig The configuration of the scanner families within a scan config
	ScanFamiliesConfig *ScanFamiliesConfig            `json:"scanFamiliesConfig,omitempty"`
	Scheduled          *RuntimeScheduleScanConfigType `json:"scheduled,omitempty"`
	Scope              *ScanScopeType                 `json:"scope,omitempty"`
}

// ScanConfigs defines model for ScanConfigs.
type ScanConfigs struct {
	// Items List of scan configs according to the given filters and page. List length must be lower or equal to pageSize.
	Items *[]ScanConfig `json:"items,omitempty"`

	// Total Total scan config count according to the given filters
	Total int `json:"total"`
}

// ScanFamiliesConfig The configuration of the scanner families within a scan config
type ScanFamiliesConfig struct {
	Exploits          *ExploitsConfig          `json:"exploits,omitempty"`
	Malware           *MalwareConfig           `json:"malware,omitempty"`
	Misconfigurations *MisconfigurationsConfig `json:"misconfigurations,omitempty"`
	Rootkits          *RootkitsConfig          `json:"rootkits,omitempty"`
	Sbom              *SBOMConfig              `json:"sbom,omitempty"`
	Secrets           *SecretsConfig           `json:"secrets,omitempty"`
	Vulnerabilities   *VulnerabilitiesConfig   `json:"vulnerabilities,omitempty"`
}

// ScanScopeType defines model for ScanScopeType.
type ScanScopeType struct {
	union json.RawMessage
}

// ScanType defines model for ScanType.
type ScanType string

// Scans defines model for Scans.
type Scans struct {
	// Items List of scans according to the given filters and page. List length must be lower or equal to pageSize.
	Items *[]Scan `json:"items,omitempty"`

	// Total Total scans count according to the given filters
	Total int `json:"total"`
}

// SecretInfo defines model for SecretInfo.
type SecretInfo struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`

	// Path Path of the file that contains secrets
	Path *string `json:"path,omitempty"`
}

// SecretScan defines model for SecretScan.
type SecretScan struct {
	Secrets *[]SecretInfo `json:"secrets,omitempty"`
}

// SecretsConfig defines model for SecretsConfig.
type SecretsConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// SingleScheduleScanConfig defines model for SingleScheduleScanConfig.
type SingleScheduleScanConfig struct {
	ObjectType    string    `json:"objectType"`
	OperationTime time.Time `json:"operationTime"`
}

// SuccessResponse An object that is returned in cases of success that returns nothing.
type SuccessResponse struct {
	Message *string `json:"message,omitempty"`
}

// Tag AWS tag
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Target defines model for Target.
type Target struct {
	Id         *string     `json:"id,omitempty"`
	TargetInfo *TargetType `json:"targetInfo,omitempty"`
}

// TargetScanResult defines model for TargetScanResult.
type TargetScanResult struct {
	Exploits          *ExploitScan          `json:"exploits,omitempty"`
	Id                *string               `json:"id,omitempty"`
	Malware           *MalwareScan          `json:"malware,omitempty"`
	Misconfigurations *MisconfigurationScan `json:"misconfigurations,omitempty"`
	Rootkits          *RootkitScan          `json:"rootkits,omitempty"`
	Sboms             *SbomScan             `json:"sboms,omitempty"`
	ScanId            string                `json:"scanId"`
	Secrets           *SecretScan           `json:"secrets,omitempty"`
	Status            *TargetScanStatus     `json:"status,omitempty"`
	TargetId          string                `json:"targetId"`
	Vulnerabilities   *VulnerabilityScan    `json:"vulnerabilities,omitempty"`
}

// TargetScanResults defines model for TargetScanResults.
type TargetScanResults struct {
	// Items List of scan results according to the given filters and page. List length must be lower or equal to pageSize.
	Items *[]TargetScanResult `json:"items,omitempty"`

	// Total Total scan results count according to the given filters
	Total int `json:"total"`
}

// TargetScanState defines model for TargetScanState.
type TargetScanState struct {
	Errors *[]string             `json:"errors,omitempty"`
	State  *TargetScanStateState `json:"state,omitempty"`
}

// TargetScanStateState defines model for TargetScanState.State.
type TargetScanStateState string

// TargetScanStatus defines model for TargetScanStatus.
type TargetScanStatus struct {
	Exploits          *TargetScanState `json:"exploits,omitempty"`
	General           *TargetScanState `json:"general,omitempty"`
	Malware           *TargetScanState `json:"malware,omitempty"`
	Misconfigurations *TargetScanState `json:"misconfigurations,omitempty"`
	Rootkits          *TargetScanState `json:"rootkits,omitempty"`
	Sbom              *TargetScanState `json:"sbom,omitempty"`
	Secrets           *TargetScanState `json:"secrets,omitempty"`
	Vulnerabilities   *TargetScanState `json:"vulnerabilities,omitempty"`
}

// TargetType defines model for TargetType.
type TargetType struct {
	union json.RawMessage
}

// Targets defines model for Targets.
type Targets struct {
	// Items List of targets in the given filters and page. List length must be lower or equal to pageSize.
	Items *[]Target `json:"items,omitempty"`

	// Total Total targets count according the given filters
	Total int `json:"total"`
}

// TimeOfDay defines model for TimeOfDay.
type TimeOfDay struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

// VMInfo defines model for VMInfo.
type VMInfo struct {
	InstanceID       *string        `json:"instanceID,omitempty"`
	InstanceProvider *CloudProvider `json:"instanceProvider,omitempty"`
	Location         *string        `json:"location,omitempty"`
	ObjectType       string         `json:"objectType"`
}

// VulnerabilitiesConfig defines model for VulnerabilitiesConfig.
type VulnerabilitiesConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// Vulnerability defines model for Vulnerability.
type Vulnerability struct {
	VulnerabilityInfo *VulnerabilityInfo `json:"VulnerabilityInfo,omitempty"`
	Id                *string            `json:"id,omitempty"`
}

// VulnerabilityInfo defines model for VulnerabilityInfo.
type VulnerabilityInfo struct {
	Description       *string `json:"description,omitempty"`
	Id                *string `json:"id,omitempty"`
	VulnerabilityName *string `json:"vulnerabilityName,omitempty"`
}

// VulnerabilityScan defines model for VulnerabilityScan.
type VulnerabilityScan struct {
	Vulnerabilities *[]Vulnerability `json:"vulnerabilities,omitempty"`
}

// WeeklyScheduleScanConfig defines model for WeeklyScheduleScanConfig.
type WeeklyScheduleScanConfig struct {
	// DayInWeek 1 - 7 which represents sun- sat
	DayInWeek  *int       `json:"dayInWeek,omitempty"`
	ObjectType string     `json:"objectType"`
	TimeOfDay  *TimeOfDay `json:"timeOfDay,omitempty"`
}

// OdataFilter defines model for odataFilter.
type OdataFilter = string

// OdataSelect defines model for odataSelect.
type OdataSelect = string

// Page defines model for page.
type Page = int

// PageSize defines model for pageSize.
type PageSize = int

// ScanConfigID defines model for scanConfigID.
type ScanConfigID = string

// ScanID defines model for scanID.
type ScanID = string

// ScanResultID defines model for scanResultID.
type ScanResultID = string

// TargetID defines model for targetID.
type TargetID = string

// Success An object that is returned in cases of success that returns nothing.
type Success = SuccessResponse

// UnknownError An object that is returned in all cases of failures.
type UnknownError = ApiResponse

// GetScanConfigsParams defines parameters for GetScanConfigs.
type GetScanConfigsParams struct {
	Filter *OdataFilter `form:"$filter,omitempty" json:"$filter,omitempty"`

	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// GetScanResultsParams defines parameters for GetScanResults.
type GetScanResultsParams struct {
	Filter *OdataFilter `form:"$filter,omitempty" json:"$filter,omitempty"`
	Select *OdataSelect `form:"$select,omitempty" json:"$select,omitempty"`

	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// GetScanResultsScanResultIDParams defines parameters for GetScanResultsScanResultID.
type GetScanResultsScanResultIDParams struct {
	Select *OdataSelect `form:"$select,omitempty" json:"$select,omitempty"`
}

// GetScansParams defines parameters for GetScans.
type GetScansParams struct {
	Filter *OdataFilter `form:"$filter,omitempty" json:"$filter,omitempty"`

	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// GetTargetsParams defines parameters for GetTargets.
type GetTargetsParams struct {
	Filter *OdataFilter `form:"$filter,omitempty" json:"$filter,omitempty"`

	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// PostScanConfigsJSONRequestBody defines body for PostScanConfigs for application/json ContentType.
type PostScanConfigsJSONRequestBody = ScanConfig

// PatchScanConfigsScanConfigIDJSONRequestBody defines body for PatchScanConfigsScanConfigID for application/json ContentType.
type PatchScanConfigsScanConfigIDJSONRequestBody = ScanConfig

// PutScanConfigsScanConfigIDJSONRequestBody defines body for PutScanConfigsScanConfigID for application/json ContentType.
type PutScanConfigsScanConfigIDJSONRequestBody = ScanConfig

// PostScanResultsJSONRequestBody defines body for PostScanResults for application/json ContentType.
type PostScanResultsJSONRequestBody = TargetScanResult

// PatchScanResultsScanResultIDJSONRequestBody defines body for PatchScanResultsScanResultID for application/json ContentType.
type PatchScanResultsScanResultIDJSONRequestBody = TargetScanResult

// PutScanResultsScanResultIDJSONRequestBody defines body for PutScanResultsScanResultID for application/json ContentType.
type PutScanResultsScanResultIDJSONRequestBody = TargetScanResult

// PostScansJSONRequestBody defines body for PostScans for application/json ContentType.
type PostScansJSONRequestBody = Scan

// PatchScansScanIDJSONRequestBody defines body for PatchScansScanID for application/json ContentType.
type PatchScansScanIDJSONRequestBody = Scan

// PutScansScanIDJSONRequestBody defines body for PutScansScanID for application/json ContentType.
type PutScansScanIDJSONRequestBody = Scan

// PostTargetsJSONRequestBody defines body for PostTargets for application/json ContentType.
type PostTargetsJSONRequestBody = Target

// PutTargetsTargetIDJSONRequestBody defines body for PutTargetsTargetID for application/json ContentType.
type PutTargetsTargetIDJSONRequestBody = Target

// AsSingleScheduleScanConfig returns the union data inside the RuntimeScheduleScanConfigType as a SingleScheduleScanConfig
func (t RuntimeScheduleScanConfigType) AsSingleScheduleScanConfig() (SingleScheduleScanConfig, error) {
	var body SingleScheduleScanConfig
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSingleScheduleScanConfig overwrites any union data inside the RuntimeScheduleScanConfigType as the provided SingleScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) FromSingleScheduleScanConfig(v SingleScheduleScanConfig) error {
	v.ObjectType = "SingleScheduleScanConfig"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSingleScheduleScanConfig performs a merge with any union data inside the RuntimeScheduleScanConfigType, using the provided SingleScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) MergeSingleScheduleScanConfig(v SingleScheduleScanConfig) error {
	v.ObjectType = "SingleScheduleScanConfig"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsByHoursScheduleScanConfig returns the union data inside the RuntimeScheduleScanConfigType as a ByHoursScheduleScanConfig
func (t RuntimeScheduleScanConfigType) AsByHoursScheduleScanConfig() (ByHoursScheduleScanConfig, error) {
	var body ByHoursScheduleScanConfig
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromByHoursScheduleScanConfig overwrites any union data inside the RuntimeScheduleScanConfigType as the provided ByHoursScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) FromByHoursScheduleScanConfig(v ByHoursScheduleScanConfig) error {
	v.ObjectType = "ByHoursScheduleScanConfig"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeByHoursScheduleScanConfig performs a merge with any union data inside the RuntimeScheduleScanConfigType, using the provided ByHoursScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) MergeByHoursScheduleScanConfig(v ByHoursScheduleScanConfig) error {
	v.ObjectType = "ByHoursScheduleScanConfig"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsByDaysScheduleScanConfig returns the union data inside the RuntimeScheduleScanConfigType as a ByDaysScheduleScanConfig
func (t RuntimeScheduleScanConfigType) AsByDaysScheduleScanConfig() (ByDaysScheduleScanConfig, error) {
	var body ByDaysScheduleScanConfig
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromByDaysScheduleScanConfig overwrites any union data inside the RuntimeScheduleScanConfigType as the provided ByDaysScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) FromByDaysScheduleScanConfig(v ByDaysScheduleScanConfig) error {
	v.ObjectType = "ByDaysScheduleScanConfig"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeByDaysScheduleScanConfig performs a merge with any union data inside the RuntimeScheduleScanConfigType, using the provided ByDaysScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) MergeByDaysScheduleScanConfig(v ByDaysScheduleScanConfig) error {
	v.ObjectType = "ByDaysScheduleScanConfig"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsWeeklyScheduleScanConfig returns the union data inside the RuntimeScheduleScanConfigType as a WeeklyScheduleScanConfig
func (t RuntimeScheduleScanConfigType) AsWeeklyScheduleScanConfig() (WeeklyScheduleScanConfig, error) {
	var body WeeklyScheduleScanConfig
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWeeklyScheduleScanConfig overwrites any union data inside the RuntimeScheduleScanConfigType as the provided WeeklyScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) FromWeeklyScheduleScanConfig(v WeeklyScheduleScanConfig) error {
	v.ObjectType = "WeeklyScheduleScanConfig"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWeeklyScheduleScanConfig performs a merge with any union data inside the RuntimeScheduleScanConfigType, using the provided WeeklyScheduleScanConfig
func (t *RuntimeScheduleScanConfigType) MergeWeeklyScheduleScanConfig(v WeeklyScheduleScanConfig) error {
	v.ObjectType = "WeeklyScheduleScanConfig"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t RuntimeScheduleScanConfigType) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"objectType"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t RuntimeScheduleScanConfigType) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "ByDaysScheduleScanConfig":
		return t.AsByDaysScheduleScanConfig()
	case "ByHoursScheduleScanConfig":
		return t.AsByHoursScheduleScanConfig()
	case "SingleScheduleScanConfig":
		return t.AsSingleScheduleScanConfig()
	case "WeeklyScheduleScanConfig":
		return t.AsWeeklyScheduleScanConfig()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t RuntimeScheduleScanConfigType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *RuntimeScheduleScanConfigType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAwsScanScope returns the union data inside the ScanScopeType as a AwsScanScope
func (t ScanScopeType) AsAwsScanScope() (AwsScanScope, error) {
	var body AwsScanScope
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAwsScanScope overwrites any union data inside the ScanScopeType as the provided AwsScanScope
func (t *ScanScopeType) FromAwsScanScope(v AwsScanScope) error {
	v.ObjectType = "AwsScanScope"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAwsScanScope performs a merge with any union data inside the ScanScopeType, using the provided AwsScanScope
func (t *ScanScopeType) MergeAwsScanScope(v AwsScanScope) error {
	v.ObjectType = "AwsScanScope"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t ScanScopeType) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"objectType"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t ScanScopeType) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "AwsScanScope":
		return t.AsAwsScanScope()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t ScanScopeType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ScanScopeType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsVMInfo returns the union data inside the TargetType as a VMInfo
func (t TargetType) AsVMInfo() (VMInfo, error) {
	var body VMInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVMInfo overwrites any union data inside the TargetType as the provided VMInfo
func (t *TargetType) FromVMInfo(v VMInfo) error {
	v.ObjectType = "VMInfo"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVMInfo performs a merge with any union data inside the TargetType, using the provided VMInfo
func (t *TargetType) MergeVMInfo(v VMInfo) error {
	v.ObjectType = "VMInfo"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsPodInfo returns the union data inside the TargetType as a PodInfo
func (t TargetType) AsPodInfo() (PodInfo, error) {
	var body PodInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPodInfo overwrites any union data inside the TargetType as the provided PodInfo
func (t *TargetType) FromPodInfo(v PodInfo) error {
	v.ObjectType = "PodInfo"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePodInfo performs a merge with any union data inside the TargetType, using the provided PodInfo
func (t *TargetType) MergePodInfo(v PodInfo) error {
	v.ObjectType = "PodInfo"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsDirInfo returns the union data inside the TargetType as a DirInfo
func (t TargetType) AsDirInfo() (DirInfo, error) {
	var body DirInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDirInfo overwrites any union data inside the TargetType as the provided DirInfo
func (t *TargetType) FromDirInfo(v DirInfo) error {
	v.ObjectType = "DirInfo"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDirInfo performs a merge with any union data inside the TargetType, using the provided DirInfo
func (t *TargetType) MergeDirInfo(v DirInfo) error {
	v.ObjectType = "DirInfo"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t TargetType) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"objectType"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t TargetType) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "DirInfo":
		return t.AsDirInfo()
	case "PodInfo":
		return t.AsPodInfo()
	case "VMInfo":
		return t.AsVMInfo()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t TargetType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TargetType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
