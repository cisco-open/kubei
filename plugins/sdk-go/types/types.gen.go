// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.1-0.20240819063527-fa048ac69eea DO NOT EDIT.
package types

import (
	"time"
)

// Defines values for InfoFinderType.
const (
	InfoFinderTypeSSHAuthorizedKeyFingerprint InfoFinderType = "SSHAuthorizedKeyFingerprint"
	InfoFinderTypeSSHDaemonKeyFingerprint     InfoFinderType = "SSHDaemonKeyFingerprint"
	InfoFinderTypeSSHKnownHostFingerprint     InfoFinderType = "SSHKnownHostFingerprint"
	InfoFinderTypeSSHPrivateKeyFingerprint    InfoFinderType = "SSHPrivateKeyFingerprint"
	InfoFinderTypeUnknown                     InfoFinderType = "UNKNOWN"
)

// Defines values for MisconfigurationSeverity.
const (
	MisconfigurationSeverityHigh   MisconfigurationSeverity = "HIGH"
	MisconfigurationSeverityInfo   MisconfigurationSeverity = "INFO"
	MisconfigurationSeverityLow    MisconfigurationSeverity = "LOW"
	MisconfigurationSeverityMedium MisconfigurationSeverity = "MEDIUM"
)

// Defines values for RootkitType.
const (
	RootkitTypeApplication RootkitType = "APPLICATION"
	RootkitTypeFirmware    RootkitType = "FIRMWARE"
	RootkitTypeKernel      RootkitType = "KERNEL"
	RootkitTypeMemory      RootkitType = "MEMORY"
	RootkitTypeUnknown     RootkitType = "UNKNOWN"
)

// Defines values for State.
const (
	StateDone     State = "Done"
	StateFailed   State = "Failed"
	StateNotReady State = "NotReady"
	StateReady    State = "Ready"
	StateRunning  State = "Running"
)

// Defines values for VulnerabilitySeverity.
const (
	VulnerabilitySeverityCritical   VulnerabilitySeverity = "CRITICAL"
	VulnerabilitySeverityHigh       VulnerabilitySeverity = "HIGH"
	VulnerabilitySeverityLow        VulnerabilitySeverity = "LOW"
	VulnerabilitySeverityMedium     VulnerabilitySeverity = "MEDIUM"
	VulnerabilitySeverityNegligible VulnerabilitySeverity = "NEGLIGIBLE"
)

// Annotations Generic map of string keys and string values to attach arbitrary non-identifying metadata to objects.
type Annotations map[string]string

// Config Describes config for scanner to start the scanning process.
type Config struct {
	// InputDir The directory which should be scanned by the scanner plugin.
	InputDir string `json:"inputDir" validate:"required"`

	// OutputFile Path to JSON file where the scanner plugin should store its results.
	OutputFile string `json:"outputFile" validate:"required"`

	// ScannerConfig Optional JSON string of internal scanner configuration used to override default scanner behaviour.
	// The config schema needs to be documented and payload manually parsed by the developer of scanner plugin.
	ScannerConfig *string `json:"scannerConfig,omitempty"`

	// TimeoutSeconds The maximum time in seconds that a scan started from this scan
	// should run for before being automatically aborted.
	TimeoutSeconds int `json:"timeoutSeconds" validate:"required,gt=0"`
}

// ErrorResponse An object that is returned for a failed API request.
type ErrorResponse struct {
	Message *string `json:"message,omitempty"`
}

// Exploit defines model for Exploit.
type Exploit struct {
	CveID       *string   `json:"cveID,omitempty"`
	Description *string   `json:"description,omitempty"`
	Name        *string   `json:"name,omitempty"`
	SourceDB    *string   `json:"sourceDB,omitempty"`
	Title       *string   `json:"title,omitempty"`
	Urls        *[]string `json:"urls"`
}

// InfoFinder defines model for InfoFinder.
type InfoFinder struct {
	// Data The data found by the scanner in the specific path for a specific type. See example for SSHKnownHostFingerprint info type
	Data *string `json:"data,omitempty"`

	// Path File path containing the info
	Path *string         `json:"path,omitempty"`
	Type *InfoFinderType `json:"type,omitempty"`
}

