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

import (
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this KubernetesCluster
func (mg *KubernetesCluster) GetTerraformResourceType() string {
	return "azurerm_kubernetes_cluster"
}

// GetTerraformResourceIDField returns Terraform identifier field for this KubernetesCluster
func (tr *KubernetesCluster) GetTerraformResourceIDField() string {
	return "id"
}

// GetConnectionDetailsMapping for this KubernetesCluster
func (tr *KubernetesCluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"kube_admin_config[*].client_certificate": "status.atProvider.kubeAdminConfig[*].clientCertificate", "kube_admin_config[*].client_key": "status.atProvider.kubeAdminConfig[*].clientKey", "kube_admin_config[*].cluster_ca_certificate": "status.atProvider.kubeAdminConfig[*].clusterCaCertificate", "kube_admin_config[*].password": "status.atProvider.kubeAdminConfig[*].password", "kube_admin_config_raw": "status.atProvider.kubeAdminConfigRaw", "kube_config[*].client_certificate": "status.atProvider.kubeConfig[*].clientCertificate", "kube_config[*].client_key": "status.atProvider.kubeConfig[*].clientKey", "kube_config[*].cluster_ca_certificate": "status.atProvider.kubeConfig[*].clusterCaCertificate", "kube_config[*].password": "status.atProvider.kubeConfig[*].password", "kube_config_raw": "status.atProvider.kubeConfigRaw", "role_based_access_control[*].azure_active_directory[*].server_app_secret": "spec.forProvider.roleBasedAccessControl[*].azureActiveDirectory[*].serverAppSecretSecretRef", "service_principal[*].client_secret": "spec.forProvider.servicePrincipal[*].clientSecretSecretRef", "windows_profile[*].admin_password": "spec.forProvider.windowsProfile[*].adminPasswordSecretRef"}
}

// GetObservation of this KubernetesCluster
func (tr *KubernetesCluster) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this KubernetesCluster
func (tr *KubernetesCluster) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this KubernetesCluster
func (tr *KubernetesCluster) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this KubernetesCluster
func (tr *KubernetesCluster) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this KubernetesCluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *KubernetesCluster) LateInitialize(attrs []byte) (bool, error) {
	params := &KubernetesClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("KubeletIdentity"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}