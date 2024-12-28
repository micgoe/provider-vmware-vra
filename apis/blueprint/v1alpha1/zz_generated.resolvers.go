// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/micgoe/provider-vmware-vra/apis/project/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Blueprint.
func (mg *Blueprint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectIDRef,
		Selector:     mg.Spec.ForProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectID")
	}
	mg.Spec.ForProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectIDRef,
		Selector:     mg.Spec.InitProvider.ProjectIDSelector,
		To: reference.To{
			List:    &v1alpha1.ProjectList{},
			Managed: &v1alpha1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProjectID")
	}
	mg.Spec.InitProvider.ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Version.
func (mg *Version) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BlueprintID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BlueprintIDRef,
		Selector:     mg.Spec.ForProvider.BlueprintIDSelector,
		To: reference.To{
			List:    &BlueprintList{},
			Managed: &Blueprint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BlueprintID")
	}
	mg.Spec.ForProvider.BlueprintID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BlueprintIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BlueprintID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BlueprintIDRef,
		Selector:     mg.Spec.InitProvider.BlueprintIDSelector,
		To: reference.To{
			List:    &BlueprintList{},
			Managed: &Blueprint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BlueprintID")
	}
	mg.Spec.InitProvider.BlueprintID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BlueprintIDRef = rsp.ResolvedReference

	return nil
}