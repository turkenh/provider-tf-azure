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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CosmosdbSqlStoredProcedureObservation struct {
}

type CosmosdbSqlStoredProcedureParameters struct {

	// +crossplane:generate:reference:type=CosmosdbAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body"`

	// +crossplane:generate:reference:type=CosmosdbSqlContainer
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`

	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=CosmosdbSqlDatabase
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`

	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// CosmosdbSqlStoredProcedureSpec defines the desired state of CosmosdbSqlStoredProcedure
type CosmosdbSqlStoredProcedureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CosmosdbSqlStoredProcedureParameters `json:"forProvider"`
}

// CosmosdbSqlStoredProcedureStatus defines the observed state of CosmosdbSqlStoredProcedure.
type CosmosdbSqlStoredProcedureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CosmosdbSqlStoredProcedureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbSqlStoredProcedure is the Schema for the CosmosdbSqlStoredProcedures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CosmosdbSqlStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbSqlStoredProcedureSpec   `json:"spec"`
	Status            CosmosdbSqlStoredProcedureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbSqlStoredProcedureList contains a list of CosmosdbSqlStoredProcedures
type CosmosdbSqlStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosdbSqlStoredProcedure `json:"items"`
}

// Repository type metadata.
var (
	CosmosdbSqlStoredProcedureKind             = "CosmosdbSqlStoredProcedure"
	CosmosdbSqlStoredProcedureGroupKind        = schema.GroupKind{Group: Group, Kind: CosmosdbSqlStoredProcedureKind}.String()
	CosmosdbSqlStoredProcedureKindAPIVersion   = CosmosdbSqlStoredProcedureKind + "." + GroupVersion.String()
	CosmosdbSqlStoredProcedureGroupVersionKind = GroupVersion.WithKind(CosmosdbSqlStoredProcedureKind)
)

func init() {
	SchemeBuilder.Register(&CosmosdbSqlStoredProcedure{}, &CosmosdbSqlStoredProcedureList{})
}
