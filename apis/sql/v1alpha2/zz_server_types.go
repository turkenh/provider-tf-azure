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

type ServerExtendedAuditingPolicyObservation struct {
}

type ServerExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
}

type ServerIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ServerIdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ServerObservation struct {
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty" tf:"fully_qualified_domain_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServerParameters struct {

	// +kubebuilder:validation:Required
	AdministratorLogin *string `json:"administratorLogin" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Required
	AdministratorLoginPasswordSecretRef v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionPolicy *string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendedAuditingPolicy []ServerExtendedAuditingPolicyParameters `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []ServerIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThreatDetectionPolicy []ThreatDetectionPolicyParameters `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ThreatDetectionPolicyObservation struct {
}

type ThreatDetectionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
