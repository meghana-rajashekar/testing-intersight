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

type ClusterAddonProfileObservation struct {
}

type ClusterAddonProfileParameters struct {

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

type ClusterAncestorsObservation struct {
}

type ClusterAncestorsParameters struct {

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

type ClusterObservation struct {
}

type ClusterOrganizationObservation struct {
}

type ClusterOrganizationParameters struct {

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

type ClusterParameters struct {

	// The Account ID for this managed object.
	// +kubebuilder:validation:Optional
	AccountMoid *string `json:"accountMoid,omitempty" tf:"account_moid,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	Ancestors []ClusterAncestorsParameters `json:"ancestors,omitempty" tf:"ancestors,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// A reference to a kubernetesClusterAddonProfile resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	ClusterAddonProfile []ClusterAddonProfileParameters `json:"clusterAddonProfile,omitempty" tf:"cluster_addon_profile,omitempty"`

	// Status of the endpoint connection of this Kubernetes cluster.
	// * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.
	// * `Connected` - Intersight is able to establish a connection to the target and initiate management activities.
	// * `NotConnected` - Intersight is unable to establish a connection to the target.
	// * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.
	// * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.
	// * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	// +kubebuilder:validation:Optional
	ConnectionStatus *string `json:"connectionStatus,omitempty" tf:"connection_status,omitempty"`

	// The time when this managed object was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The DomainGroup ID for this managed object.
	// +kubebuilder:validation:Optional
	DomainGroupMoid *string `json:"domainGroupMoid,omitempty" tf:"domain_group_moid,omitempty"`

	// Kubeconfig for the cluster to collect inventory for.
	// +kubebuilder:validation:Optional
	KubeConfig *string `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`

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
	Organization []ClusterOrganizationParameters `json:"organization,omitempty" tf:"organization,omitempty"`

	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// A reference to a moBaseMo resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	Parent []ClusterParentParameters `json:"parent,omitempty" tf:"parent,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	PermissionResources []ClusterPermissionResourcesParameters `json:"permissionResources,omitempty" tf:"permission_resources,omitempty"`

	// An array of relationships to assetDeviceRegistration resources.
	// +kubebuilder:validation:Optional
	RegisteredDevices []RegisteredDevicesParameters `json:"registeredDevices,omitempty" tf:"registered_devices,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// Cluster health status as reported by the hypervisor platform.
	// * `Unknown` - Entity status is unknown.
	// * `Degraded` - State is degraded, and might impact normal operation of the entity.
	// * `Critical` - Entity is in a critical state, impacting operations.
	// * `Ok` - Entity status is in a stable state, operating normally.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// An array of relationships to storageBaseCluster resources.
	// +kubebuilder:validation:Optional
	StorageClusters []StorageClustersParameters `json:"storageClusters,omitempty" tf:"storage_clusters,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []ClusterTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// The versioning info for this managed object.
	// +kubebuilder:validation:Optional
	VersionContext []ClusterVersionContextParameters `json:"versionContext,omitempty" tf:"version_context,omitempty"`
}

type ClusterParentObservation struct {
}

type ClusterParentParameters struct {

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

type ClusterPermissionResourcesObservation struct {
}

type ClusterPermissionResourcesParameters struct {

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

type ClusterTagsObservation struct {
}

type ClusterTagsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The string representation of a tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The string representation of a tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ClusterVersionContextInterestedMosObservation struct {
}

type ClusterVersionContextInterestedMosParameters struct {

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

type ClusterVersionContextObservation struct {
}

type ClusterVersionContextParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// +kubebuilder:validation:Optional
	InterestedMos []ClusterVersionContextInterestedMosParameters `json:"interestedMos,omitempty" tf:"interested_mos,omitempty"`

	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	// +kubebuilder:validation:Optional
	NrVersion *string `json:"nrVersion,omitempty" tf:"nr_version,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// A reference to the original Managed Object.
	// +kubebuilder:validation:Optional
	RefMo []ClusterVersionContextRefMoParameters `json:"refMo,omitempty" tf:"ref_mo,omitempty"`

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

type ClusterVersionContextRefMoObservation struct {
}

type ClusterVersionContextRefMoParameters struct {

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

type RegisteredDevicesObservation struct {
}

type RegisteredDevicesParameters struct {

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

type StorageClustersObservation struct {
}

type StorageClustersParameters struct {

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

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
