/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedStorageAccountObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagedStorageAccountParameters struct {

	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RegenerateKeyAutomatically *bool `json:"regenerateKeyAutomatically,omitempty" tf:"regenerate_key_automatically,omitempty"`

	// +kubebuilder:validation:Optional
	RegenerationPeriod *string `json:"regenerationPeriod,omitempty" tf:"regeneration_period,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/storage/v1alpha2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	StorageAccountKey *string `json:"storageAccountKey" tf:"storage_account_key,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ManagedStorageAccountSpec defines the desired state of ManagedStorageAccount
type ManagedStorageAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedStorageAccountParameters `json:"forProvider"`
}

// ManagedStorageAccountStatus defines the observed state of ManagedStorageAccount.
type ManagedStorageAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedStorageAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedStorageAccount is the Schema for the ManagedStorageAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagedStorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedStorageAccountSpec   `json:"spec"`
	Status            ManagedStorageAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedStorageAccountList contains a list of ManagedStorageAccounts
type ManagedStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedStorageAccount `json:"items"`
}

// Repository type metadata.
var (
	ManagedStorageAccount_Kind             = "ManagedStorageAccount"
	ManagedStorageAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedStorageAccount_Kind}.String()
	ManagedStorageAccount_KindAPIVersion   = ManagedStorageAccount_Kind + "." + CRDGroupVersion.String()
	ManagedStorageAccount_GroupVersionKind = CRDGroupVersion.WithKind(ManagedStorageAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedStorageAccount{}, &ManagedStorageAccountList{})
}
