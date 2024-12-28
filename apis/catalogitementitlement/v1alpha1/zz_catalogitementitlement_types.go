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

type CatalogItemEntitlementInitParameters struct {

	// The id of the catalog item to create the entitlement.
	// Catalog item id.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type CatalogItemEntitlementObservation struct {

	// The id of the catalog item to create the entitlement.
	// Catalog item id.
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// Represents a catalog item that is linked to a project via an entitlement.
	Definition []DefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	// Id of the catalog item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type CatalogItemEntitlementParameters struct {

	// The id of the catalog item to create the entitlement.
	// Catalog item id.
	// +kubebuilder:validation:Optional
	CatalogItemID *string `json:"catalogItemId,omitempty" tf:"catalog_item_id,omitempty"`

	// The id of the project this entity belongs to.
	// Project id.
	// +crossplane:generate:reference:type=github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type DefinitionInitParameters struct {
}

type DefinitionObservation struct {

	// Description of the catalog item.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Id of the catalog item.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Icon id of associated catalog item.
	IconID *string `json:"iconId,omitempty" tf:"icon_id,omitempty"`

	// Name of the catalog item.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of items in the associated catalog source.
	NumberOfItems *float64 `json:"numberOfItems,omitempty" tf:"number_of_items,omitempty"`

	// Catalog source name.
	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`

	// Catalog source type.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Content definition type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefinitionParameters struct {
}

// CatalogItemEntitlementSpec defines the desired state of CatalogItemEntitlement
type CatalogItemEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogItemEntitlementParameters `json:"forProvider"`
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
	InitProvider CatalogItemEntitlementInitParameters `json:"initProvider,omitempty"`
}

// CatalogItemEntitlementStatus defines the observed state of CatalogItemEntitlement.
type CatalogItemEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogItemEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CatalogItemEntitlement is the Schema for the CatalogItemEntitlements API. A resource that can be used to create a vRealize Automation catalog item entitlement.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vmware-vra}
type CatalogItemEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogItemId) || (has(self.initProvider) && has(self.initProvider.catalogItemId))",message="spec.forProvider.catalogItemId is a required parameter"
	Spec   CatalogItemEntitlementSpec   `json:"spec"`
	Status CatalogItemEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogItemEntitlementList contains a list of CatalogItemEntitlements
type CatalogItemEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogItemEntitlement `json:"items"`
}

// Repository type metadata.
var (
	CatalogItemEntitlement_Kind             = "CatalogItemEntitlement"
	CatalogItemEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogItemEntitlement_Kind}.String()
	CatalogItemEntitlement_KindAPIVersion   = CatalogItemEntitlement_Kind + "." + CRDGroupVersion.String()
	CatalogItemEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(CatalogItemEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogItemEntitlement{}, &CatalogItemEntitlementList{})
}