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

type ArfsSettingsObservation struct {
}

type ArfsSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`
}

type CompletionQueueSettingsObservation struct {
}

type CompletionQueueSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources.
	// +kubebuilder:validation:Optional
	NrCount *int64 `json:"nrCount,omitempty" tf:"nr_count,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The number of descriptors in each completion queue.
	// +kubebuilder:validation:Optional
	RingSize *int64 `json:"ringSize,omitempty" tf:"ring_size,omitempty"`
}

type EthAdapterPolicyObservation struct {
}

type EthAdapterPolicyParameters struct {

	// The Account ID for this managed object.
	// +kubebuilder:validation:Optional
	AccountMoid *string `json:"accountMoid,omitempty" tf:"account_moid,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// Enables advanced filtering on the interface.
	// +kubebuilder:validation:Optional
	AdvancedFilter *bool `json:"advancedFilter,omitempty" tf:"advanced_filter,omitempty"`

	// An array of relationships to moBaseMo resources.
	// +kubebuilder:validation:Optional
	Ancestors []AncestorsParameters `json:"ancestors,omitempty" tf:"ancestors,omitempty"`

	// Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.
	// +kubebuilder:validation:Optional
	ArfsSettings []ArfsSettingsParameters `json:"arfsSettings,omitempty" tf:"arfs_settings,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Completion Queue resource settings.
	// +kubebuilder:validation:Optional
	CompletionQueueSettings []CompletionQueueSettingsParameters `json:"completionQueueSettings,omitempty" tf:"completion_queue_settings,omitempty"`

	// The time when this managed object was created.
	// +kubebuilder:validation:Optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of the policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The DomainGroup ID for this managed object.
	// +kubebuilder:validation:Optional
	DomainGroupMoid *string `json:"domainGroupMoid,omitempty" tf:"domain_group_moid,omitempty"`

	// GENEVE offload protocol allows you to create logical networks that span physical network boundaries by allowing any information to be encoded in a packet and passed between tunnel endpoints.
	// +kubebuilder:validation:Optional
	GeneveEnabled *bool `json:"geneveEnabled,omitempty" tf:"geneve_enabled,omitempty"`

	// Enables Interrupt Scaling on the interface.
	// +kubebuilder:validation:Optional
	InterruptScaling *bool `json:"interruptScaling,omitempty" tf:"interrupt_scaling,omitempty"`

	// Interrupt Settings for the virtual ethernet interface.
	// +kubebuilder:validation:Optional
	InterruptSettings []InterruptSettingsParameters `json:"interruptSettings,omitempty" tf:"interrupt_settings,omitempty"`

	// The time when this managed object was last modified.
	// +kubebuilder:validation:Optional
	ModTime *string `json:"modTime,omitempty" tf:"mod_time,omitempty"`

	// The unique identifier of this Managed Object instance.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// Network Virtualization using Generic Routing Encapsulation Settings.
	// +kubebuilder:validation:Optional
	NvgreSettings []NvgreSettingsParameters `json:"nvgreSettings,omitempty" tf:"nvgre_settings,omitempty"`

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

	// Settings for RDMA over Converged Ethernet.
	// +kubebuilder:validation:Optional
	RoceSettings []RoceSettingsParameters `json:"roceSettings,omitempty" tf:"roce_settings,omitempty"`

	// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
	// +kubebuilder:validation:Optional
	RssHashSettings []RssHashSettingsParameters `json:"rssHashSettings,omitempty" tf:"rss_hash_settings,omitempty"`

	// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
	// +kubebuilder:validation:Optional
	RssSettings *bool `json:"rssSettings,omitempty" tf:"rss_settings,omitempty"`

	// Receive Queue resouce settings.
	// +kubebuilder:validation:Optional
	RxQueueSettings []RxQueueSettingsParameters `json:"rxQueueSettings,omitempty" tf:"rx_queue_settings,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.
	// +kubebuilder:validation:Optional
	TCPOffloadSettings []TCPOffloadSettingsParameters `json:"tcpOffloadSettings,omitempty" tf:"tcp_offload_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// Transmit Queue resource settings.
	// +kubebuilder:validation:Optional
	TxQueueSettings []TxQueueSettingsParameters `json:"txQueueSettings,omitempty" tf:"tx_queue_settings,omitempty"`

	// Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.
	// +kubebuilder:validation:Optional
	UplinkFailbackTimeout *int64 `json:"uplinkFailbackTimeout,omitempty" tf:"uplink_failback_timeout,omitempty"`

	// The versioning info for this managed object.
	// +kubebuilder:validation:Optional
	VersionContext []VersionContextParameters `json:"versionContext,omitempty" tf:"version_context,omitempty"`

	// Virtual Extensible LAN Protocol Settings.
	// +kubebuilder:validation:Optional
	VxlanSettings []VxlanSettingsParameters `json:"vxlanSettings,omitempty" tf:"vxlan_settings,omitempty"`
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

type InterruptSettingsObservation struct {
}

type InterruptSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.
	// +kubebuilder:validation:Optional
	CoalescingTime *int64 `json:"coalescingTime,omitempty" tf:"coalescing_time,omitempty"`

	// Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.
	// * `MIN` - The system waits for the time specified in the Coalescing Time field before sending another interrupt event.
	// * `IDLE` - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.
	// +kubebuilder:validation:Optional
	CoalescingType *string `json:"coalescingType,omitempty" tf:"coalescing_type,omitempty"`

	// Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.
	// * `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option.
	// * `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts.
	// * `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.
	// +kubebuilder:validation:Optional
	NrCount *int64 `json:"nrCount,omitempty" tf:"nr_count,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`
}

