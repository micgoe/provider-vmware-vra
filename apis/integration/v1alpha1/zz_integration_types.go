// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IntegrationInitParameters struct {

	// Ids of the cloud accounts to associate with this integration.
	// Ids of the cloud accounts to associate with this integration.
	// +listType=set
	AssociatedCloudAccountIds []*string `json:"associatedCloudAccountIds,omitempty" tf:"associated_cloud_account_ids,omitempty"`

	// Certificate to be used to connect to the integration.
	// Certificate to be used to connect to the integration.
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// Additional custom properties that may be used to extend the Integration.
	// Additional custom properties that may be used to extend the Integration.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Integration specific properties supplied in as name value pairs.
	// Integration specific properties supplied in as name value pairs.
	// +mapType=granular
	IntegrationProperties map[string]*string `json:"integrationProperties,omitempty" tf:"integration_properties,omitempty"`

	// Integration type.
	// Integration type.
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// The name of the integration.
	// The name of the integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Access key id or username to be used to authenticate with the integration.
	// Access key id or username to be used to authenticate with the integration.
	PrivateKeyIDSecretRef *v1.SecretKeySelector `json:"privateKeyIdSecretRef,omitempty" tf:"-"`

	// Secret access key or password to be used to authenticate with the integration.
	// Secret access key or password to be used to authenticate with the integration.
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// A set of tag keys and optional values to apply to the integration. Example: [ { "key" : "provider", "value": "vmware" } ].
	Tags []TagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IntegrationObservation struct {

	// Ids of the cloud accounts to associate with this integration.
	// Ids of the cloud accounts to associate with this integration.
	// +listType=set
	AssociatedCloudAccountIds []*string `json:"associatedCloudAccountIds,omitempty" tf:"associated_cloud_account_ids,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Additional custom properties that may be used to extend the Integration.
	// Additional custom properties that may be used to extend the Integration.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the integration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Integration specific properties supplied in as name value pairs.
	// Integration specific properties supplied in as name value pairs.
	// +mapType=granular
	IntegrationProperties map[string]*string `json:"integrationProperties,omitempty" tf:"integration_properties,omitempty"`

	// Integration type.
	// Integration type.
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// HATEOAS of entity.
	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// The name of the integration.
	// The name of the integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// A set of tag keys and optional values to apply to the integration. Example: [ { "key" : "provider", "value": "vmware" } ].
	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type IntegrationParameters struct {

	// Ids of the cloud accounts to associate with this integration.
	// Ids of the cloud accounts to associate with this integration.
	// +kubebuilder:validation:Optional
	// +listType=set
	AssociatedCloudAccountIds []*string `json:"associatedCloudAccountIds,omitempty" tf:"associated_cloud_account_ids,omitempty"`

	// Certificate to be used to connect to the integration.
	// Certificate to be used to connect to the integration.
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// Additional custom properties that may be used to extend the Integration.
	// Additional custom properties that may be used to extend the Integration.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Integration specific properties supplied in as name value pairs.
	// Integration specific properties supplied in as name value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	IntegrationProperties map[string]*string `json:"integrationProperties,omitempty" tf:"integration_properties,omitempty"`

	// Integration type.
	// Integration type.
	// +kubebuilder:validation:Optional
	IntegrationType *string `json:"integrationType,omitempty" tf:"integration_type,omitempty"`

	// The name of the integration.
	// The name of the integration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Access key id or username to be used to authenticate with the integration.
	// Access key id or username to be used to authenticate with the integration.
	// +kubebuilder:validation:Optional
	PrivateKeyIDSecretRef *v1.SecretKeySelector `json:"privateKeyIdSecretRef,omitempty" tf:"-"`

	// Secret access key or password to be used to authenticate with the integration.
	// Secret access key or password to be used to authenticate with the integration.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// A set of tag keys and optional values to apply to the integration. Example: [ { "key" : "provider", "value": "vmware" } ].
	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LinksInitParameters struct {
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

type TagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IntegrationInitParameters `json:"initProvider,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Integration is the Schema for the Integrations API. Creates a vra_integration resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationProperties) || (has(self.initProvider) && has(self.initProvider.integrationProperties))",message="spec.forProvider.integrationProperties is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationType) || (has(self.initProvider) && has(self.initProvider.integrationType))",message="spec.forProvider.integrationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IntegrationSpec   `json:"spec"`
	Status IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
