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

type KvmSessionObservation struct {
}

type KvmSessionParameters struct {

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

type TunnelAncestorsObservation struct {
}

type TunnelAncestorsParameters struct {

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

type TunnelDeviceObservation struct {
}

type TunnelDeviceParameters struct {

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

type TunnelObservation_2 struct {
}

type TunnelParameters_2 struct {

	// The Account ID for this managed object.
	// +kubebuilder:validation:Optional
	AccountMoid *string `json:"accountMoid,omitempty" tf:"account_moid,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	Ancestors []TunnelAncestorsParameters `json:"ancestors,omitempty" tf:"ancestors,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The user agent IP address from which the session is launched.
	// +kubebuilder:validation:Optional
	ClientIPAddress *string `json:"clientIpAddress,omitempty" tf:"client_ip_address,omitempty"`

	// The multiplexer URL for the client to connect on.
	// +kubebuilder:validation:Optional
	ClientURL *string `json:"clientUrl,omitempty" tf:"client_url,omitempty"`

	// The time when this managed object was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// A reference to a assetDeviceRegistration resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Device []TunnelDeviceParameters `json:"device,omitempty" tf:"device,omitempty"`

	// The DomainGroup ID for this managed object.
	// +kubebuilder:validation:Optional
	DomainGroupMoid *string `json:"domainGroupMoid,omitempty" tf:"domain_group_moid,omitempty"`

	// The time at which the session ended.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A reference to a kvmSession resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	KvmSession []KvmSessionParameters `json:"kvmSession,omitempty" tf:"kvm_session,omitempty"`

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

	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// A reference to a moBaseMo resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Parent []TunnelParentParameters `json:"parent,omitempty" tf:"parent,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	PermissionResources []TunnelPermissionResourcesParameters `json:"permissionResources,omitempty" tf:"permission_resources,omitempty"`

	// Role of the user who launched the session.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// A reference to a computePhysical resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Server []TunnelServerParameters `json:"server,omitempty" tf:"server,omitempty"`

	// A reference to a sessionAbstractSession resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Session []TunnelSessionParameters `json:"session,omitempty" tf:"session,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// The status of the session.
	// * `Active` - The session is currently active.
	// * `Ended` - The session has ended normally.
	// * `Terminated` - The session was terminated by an admin.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TunnelTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// A reference to a moBaseMo resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Target []TunnelTargetParameters `json:"target,omitempty" tf:"target,omitempty"`

	// Name of target on which session is initiated.
	// +kubebuilder:validation:Optional
	TargetName *string `json:"targetName,omitempty" tf:"target_name,omitempty"`

	// A reference to a iamUser resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	User []TunnelUserParameters `json:"user,omitempty" tf:"user,omitempty"`

	// User ID or E-mail Address of the user who launched the session.
	// +kubebuilder:validation:Optional
	UserIDOrEmail *string `json:"userIdOrEmail,omitempty" tf:"user_id_or_email,omitempty"`

	// The versioning info for this managed object.
	// +kubebuilder:validation:Optional
	VersionContext []TunnelVersionContextParameters `json:"versionContext,omitempty" tf:"version_context,omitempty"`
}

type TunnelParentObservation struct {
}

type TunnelParentParameters struct {

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

type TunnelPermissionResourcesObservation struct {
}

type TunnelPermissionResourcesParameters struct {

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

type TunnelServerObservation struct {
}

type TunnelServerParameters struct {

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

type TunnelSessionObservation struct {
}

type TunnelSessionParameters struct {

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

type TunnelTagsObservation struct {
}

type TunnelTagsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The string representation of a tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The string representation of a tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TunnelTargetObservation struct {
}

type TunnelTargetParameters struct {

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

type TunnelUserObservation struct {
}

type TunnelUserParameters struct {

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

type TunnelVersionContextInterestedMosObservation struct {
}

type TunnelVersionContextInterestedMosParameters struct {

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

type TunnelVersionContextObservation struct {
}

type TunnelVersionContextParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// +kubebuilder:validation:Optional
	InterestedMos []TunnelVersionContextInterestedMosParameters `json:"interestedMos,omitempty" tf:"interested_mos,omitempty"`

	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	// +kubebuilder:validation:Optional
	NrVersion *string `json:"nrVersion,omitempty" tf:"nr_version,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// A reference to the original Managed Object.
	// +kubebuilder:validation:Optional
	RefMo []TunnelVersionContextRefMoParameters `json:"refMo,omitempty" tf:"ref_mo,omitempty"`

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

type TunnelVersionContextRefMoObservation struct {
}

type TunnelVersionContextRefMoParameters struct {

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

// TunnelSpec defines the desired state of Tunnel
type TunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelParameters_2 `json:"forProvider"`
}

// TunnelStatus defines the observed state of Tunnel.
type TunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tunnel is the Schema for the Tunnels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type Tunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TunnelSpec   `json:"spec"`
	Status            TunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelList contains a list of Tunnels
type TunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tunnel `json:"items"`
}

// Repository type metadata.
var (
	Tunnel_Kind             = "Tunnel"
	Tunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tunnel_Kind}.String()
	Tunnel_KindAPIVersion   = Tunnel_Kind + "." + CRDGroupVersion.String()
	Tunnel_GroupVersionKind = CRDGroupVersion.WithKind(Tunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&Tunnel{}, &TunnelList{})
}
