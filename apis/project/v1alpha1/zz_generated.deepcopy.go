//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdministratorRolesInitParameters) DeepCopyInto(out *AdministratorRolesInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdministratorRolesInitParameters.
func (in *AdministratorRolesInitParameters) DeepCopy() *AdministratorRolesInitParameters {
	if in == nil {
		return nil
	}
	out := new(AdministratorRolesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdministratorRolesObservation) DeepCopyInto(out *AdministratorRolesObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdministratorRolesObservation.
func (in *AdministratorRolesObservation) DeepCopy() *AdministratorRolesObservation {
	if in == nil {
		return nil
	}
	out := new(AdministratorRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdministratorRolesParameters) DeepCopyInto(out *AdministratorRolesParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdministratorRolesParameters.
func (in *AdministratorRolesParameters) DeepCopy() *AdministratorRolesParameters {
	if in == nil {
		return nil
	}
	out := new(AdministratorRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsInitParameters) DeepCopyInto(out *ConstraintsInitParameters) {
	*out = *in
	if in.Extensibility != nil {
		in, out := &in.Extensibility, &out.Extensibility
		*out = make([]ExtensibilityInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsInitParameters.
func (in *ConstraintsInitParameters) DeepCopy() *ConstraintsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsObservation) DeepCopyInto(out *ConstraintsObservation) {
	*out = *in
	if in.Extensibility != nil {
		in, out := &in.Extensibility, &out.Extensibility
		*out = make([]ExtensibilityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsObservation.
func (in *ConstraintsObservation) DeepCopy() *ConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(ConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsParameters) DeepCopyInto(out *ConstraintsParameters) {
	*out = *in
	if in.Extensibility != nil {
		in, out := &in.Extensibility, &out.Extensibility
		*out = make([]ExtensibilityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]StorageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsParameters.
func (in *ConstraintsParameters) DeepCopy() *ConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensibilityInitParameters) DeepCopyInto(out *ExtensibilityInitParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensibilityInitParameters.
func (in *ExtensibilityInitParameters) DeepCopy() *ExtensibilityInitParameters {
	if in == nil {
		return nil
	}
	out := new(ExtensibilityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensibilityObservation) DeepCopyInto(out *ExtensibilityObservation) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensibilityObservation.
func (in *ExtensibilityObservation) DeepCopy() *ExtensibilityObservation {
	if in == nil {
		return nil
	}
	out := new(ExtensibilityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensibilityParameters) DeepCopyInto(out *ExtensibilityParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensibilityParameters.
func (in *ExtensibilityParameters) DeepCopy() *ExtensibilityParameters {
	if in == nil {
		return nil
	}
	out := new(ExtensibilityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberRolesInitParameters) DeepCopyInto(out *MemberRolesInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberRolesInitParameters.
func (in *MemberRolesInitParameters) DeepCopy() *MemberRolesInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemberRolesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberRolesObservation) DeepCopyInto(out *MemberRolesObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberRolesObservation.
func (in *MemberRolesObservation) DeepCopy() *MemberRolesObservation {
	if in == nil {
		return nil
	}
	out := new(MemberRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberRolesParameters) DeepCopyInto(out *MemberRolesParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberRolesParameters.
func (in *MemberRolesParameters) DeepCopy() *MemberRolesParameters {
	if in == nil {
		return nil
	}
	out := new(MemberRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInitParameters) DeepCopyInto(out *NetworkInitParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInitParameters.
func (in *NetworkInitParameters) DeepCopy() *NetworkInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectInitParameters) DeepCopyInto(out *ProjectInitParameters) {
	*out = *in
	if in.AdministratorRoles != nil {
		in, out := &in.AdministratorRoles, &out.AdministratorRoles
		*out = make([]AdministratorRolesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Administrators != nil {
		in, out := &in.Administrators, &out.Administrators
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.MachineNamingTemplate != nil {
		in, out := &in.MachineNamingTemplate, &out.MachineNamingTemplate
		*out = new(string)
		**out = **in
	}
	if in.MemberRoles != nil {
		in, out := &in.MemberRoles, &out.MemberRoles
		*out = make([]MemberRolesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OperationTimeout != nil {
		in, out := &in.OperationTimeout, &out.OperationTimeout
		*out = new(float64)
		**out = **in
	}
	if in.PlacementPolicy != nil {
		in, out := &in.PlacementPolicy, &out.PlacementPolicy
		*out = new(string)
		**out = **in
	}
	if in.SharedResources != nil {
		in, out := &in.SharedResources, &out.SharedResources
		*out = new(bool)
		**out = **in
	}
	if in.SupervisorRoles != nil {
		in, out := &in.SupervisorRoles, &out.SupervisorRoles
		*out = make([]SupervisorRolesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ViewerRoles != nil {
		in, out := &in.ViewerRoles, &out.ViewerRoles
		*out = make([]ViewerRolesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Viewers != nil {
		in, out := &in.Viewers, &out.Viewers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ZoneAssignments != nil {
		in, out := &in.ZoneAssignments, &out.ZoneAssignments
		*out = make([]ZoneAssignmentsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectInitParameters.
func (in *ProjectInitParameters) DeepCopy() *ProjectInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.AdministratorRoles != nil {
		in, out := &in.AdministratorRoles, &out.AdministratorRoles
		*out = make([]AdministratorRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Administrators != nil {
		in, out := &in.Administrators, &out.Administrators
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.MachineNamingTemplate != nil {
		in, out := &in.MachineNamingTemplate, &out.MachineNamingTemplate
		*out = new(string)
		**out = **in
	}
	if in.MemberRoles != nil {
		in, out := &in.MemberRoles, &out.MemberRoles
		*out = make([]MemberRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OperationTimeout != nil {
		in, out := &in.OperationTimeout, &out.OperationTimeout
		*out = new(float64)
		**out = **in
	}
	if in.PlacementPolicy != nil {
		in, out := &in.PlacementPolicy, &out.PlacementPolicy
		*out = new(string)
		**out = **in
	}
	if in.SharedResources != nil {
		in, out := &in.SharedResources, &out.SharedResources
		*out = new(bool)
		**out = **in
	}
	if in.SupervisorRoles != nil {
		in, out := &in.SupervisorRoles, &out.SupervisorRoles
		*out = make([]SupervisorRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ViewerRoles != nil {
		in, out := &in.ViewerRoles, &out.ViewerRoles
		*out = make([]ViewerRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Viewers != nil {
		in, out := &in.Viewers, &out.Viewers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ZoneAssignments != nil {
		in, out := &in.ZoneAssignments, &out.ZoneAssignments
		*out = make([]ZoneAssignmentsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.AdministratorRoles != nil {
		in, out := &in.AdministratorRoles, &out.AdministratorRoles
		*out = make([]AdministratorRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Administrators != nil {
		in, out := &in.Administrators, &out.Administrators
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.MachineNamingTemplate != nil {
		in, out := &in.MachineNamingTemplate, &out.MachineNamingTemplate
		*out = new(string)
		**out = **in
	}
	if in.MemberRoles != nil {
		in, out := &in.MemberRoles, &out.MemberRoles
		*out = make([]MemberRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OperationTimeout != nil {
		in, out := &in.OperationTimeout, &out.OperationTimeout
		*out = new(float64)
		**out = **in
	}
	if in.PlacementPolicy != nil {
		in, out := &in.PlacementPolicy, &out.PlacementPolicy
		*out = new(string)
		**out = **in
	}
	if in.SharedResources != nil {
		in, out := &in.SharedResources, &out.SharedResources
		*out = new(bool)
		**out = **in
	}
	if in.SupervisorRoles != nil {
		in, out := &in.SupervisorRoles, &out.SupervisorRoles
		*out = make([]SupervisorRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ViewerRoles != nil {
		in, out := &in.ViewerRoles, &out.ViewerRoles
		*out = make([]ViewerRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Viewers != nil {
		in, out := &in.Viewers, &out.Viewers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ZoneAssignments != nil {
		in, out := &in.ZoneAssignments, &out.ZoneAssignments
		*out = make([]ZoneAssignmentsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageInitParameters) DeepCopyInto(out *StorageInitParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageInitParameters.
func (in *StorageInitParameters) DeepCopy() *StorageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupervisorRolesInitParameters) DeepCopyInto(out *SupervisorRolesInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupervisorRolesInitParameters.
func (in *SupervisorRolesInitParameters) DeepCopy() *SupervisorRolesInitParameters {
	if in == nil {
		return nil
	}
	out := new(SupervisorRolesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupervisorRolesObservation) DeepCopyInto(out *SupervisorRolesObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupervisorRolesObservation.
func (in *SupervisorRolesObservation) DeepCopy() *SupervisorRolesObservation {
	if in == nil {
		return nil
	}
	out := new(SupervisorRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupervisorRolesParameters) DeepCopyInto(out *SupervisorRolesParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupervisorRolesParameters.
func (in *SupervisorRolesParameters) DeepCopy() *SupervisorRolesParameters {
	if in == nil {
		return nil
	}
	out := new(SupervisorRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewerRolesInitParameters) DeepCopyInto(out *ViewerRolesInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewerRolesInitParameters.
func (in *ViewerRolesInitParameters) DeepCopy() *ViewerRolesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ViewerRolesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewerRolesObservation) DeepCopyInto(out *ViewerRolesObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewerRolesObservation.
func (in *ViewerRolesObservation) DeepCopy() *ViewerRolesObservation {
	if in == nil {
		return nil
	}
	out := new(ViewerRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewerRolesParameters) DeepCopyInto(out *ViewerRolesParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewerRolesParameters.
func (in *ViewerRolesParameters) DeepCopy() *ViewerRolesParameters {
	if in == nil {
		return nil
	}
	out := new(ViewerRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAssignmentsInitParameters) DeepCopyInto(out *ZoneAssignmentsInitParameters) {
	*out = *in
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.MaxInstances != nil {
		in, out := &in.MaxInstances, &out.MaxInstances
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimitMb != nil {
		in, out := &in.MemoryLimitMb, &out.MemoryLimitMb
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.StorageLimitGb != nil {
		in, out := &in.StorageLimitGb, &out.StorageLimitGb
		*out = new(float64)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAssignmentsInitParameters.
func (in *ZoneAssignmentsInitParameters) DeepCopy() *ZoneAssignmentsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ZoneAssignmentsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAssignmentsObservation) DeepCopyInto(out *ZoneAssignmentsObservation) {
	*out = *in
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.MaxInstances != nil {
		in, out := &in.MaxInstances, &out.MaxInstances
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimitMb != nil {
		in, out := &in.MemoryLimitMb, &out.MemoryLimitMb
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.StorageLimitGb != nil {
		in, out := &in.StorageLimitGb, &out.StorageLimitGb
		*out = new(float64)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAssignmentsObservation.
func (in *ZoneAssignmentsObservation) DeepCopy() *ZoneAssignmentsObservation {
	if in == nil {
		return nil
	}
	out := new(ZoneAssignmentsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneAssignmentsParameters) DeepCopyInto(out *ZoneAssignmentsParameters) {
	*out = *in
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.MaxInstances != nil {
		in, out := &in.MaxInstances, &out.MaxInstances
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimitMb != nil {
		in, out := &in.MemoryLimitMb, &out.MemoryLimitMb
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.StorageLimitGb != nil {
		in, out := &in.StorageLimitGb, &out.StorageLimitGb
		*out = new(float64)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneAssignmentsParameters.
func (in *ZoneAssignmentsParameters) DeepCopy() *ZoneAssignmentsParameters {
	if in == nil {
		return nil
	}
	out := new(ZoneAssignmentsParameters)
	in.DeepCopyInto(out)
	return out
}