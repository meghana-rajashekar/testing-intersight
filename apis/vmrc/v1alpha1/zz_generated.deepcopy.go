//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AncestorsObservation) DeepCopyInto(out *AncestorsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AncestorsObservation.
func (in *AncestorsObservation) DeepCopy() *AncestorsObservation {
	if in == nil {
		return nil
	}
	out := new(AncestorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AncestorsParameters) DeepCopyInto(out *AncestorsParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AncestorsParameters.
func (in *AncestorsParameters) DeepCopy() *AncestorsParameters {
	if in == nil {
		return nil
	}
	out := new(AncestorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Console) DeepCopyInto(out *Console) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Console.
func (in *Console) DeepCopy() *Console {
	if in == nil {
		return nil
	}
	out := new(Console)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Console) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleList) DeepCopyInto(out *ConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Console, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleList.
func (in *ConsoleList) DeepCopy() *ConsoleList {
	if in == nil {
		return nil
	}
	out := new(ConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleObservation) DeepCopyInto(out *ConsoleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleObservation.
func (in *ConsoleObservation) DeepCopy() *ConsoleObservation {
	if in == nil {
		return nil
	}
	out := new(ConsoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleParameters) DeepCopyInto(out *ConsoleParameters) {
	*out = *in
	if in.AccountMoid != nil {
		in, out := &in.AccountMoid, &out.AccountMoid
		*out = new(string)
		**out = **in
	}
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.Ancestors != nil {
		in, out := &in.Ancestors, &out.Ancestors
		*out = make([]AncestorsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.ClientIPAddress != nil {
		in, out := &in.ClientIPAddress, &out.ClientIPAddress
		*out = new(string)
		**out = **in
	}
	if in.ClientURL != nil {
		in, out := &in.ClientURL, &out.ClientURL
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DomainGroupMoid != nil {
		in, out := &in.DomainGroupMoid, &out.DomainGroupMoid
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ModTime != nil {
		in, out := &in.ModTime, &out.ModTime
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = make([]ParentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PermissionResources != nil {
		in, out := &in.PermissionResources, &out.PermissionResources
		*out = make([]PermissionResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = make([]SessionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SharedScope != nil {
		in, out := &in.SharedScope, &out.SharedScope
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetName != nil {
		in, out := &in.TargetName, &out.TargetName
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = make([]UserParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserIDOrEmail != nil {
		in, out := &in.UserIDOrEmail, &out.UserIDOrEmail
		*out = new(string)
		**out = **in
	}
	if in.Vcenter != nil {
		in, out := &in.Vcenter, &out.Vcenter
		*out = make([]VcenterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VersionContext != nil {
		in, out := &in.VersionContext, &out.VersionContext
		*out = make([]VersionContextParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualMachine != nil {
		in, out := &in.VirtualMachine, &out.VirtualMachine
		*out = make([]VirtualMachineParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleParameters.
func (in *ConsoleParameters) DeepCopy() *ConsoleParameters {
	if in == nil {
		return nil
	}
	out := new(ConsoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSpec) DeepCopyInto(out *ConsoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSpec.
func (in *ConsoleSpec) DeepCopy() *ConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleStatus) DeepCopyInto(out *ConsoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleStatus.
func (in *ConsoleStatus) DeepCopy() *ConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(ConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterestedMosObservation) DeepCopyInto(out *InterestedMosObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterestedMosObservation.
func (in *InterestedMosObservation) DeepCopy() *InterestedMosObservation {
	if in == nil {
		return nil
	}
	out := new(InterestedMosObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterestedMosParameters) DeepCopyInto(out *InterestedMosParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterestedMosParameters.
func (in *InterestedMosParameters) DeepCopy() *InterestedMosParameters {
	if in == nil {
		return nil
	}
	out := new(InterestedMosParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParentObservation) DeepCopyInto(out *ParentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParentObservation.
func (in *ParentObservation) DeepCopy() *ParentObservation {
	if in == nil {
		return nil
	}
	out := new(ParentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParentParameters) DeepCopyInto(out *ParentParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParentParameters.
func (in *ParentParameters) DeepCopy() *ParentParameters {
	if in == nil {
		return nil
	}
	out := new(ParentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionResourcesObservation) DeepCopyInto(out *PermissionResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionResourcesObservation.
func (in *PermissionResourcesObservation) DeepCopy() *PermissionResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(PermissionResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionResourcesParameters) DeepCopyInto(out *PermissionResourcesParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionResourcesParameters.
func (in *PermissionResourcesParameters) DeepCopy() *PermissionResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(PermissionResourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefMoObservation) DeepCopyInto(out *RefMoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefMoObservation.
func (in *RefMoObservation) DeepCopy() *RefMoObservation {
	if in == nil {
		return nil
	}
	out := new(RefMoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefMoParameters) DeepCopyInto(out *RefMoParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefMoParameters.
func (in *RefMoParameters) DeepCopy() *RefMoParameters {
	if in == nil {
		return nil
	}
	out := new(RefMoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionObservation) DeepCopyInto(out *SessionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionObservation.
func (in *SessionObservation) DeepCopy() *SessionObservation {
	if in == nil {
		return nil
	}
	out := new(SessionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionParameters) DeepCopyInto(out *SessionParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionParameters.
func (in *SessionParameters) DeepCopy() *SessionParameters {
	if in == nil {
		return nil
	}
	out := new(SessionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsObservation) DeepCopyInto(out *TagsObservation) {
	*out = *in
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
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetObservation) DeepCopyInto(out *TargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetObservation.
func (in *TargetObservation) DeepCopy() *TargetObservation {
	if in == nil {
		return nil
	}
	out := new(TargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParameters) DeepCopyInto(out *TargetParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParameters.
func (in *TargetParameters) DeepCopy() *TargetParameters {
	if in == nil {
		return nil
	}
	out := new(TargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserObservation) DeepCopyInto(out *UserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserObservation.
func (in *UserObservation) DeepCopy() *UserObservation {
	if in == nil {
		return nil
	}
	out := new(UserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserParameters) DeepCopyInto(out *UserParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserParameters.
func (in *UserParameters) DeepCopy() *UserParameters {
	if in == nil {
		return nil
	}
	out := new(UserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcenterObservation) DeepCopyInto(out *VcenterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcenterObservation.
func (in *VcenterObservation) DeepCopy() *VcenterObservation {
	if in == nil {
		return nil
	}
	out := new(VcenterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcenterParameters) DeepCopyInto(out *VcenterParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcenterParameters.
func (in *VcenterParameters) DeepCopy() *VcenterParameters {
	if in == nil {
		return nil
	}
	out := new(VcenterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionContextObservation) DeepCopyInto(out *VersionContextObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionContextObservation.
func (in *VersionContextObservation) DeepCopy() *VersionContextObservation {
	if in == nil {
		return nil
	}
	out := new(VersionContextObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionContextParameters) DeepCopyInto(out *VersionContextParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.InterestedMos != nil {
		in, out := &in.InterestedMos, &out.InterestedMos
		*out = make([]InterestedMosParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NrVersion != nil {
		in, out := &in.NrVersion, &out.NrVersion
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.RefMo != nil {
		in, out := &in.RefMo, &out.RefMo
		*out = make([]RefMoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(string)
		**out = **in
	}
	if in.VersionType != nil {
		in, out := &in.VersionType, &out.VersionType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionContextParameters.
func (in *VersionContextParameters) DeepCopy() *VersionContextParameters {
	if in == nil {
		return nil
	}
	out := new(VersionContextParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineObservation) DeepCopyInto(out *VirtualMachineObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineObservation.
func (in *VirtualMachineObservation) DeepCopy() *VirtualMachineObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineParameters) DeepCopyInto(out *VirtualMachineParameters) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineParameters.
func (in *VirtualMachineParameters) DeepCopy() *VirtualMachineParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineParameters)
	in.DeepCopyInto(out)
	return out
}
