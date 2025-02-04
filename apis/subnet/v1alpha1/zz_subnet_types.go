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

type DelegationObservation struct {
}

type DelegationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceDelegation []ServiceDelegationParameters `json:"serviceDelegation" tf:"service_delegation,omitempty"`
}

type ServiceDelegationObservation struct {
}

type ServiceDelegationParameters struct {

	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type SubnetObservation struct {
}

type SubnetParameters struct {

	// +kubebuilder:validation:Optional
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	Delegation []DelegationParameters `json:"delegation,omitempty" tf:"delegation,omitempty"`

	// +kubebuilder:validation:Optional
	EnforcePrivateLinkEndpointNetworkPolicies *bool `json:"enforcePrivateLinkEndpointNetworkPolicies,omitempty" tf:"enforce_private_link_endpoint_network_policies,omitempty"`

	// +kubebuilder:validation:Optional
	EnforcePrivateLinkServiceNetworkPolicies *bool `json:"enforcePrivateLinkServiceNetworkPolicies,omitempty" tf:"enforce_private_link_service_network_policies,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceEndpointPolicyIds []*string `json:"serviceEndpointPolicyIds,omitempty" tf:"service_endpoint_policy_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceEndpoints []*string `json:"serviceEndpoints,omitempty" tf:"service_endpoints,omitempty"`

	// +kubebuilder:validation:Required
	VirtualNetworkName *string `json:"virtualNetworkName" tf:"virtual_network_name,omitempty"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	SubnetKind             = "Subnet"
	SubnetGroupKind        = schema.GroupKind{Group: Group, Kind: SubnetKind}.String()
	SubnetKindAPIVersion   = SubnetKind + "." + GroupVersion.String()
	SubnetGroupVersionKind = GroupVersion.WithKind(SubnetKind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
