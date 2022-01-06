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
func (in *HealthCheckDefinition) DeepCopyInto(out *HealthCheckDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinition.
func (in *HealthCheckDefinition) DeepCopy() *HealthCheckDefinition {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthCheckDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckDefinitionList) DeepCopyInto(out *HealthCheckDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthCheckDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinitionList.
func (in *HealthCheckDefinitionList) DeepCopy() *HealthCheckDefinitionList {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthCheckDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckDefinitionObservation) DeepCopyInto(out *HealthCheckDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinitionObservation.
func (in *HealthCheckDefinitionObservation) DeepCopy() *HealthCheckDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckDefinitionParameters) DeepCopyInto(out *HealthCheckDefinitionParameters) {
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
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.ClassID != nil {
		in, out := &in.ClassID, &out.ClassID
		*out = new(string)
		**out = **in
	}
	if in.CommonCauses != nil {
		in, out := &in.CommonCauses, &out.CommonCauses
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DomainGroupMoid != nil {
		in, out := &in.DomainGroupMoid, &out.DomainGroupMoid
		*out = new(string)
		**out = **in
	}
	if in.ExecutionMode != nil {
		in, out := &in.ExecutionMode, &out.ExecutionMode
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
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
	if in.SharedScope != nil {
		in, out := &in.SharedScope, &out.SharedScope
		*out = new(string)
		**out = **in
	}
	if in.SuggestedResolution != nil {
		in, out := &in.SuggestedResolution, &out.SuggestedResolution
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
	if in.VersionContext != nil {
		in, out := &in.VersionContext, &out.VersionContext
		*out = make([]VersionContextParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinitionParameters.
func (in *HealthCheckDefinitionParameters) DeepCopy() *HealthCheckDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckDefinitionSpec) DeepCopyInto(out *HealthCheckDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinitionSpec.
func (in *HealthCheckDefinitionSpec) DeepCopy() *HealthCheckDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckDefinitionStatus) DeepCopyInto(out *HealthCheckDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckDefinitionStatus.
func (in *HealthCheckDefinitionStatus) DeepCopy() *HealthCheckDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(HealthCheckDefinitionStatus)
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