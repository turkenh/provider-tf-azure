/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/conversion"

// GetTerraformResourceType returns Terraform resource type for this ApiManagementApiDiagnostic
func (mg *ApiManagementApiDiagnostic) GetTerraformResourceType() string {
	return "azurerm_api_management_api_diagnostic"
}

// GetTerraformResourceIdField returns Terraform identifier field for this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this ApiManagementApiDiagnostic
func (tr *ApiManagementApiDiagnostic) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
