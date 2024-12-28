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

type DeploymentInitParameters struct {

	// The content of the the cloud template to be used to request the deployment. Conflicts with blueprint_id and catalog_item_id.
	// The content of the the cloud template to be used to request the deployment.
	BlueprintContent *string `json:"blueprintContent,omitempty" tf:"blueprint_content,omitempty"`

	// The id of the cloud template to be used to request the deployment. Conflicts with blueprint_content and catalog_item_id.
	// The id of the cloud template to be used to request the deployment.
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The version of the cloud template to be used to request the deployment. Used only when blueprint_id is provided.
	// The version of the cloud template to be used to request the deployment.
	BlueprintVersion *string `json:"blueprintVersion,omitempty" tf:"blueprint_version,omitempty"`

	// The id of the catalog item to be used to request the deployment. Conflicts with blueprint_id and blueprint_content.
	// The id of the catalog item to be used to request the deployment.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// The version of the catalog item to be used to request the deployment. Used only when catalog_item_id is provided.
	// The version of the catalog item to be used to request the deployment.
	CatalogItemVersion *string `json:"catalogItemVersion,omitempty" tf:"catalog_item_version,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Represents deployment requests.
	ExpandLastRequest *bool `json:"expandLastRequest,omitempty" tf:"expand_last_request,omitempty"`

	// Flag to indicate whether to expand project information.
	// Flag to indicate whether to expand project information.
	ExpandProject *bool `json:"expandProject,omitempty" tf:"expand_project,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
	ExpandResources *bool `json:"expandResources,omitempty" tf:"expand_resources,omitempty"`

	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// +mapType=granular
	Inputs map[string]*string `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// The name of the deployment.
	// The name of the deployment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
	// The user this deployment belongs to.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The id of the project this deployment belongs to.
	// The id of the project this deployment belongs to.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Reason for requesting/updating a blueprint.
	// Reason for requesting/updating a blueprint.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type DeploymentObservation struct {

	// The content of the the cloud template to be used to request the deployment. Conflicts with blueprint_id and catalog_item_id.
	// The content of the the cloud template to be used to request the deployment.
	BlueprintContent *string `json:"blueprintContent,omitempty" tf:"blueprint_content,omitempty"`

	// The id of the cloud template to be used to request the deployment. Conflicts with blueprint_content and catalog_item_id.
	// The id of the cloud template to be used to request the deployment.
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The version of the cloud template to be used to request the deployment. Used only when blueprint_id is provided.
	// The version of the cloud template to be used to request the deployment.
	BlueprintVersion *string `json:"blueprintVersion,omitempty" tf:"blueprint_version,omitempty"`

	// The id of the catalog item to be used to request the deployment. Conflicts with blueprint_id and blueprint_content.
	// The id of the catalog item to be used to request the deployment.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// The version of the catalog item to be used to request the deployment. Used only when catalog_item_id is provided.
	// The version of the catalog item to be used to request the deployment.
	CatalogItemVersion *string `json:"catalogItemVersion,omitempty" tf:"catalog_item_version,omitempty"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The user the entity was created by.
	// The user the entity was created by.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Represents deployment requests.
	ExpandLastRequest *bool `json:"expandLastRequest,omitempty" tf:"expand_last_request,omitempty"`

	// Flag to indicate whether to expand project information.
	// Flag to indicate whether to expand project information.
	ExpandProject *bool `json:"expandProject,omitempty" tf:"expand_project,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
	ExpandResources *bool `json:"expandResources,omitempty" tf:"expand_resources,omitempty"`

	// Expense incurred for the deployment.
	Expense []ExpenseObservation `json:"expense,omitempty" tf:"expense,omitempty"`

	// The id of the deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// +mapType=granular
	Inputs map[string]*string `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// All the inputs applied during last create/update operation, including those with default values. For the list of inputs provided by the user in the configuration, refer to inputs.
	// All the inputs applied during last create/update operation, including those with default values.
	// +mapType=granular
	InputsIncludingDefaults map[string]*string `json:"inputsIncludingDefaults,omitempty" tf:"inputs_including_defaults,omitempty"`

	// Represents deployment requests.
	LastRequest []LastRequestObservation `json:"lastRequest,omitempty" tf:"last_request,omitempty"`

	// TDate when the entity was last updated. The date is in ISO 6801 and UTC.
	// Date when the entity was last updated. The date is in ISO 6801 and UTC.
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" tf:"last_updated_at,omitempty"`

	// The user that last updated the deployment.
	// The user that last updated the deployment.
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty" tf:"last_updated_by,omitempty"`

	// Date when the deployment lease expire. The date is in ISO 6801 and UTC.
	// Date when the deployment lease expire. The date is in ISO 6801 and UTC.
	LeaseExpireAt *string `json:"leaseExpireAt,omitempty" tf:"lease_expire_at,omitempty"`

	// The name of the deployment.
	// The name of the deployment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Id of the organization this deployment belongs to.
	// The Id of the organization this deployment belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
	// The user this deployment belongs to.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The project this entity belongs to.
	Project []ProjectObservation `json:"project,omitempty" tf:"project,omitempty"`

	// The id of the project this deployment belongs to.
	// The id of the project this deployment belongs to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reason for requesting/updating a blueprint.
	// Reason for requesting/updating a blueprint.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Request overall execution status. Supported values: CREATED, PENDING, INITIALIZATION, CHECKING_APPROVAL, APPROVAL_PENDING, INPROGRESS, COMPLETION, APPROVAL_REJECTED, ABORTED, SUCCESSFUL, FAILED.
	// The status of the deployment with respect to its life cycle operations.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DeploymentParameters struct {

	// The content of the the cloud template to be used to request the deployment. Conflicts with blueprint_id and catalog_item_id.
	// The content of the the cloud template to be used to request the deployment.
	// +kubebuilder:validation:Optional
	BlueprintContent *string `json:"blueprintContent,omitempty" tf:"blueprint_content,omitempty"`

	// The id of the cloud template to be used to request the deployment. Conflicts with blueprint_content and catalog_item_id.
	// The id of the cloud template to be used to request the deployment.
	// +kubebuilder:validation:Optional
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// The version of the cloud template to be used to request the deployment. Used only when blueprint_id is provided.
	// The version of the cloud template to be used to request the deployment.
	// +kubebuilder:validation:Optional
	BlueprintVersion *string `json:"blueprintVersion,omitempty" tf:"blueprint_version,omitempty"`

	// The id of the catalog item to be used to request the deployment. Conflicts with blueprint_id and blueprint_content.
	// The id of the catalog item to be used to request the deployment.
	// +kubebuilder:validation:Optional
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// The version of the catalog item to be used to request the deployment. Used only when catalog_item_id is provided.
	// The version of the catalog item to be used to request the deployment.
	// +kubebuilder:validation:Optional
	CatalogItemVersion *string `json:"catalogItemVersion,omitempty" tf:"catalog_item_version,omitempty"`

	// A human-friendly description.
	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Represents deployment requests.
	// +kubebuilder:validation:Optional
	ExpandLastRequest *bool `json:"expandLastRequest,omitempty" tf:"expand_last_request,omitempty"`

	// Flag to indicate whether to expand project information.
	// Flag to indicate whether to expand project information.
	// +kubebuilder:validation:Optional
	ExpandProject *bool `json:"expandProject,omitempty" tf:"expand_project,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible.
	// +kubebuilder:validation:Optional
	ExpandResources *bool `json:"expandResources,omitempty" tf:"expand_resources,omitempty"`

	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Inputs map[string]*string `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// The name of the deployment.
	// The name of the deployment.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.
	// The user this deployment belongs to.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The id of the project this deployment belongs to.
	// The id of the project this deployment belongs to.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Reason for requesting/updating a blueprint.
	// Reason for requesting/updating a blueprint.
	// +kubebuilder:validation:Optional
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type ExpenseInitParameters struct {
}

type ExpenseObservation struct {

	// Additional expense incurred for the resource.
	AdditionalExpense *float64 `json:"additionalExpense,omitempty" tf:"additional_expense,omitempty"`

	// Expense sync message code if any.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Compute expense of the entity.
	ComputeExpense *float64 `json:"computeExpense,omitempty" tf:"compute_expense,omitempty"`

	// Last expense sync time.
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" tf:"last_update_time,omitempty"`

	// Expense sync message if any.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Network expense of the entity.
	NetworkExpense *float64 `json:"networkExpense,omitempty" tf:"network_expense,omitempty"`

	// Storage expense of the entity.
	StorageExpense *float64 `json:"storageExpense,omitempty" tf:"storage_expense,omitempty"`

	// Total expense of the entity.
	TotalExpense *float64 `json:"totalExpense,omitempty" tf:"total_expense,omitempty"`

	// Monetary unit.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type ExpenseParameters struct {
}

type LastRequestInitParameters struct {
}

type LastRequestObservation struct {

	// Identifier of the requested action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Time at which the request was approved.
	ApprovedAt *string `json:"approvedAt,omitempty" tf:"approved_at,omitempty"`

	// The id of the cloud template to be used to request the deployment. Conflicts with blueprint_content and catalog_item_id.
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// Indicates whether request can be canceled or not.
	Cancelable *bool `json:"cancelable,omitempty" tf:"cancelable,omitempty"`

	// The id of the catalog item to be used to request the deployment. Conflicts with blueprint_id and blueprint_content.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// Time at which the request completed.
	CompletedAt *string `json:"completedAt,omitempty" tf:"completed_at,omitempty"`

	// The number of tasks completed while fulfilling this request.
	CompletedTasks *float64 `json:"completedTasks,omitempty" tf:"completed_tasks,omitempty"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Longer user-friendly details of the request.
	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	// Indicates whether request is in dismissed state.
	Dismissed *bool `json:"dismissed,omitempty" tf:"dismissed,omitempty"`

	// The id of the deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time at which the request was initialized.
	InitializedAt *string `json:"initializedAt,omitempty" tf:"initialized_at,omitempty"`

	// Inputs provided by the user. For inputs including those with default values, refer to inputs_including_defaults.
	// +mapType=granular
	Inputs map[string]*string `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// The name of the deployment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Request outputs.
	// +mapType=granular
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// The user that initiated the request.
	RequestedBy *string `json:"requestedBy,omitempty" tf:"requested_by,omitempty"`

	// +listType=set
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// Request overall execution status. Supported values: CREATED, PENDING, INITIALIZATION, CHECKING_APPROVAL, APPROVAL_PENDING, INPROGRESS, COMPLETION, APPROVAL_REJECTED, ABORTED, SUCCESSFUL, FAILED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The total number of tasks need to be completed to fulfil this request.
	TotalTasks *float64 `json:"totalTasks,omitempty" tf:"total_tasks,omitempty"`

	// Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type LastRequestParameters struct {
}

type ProjectInitParameters struct {
}

type ProjectObservation struct {

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the deployment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Version of the entity, if applicable.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectParameters struct {
}

type ResourcesExpenseInitParameters struct {
}

type ResourcesExpenseObservation struct {

	// Additional expense incurred for the resource.
	AdditionalExpense *float64 `json:"additionalExpense,omitempty" tf:"additional_expense,omitempty"`

	// Expense sync message code if any.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Compute expense of the entity.
	ComputeExpense *float64 `json:"computeExpense,omitempty" tf:"compute_expense,omitempty"`

	// Last expense sync time.
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" tf:"last_update_time,omitempty"`

	// Expense sync message if any.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Network expense of the entity.
	NetworkExpense *float64 `json:"networkExpense,omitempty" tf:"network_expense,omitempty"`

	// Storage expense of the entity.
	StorageExpense *float64 `json:"storageExpense,omitempty" tf:"storage_expense,omitempty"`

	// Total expense of the entity.
	TotalExpense *float64 `json:"totalExpense,omitempty" tf:"total_expense,omitempty"`

	// Monetary unit.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type ResourcesExpenseParameters struct {
}

type ResourcesInitParameters struct {
}

type ResourcesObservation struct {

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A list of other resources this resource depends on.
	// +listType=set
	DependsOn []*string `json:"dependsOn,omitempty" tf:"depends_on,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Expense incurred for the deployment.
	Expense []ResourcesExpenseObservation `json:"expense,omitempty" tf:"expense,omitempty"`

	// The id of the deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the deployment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of properties in the encoded JSON string format.
	PropertiesJSON *string `json:"propertiesJson,omitempty" tf:"properties_json,omitempty"`

	// The current state of the resource. Supported values are PARTIAL, TAINTED, OK.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The current sync status. Supported values are SUCCESS, MISSING, STALE.
	SyncStatus *string `json:"syncStatus,omitempty" tf:"sync_status,omitempty"`

	// Type of the resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourcesParameters struct {
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
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
	InitProvider DeploymentInitParameters `json:"initProvider,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Deployment is the Schema for the Deployments API. A resource that can be used to create a vRealize Automation deployment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DeploymentSpec   `json:"spec"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
