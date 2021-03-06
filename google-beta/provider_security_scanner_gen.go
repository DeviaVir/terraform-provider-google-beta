// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "github.com/hashicorp/terraform/helper/schema"

// If the base path has changed as a result of your PR, make sure to update
// the provider_reference page!
var SecurityScannerDefaultBasePath = "https://websecurityscanner.googleapis.com/v1beta/"
var SecurityScannerCustomEndpointEntryKey = "security_scanner_custom_endpoint"
var SecurityScannerCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_SECURITY_SCANNER_CUSTOM_ENDPOINT",
	}, SecurityScannerDefaultBasePath),
}

var GeneratedSecurityScannerResourcesMap = map[string]*schema.Resource{
	"google_security_scanner_scan_config": resourceSecurityScannerScanConfig(),
}
