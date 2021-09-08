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

// GetTerraformResourceType returns Terraform resource type for this LogAnalyticsSavedSearch
func (mg *LogAnalyticsSavedSearch) GetTerraformResourceType() string {
	return "azurerm_log_analytics_saved_search"
}

// GetTerraformResourceIdField returns Terraform identifier field for this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this LogAnalyticsSavedSearch
func (tr *LogAnalyticsSavedSearch) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
