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

// GetTerraformResourceType returns Terraform resource type for this SqlActiveDirectoryAdministrator
func (mg *SqlActiveDirectoryAdministrator) GetTerraformResourceType() string {
	return "azurerm_sql_active_directory_administrator"
}

// GetTerraformResourceIdField returns Terraform identifier field for this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this SqlActiveDirectoryAdministrator
func (tr *SqlActiveDirectoryAdministrator) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