// InfoFinderType defines model for InfoFinderType.
type InfoFinderType string

// Malware defines model for Malware.
type Malware struct {
	MalwareName *string `json:"malwareName,omitempty"`
	MalwareType *string `json:"malwareType,omitempty"`

	// Path Path of the file that contains malware
	Path     *string `json:"path,omitempty"`
	RuleName *string `json:"ruleName,omitempty"`
}

// Metadata Describes the scanner plugin.
type Metadata struct {
	// Annotations Generic map of string keys and string values to attach arbitrary non-identifying metadata to objects.
	Annotations *Annotations `json:"annotations,omitempty"`

	// ApiVersion This value will be automatically set by the SDK.
	ApiVersion *string `json:"apiVersion,omitempty"`
	Name       *string `json:"name,omitempty"`
	Version    *string `json:"version,omitempty"`
}

// Misconfiguration defines model for Misconfiguration.
type Misconfiguration struct {
	// Category Specifies misconfiguration impact category
	Category *string `json:"category,omitempty"`

	// Description Additional context such as the potential impact
	Description *string `json:"description,omitempty"`

	// Id Check or test ID, if applicable (e.g. Lynis TestID, CIS Docker Benchmark checkpoint code, etc)
	Id *string `json:"id,omitempty"`

	// Location Location within the asset where the misconfiguration was recorded (e.g. filesystem path)
	Location *string `json:"location,omitempty"`

	// Message Short info about the misconfiguration
	Message *string `json:"message,omitempty"`

	// Remediation Possible fix for the misconfiguration
	Remediation *string                   `json:"remediation,omitempty"`
	Severity    *MisconfigurationSeverity `json:"severity,omitempty"`
}

// MisconfigurationSeverity defines model for MisconfigurationSeverity.
type MisconfigurationSeverity string

// Package defines model for Package.
type Package struct {
	Cpes     *[]string `json:"cpes"`
	Language *string   `json:"language,omitempty"`
	Licenses *[]string `json:"licenses"`
	Name     *string   `json:"name,omitempty"`
	Purl     *string   `json:"purl,omitempty"`
	Type     *string   `json:"type,omitempty"`
	Version  *string   `json:"version,omitempty"`
}

// Result Describes data saved to a JSON file when a scan finishes successfully.
type Result struct {
	// Annotations Generic map of string keys and string values to attach arbitrary non-identifying metadata to objects.
	Annotations *Annotations `json:"annotations,omitempty"`

	// RawJSON Defines scan result data that is not consumed by VMClarity API.
	RawJSON interface{} `json:"rawJSON"`

	// RawSarif Defines scan result data in that is not consumed by the VMClarity API.
	RawSarif *interface{} `json:"rawSarif,omitempty"`

	// Vmclarity Defines scan result data that can be consumed by VMClarity API.
	Vmclarity VMClarityData `json:"vmclarity"`
}

// Rootkit defines model for Rootkit.
type Rootkit struct {
	Message     *string      `json:"message,omitempty"`
	RootkitName *string      `json:"rootkitName,omitempty"`
	RootkitType *RootkitType `json:"rootkitType,omitempty"`
}

// RootkitType defines model for RootkitType.
type RootkitType string

// Secret defines model for Secret.
type Secret struct {
	Description *string `json:"description,omitempty"`
	EndColumn   *int    `json:"endColumn,omitempty"`
	EndLine     *int    `json:"endLine,omitempty"`

	// FilePath Name of the file containing the secret
	FilePath *string `json:"filePath,omitempty"`

	// Fingerprint Note: this is not unique
	Fingerprint *string `json:"fingerprint,omitempty"`
	StartColumn *int    `json:"startColumn,omitempty"`
	StartLine   *int    `json:"startLine,omitempty"`
}

// State Describes the status of scanner.
// | Status         | Description                                                   |
// | -------------- | ------------------------------------------------------------- |
// | NotReady       | Initial state when the scanner container starts               |
// | Ready          | Scanner setup is complete and it is ready to receive requests |
// | Running        | Scanner config was received and the scanner is running        |
// | Failed         | Scanner failed                                                |
// | Done           | Scanner completed successfully                                |
type State string

