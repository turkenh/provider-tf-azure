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

// GetTerraformResourceType returns Terraform resource type for this RouteFilter
func (mg *RouteFilter) GetTerraformResourceType() string {
	return "azurerm_route_filter"
}

// GetTerraformResourceIdField returns Terraform identifier field for this RouteFilter
func (tr *RouteFilter) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this RouteFilter
func (tr *RouteFilter) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this RouteFilter
func (tr *RouteFilter) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this RouteFilter
func (tr *RouteFilter) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this RouteFilter
func (tr *RouteFilter) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this RouteFilter
func (tr *RouteFilter) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
