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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type M2VirtualDriveObservation struct {
}

type M2VirtualDriveParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Select the M.2 RAID controller slot on which the virtual drive is to be created. Select 'MSTOR-RAID-1' to create virtual drive on the M.2 RAID controller in the first slot or in the MSTOR-RAID slot, 'MSTOR-RAID-2' for second slot, 'MSTOR-RAID-1, MSTOR-RAID-2' for both slots or either slot.
	// * `MSTOR-RAID-1` - Virtual drive  will be created on the M.2 RAID controller in the first slot.
	// * `MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available.
	// * `MSTOR-RAID-1,MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available.
	// +kubebuilder:validation:Optional
	ControllerSlot *string `json:"controllerSlot,omitempty" tf:"controller_slot,omitempty"`

	// If enabled, this will create a virtual drive on the M.2 RAID controller.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`
}

type OrganizationObservation struct {
}

type OrganizationParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type ProfilesObservation struct {
}

type ProfilesParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type Raid0DriveObservation struct {
}

type Raid0DriveParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The set of drive slots where RAID0 virtual drives must be created.
	// +kubebuilder:validation:Optional
	DriveSlots *string `json:"driveSlots,omitempty" tf:"drive_slots,omitempty"`

	// The list of drive slots where RAID0 virtual drives must be created (comma seperated).
	// +kubebuilder:validation:Optional
	DriveSlotsList *string `json:"driveSlotsList,omitempty" tf:"drive_slots_list,omitempty"`

	// If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// This defines the characteristics of a specific virtual drive.
	// +kubebuilder:validation:Optional
	VirtualDrivePolicy []Raid0DriveVirtualDrivePolicyParameters `json:"virtualDrivePolicy,omitempty" tf:"virtual_drive_policy,omitempty"`
}

type Raid0DriveVirtualDrivePolicyObservation struct {
}

