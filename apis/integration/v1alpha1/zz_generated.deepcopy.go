//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Integration) DeepCopyInto(out *Integration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Integration.
func (in *Integration) DeepCopy() *Integration {
	if in == nil {
		return nil
	}
	out := new(Integration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Integration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationInitParameters) DeepCopyInto(out *IntegrationInitParameters) {
	*out = *in
	if in.AssociatedCloudAccountIds != nil {
		in, out := &in.AssociatedCloudAccountIds, &out.AssociatedCloudAccountIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IntegrationProperties != nil {
		in, out := &in.IntegrationProperties, &out.IntegrationProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IntegrationType != nil {
		in, out := &in.IntegrationType, &out.IntegrationType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeyIDSecretRef != nil {
		in, out := &in.PrivateKeyIDSecretRef, &out.PrivateKeyIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationInitParameters.
func (in *IntegrationInitParameters) DeepCopy() *IntegrationInitParameters {
	if in == nil {
		return nil
	}
	out := new(IntegrationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationList) DeepCopyInto(out *IntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Integration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationList.
func (in *IntegrationList) DeepCopy() *IntegrationList {
	if in == nil {
		return nil
	}
	out := new(IntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationObservation) DeepCopyInto(out *IntegrationObservation) {
	*out = *in
	if in.AssociatedCloudAccountIds != nil {
		in, out := &in.AssociatedCloudAccountIds, &out.AssociatedCloudAccountIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IntegrationProperties != nil {
		in, out := &in.IntegrationProperties, &out.IntegrationProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IntegrationType != nil {
		in, out := &in.IntegrationType, &out.IntegrationType
		*out = new(string)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationObservation.
func (in *IntegrationObservation) DeepCopy() *IntegrationObservation {
	if in == nil {
		return nil
	}
	out := new(IntegrationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationParameters) DeepCopyInto(out *IntegrationParameters) {
	*out = *in
	if in.AssociatedCloudAccountIds != nil {
		in, out := &in.AssociatedCloudAccountIds, &out.AssociatedCloudAccountIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IntegrationProperties != nil {
		in, out := &in.IntegrationProperties, &out.IntegrationProperties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IntegrationType != nil {
		in, out := &in.IntegrationType, &out.IntegrationType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateKeyIDSecretRef != nil {
		in, out := &in.PrivateKeyIDSecretRef, &out.PrivateKeyIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationParameters.
func (in *IntegrationParameters) DeepCopy() *IntegrationParameters {
	if in == nil {
		return nil
	}
	out := new(IntegrationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationSpec) DeepCopyInto(out *IntegrationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationSpec.
func (in *IntegrationSpec) DeepCopy() *IntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(IntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationStatus) DeepCopyInto(out *IntegrationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationStatus.
func (in *IntegrationStatus) DeepCopy() *IntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(IntegrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksInitParameters) DeepCopyInto(out *LinksInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksInitParameters.
func (in *LinksInitParameters) DeepCopy() *LinksInitParameters {
	if in == nil {
		return nil
	}
	out := new(LinksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksObservation) DeepCopyInto(out *LinksObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
	if in.Hrefs != nil {
		in, out := &in.Hrefs, &out.Hrefs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Rel != nil {
		in, out := &in.Rel, &out.Rel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksObservation.
func (in *LinksObservation) DeepCopy() *LinksObservation {
	if in == nil {
		return nil
	}
	out := new(LinksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinksParameters) DeepCopyInto(out *LinksParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinksParameters.
func (in *LinksParameters) DeepCopy() *LinksParameters {
	if in == nil {
		return nil
	}
	out := new(LinksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsInitParameters) DeepCopyInto(out *TagsInitParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsInitParameters.
func (in *TagsInitParameters) DeepCopy() *TagsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TagsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsObservation) DeepCopyInto(out *TagsObservation) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsObservation.
func (in *TagsObservation) DeepCopy() *TagsObservation {
	if in == nil {
		return nil
	}
	out := new(TagsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsParameters) DeepCopyInto(out *TagsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsParameters.
func (in *TagsParameters) DeepCopy() *TagsParameters {
	if in == nil {
		return nil
	}
	out := new(TagsParameters)
	in.DeepCopyInto(out)
	return out
}