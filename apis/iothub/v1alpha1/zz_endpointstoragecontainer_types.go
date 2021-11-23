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

type EndpointStorageContainerObservation struct {
}

type EndpointStorageContainerParameters struct {

	// +kubebuilder:validation:Optional
	BatchFrequencyInSeconds *int64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Optional
	MaxChunkSizeInBytes *int64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// EndpointStorageContainerSpec defines the desired state of EndpointStorageContainer
type EndpointStorageContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointStorageContainerParameters `json:"forProvider"`
}

// EndpointStorageContainerStatus defines the observed state of EndpointStorageContainer.
type EndpointStorageContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointStorageContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointStorageContainer is the Schema for the EndpointStorageContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EndpointStorageContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointStorageContainerSpec   `json:"spec"`
	Status            EndpointStorageContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointStorageContainerList contains a list of EndpointStorageContainers
type EndpointStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointStorageContainer `json:"items"`
}

// Repository type metadata.
var (
	EndpointStorageContainer_Kind             = "EndpointStorageContainer"
	EndpointStorageContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointStorageContainer_Kind}.String()
	EndpointStorageContainer_KindAPIVersion   = EndpointStorageContainer_Kind + "." + CRDGroupVersion.String()
	EndpointStorageContainer_GroupVersionKind = CRDGroupVersion.WithKind(EndpointStorageContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointStorageContainer{}, &EndpointStorageContainerList{})
}
