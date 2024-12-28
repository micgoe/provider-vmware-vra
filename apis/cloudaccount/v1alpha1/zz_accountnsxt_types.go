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

type AccountNsxtInitParameters struct {

	// Accept self-signed certificate when connecting to the cloud account.
	// Accept self signed certificate when connecting.
	AcceptSelfSignedCert *bool `json:"acceptSelfSignedCert,omitempty" tf:"accept_self_signed_cert,omitempty"`

	// Identifier of a data collector VM deployed in the on premise infrastructure.
	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Host name for NSX-T cloud account.
	// Host name for the NSX-T endpoint.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode. Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode *bool `json:"managerMode,omitempty" tf:"manager_mode,omitempty"`

	// Name of NSX-T cloud account.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Password used to authenticate to the cloud Account.
	// Password for the user used to authenticate with the cloud Account.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of tag keys and values to apply to the cloud account.
	// Example: [ { "key" : "vmware", "value": "provider" } ]
	Tags []AccountNsxtTagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// Username used to authenticate to the cloud account.
	// Username to authenticate with the cloud account.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AccountNsxtLinksInitParameters struct {
}

type AccountNsxtLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type AccountNsxtLinksParameters struct {
}

type AccountNsxtObservation struct {

	// Accept self-signed certificate when connecting to the cloud account.
	// Accept self signed certificate when connecting.
	AcceptSelfSignedCert *bool `json:"acceptSelfSignedCert,omitempty" tf:"accept_self_signed_cert,omitempty"`

	// Cloud accounts associated with the cloud account.
	// +listType=set
	AssociatedCloudAccountIds []*string `json:"associatedCloudAccountIds,omitempty" tf:"associated_cloud_account_ids,omitempty"`

	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Identifier of a data collector VM deployed in the on premise infrastructure.
	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Host name for NSX-T cloud account.
	// Host name for the NSX-T endpoint.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// ID of NSX-T cloud account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// HATEOAS of entity.
	Links []AccountNsxtLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode. Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode *bool `json:"managerMode,omitempty" tf:"manager_mode,omitempty"`

	// Name of NSX-T cloud account.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of organization that entity belongs to.
	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of entity owner.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Set of tag keys and values to apply to the cloud account.
	// Example: [ { "key" : "vmware", "value": "provider" } ]
	Tags []AccountNsxtTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// Username used to authenticate to the cloud account.
	// Username to authenticate with the cloud account.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AccountNsxtParameters struct {

	// Accept self-signed certificate when connecting to the cloud account.
	// Accept self signed certificate when connecting.
	// +kubebuilder:validation:Optional
	AcceptSelfSignedCert *bool `json:"acceptSelfSignedCert,omitempty" tf:"accept_self_signed_cert,omitempty"`

	// Identifier of a data collector VM deployed in the on premise infrastructure.
	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// +kubebuilder:validation:Optional
	DcID *string `json:"dcId,omitempty" tf:"dc_id,omitempty"`

	// Human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Host name for NSX-T cloud account.
	// Host name for the NSX-T endpoint.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode. Mode cannot be changed after cloud account is created. Default value is false.
	// +kubebuilder:validation:Optional
	ManagerMode *bool `json:"managerMode,omitempty" tf:"manager_mode,omitempty"`

	// Name of NSX-T cloud account.
	// A human-friendly name used as an identifier in APIs that support this option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Password used to authenticate to the cloud Account.
	// Password for the user used to authenticate with the cloud Account.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of tag keys and values to apply to the cloud account.
	// Example: [ { "key" : "vmware", "value": "provider" } ]
	// +kubebuilder:validation:Optional
	Tags []AccountNsxtTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// Username used to authenticate to the cloud account.
	// Username to authenticate with the cloud account.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AccountNsxtTagsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountNsxtTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AccountNsxtTagsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AccountNsxtSpec defines the desired state of AccountNsxt
type AccountNsxtSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountNsxtParameters `json:"forProvider"`
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
	InitProvider AccountNsxtInitParameters `json:"initProvider,omitempty"`
}

// AccountNsxtStatus defines the observed state of AccountNsxt.
type AccountNsxtStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountNsxtObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccountNsxt is the Schema for the AccountNsxts API. Creates a vra_cloud_account_nsxt resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type AccountNsxt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostname) || (has(self.initProvider) && has(self.initProvider.hostname))",message="spec.forProvider.hostname is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   AccountNsxtSpec   `json:"spec"`
	Status AccountNsxtStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNsxtList contains a list of AccountNsxts
type AccountNsxtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountNsxt `json:"items"`
}

// Repository type metadata.
var (
	AccountNsxt_Kind             = "AccountNsxt"
	AccountNsxt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountNsxt_Kind}.String()
	AccountNsxt_KindAPIVersion   = AccountNsxt_Kind + "." + CRDGroupVersion.String()
	AccountNsxt_GroupVersionKind = CRDGroupVersion.WithKind(AccountNsxt_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountNsxt{}, &AccountNsxtList{})
}