type NvgreSettingsObservation struct {
}

type NvgreSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

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

type RoceSettingsObservation struct {
}

type RoceSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The Class of Service for RoCE on this virtual interface.
	// * `5` - RDMA CoS Service Level 5.
	// * `1` - RDMA CoS Service Level 1.
	// * `2` - RDMA CoS Service Level 2.
	// * `4` - RDMA CoS Service Level 4.
	// * `6` - RDMA CoS Service Level 6.
	// +kubebuilder:validation:Optional
	ClassOfService *int64 `json:"classOfService,omitempty" tf:"class_of_service,omitempty"`

	// If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The number of memory regions per adapter. Recommended value = integer power of 2.
	// +kubebuilder:validation:Optional
	MemoryRegions *int64 `json:"memoryRegions,omitempty" tf:"memory_regions,omitempty"`

	// Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters.
	// * `1` - RDMA over Converged Ethernet Protocol Version 1.
	// * `2` - RDMA over Converged Ethernet Protocol Version 2.
	// +kubebuilder:validation:Optional
	NrVersion *int64 `json:"nrVersion,omitempty" tf:"nr_version,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The number of queue pairs per adapter. Recommended value = integer power of 2.
	// +kubebuilder:validation:Optional
	QueuePairs *int64 `json:"queuePairs,omitempty" tf:"queue_pairs,omitempty"`

	// The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance.
	// +kubebuilder:validation:Optional
	ResourceGroups *int64 `json:"resourceGroups,omitempty" tf:"resource_groups,omitempty"`
}

type RssHashSettingsObservation struct {
}

type RssHashSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// When enabled, the IPv4 address is used for traffic distribution.
	// +kubebuilder:validation:Optional
	IPv4Hash *bool `json:"ipv4Hash,omitempty" tf:"ipv4_hash,omitempty"`

	// When enabled, the IPv6 extensions are used for traffic distribution.
	// +kubebuilder:validation:Optional
	IPv6ExtHash *bool `json:"ipv6ExtHash,omitempty" tf:"ipv6_ext_hash,omitempty"`

	// When enabled, the IPv6 address is used for traffic distribution.
	// +kubebuilder:validation:Optional
	IPv6Hash *bool `json:"ipv6Hash,omitempty" tf:"ipv6_hash,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// When enabled, both the IPv4 address and TCP port number are used for traffic distribution.
	// +kubebuilder:validation:Optional
	TCPIPv4Hash *bool `json:"tcpIpv4Hash,omitempty" tf:"tcp_ipv4_hash,omitempty"`

	// When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution.
	// +kubebuilder:validation:Optional
	TCPIPv6ExtHash *bool `json:"tcpIpv6ExtHash,omitempty" tf:"tcp_ipv6_ext_hash,omitempty"`

	// When enabled, both the IPv6 address and TCP port number are used for traffic distribution.
	// +kubebuilder:validation:Optional
	TCPIPv6Hash *bool `json:"tcpIpv6Hash,omitempty" tf:"tcp_ipv6_hash,omitempty"`

	// When enabled, both the IPv4 address and UDP port number are used for traffic distribution.
	// +kubebuilder:validation:Optional
	UDPIPv4Hash *bool `json:"udpIpv4Hash,omitempty" tf:"udp_ipv4_hash,omitempty"`

	// When enabled, both the IPv6 address and UDP port number are used for traffic distribution.
	// +kubebuilder:validation:Optional
	UDPIPv6Hash *bool `json:"udpIpv6Hash,omitempty" tf:"udp_ipv6_hash,omitempty"`
}

