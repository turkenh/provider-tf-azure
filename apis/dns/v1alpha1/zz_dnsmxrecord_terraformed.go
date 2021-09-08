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

// GetTerraformResourceType returns Terraform resource type for this DnsMxRecord
func (mg *DnsMxRecord) GetTerraformResourceType() string {
	return "azurerm_dns_mx_record"
}

// GetTerraformResourceIdField returns Terraform identifier field for this DnsMxRecord
func (tr *DnsMxRecord) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this DnsMxRecord
func (tr *DnsMxRecord) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this DnsMxRecord
func (tr *DnsMxRecord) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this DnsMxRecord
func (tr *DnsMxRecord) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this DnsMxRecord
func (tr *DnsMxRecord) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}

// GetForProvider of this DnsMxRecord
func (tr *DnsMxRecord) GetForProvider() interface{} {
	return &tr.Spec.ForProvider
}
