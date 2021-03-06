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
func (in *ConfigPolicy) DeepCopyInto(out *ConfigPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicy.
func (in *ConfigPolicy) DeepCopy() *ConfigPolicy {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPolicyList) DeepCopyInto(out *ConfigPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicyList.
func (in *ConfigPolicyList) DeepCopy() *ConfigPolicyList {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPolicyObservation) DeepCopyInto(out *ConfigPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicyObservation.
func (in *ConfigPolicyObservation) DeepCopy() *ConfigPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPolicyParameters) DeepCopyInto(out *ConfigPolicyParameters) {
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
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = make([]OrganizationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]ProfilesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make([]SettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SharedScope != nil {
		in, out := &in.SharedScope, &out.SharedScope
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicyParameters.
func (in *ConfigPolicyParameters) DeepCopy() *ConfigPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPolicySpec) DeepCopyInto(out *ConfigPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicySpec.
func (in *ConfigPolicySpec) DeepCopy() *ConfigPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigPolicyStatus) DeepCopyInto(out *ConfigPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigPolicyStatus.
func (in *ConfigPolicyStatus) DeepCopy() *ConfigPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DceInterfaceSettingsObservation) DeepCopyInto(out *DceInterfaceSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DceInterfaceSettingsObservation.
func (in *DceInterfaceSettingsObservation) DeepCopy() *DceInterfaceSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(DceInterfaceSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DceInterfaceSettingsParameters) DeepCopyInto(out *DceInterfaceSettingsParameters) {
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
	if in.FecMode != nil {
		in, out := &in.FecMode, &out.FecMode
		*out = new(string)
		**out = **in
	}
	if in.InterfaceID != nil {
		in, out := &in.InterfaceID, &out.InterfaceID
		*out = new(int64)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DceInterfaceSettingsParameters.
func (in *DceInterfaceSettingsParameters) DeepCopy() *DceInterfaceSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(DceInterfaceSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EthSettingsObservation) DeepCopyInto(out *EthSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EthSettingsObservation.
func (in *EthSettingsObservation) DeepCopy() *EthSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(EthSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EthSettingsParameters) DeepCopyInto(out *EthSettingsParameters) {
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
	if in.LldpEnabled != nil {
		in, out := &in.LldpEnabled, &out.LldpEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EthSettingsParameters.
func (in *EthSettingsParameters) DeepCopy() *EthSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(EthSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FcSettingsObservation) DeepCopyInto(out *FcSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FcSettingsObservation.
func (in *FcSettingsObservation) DeepCopy() *FcSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(FcSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FcSettingsParameters) DeepCopyInto(out *FcSettingsParameters) {
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
	if in.FipEnabled != nil {
		in, out := &in.FipEnabled, &out.FipEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FcSettingsParameters.
func (in *FcSettingsParameters) DeepCopy() *FcSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(FcSettingsParameters)
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
func (in *OrganizationObservation) DeepCopyInto(out *OrganizationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationObservation.
func (in *OrganizationObservation) DeepCopy() *OrganizationObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationParameters) DeepCopyInto(out *OrganizationParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationParameters.
func (in *OrganizationParameters) DeepCopy() *OrganizationParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationParameters)
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
func (in *PortChannelSettingsObservation) DeepCopyInto(out *PortChannelSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortChannelSettingsObservation.
func (in *PortChannelSettingsObservation) DeepCopy() *PortChannelSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(PortChannelSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortChannelSettingsParameters) DeepCopyInto(out *PortChannelSettingsParameters) {
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortChannelSettingsParameters.
func (in *PortChannelSettingsParameters) DeepCopy() *PortChannelSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(PortChannelSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfilesObservation) DeepCopyInto(out *ProfilesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfilesObservation.
func (in *ProfilesObservation) DeepCopy() *ProfilesObservation {
	if in == nil {
		return nil
	}
	out := new(ProfilesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfilesParameters) DeepCopyInto(out *ProfilesParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfilesParameters.
func (in *ProfilesParameters) DeepCopy() *ProfilesParameters {
	if in == nil {
		return nil
	}
	out := new(ProfilesParameters)
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
func (in *SettingsObservation) DeepCopyInto(out *SettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsObservation.
func (in *SettingsObservation) DeepCopy() *SettingsObservation {
	if in == nil {
		return nil
	}
	out := new(SettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsParameters) DeepCopyInto(out *SettingsParameters) {
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
	if in.DceInterfaceSettings != nil {
		in, out := &in.DceInterfaceSettings, &out.DceInterfaceSettings
		*out = make([]DceInterfaceSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EthSettings != nil {
		in, out := &in.EthSettings, &out.EthSettings
		*out = make([]EthSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FcSettings != nil {
		in, out := &in.FcSettings, &out.FcSettings
		*out = make([]FcSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.PortChannelSettings != nil {
		in, out := &in.PortChannelSettings, &out.PortChannelSettings
		*out = make([]PortChannelSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SlotID != nil {
		in, out := &in.SlotID, &out.SlotID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsParameters.
func (in *SettingsParameters) DeepCopy() *SettingsParameters {
	if in == nil {
		return nil
	}
	out := new(SettingsParameters)
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
