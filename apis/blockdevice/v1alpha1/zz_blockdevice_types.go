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

type BlockDeviceInitParameters struct {

	// Capacity of block device in GB.
	// Capacity of the block device in GB.
	CapacityInGb *float64 `json:"capacityInGb,omitempty" tf:"capacity_in_gb,omitempty"`

	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsInitParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Additional custom properties that may be used to extend the machine.
	// Additional custom properties that may be used to extend the block device. Following disk custom properties can be passed while creating a block device:
	//
	// 1. dataStore: Defines name of the datastore in which the disk has to be provisioned.
	// 2. storagePolicy: Defines name of the storage policy in which the disk has to be provisioned. If name of the datastore is specified in the custom properties then, datastore takes precedence.
	// 3. provisioningType: Defines the type of provisioning. For eg. thick/thin.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// ID of deployment associated with resource.
	// The id of the deployment that is associated with this resource.
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Content of a disk, base64 encoded.
	// Content of a disk, base64 encoded.
	DiskContentBase64 *string `json:"diskContentBase64,omitempty" tf:"disk_content_base_64,omitempty"`

	// Indicates whether block device should be encrypted or not.
	// Indicates whether the block device should be encrypted or not.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	// Indicates whether the snapshots of the block-devices should be included in the resource state. Applicable only for first class block devices.
	ExpandSnapshots *bool `json:"expandSnapshots,omitempty" tf:"expand_snapshots,omitempty"`

	// Human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for the block device.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether block device survives a delete action.
	// Indicates whether the block device survives a delete action.
	Persistent *bool `json:"persistent,omitempty" tf:"persistent,omitempty"`

	// ID of project that current user belongs to.
	// The id of the project this resource belongs to.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	// Indicates if the disk has to be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true, only used for destroy the resource
	Purge *bool `json:"purge,omitempty" tf:"purge,omitempty"`

	// URI to use for block device. Example: ami-0d4cfd66
	// Reference to URI using which the block device has to be created. Example: ami-0d4cfd66
	SourceReference *string `json:"sourceReference,omitempty" tf:"source_reference,omitempty"`

	// Set of tag keys and values to apply to the resource instance.
	// Example: [ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]
	Tags []TagsInitParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BlockDeviceObservation struct {

	// Capacity of block device in GB.
	// Capacity of the block device in GB.
	CapacityInGb *float64 `json:"capacityInGb,omitempty" tf:"capacity_in_gb,omitempty"`

	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Additional custom properties that may be used to extend the machine.
	// Additional custom properties that may be used to extend the block device. Following disk custom properties can be passed while creating a block device:
	//
	// 1. dataStore: Defines name of the datastore in which the disk has to be provisioned.
	// 2. storagePolicy: Defines name of the storage policy in which the disk has to be provisioned. If name of the datastore is specified in the custom properties then, datastore takes precedence.
	// 3. provisioningType: Defines the type of provisioning. For eg. thick/thin.
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// ID of deployment associated with resource.
	// The id of the deployment that is associated with this resource.
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Content of a disk, base64 encoded.
	// Content of a disk, base64 encoded.
	DiskContentBase64 *string `json:"diskContentBase64,omitempty" tf:"disk_content_base_64,omitempty"`

	// Indicates whether block device should be encrypted or not.
	// Indicates whether the block device should be encrypted or not.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	// Indicates whether the snapshots of the block-devices should be included in the resource state. Applicable only for first class block devices.
	ExpandSnapshots *bool `json:"expandSnapshots,omitempty" tf:"expand_snapshots,omitempty"`

	// External entity ID on provider side.
	// External entity id on the cloud provider side.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// External regionId of resource.
	// The external regionId of the resource.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	// External zoneId of resource.
	// The external zoneId of the resource.
	ExternalZoneID *string `json:"externalZoneId,omitempty" tf:"external_zone_id,omitempty"`

	// ID of the block device snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// HATEOAS of entity.
	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// Human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for the block device.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of organization that entity belongs to.
	// The id of the organization this block device belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of entity owner.
	// Email of the user that owns this block device.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Indicates whether block device survives a delete action.
	// Indicates whether the block device survives a delete action.
	Persistent *bool `json:"persistent,omitempty" tf:"persistent,omitempty"`

	// ID of project that current user belongs to.
	// The id of the project this resource belongs to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	// Indicates if the disk has to be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true, only used for destroy the resource
	Purge *bool `json:"purge,omitempty" tf:"purge,omitempty"`

	// Represents a machine snapshot.
	Snapshots []SnapshotsObservation `json:"snapshots,omitempty" tf:"snapshots,omitempty"`

	// URI to use for block device. Example: ami-0d4cfd66
	// Reference to URI using which the block device has to be created. Example: ami-0d4cfd66
	SourceReference *string `json:"sourceReference,omitempty" tf:"source_reference,omitempty"`

	// Status of block device.
	// Status of the block device.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Set of tag keys and values to apply to the resource instance.
	// Example: [ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]
	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type BlockDeviceParameters struct {

	// Capacity of block device in GB.
	// Capacity of the block device in GB.
	// +kubebuilder:validation:Optional
	CapacityInGb *float64 `json:"capacityInGb,omitempty" tf:"capacity_in_gb,omitempty"`

	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Additional custom properties that may be used to extend the machine.
	// Additional custom properties that may be used to extend the block device. Following disk custom properties can be passed while creating a block device:
	//
	// 1. dataStore: Defines name of the datastore in which the disk has to be provisioned.
	// 2. storagePolicy: Defines name of the storage policy in which the disk has to be provisioned. If name of the datastore is specified in the custom properties then, datastore takes precedence.
	// 3. provisioningType: Defines the type of provisioning. For eg. thick/thin.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// ID of deployment associated with resource.
	// The id of the deployment that is associated with this resource.
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Content of a disk, base64 encoded.
	// Content of a disk, base64 encoded.
	// +kubebuilder:validation:Optional
	DiskContentBase64 *string `json:"diskContentBase64,omitempty" tf:"disk_content_base_64,omitempty"`

	// Indicates whether block device should be encrypted or not.
	// Indicates whether the block device should be encrypted or not.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	// Indicates whether the snapshots of the block-devices should be included in the resource state. Applicable only for first class block devices.
	// +kubebuilder:validation:Optional
	ExpandSnapshots *bool `json:"expandSnapshots,omitempty" tf:"expand_snapshots,omitempty"`

	// Human-friendly name used as an identifier in APIs that support this option.
	// A human-friendly name for the block device.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether block device survives a delete action.
	// Indicates whether the block device survives a delete action.
	// +kubebuilder:validation:Optional
	Persistent *bool `json:"persistent,omitempty" tf:"persistent,omitempty"`

	// ID of project that current user belongs to.
	// The id of the project this resource belongs to.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	// Indicates if the disk has to be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true, only used for destroy the resource
	// +kubebuilder:validation:Optional
	Purge *bool `json:"purge,omitempty" tf:"purge,omitempty"`

	// URI to use for block device. Example: ami-0d4cfd66
	// Reference to URI using which the block device has to be created. Example: ami-0d4cfd66
	// +kubebuilder:validation:Optional
	SourceReference *string `json:"sourceReference,omitempty" tf:"source_reference,omitempty"`

	// Set of tag keys and values to apply to the resource instance.
	// Example: [ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]
	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

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

type SnapshotsInitParameters struct {
}

type SnapshotsLinksInitParameters struct {
}

type SnapshotsLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// +listType=set
	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type SnapshotsLinksParameters struct {
}

type SnapshotsObservation struct {

	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the block device snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether snapshot on block device is current.
	IsCurrent *bool `json:"isCurrent,omitempty" tf:"is_current,omitempty"`

	// HATEOAS of entity.
	Links []SnapshotsLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of organization that entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of entity owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SnapshotsParameters struct {
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

// BlockDeviceSpec defines the desired state of BlockDevice
type BlockDeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlockDeviceParameters `json:"forProvider"`
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
	InitProvider BlockDeviceInitParameters `json:"initProvider,omitempty"`
}

// BlockDeviceStatus defines the observed state of BlockDevice.
type BlockDeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlockDeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BlockDevice is the Schema for the BlockDevices API. Creates a vra_block_device resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type BlockDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.capacityInGb) || (has(self.initProvider) && has(self.initProvider.capacityInGb))",message="spec.forProvider.capacityInGb is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BlockDeviceSpec   `json:"spec"`
	Status BlockDeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlockDeviceList contains a list of BlockDevices
type BlockDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlockDevice `json:"items"`
}

// Repository type metadata.
var (
	BlockDevice_Kind             = "BlockDevice"
	BlockDevice_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BlockDevice_Kind}.String()
	BlockDevice_KindAPIVersion   = BlockDevice_Kind + "." + CRDGroupVersion.String()
	BlockDevice_GroupVersionKind = CRDGroupVersion.WithKind(BlockDevice_Kind)
)

func init() {
	SchemeBuilder.Register(&BlockDevice{}, &BlockDeviceList{})
}
