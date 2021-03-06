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
func (in *AccountObservation) DeepCopyInto(out *AccountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountObservation.
func (in *AccountObservation) DeepCopy() *AccountObservation {
	if in == nil {
		return nil
	}
	out := new(AccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountParameters) DeepCopyInto(out *AccountParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountParameters.
func (in *AccountParameters) DeepCopy() *AccountParameters {
	if in == nil {
		return nil
	}
	out := new(AccountParameters)
	in.DeepCopyInto(out)
	return out
}

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
func (in *CloudResourceObservation) DeepCopyInto(out *CloudResourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourceObservation.
func (in *CloudResourceObservation) DeepCopy() *CloudResourceObservation {
	if in == nil {
		return nil
	}
	out := new(CloudResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourceParameters) DeepCopyInto(out *CloudResourceParameters) {
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
	if in.CurrentStatus != nil {
		in, out := &in.CurrentStatus, &out.CurrentStatus
		*out = new(string)
		**out = **in
	}
	if in.DesiredStatus != nil {
		in, out := &in.DesiredStatus, &out.DesiredStatus
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourceParameters.
func (in *CloudResourceParameters) DeepCopy() *CloudResourceParameters {
	if in == nil {
		return nil
	}
	out := new(CloudResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Executor) DeepCopyInto(out *Executor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Executor.
func (in *Executor) DeepCopy() *Executor {
	if in == nil {
		return nil
	}
	out := new(Executor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Executor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorList) DeepCopyInto(out *ExecutorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Executor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorList.
func (in *ExecutorList) DeepCopy() *ExecutorList {
	if in == nil {
		return nil
	}
	out := new(ExecutorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExecutorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorObservation) DeepCopyInto(out *ExecutorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorObservation.
func (in *ExecutorObservation) DeepCopy() *ExecutorObservation {
	if in == nil {
		return nil
	}
	out := new(ExecutorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorParameters) DeepCopyInto(out *ExecutorParameters) {
	*out = *in
	if in.Account != nil {
		in, out := &in.Account, &out.Account
		*out = make([]AccountParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
	if in.CloudResource != nil {
		in, out := &in.CloudResource, &out.CloudResource
		*out = make([]CloudResourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Operation != nil {
		in, out := &in.Operation, &out.Operation
		*out = new(string)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
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
	if in.PlatformType != nil {
		in, out := &in.PlatformType, &out.PlatformType
		*out = new(string)
		**out = **in
	}
	if in.RegisteredDevice != nil {
		in, out := &in.RegisteredDevice, &out.RegisteredDevice
		*out = make([]RegisteredDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunState != nil {
		in, out := &in.RunState, &out.RunState
		*out = make([]RunStateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SharedScope != nil {
		in, out := &in.SharedScope, &out.SharedScope
		*out = new(string)
		**out = **in
	}
	if in.SourceFolderName != nil {
		in, out := &in.SourceFolderName, &out.SourceFolderName
		*out = new(string)
		**out = **in
	}
	if in.SourceFolderPath != nil {
		in, out := &in.SourceFolderPath, &out.SourceFolderPath
		*out = new(string)
		**out = **in
	}
	if in.SourceLocation != nil {
		in, out := &in.SourceLocation, &out.SourceLocation
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Stderr != nil {
		in, out := &in.Stderr, &out.Stderr
		*out = new(string)
		**out = **in
	}
	if in.Stdout != nil {
		in, out := &in.Stdout, &out.Stdout
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
	if in.TaskID != nil {
		in, out := &in.TaskID, &out.TaskID
		*out = new(string)
		**out = **in
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = new(string)
		**out = **in
	}
	if in.VersionContext != nil {
		in, out := &in.VersionContext, &out.VersionContext
		*out = make([]VersionContextParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WaitForCompletion != nil {
		in, out := &in.WaitForCompletion, &out.WaitForCompletion
		*out = new(bool)
		**out = **in
	}
	if in.WorkflowInfo != nil {
		in, out := &in.WorkflowInfo, &out.WorkflowInfo
		*out = make([]WorkflowInfoParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorParameters.
func (in *ExecutorParameters) DeepCopy() *ExecutorParameters {
	if in == nil {
		return nil
	}
	out := new(ExecutorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorSpec) DeepCopyInto(out *ExecutorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorSpec.
func (in *ExecutorSpec) DeepCopy() *ExecutorSpec {
	if in == nil {
		return nil
	}
	out := new(ExecutorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorStatus) DeepCopyInto(out *ExecutorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorStatus.
func (in *ExecutorStatus) DeepCopy() *ExecutorStatus {
	if in == nil {
		return nil
	}
	out := new(ExecutorStatus)
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
func (in *RegisteredDeviceObservation) DeepCopyInto(out *RegisteredDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisteredDeviceObservation.
func (in *RegisteredDeviceObservation) DeepCopy() *RegisteredDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(RegisteredDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisteredDeviceParameters) DeepCopyInto(out *RegisteredDeviceParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisteredDeviceParameters.
func (in *RegisteredDeviceParameters) DeepCopy() *RegisteredDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(RegisteredDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStateObservation) DeepCopyInto(out *RunStateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStateObservation.
func (in *RunStateObservation) DeepCopy() *RunStateObservation {
	if in == nil {
		return nil
	}
	out := new(RunStateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStateParameters) DeepCopyInto(out *RunStateParameters) {
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
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.RunID != nil {
		in, out := &in.RunID, &out.RunID
		*out = new(string)
		**out = **in
	}
	if in.StateFile != nil {
		in, out := &in.StateFile, &out.StateFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStateParameters.
func (in *RunStateParameters) DeepCopy() *RunStateParameters {
	if in == nil {
		return nil
	}
	out := new(RunStateParameters)
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
func (in *WorkflowInfoObservation) DeepCopyInto(out *WorkflowInfoObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowInfoObservation.
func (in *WorkflowInfoObservation) DeepCopy() *WorkflowInfoObservation {
	if in == nil {
		return nil
	}
	out := new(WorkflowInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowInfoParameters) DeepCopyInto(out *WorkflowInfoParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowInfoParameters.
func (in *WorkflowInfoParameters) DeepCopy() *WorkflowInfoParameters {
	if in == nil {
		return nil
	}
	out := new(WorkflowInfoParameters)
	in.DeepCopyInto(out)
	return out
}