type RxQueueSettingsObservation struct {
}

type RxQueueSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The number of queue resources to allocate.
	// +kubebuilder:validation:Optional
	NrCount *int64 `json:"nrCount,omitempty" tf:"nr_count,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The number of descriptors in each queue.
	// +kubebuilder:validation:Optional
	RingSize *int64 `json:"ringSize,omitempty" tf:"ring_size,omitempty"`
}

type TCPOffloadSettingsObservation struct {
}

type TCPOffloadSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Enables the reassembly of segmented packets in hardware before sending them to the CPU.
	// +kubebuilder:validation:Optional
	LargeReceive *bool `json:"largeReceive,omitempty" tf:"large_receive,omitempty"`

	// Enables the CPU to send large packets to the hardware for segmentation.
	// +kubebuilder:validation:Optional
	LargeSend *bool `json:"largeSend,omitempty" tf:"large_send,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// When enabled, the CPU sends all packet checksums to the hardware for validation.
	// +kubebuilder:validation:Optional
	RxChecksum *bool `json:"rxChecksum,omitempty" tf:"rx_checksum,omitempty"`

	// When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.
	// +kubebuilder:validation:Optional
	TxChecksum *bool `json:"txChecksum,omitempty" tf:"tx_checksum,omitempty"`
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

type TxQueueSettingsObservation struct {
}

type TxQueueSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// The number of queue resources to allocate.
	// +kubebuilder:validation:Optional
	NrCount *int64 `json:"nrCount,omitempty" tf:"nr_count,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The number of descriptors in each queue.
	// +kubebuilder:validation:Optional
	RingSize *int64 `json:"ringSize,omitempty" tf:"ring_size,omitempty"`
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

type VxlanSettingsObservation struct {
}

type VxlanSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`
}

// EthAdapterPolicySpec defines the desired state of EthAdapterPolicy
type EthAdapterPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EthAdapterPolicyParameters `json:"forProvider"`
}

// EthAdapterPolicyStatus defines the observed state of EthAdapterPolicy.
type EthAdapterPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EthAdapterPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EthAdapterPolicy is the Schema for the EthAdapterPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,intersightjet}
type EthAdapterPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EthAdapterPolicySpec   `json:"spec"`
	Status            EthAdapterPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EthAdapterPolicyList contains a list of EthAdapterPolicys
type EthAdapterPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EthAdapterPolicy `json:"items"`
}

// Repository type metadata.
var (
	EthAdapterPolicy_Kind             = "EthAdapterPolicy"
	EthAdapterPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EthAdapterPolicy_Kind}.String()
	EthAdapterPolicy_KindAPIVersion   = EthAdapterPolicy_Kind + "." + CRDGroupVersion.String()
	EthAdapterPolicy_GroupVersionKind = CRDGroupVersion.WithKind(EthAdapterPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&EthAdapterPolicy{}, &EthAdapterPolicyList{})
}