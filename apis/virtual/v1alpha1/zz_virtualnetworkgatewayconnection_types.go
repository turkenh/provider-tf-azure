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

type IpsecPolicyObservation struct {
}

type IpsecPolicyParameters struct {

	// +kubebuilder:validation:Required
	DhGroup *string `json:"dhGroup" tf:"dh_group"`

	// +kubebuilder:validation:Required
	IkeEncryption *string `json:"ikeEncryption" tf:"ike_encryption"`

	// +kubebuilder:validation:Required
	IkeIntegrity *string `json:"ikeIntegrity" tf:"ike_integrity"`

	// +kubebuilder:validation:Required
	IpsecEncryption *string `json:"ipsecEncryption" tf:"ipsec_encryption"`

	// +kubebuilder:validation:Required
	IpsecIntegrity *string `json:"ipsecIntegrity" tf:"ipsec_integrity"`

	// +kubebuilder:validation:Required
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group"`

	// +kubebuilder:validation:Optional
	SaDatasize *int64 `json:"saDatasize,omitempty" tf:"sa_datasize"`

	// +kubebuilder:validation:Optional
	SaLifetime *int64 `json:"saLifetime,omitempty" tf:"sa_lifetime"`
}

type TrafficSelectorPolicyObservation struct {
}

type TrafficSelectorPolicyParameters struct {

	// +kubebuilder:validation:Required
	LocalAddressCidrs []*string `json:"localAddressCidrs" tf:"local_address_cidrs"`

	// +kubebuilder:validation:Required
	RemoteAddressCidrs []*string `json:"remoteAddressCidrs" tf:"remote_address_cidrs"`
}

type VirtualNetworkGatewayConnectionObservation struct {
}

type VirtualNetworkGatewayConnectionParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionProtocol *string `json:"connectionProtocol,omitempty" tf:"connection_protocol"`

	// +kubebuilder:validation:Optional
	DpdTimeoutSeconds *int64 `json:"dpdTimeoutSeconds,omitempty" tf:"dpd_timeout_seconds"`

	// +kubebuilder:validation:Optional
	EnableBgp *bool `json:"enableBgp,omitempty" tf:"enable_bgp"`

	// +kubebuilder:validation:Optional
	ExpressRouteCircuitID *string `json:"expressRouteCircuitId,omitempty" tf:"express_route_circuit_id"`

	// +kubebuilder:validation:Optional
	ExpressRouteGatewayBypass *bool `json:"expressRouteGatewayBypass,omitempty" tf:"express_route_gateway_bypass"`

	// +kubebuilder:validation:Optional
	IpsecPolicy []IpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy"`

	// +kubebuilder:validation:Optional
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled"`

	// +kubebuilder:validation:Optional
	LocalNetworkGatewayID *string `json:"localNetworkGatewayId,omitempty" tf:"local_network_gateway_id"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayID *string `json:"peerVirtualNetworkGatewayId,omitempty" tf:"peer_virtual_network_gateway_id"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDRef *v1.Reference `json:"peerVirtualNetworkGatewayIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDSelector *v1.Selector `json:"peerVirtualNetworkGatewayIDSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoutingWeight *int64 `json:"routingWeight,omitempty" tf:"routing_weight"`

	// +kubebuilder:validation:Optional
	SharedKeySecretRef v1.SecretKeySelector `json:"sharedKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TrafficSelectorPolicy []TrafficSelectorPolicyParameters `json:"trafficSelectorPolicy,omitempty" tf:"traffic_selector_policy"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	UsePolicyBasedTrafficSelectors *bool `json:"usePolicyBasedTrafficSelectors,omitempty" tf:"use_policy_based_traffic_selectors"`

	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayID *string `json:"virtualNetworkGatewayId,omitempty" tf:"virtual_network_gateway_id"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDRef *v1.Reference `json:"virtualNetworkGatewayIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDSelector *v1.Selector `json:"virtualNetworkGatewayIDSelector,omitempty" tf:"-"`
}

// VirtualNetworkGatewayConnectionSpec defines the desired state of VirtualNetworkGatewayConnection
type VirtualNetworkGatewayConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayConnectionParameters `json:"forProvider"`
}

// VirtualNetworkGatewayConnectionStatus defines the observed state of VirtualNetworkGatewayConnection.
type VirtualNetworkGatewayConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnection is the Schema for the VirtualNetworkGatewayConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewayConnectionSpec   `json:"spec"`
	Status            VirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnectionList contains a list of VirtualNetworkGatewayConnections
type VirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGatewayConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGatewayConnectionKind             = "VirtualNetworkGatewayConnection"
	VirtualNetworkGatewayConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkGatewayConnectionKind}.String()
	VirtualNetworkGatewayConnectionKindAPIVersion   = VirtualNetworkGatewayConnectionKind + "." + GroupVersion.String()
	VirtualNetworkGatewayConnectionGroupVersionKind = GroupVersion.WithKind(VirtualNetworkGatewayConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGatewayConnection{}, &VirtualNetworkGatewayConnectionList{})
}