// Status Describes the scanner status.
type Status struct {
	// LastTransitionTime Last date time when the status has changed.
	LastTransitionTime time.Time `json:"lastTransitionTime"`

	// Message Human readable message.
	Message *string `json:"message,omitempty"`

	// State Describes the status of scanner.
	// | Status         | Description                                                   |
	// | -------------- | ------------------------------------------------------------- |
	// | NotReady       | Initial state when the scanner container starts               |
	// | Ready          | Scanner setup is complete and it is ready to receive requests |
	// | Running        | Scanner config was received and the scanner is running        |
	// | Failed         | Scanner failed                                                |
	// | Done           | Scanner completed successfully                                |
	State State `json:"state"`
}

// Stop Describes data for scanner to stop the scanning process.
type Stop struct {
	// TimeoutSeconds After this timeout the server will be stopped.
	TimeoutSeconds int `json:"timeoutSeconds" validate:"required,gt=0"`
}

// VMClarityData Defines scan result data that can be consumed by VMClarity API.
type VMClarityData struct {
	Exploits          *[]Exploit          `json:"exploits,omitempty"`
	InfoFinder        *[]InfoFinder       `json:"infoFinder,omitempty"`
	Malware           *[]Malware          `json:"malware,omitempty"`
	Misconfigurations *[]Misconfiguration `json:"misconfigurations,omitempty"`
	Packages          *[]Package          `json:"packages,omitempty"`
	Rootkits          *[]Rootkit          `json:"rootkits,omitempty"`
	Secrets           *[]Secret           `json:"secrets,omitempty"`
	Vulnerabilities   *[]Vulnerability    `json:"vulnerabilities,omitempty"`
}

// Vulnerability defines model for Vulnerability.
type Vulnerability struct {
	Cvss        *[]VulnerabilityCvss `json:"cvss"`
	Description *string              `json:"description,omitempty"`

	// Distro Distro provides information about a detected Linux distribution.
	Distro            *VulnerabilityDistro   `json:"distro,omitempty"`
	Fix               *VulnerabilityFix      `json:"fix,omitempty"`
	LayerId           *string                `json:"layerId,omitempty"`
	Links             *[]string              `json:"links"`
	Package           *Package               `json:"package,omitempty"`
	Path              *string                `json:"path,omitempty"`
	Severity          *VulnerabilitySeverity `json:"severity,omitempty"`
	VulnerabilityName *string                `json:"vulnerabilityName,omitempty"`
}

// VulnerabilityCvss defines model for VulnerabilityCvss.
type VulnerabilityCvss struct {
	BaseScore           *float32 `json:"baseScore,omitempty"`
	ExploitabilityScore *float32 `json:"exploitabilityScore,omitempty"`
	ImpactScore         *float32 `json:"impactScore,omitempty"`
	Vector              *string  `json:"vector,omitempty"`
	Version             *string  `json:"version,omitempty"`
}

// VulnerabilityDistro Distro provides information about a detected Linux distribution.
type VulnerabilityDistro struct {
	// IDLike the ID_LIKE field found within the /etc/os-release file
	IDLike *[]string `json:"IDLike"`

	// Name Name of the Linux distribution
	Name *string `json:"name,omitempty"`

	// Version Version of the Linux distribution (major or major.minor version)
	Version *string `json:"version,omitempty"`
}

// VulnerabilityFix defines model for VulnerabilityFix.
type VulnerabilityFix struct {
	State    *string   `json:"state,omitempty"`
	Versions *[]string `json:"versions"`
}

// VulnerabilitySeverity defines model for VulnerabilitySeverity.
type VulnerabilitySeverity string

// PostConfigJSONRequestBody defines body for PostConfig for application/json ContentType.
type PostConfigJSONRequestBody = Config

// PostStopJSONRequestBody defines body for PostStop for application/json ContentType.
type PostStopJSONRequestBody = Stop
