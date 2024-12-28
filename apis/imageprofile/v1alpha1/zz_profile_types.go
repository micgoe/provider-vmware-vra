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

type ConstraintsInitParameters struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`
}

type ConstraintsObservation struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`
}

type ConstraintsParameters struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	// +kubebuilder:validation:Optional
	Mandatory *bool `json:"mandatory" tf:"mandatory,omitempty"`
}

type ImageMappingInitParameters struct {

	// Cloud config for this image. This cloud config will be merged during provisioning with other cloud configurations such as the bootConfig provided in MachineSpecification or vRA cloud templates.
	CloudConfig *string `json:"cloudConfig,omitempty" tf:"cloud_config,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsInitParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// The id of this resource instance.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly image name as seen on the cloud provider side.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name of the image mapping.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ImageMappingObservation struct {

	// Cloud config for this image. This cloud config will be merged during provisioning with other cloud configurations such as the bootConfig provided in MachineSpecification or vRA cloud templates.
	CloudConfig *string `json:"cloudConfig,omitempty" tf:"cloud_config,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External entity id on the cloud provider side.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// The external regionId of the resource.
	// External region id on the cloud provider side.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	// The id of this resource instance.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly image name as seen on the cloud provider side.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name of the image mapping.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A human-friendly description.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Operating system family of the image.
	OsFamily *string `json:"osFamily,omitempty" tf:"os_family,omitempty"`

	// Email of the user that owns the entity.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Indicates whether this fabric image is private. For vSphere, private images are considered to be templates and snapshots and public are Content Library items.
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`
}

type ImageMappingParameters struct {

	// Cloud config for this image. This cloud config will be merged during provisioning with other cloud configurations such as the bootConfig provided in MachineSpecification or vRA cloud templates.
	// +kubebuilder:validation:Optional
	CloudConfig *string `json:"cloudConfig,omitempty" tf:"cloud_config,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// The id of this resource instance.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly image name as seen on the cloud provider side.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name of the image mapping.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ProfileInitParameters struct {

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Image mapping defined for the corresponding region.
	ImageMapping []ImageMappingInitParameters `json:"imageMapping,omitempty" tf:"image_mapping,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The if of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`
}

type ProfileObservation struct {

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The external regionId of the resource.
	// The id of the region for which this profile is defined as in the cloud provider.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Image mapping defined for the corresponding region.
	ImageMapping []ImageMappingObservation `json:"imageMapping,omitempty" tf:"image_mapping,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email of the user that owns the entity.
	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The if of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ProfileParameters struct {

	// A human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Image mapping defined for the corresponding region.
	// +kubebuilder:validation:Optional
	ImageMapping []ImageMappingParameters `json:"imageMapping,omitempty" tf:"image_mapping,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name used as an identifier in APIs that support this option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	// The if of the region for which this profile is defined as in vRealize Automation(vRA).
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`
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

// Profile is the Schema for the Profiles API. Provides a data lookup for vra_image_profile.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
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