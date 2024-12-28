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

type ProfileInitParameters struct {

	// Indicates if this storage profile is a default profile.
	// Indicates if this storage profile is a default profile.
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Map of storage properties that are to be applied on disk while provisioning.
	// +mapType=granular
	DiskProperties map[string]*string `json:"diskProperties,omitempty" tf:"disk_properties,omitempty"`

	// Map of storage placements to know where the disk is provisioned.
	// Map of storage placements to know where the disk is provisioned.
	// +mapType=granular
	DiskTargetProperties map[string]*string `json:"diskTargetProperties,omitempty" tf:"disk_target_properties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for storage profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The id of the region that is associated with the storage profile.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// A set of tag keys and optional values that were set on this Network Profile.
	// example: [ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []TagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProfileObservation struct {

	// Id of the cloud account this storage profile belongs to.
	// Id of the cloud account this storage profile belongs to.
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Indicates if this storage profile is a default profile.
	// Indicates if this storage profile is a default profile.
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Map of storage properties that are to be applied on disk while provisioning.
	// +mapType=granular
	DiskProperties map[string]*string `json:"diskProperties,omitempty" tf:"disk_properties,omitempty"`

	// Map of storage placements to know where the disk is provisioned.
	// Map of storage placements to know where the disk is provisioned.
	// +mapType=granular
	DiskTargetProperties map[string]*string `json:"diskTargetProperties,omitempty" tf:"disk_target_properties,omitempty"`

	// The id of the region as seen in the cloud provider for which this profile is defined.
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// HATEOAS of the entity
	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for storage profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The id of the region that is associated with the storage profile.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// A set of tag keys and optional values that were set on this Network Profile.
	// example: [ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ProfileParameters struct {

	// Indicates if this storage profile is a default profile.
	// Indicates if this storage profile is a default profile.
	// +kubebuilder:validation:Optional
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Map of storage properties that are to be applied on disk while provisioning.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DiskProperties map[string]*string `json:"diskProperties,omitempty" tf:"disk_properties,omitempty"`

	// Map of storage placements to know where the disk is provisioned.
	// Map of storage placements to know where the disk is provisioned.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DiskTargetProperties map[string]*string `json:"diskTargetProperties,omitempty" tf:"disk_target_properties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for storage profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The id of the region that is associated with the storage profile.
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Indicates whether this storage profile supports encryption or not.
	// +kubebuilder:validation:Optional
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// A set of tag keys and optional values that were set on this Network Profile.
	// example: [ { "key" : "ownedBy", "value": "Rainpole" } ]
	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
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

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileParameters `json:"forProvider"`
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
	InitProvider ProfileInitParameters `json:"initProvider,omitempty"`
}

// ProfileStatus defines the observed state of Profile.
type ProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Profile is the Schema for the Profiles API. Provides a data lookup for vra_storage_profile.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultItem) || (has(self.initProvider) && has(self.initProvider.defaultItem))",message="spec.forProvider.defaultItem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionId) || (has(self.initProvider) && has(self.initProvider.regionId))",message="spec.forProvider.regionId is a required parameter"
	Spec   ProfileSpec   `json:"spec"`
	Status ProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileList contains a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Repository type metadata.
var (
	Profile_Kind             = "Profile"
	Profile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Profile_Kind}.String()
	Profile_KindAPIVersion   = Profile_Kind + "." + CRDGroupVersion.String()
	Profile_GroupVersionKind = CRDGroupVersion.WithKind(Profile_Kind)
)

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}