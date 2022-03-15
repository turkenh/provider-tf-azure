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

type MongoCollectionAutoscaleSettingsObservation struct {
}

type MongoCollectionAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoCollectionIndexObservation struct {
}

type MongoCollectionIndexParameters struct {

	// +kubebuilder:validation:Required
	Keys []*string `json:"keys" tf:"keys,omitempty"`

	// +kubebuilder:validation:Optional
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type MongoCollectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SystemIndexes []SystemIndexesObservation `json:"systemIndexes,omitempty" tf:"system_indexes,omitempty"`
}

type MongoCollectionParameters struct {

	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *int64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []MongoCollectionAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +crossplane:generate:reference:type=MongoDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultTTLSeconds *int64 `json:"defaultTtlSeconds,omitempty" tf:"default_ttl_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Index []MongoCollectionIndexParameters `json:"index,omitempty" tf:"index,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ShardKey *string `json:"shardKey,omitempty" tf:"shard_key,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type SystemIndexesObservation struct {
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`
}

type SystemIndexesParameters struct {
}

// MongoCollectionSpec defines the desired state of MongoCollection
type MongoCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoCollectionParameters `json:"forProvider"`
}

// MongoCollectionStatus defines the observed state of MongoCollection.
type MongoCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollection is the Schema for the MongoCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MongoCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoCollectionSpec   `json:"spec"`
	Status            MongoCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoCollectionList contains a list of MongoCollections
type MongoCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoCollection `json:"items"`
}

// Repository type metadata.
var (
	MongoCollection_Kind             = "MongoCollection"
	MongoCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongoCollection_Kind}.String()
	MongoCollection_KindAPIVersion   = MongoCollection_Kind + "." + CRDGroupVersion.String()
	MongoCollection_GroupVersionKind = CRDGroupVersion.WithKind(MongoCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&MongoCollection{}, &MongoCollectionList{})
}
