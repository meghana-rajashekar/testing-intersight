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

type DeviceObservation struct {
}

type DeviceParameters struct {

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

type ServerObservation struct {
}

type ServerParameters struct {

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

type SessionAncestorsObservation struct {
}

type SessionAncestorsParameters struct {

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

type SessionObservation struct {
}

type SessionParameters struct {

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

type SessionParentObservation struct {
}

type SessionParentParameters struct {

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

type SessionPermissionResourcesObservation struct {
}

type SessionPermissionResourcesParameters struct {

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

type SessionTagsObservation struct {
}

type SessionTagsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The string representation of a tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The string representation of a tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SessionVersionContextObservation struct {
}

type SessionVersionContextParameters struct {

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

type TargetObservation struct {
}

type TargetParameters struct {

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

type TunnelObservation struct {
}

type TunnelParameters struct {

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

type UserObservation struct {
}

type UserParameters struct {

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

// SessionSpec defines the desired state of Session
type SessionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SessionParameters `json:"forProvider"`
}

// SessionStatus defines the observed state of Session.
type SessionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SessionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Session is the Schema for the Sessions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type Session struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SessionSpec   `json:"spec"`
	Status            SessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SessionList contains a list of Sessions
type SessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Session `json:"items"`
}

// Repository type metadata.
var (
	Session_Kind             = "Session"
	Session_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Session_Kind}.String()
	Session_KindAPIVersion   = Session_Kind + "." + CRDGroupVersion.String()
	Session_GroupVersionKind = CRDGroupVersion.WithKind(Session_Kind)
)

func init() {
	SchemeBuilder.Register(&Session{}, &SessionList{})
}