type Raid0DriveVirtualDrivePolicyParameters struct {

	// Access policy that host has on this virtual drive.
	// * `Default` - Use platform default access mode.
	// * `ReadWrite` - Enables host to perform read-write on the VD.
	// * `ReadOnly` - Host can only read from the VD.
	// * `Blocked` - Host can neither read nor write to the VD.
	// +kubebuilder:validation:Optional
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Disk cache policy for the virtual drive.
	// * `Default` - Use platform default drive cache mode.
	// * `NoChange` - Drive cache policy is unchanged.
	// * `Enable` - Enables IO caching on the drive.
	// * `Disable` - Disables IO caching on the drive.
	// +kubebuilder:validation:Optional
	DriveCache *string `json:"driveCache,omitempty" tf:"drive_cache,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// Read ahead mode to be used to read data from this virtual drive.
	// * `Default` - Use platform default read ahead mode.
	// * `ReadAhead` - Use read ahead mode for the policy.
	// * `NoReadAhead` - Do not use read ahead mode for the policy.
	// +kubebuilder:validation:Optional
	ReadPolicy *string `json:"readPolicy,omitempty" tf:"read_policy,omitempty"`

	// Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB.
	// * `64` - Number of bytes in a strip is 64 Kibibytes.
	// * `128` - Number of bytes in a strip is 128 Kibibytes.
	// * `256` - Number of bytes in a strip is 256 Kibibytes.
	// * `512` - Number of bytes in a strip is 512 Kibibytes.
	// * `1024` - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte.
	// +kubebuilder:validation:Optional
	StripSize *int64 `json:"stripSize,omitempty" tf:"strip_size,omitempty"`

	// Write mode to be used to write data to this virtual drive.
	// * `Default` - Use platform default write mode.
	// * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache.
	// * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure.
	// * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
	// +kubebuilder:validation:Optional
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

type StoragePolicyAncestorsObservation struct {
}

type StoragePolicyAncestorsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type StoragePolicyDriveGroupObservation struct {
}

type StoragePolicyDriveGroupParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type StoragePolicyObservation_2 struct {
}

type StoragePolicyParameters_2 struct {

	// The Account ID for this managed object.
	// +kubebuilder:validation:Optional
	AccountMoid *string `json:"accountMoid,omitempty" tf:"account_moid,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	Ancestors []StoragePolicyAncestorsParameters `json:"ancestors,omitempty" tf:"ancestors,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The time when this managed object was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The DomainGroup ID for this managed object.
	// +kubebuilder:validation:Optional
	DomainGroupMoid *string `json:"domainGroupMoid,omitempty" tf:"domain_group_moid,omitempty"`

	// An array of relationships to storageDriveGroup resources.
	// +kubebuilder:validation:Optional
	DriveGroup []StoragePolicyDriveGroupParameters `json:"driveGroup,omitempty" tf:"drive_group,omitempty"`

	// A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen.
	// +kubebuilder:validation:Optional
	GlobalHotSpares *string `json:"globalHotSpares,omitempty" tf:"global_hot_spares,omitempty"`

	// Virtual Drive configuration for M.2 RAID controller.
	// +kubebuilder:validation:Optional
	M2VirtualDrive []M2VirtualDriveParameters `json:"m2VirtualDrive,omitempty" tf:"m2_virtual_drive,omitempty"`

	// The time when this managed object was last modified.
	// +kubebuilder:validation:Optional
	ModTime *string `json:"modTime,omitempty" tf:"mod_time,omitempty"`

	// The unique identifier of this Managed Object instance.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// A reference to a organizationOrganization resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Organization []OrganizationParameters `json:"organization,omitempty" tf:"organization,omitempty"`

	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// A reference to a moBaseMo resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Parent []StoragePolicyParentParameters `json:"parent,omitempty" tf:"parent,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	PermissionResources []StoragePolicyPermissionResourcesParameters `json:"permissionResources,omitempty" tf:"permission_resources,omitempty"`

	// An array of relationships to policyAbstractConfigProfile resources.
	// +kubebuilder:validation:Optional
	Profiles []ProfilesParameters `json:"profiles,omitempty" tf:"profiles,omitempty"`

	// The list of disks where RAID0 virtual drives must be created on each individual drive.
	// +kubebuilder:validation:Optional
	Raid0Drive []Raid0DriveParameters `json:"raid0Drive,omitempty" tf:"raid0_drive,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []StoragePolicyTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// State to which disks, not used in this policy, are to be moved. NoChange will not change the drive state.
	// * `NoChange` - Drive state will not be modified by Storage Policy.
	// * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.
	// * `Jbod` - JBOD state where the disks start showing up to Host OS.
	// +kubebuilder:validation:Optional
	UnusedDisksState *string `json:"unusedDisksState,omitempty" tf:"unused_disks_state,omitempty"`

	// Disks in JBOD State are used to create virtual drives.
	// +kubebuilder:validation:Optional
	UseJbodForVdCreation *bool `json:"useJbodForVdCreation,omitempty" tf:"use_jbod_for_vd_creation,omitempty"`

	// The versioning info for this managed object.
	// +kubebuilder:validation:Optional
	VersionContext []StoragePolicyVersionContextParameters `json:"versionContext,omitempty" tf:"version_context,omitempty"`
}

type StoragePolicyParentObservation struct {
}

type StoragePolicyParentParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type StoragePolicyPermissionResourcesObservation struct {
}

type StoragePolicyPermissionResourcesParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type StoragePolicyTagsObservation struct {
}

type StoragePolicyTagsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The string representation of a tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The string representation of a tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StoragePolicyVersionContextObservation struct {
}

type StoragePolicyVersionContextParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// +kubebuilder:validation:Optional
	InterestedMos []VersionContextInterestedMosParameters `json:"interestedMos,omitempty" tf:"interested_mos,omitempty"`

	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	// +kubebuilder:validation:Optional
	NrVersion *string `json:"nrVersion,omitempty" tf:"nr_version,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// A reference to the original Managed Object.
	// +kubebuilder:validation:Optional
	RefMo []VersionContextRefMoParameters `json:"refMo,omitempty" tf:"ref_mo,omitempty"`

	// The time this versioned Managed Object was created.
	// +kubebuilder:validation:Optional
	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp,omitempty"`

	// Specifies type of version. Currently the only supported value is "Configured"
	// that is used to keep track of snapshots of policies and profiles that are intended
	// to be configured to target endpoints.
	// * `Modified` - Version created every time an object is modified.
	// * `Configured` - Version created every time an object is configured to the service profile.
	// * `Deployed` - Version created for objects related to a service profile when it is deployed.
	// +kubebuilder:validation:Optional
	VersionType *string `json:"versionType,omitempty" tf:"version_type,omitempty"`
}

type VersionContextInterestedMosObservation struct {
}

type VersionContextInterestedMosParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type VersionContextRefMoObservation struct {
}

type VersionContextRefMoParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Moid of the referenced REST resource.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// The fully-qualified name of the remote type referred by this relationship.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// An OData $filter expression which describes the REST resource to be referenced. This field may
	// be set instead of 'moid' by clients.
	// 1. If 'moid' is set this field is ignored.
	// 1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the
	// resource matching the filter expression and populates it in the MoRef that is part of the object
	// instance being inserted/updated to fulfill the REST request.
	// An error is returned if the filter matches zero or more than one REST resource.
	// An example filter string is: Serial eq '3AA8B7T11'.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

// StoragePolicySpec defines the desired state of StoragePolicy
type StoragePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StoragePolicyParameters_2 `json:"forProvider"`
}

// StoragePolicyStatus defines the observed state of StoragePolicy.
type StoragePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StoragePolicyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragePolicy is the Schema for the StoragePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type StoragePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragePolicySpec   `json:"spec"`
	Status            StoragePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragePolicyList contains a list of StoragePolicys
type StoragePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragePolicy `json:"items"`
}

// Repository type metadata.
var (
	StoragePolicy_Kind             = "StoragePolicy"
	StoragePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StoragePolicy_Kind}.String()
	StoragePolicy_KindAPIVersion   = StoragePolicy_Kind + "." + CRDGroupVersion.String()
	StoragePolicy_GroupVersionKind = CRDGroupVersion.WithKind(StoragePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&StoragePolicy{}, &StoragePolicyList{})
}