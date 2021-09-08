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

// GetTerraformResourceType returns Terraform resource type for this SpringCloudActiveDeployment
func (mg *SpringCloudActiveDeployment) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_active_deployment"
}

// GetTerraformResourceIdField returns Terraform identifier field for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
