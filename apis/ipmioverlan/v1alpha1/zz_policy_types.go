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

type AncestorsObservation struct {
}

type AncestorsParameters struct {

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

type InterestedMosObservation struct {
}

type InterestedMosParameters struct {

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

type ParentObservation struct {
}

type ParentParameters struct {

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

type PermissionResourcesObservation struct {
}

type PermissionResourcesParameters struct {

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

type PolicyObservation struct {
}

type PolicyParameters struct {

	// The Account ID for this managed object.
	// +kubebuilder:validation:Optional
	AccountMoid *string `json:"accountMoid,omitempty" tf:"account_moid,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	Ancestors []AncestorsParameters `json:"ancestors,omitempty" tf:"ancestors,omitempty"`

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

	// State of the IPMI Over LAN service on the endpoint.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters.
	// +kubebuilder:validation:Optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// Indicates whether the value of the 'encryptionKey' property has been set.
	// +kubebuilder:validation:Optional
	IsEncryptionKeySet *bool `json:"isEncryptionKeySet,omitempty" tf:"is_encryption_key_set,omitempty"`

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
	Parent []ParentParameters `json:"parent,omitempty" tf:"parent,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	PermissionResources []PermissionResourcesParameters `json:"permissionResources,omitempty" tf:"permission_resources,omitempty"`

	// The highest privilege level that can be assigned to an IPMI session on a server.
	// * `admin` - Privilege to perform all actions available through IPMI.
	// * `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks.
	// * `read-only` - Privilege to view information throught IPMI but restriction on making any changes.
	// +kubebuilder:validation:Optional
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// An array of relationships to policyAbstractConfigProfile resources.
	// +kubebuilder:validation:Optional
	Profiles []ProfilesParameters `json:"profiles,omitempty" tf:"profiles,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// The versioning info for this managed object.
	// +kubebuilder:validation:Optional
	VersionContext []VersionContextParameters `json:"versionContext,omitempty" tf:"version_context,omitempty"`
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

type RefMoObservation struct {
}

type RefMoParameters struct {

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

type TagsObservation struct {
}

type TagsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The string representation of a tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The string representation of a tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VersionContextObservation struct {
}

type VersionContextParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// +kubebuilder:validation:Optional
	InterestedMos []InterestedMosParameters `json:"interestedMos,omitempty" tf:"interested_mos,omitempty"`

	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	// +kubebuilder:validation:Optional
	NrVersion *string `json:"nrVersion,omitempty" tf:"nr_version,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// A reference to the original Managed Object.
	// +kubebuilder:validation:Optional
	RefMo []RefMoParameters `json:"refMo,omitempty" tf:"ref_mo,omitempty"`

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

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
