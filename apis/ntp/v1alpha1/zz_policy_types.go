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

type ApplianceAccountObservation struct {
}

type ApplianceAccountParameters struct {

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

type AuthenticatedNtpServersObservation struct {
}

type AuthenticatedNtpServersParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties *string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// This property is used as a discriminator to identify the type of the payload
	// when marshaling and unmarshaling data.
	// +kubebuilder:validation:Optional
	ClassID *string `json:"classId,omitempty" tf:"class_id,omitempty"`

	// Type of symmetric key to use for this server.
	// * `SHA1` - Key type used by the authentication is SHA1.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// The fully-qualified name of the instantiated, concrete type.
	// The value should be the same as the 'ClassId' property.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// Server hostname or IP address.
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The key ID is a positive integer that identifies a cryptographic key used to authenticate NTP messages.
	// +kubebuilder:validation:Optional
	SymKeyID *int64 `json:"symKeyId,omitempty" tf:"sym_key_id,omitempty"`

	// The value of the symmetric key.
	// +kubebuilder:validation:Optional
	SymKeyValue *string `json:"symKeyValue,omitempty" tf:"sym_key_value,omitempty"`
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

	// A reference to a iamAccount resource.
	// When the $expand query parameter is specified, the referenced resource is returned inline.
	// +kubebuilder:validation:Optional
	ApplianceAccount []ApplianceAccountParameters `json:"applianceAccount,omitempty" tf:"appliance_account,omitempty"`

	// +kubebuilder:validation:Optional
	AuthenticatedNtpServers []AuthenticatedNtpServersParameters `json:"authenticatedNtpServers,omitempty" tf:"authenticated_ntp_servers,omitempty"`

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

	// State of NTP service on the endpoint.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The time when this managed object was last modified.
	// +kubebuilder:validation:Optional
	ModTime *string `json:"modTime,omitempty" tf:"mod_time,omitempty"`

	// The unique identifier of this Managed Object instance.
	// +kubebuilder:validation:Optional
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`

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

	// An array of relationships to policyAbstractConfigProfile resources.
	// +kubebuilder:validation:Optional
	Profiles []ProfilesParameters `json:"profiles,omitempty" tf:"profiles,omitempty"`

	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.
	// Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	// +kubebuilder:validation:Optional
	SharedScope *string `json:"sharedScope,omitempty" tf:"shared_scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// Timezone of services on the endpoint.
	// * `Pacific/Niue` -
	// * `Pacific/Pago_Pago` -
	// * `Pacific/Honolulu` -
	// * `Pacific/Rarotonga` -
	// * `Pacific/Tahiti` -
	// * `Pacific/Marquesas` -
	// * `America/Anchorage` -
	// * `Pacific/Gambier` -
	// * `America/Los_Angeles` -
	// * `America/Tijuana` -
	// * `America/Vancouver` -
	// * `America/Whitehorse` -
	// * `Pacific/Pitcairn` -
	// * `America/Dawson_Creek` -
	// * `America/Denver` -
	// * `America/Edmonton` -
	// * `America/Hermosillo` -
	// * `America/Mazatlan` -
	// * `America/Phoenix` -
	// * `America/Yellowknife` -
	// * `America/Belize` -
	// * `America/Chicago` -
	// * `America/Costa_Rica` -
	// * `America/El_Salvador` -
	// * `America/Guatemala` -
	// * `America/Managua` -
	// * `America/Mexico_City` -
	// * `America/Regina` -
	// * `America/Tegucigalpa` -
	// * `America/Winnipeg` -
	// * `Pacific/Galapagos` -
	// * `America/Bogota` -
	// * `America/Cancun` -
	// * `America/Cayman` -
	// * `America/Guayaquil` -
	// * `America/Havana` -
	// * `America/Iqaluit` -
	// * `America/Jamaica` -
	// * `America/Lima` -
	// * `America/Nassau` -
	// * `America/New_York` -
	// * `America/Panama` -
	// * `America/Port-au-Prince` -
	// * `America/Rio_Branco` -
	// * `America/Toronto` -
	// * `Pacific/Easter` -
	// * `America/Caracas` -
	// * `America/Asuncion` -
	// * `America/Barbados` -
	// * `America/Boa_Vista` -
	// * `America/Campo_Grande` -
	// * `America/Cuiaba` -
	// * `America/Curacao` -
	// * `America/Grand_Turk` -
	// * `America/Guyana` -
	// * `America/Halifax` -
	// * `America/La_Paz` -
	// * `America/Manaus` -
	// * `America/Martinique` -
	// * `America/Port_of_Spain` -
	// * `America/Porto_Velho` -
	// * `America/Puerto_Rico` -
	// * `America/Santo_Domingo` -
	// * `America/Thule` -
	// * `Atlantic/Bermuda` -
	// * `America/St_Johns` -
	// * `America/Araguaina` -
	// * `America/Argentina/Buenos_Aires` -
	// * `America/Bahia` -
	// * `America/Belem` -
	// * `America/Cayenne` -
	// * `America/Fortaleza` -
	// * `America/Godthab` -
	// * `America/Maceio` -
	// * `America/Miquelon` -
	// * `America/Montevideo` -
	// * `America/Paramaribo` -
	// * `America/Recife` -
	// * `America/Santiago` -
	// * `America/Sao_Paulo` -
	// * `Antarctica/Palmer` -
	// * `Antarctica/Rothera` -
	// * `Atlantic/Stanley` -
	// * `America/Noronha` -
	// * `Atlantic/South_Georgia` -
	// * `America/Scoresbysund` -
	// * `Atlantic/Azores` -
	// * `Atlantic/Cape_Verde` -
	// * `Africa/Abidjan` -
	// * `Africa/Accra` -
	// * `Africa/Bissau` -
	// * `Africa/Casablanca` -
	// * `Africa/El_Aaiun` -
	// * `Africa/Monrovia` -
	// * `America/Danmarkshavn` -
	// * `Atlantic/Canary` -
	// * `Atlantic/Faroe` -
	// * `Atlantic/Reykjavik` -
	// * `Etc/GMT` -
	// * `Europe/Dublin` -
	// * `Europe/Lisbon` -
	// * `Europe/London` -
	// * `Africa/Algiers` -
	// * `Africa/Ceuta` -
	// * `Africa/Lagos` -
	// * `Africa/Ndjamena` -
	// * `Africa/Tunis` -
	// * `Africa/Windhoek` -
	// * `Europe/Amsterdam` -
	// * `Europe/Andorra` -
	// * `Europe/Belgrade` -
	// * `Europe/Berlin` -
	// * `Europe/Brussels` -
	// * `Europe/Budapest` -
	// * `Europe/Copenhagen` -
	// * `Europe/Gibraltar` -
	// * `Europe/Luxembourg` -
	// * `Europe/Madrid` -
	// * `Europe/Malta` -
	// * `Europe/Monaco` -
	// * `Europe/Oslo` -
	// * `Europe/Paris` -
	// * `Europe/Prague` -
	// * `Europe/Rome` -
	// * `Europe/Stockholm` -
	// * `Europe/Tirane` -
	// * `Europe/Vienna` -
	// * `Europe/Warsaw` -
	// * `Europe/Zurich` -
	// * `Africa/Cairo` -
	// * `Africa/Johannesburg` -
	// * `Africa/Maputo` -
	// * `Africa/Tripoli` -
	// * `Asia/Amman` -
	// * `Asia/Beirut` -
	// * `Asia/Damascus` -
	// * `Asia/Gaza` -
	// * `Asia/Jerusalem` -
	// * `Asia/Nicosia` -
	// * `Europe/Athens` -
	// * `Europe/Bucharest` -
	// * `Europe/Chisinau` -
	// * `Europe/Helsinki` -
	// * `Europe/Istanbul` -
	// * `Europe/Kaliningrad` -
	// * `Europe/Kiev` -
	// * `Europe/Riga` -
	// * `Europe/Sofia` -
	// * `Europe/Tallinn` -
	// * `Europe/Vilnius` -
	// * `Africa/Khartoum` -
	// * `Africa/Nairobi` -
	// * `Antarctica/Syowa` -
	// * `Asia/Baghdad` -
	// * `Asia/Qatar` -
	// * `Asia/Riyadh` -
	// * `Europe/Minsk` -
	// * `Europe/Moscow` -
	// * `Asia/Tehran` -
	// * `Asia/Baku` -
	// * `Asia/Dubai` -
	// * `Asia/Tbilisi` -
	// * `Asia/Yerevan` -
	// * `Europe/Samara` -
	// * `Indian/Mahe` -
	// * `Indian/Mauritius` -
	// * `Indian/Reunion` -
	// * `Asia/Kabul` -
	// * `Antarctica/Mawson` -
	// * `Asia/Aqtau` -
	// * `Asia/Aqtobe` -
	// * `Asia/Ashgabat` -
	// * `Asia/Dushanbe` -
	// * `Asia/Karachi` -
	// * `Asia/Tashkent` -
	// * `Asia/Yekaterinburg` -
	// * `Indian/Kerguelen` -
	// * `Indian/Maldives` -
	// * `Asia/Calcutta` -
	// * `Asia/Kolkata` -
	// * `Asia/Colombo` -
	// * `Asia/Katmandu` -
	// * `Antarctica/Vostok` -
	// * `Asia/Almaty` -
	// * `Asia/Bishkek` -
	// * `Asia/Dhaka` -
	// * `Asia/Omsk` -
	// * `Asia/Thimphu` -
	// * `Indian/Chagos` -
	// * `Asia/Rangoon` -
	// * `Indian/Cocos` -
	// * `Antarctica/Davis` -
	// * `Asia/Bangkok` -
	// * `Asia/Hovd` -
	// * `Asia/Jakarta` -
	// * `Asia/Krasnoyarsk` -
	// * `Asia/Saigon` -
	// * `Indian/Christmas` -
	// * `Antarctica/Casey` -
	// * `Asia/Brunei` -
	// * `Asia/Choibalsan` -
	// * `Asia/Hong_Kong` -
	// * `Asia/Irkutsk` -
	// * `Asia/Kuala_Lumpur` -
	// * `Asia/Macau` -
	// * `Asia/Makassar` -
	// * `Asia/Manila` -
	// * `Asia/Shanghai` -
	// * `Asia/Singapore` -
	// * `Asia/Taipei` -
	// * `Asia/Ulaanbaatar` -
	// * `Australia/Perth` -
	// * `Asia/Pyongyang` -
	// * `Asia/Dili` -
	// * `Asia/Jayapura` -
	// * `Asia/Seoul` -
	// * `Asia/Tokyo` -
	// * `Asia/Yakutsk` -
	// * `Pacific/Palau` -
	// * `Australia/Adelaide` -
	// * `Australia/Darwin` -
	// * `Antarctica/DumontDUrville` -
	// * `Asia/Magadan` -
	// * `Asia/Vladivostok` -
	// * `Australia/Brisbane` -
	// * `Australia/Hobart` -
	// * `Australia/Sydney` -
	// * `Pacific/Chuuk` -
	// * `Pacific/Guam` -
	// * `Pacific/Port_Moresby` -
	// * `Pacific/Efate` -
	// * `Pacific/Guadalcanal` -
	// * `Pacific/Kosrae` -
	// * `Pacific/Norfolk` -
	// * `Pacific/Noumea` -
	// * `Pacific/Pohnpei` -
	// * `Asia/Kamchatka` -
	// * `Pacific/Auckland` -
	// * `Pacific/Fiji` -
	// * `Pacific/Funafuti` -
	// * `Pacific/Kwajalein` -
	// * `Pacific/Majuro` -
	// * `Pacific/Nauru` -
	// * `Pacific/Tarawa` -
	// * `Pacific/Wake` -
	// * `Pacific/Wallis` -
	// * `Pacific/Apia` -
	// * `Pacific/Enderbury` -
	// * `Pacific/Fakaofo` -
	// * `Pacific/Tongatapu` -
	// * `Pacific/Kiritimati` -
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

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
