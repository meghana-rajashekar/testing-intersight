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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	retentionpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/aaa/retentionpolicy"
	policy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/access/policy"
	configpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/adapter/configpolicy"
	autormapolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/autormapolicy"
	backup "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/backup"
	backuppolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/backuppolicy"
	dataexportpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/dataexportpolicy"
	deviceclaim "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/deviceclaim"
	diagsetting "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/diagsetting"
	remotefileimport "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/remotefileimport"
	restore "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/appliance/restore"
	target "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/asset/target"
	policybios "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/bios/policy"
	precisionpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/boot/precisionpolicy"
	export "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/bulk/export"
	request "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/bulk/request"
	adapterunitdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/adapterunitdescriptor"
	chassisdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/chassisdescriptor"
	chassismanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/chassismanufacturingdef"
	cimcfirmwaredescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/cimcfirmwaredescriptor"
	equipmentphysicaldef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/equipmentphysicaldef"
	equipmentslotarray "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/equipmentslotarray"
	fanmoduledescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/fanmoduledescriptor"
	fanmodulemanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/fanmodulemanufacturingdef"
	iocardcapabilitydef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/iocardcapabilitydef"
	iocarddescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/iocarddescriptor"
	iocardmanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/iocardmanufacturingdef"
	portgroupaggregationdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/portgroupaggregationdef"
	psudescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/psudescriptor"
	psumanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/psumanufacturingdef"
	servermodelscapabilitydef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/servermodelscapabilitydef"
	serverschemadescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/serverschemadescriptor"
	siocmodulecapabilitydef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/siocmodulecapabilitydef"
	siocmoduledescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/siocmoduledescriptor"
	siocmodulemanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/siocmodulemanufacturingdef"
	switchcapability "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/switchcapability"
	switchdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/switchdescriptor"
	switchmanufacturingdef "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/capability/switchmanufacturingdef"
	policycertificatemanagement "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/certificatemanagement/policy"
	configimport "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/chassis/configimport"
	profile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/chassis/profile"
	httpproxypolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/comm/httpproxypolicy"
	connectorpackupgrade "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/connectorpack/connectorpackupgrade"
	healthcheckdefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/convergedinfra/healthcheckdefinition"
	customresource "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/crd/customresource"
	policydeviceconnector "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/deviceconnector/policy"
	authorization "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/externalsite/authorization"
	appliancepcrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/appliancepcrole"
	appliancerole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/appliancerole"
	ethnetworkcontrolpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/ethnetworkcontrolpolicy"
	ethnetworkgrouppolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/ethnetworkgrouppolicy"
	ethnetworkpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/ethnetworkpolicy"
	fcnetworkpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcnetworkpolicy"
	fcoeuplinkpcrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcoeuplinkpcrole"
	fcoeuplinkrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcoeuplinkrole"
	fcstoragerole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcstoragerole"
	fcuplinkpcrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcuplinkpcrole"
	fcuplinkrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/fcuplinkrole"
	flowcontrolpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/flowcontrolpolicy"
	linkaggregationpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/linkaggregationpolicy"
	linkcontrolpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/linkcontrolpolicy"
	multicastpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/multicastpolicy"
	pcoperation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/pcoperation"
	portmode "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/portmode"
	portoperation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/portoperation"
	portpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/portpolicy"
	serverrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/serverrole"
	switchclusterprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/switchclusterprofile"
	switchcontrolpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/switchcontrolpolicy"
	switchprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/switchprofile"
	systemqospolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/systemqospolicy"
	uplinkpcrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/uplinkpcrole"
	uplinkrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/uplinkrole"
	vlan "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/vlan"
	vsan "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fabric/vsan"
	pool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/fcpool/pool"
	biosdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/biosdescriptor"
	boardcontrollerdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/boardcontrollerdescriptor"
	chassisupgrade "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/chassisupgrade"
	cimcdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/cimcdescriptor"
	dimmdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/dimmdescriptor"
	distributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/distributable"
	drivedescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/drivedescriptor"
	driverdistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/driverdistributable"
	eula "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/eula"
	gpudescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/gpudescriptor"
	hbadescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/hbadescriptor"
	iomdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/iomdescriptor"
	mswitchdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/mswitchdescriptor"
	nxosdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/nxosdescriptor"
	pciedescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/pciedescriptor"
	psudescriptorfirmware "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/psudescriptor"
	sasexpanderdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/sasexpanderdescriptor"
	serverconfigurationutilitydistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/serverconfigurationutilitydistributable"
	storagecontrollerdescriptor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/storagecontrollerdescriptor"
	switchupgrade "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/switchupgrade"
	unsupportedversionupgrade "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/unsupportedversionupgrade"
	upgrade "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/firmware/upgrade"
	hyperflexsoftwarecompatibilityinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hcl/hyperflexsoftwarecompatibilityinfo"
	appcatalog "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/appcatalog"
	autosupportpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/autosupportpolicy"
	capabilityinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/capabilityinfo"
	clusterbackuppolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterbackuppolicy"
	clusterbackuppolicydeployment "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterbackuppolicydeployment"
	clusternetworkpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusternetworkpolicy"
	clusterprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterprofile"
	clusterreplicationnetworkpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterreplicationnetworkpolicy"
	clusterreplicationnetworkpolicydeployment "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterreplicationnetworkpolicydeployment"
	clusterstoragepolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/clusterstoragepolicy"
	extfcstoragepolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/extfcstoragepolicy"
	extiscsistoragepolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/extiscsistoragepolicy"
	featurelimitexternal "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/featurelimitexternal"
	featurelimitinternal "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/featurelimitinternal"
	healthcheckdefinitionhyperflex "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/healthcheckdefinition"
	healthcheckpackagechecksum "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/healthcheckpackagechecksum"
	hxdpversion "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/hxdpversion"
	localcredentialpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/localcredentialpolicy"
	nodeconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/nodeconfigpolicy"
	nodeprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/nodeprofile"
	proxysettingpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/proxysettingpolicy"
	serverfirmwareversion "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/serverfirmwareversion"
	serverfirmwareversionentry "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/serverfirmwareversionentry"
	servermodel "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/servermodel"
	serviceauthtoken "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/serviceauthtoken"
	softwaredistributioncomponent "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/softwaredistributioncomponent"
	softwaredistributionentry "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/softwaredistributionentry"
	softwaredistributionversion "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/softwaredistributionversion"
	softwareversionpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/softwareversionpolicy"
	sysconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/sysconfigpolicy"
	ucsmconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/ucsmconfigpolicy"
	vcenterconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/vcenterconfigpolicy"
	vmimportoperation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/vmimportoperation"
	vmrestoreoperation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/hyperflex/vmrestoreoperation"
	account "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/account"
	accountexperience "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/accountexperience"
	apikey "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/apikey"
	appregistration "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/appregistration"
	certificate "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/certificate"
	certificaterequest "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/certificaterequest"
	endpointuser "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/endpointuser"
	endpointuserpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/endpointuserpolicy"
	endpointuserrole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/endpointuserrole"
	idp "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/idp"
	ipaccessmanagement "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/ipaccessmanagement"
	ipaddress "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/ipaddress"
	ldapgroup "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/ldapgroup"
	ldappolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/ldappolicy"
	ldapprovider "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/ldapprovider"
	permission "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/permission"
	privatekeyspec "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/privatekeyspec"
	qualifier "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/qualifier"
	resourceroles "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/resourceroles"
	sessionlimits "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/sessionlimits"
	trustpoint "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/trustpoint"
	user "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/user"
	usergroup "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iam/usergroup"
	policyipmioverlan "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/ipmioverlan/policy"
	poolippool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/ippool/pool"
	pooliqnpool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/iqnpool/pool"
	acicniapic "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/acicniapic"
	acicniprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/acicniprofile"
	acicnitenantclusterallocation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/acicnitenantclusterallocation"
	addondefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/addondefinition"
	addonpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/addonpolicy"
	addonrepository "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/addonrepository"
	baremetalnodeprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/baremetalnodeprofile"
	cluster "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/cluster"
	clusteraddonprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/clusteraddonprofile"
	clusterprofilekubernetes "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/clusterprofile"
	containerruntimepolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/containerruntimepolicy"
	networkpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/networkpolicy"
	nodegroupprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/nodegroupprofile"
	sysconfigpolicykubernetes "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/sysconfigpolicy"
	trustedregistriespolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/trustedregistriespolicy"
	version "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/version"
	versionpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/versionpolicy"
	virtualmachineinfraconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/virtualmachineinfraconfigpolicy"
	virtualmachineinfrastructureprovider "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/virtualmachineinfrastructureprovider"
	virtualmachineinstancetype "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/virtualmachineinstancetype"
	virtualmachinenodeprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kubernetes/virtualmachinenodeprofile"
	policykvm "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kvm/policy"
	session "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kvm/session"
	tunnel "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/kvm/tunnel"
	ikslicensecount "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/license/ikslicensecount"
	iwolicensecount "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/license/iwolicensecount"
	licenseinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/license/licenseinfo"
	licensereservationop "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/license/licensereservationop"
	poolmacpool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/macpool/pool"
	persistentmemorypolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/memory/persistentmemorypolicy"
	policynetworkconfig "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/networkconfig/policy"
	accountsubscription "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/notification/accountsubscription"
	policyntp "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/ntp/policy"
	authorizationoauth "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/oauth/authorization"
	deployment "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/oprs/deployment"
	synctargetlistmessage "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/oprs/synctargetlistmessage"
	organization "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/organization/organization"
	bulkinstallinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/os/bulkinstallinfo"
	configurationfile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/os/configurationfile"
	install "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/os/install"
	policypower "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/power/policy"
	providerconfig "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/providerconfig"
	backupconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/recovery/backupconfigpolicy"
	backupprofile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/recovery/backupprofile"
	ondemandbackup "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/recovery/ondemandbackup"
	restorerecovery "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/recovery/restore"
	scheduleconfigpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/recovery/scheduleconfigpolicy"
	group "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/resource/group"
	reservation "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/resource/reservation"
	poolresourcepool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/resourcepool/pool"
	policysdcard "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sdcard/policy"
	profilesdwan "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sdwan/profile"
	routernode "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sdwan/routernode"
	routerpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sdwan/routerpolicy"
	vmanageaccountpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sdwan/vmanageaccountpolicy"
	configimportserver "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/server/configimport"
	profileserver "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/server/profile"
	profiletemplate "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/server/profiletemplate"
	policysmtp "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/smtp/policy"
	policysnmp "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/snmp/policy"
	appliancedistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/appliancedistributable"
	hclmeta "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/hclmeta"
	hyperflexbundledistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/hyperflexbundledistributable"
	hyperflexdistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/hyperflexdistributable"
	releasemeta "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/releasemeta"
	solutiondistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/solutiondistributable"
	ucsdbundledistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/ucsdbundledistributable"
	ucsddistributable "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/software/ucsddistributable"
	authorizationsoftwarerepository "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/authorization"
	categorymapper "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/categorymapper"
	categorymappermodel "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/categorymappermodel"
	categorysupportconstraint "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/categorysupportconstraint"
	operatingsystemfile "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/operatingsystemfile"
	release "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/softwarerepository/release"
	policysol "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/sol/policy"
	policyssh "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/ssh/policy"
	drivegroup "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/storage/drivegroup"
	storagepolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/storage/storagepolicy"
	policysyslog "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/syslog/policy"
	advisorycount "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/tam/advisorycount"
	advisorydefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/tam/advisorydefinition"
	advisoryinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/tam/advisoryinfo"
	advisoryinstance "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/tam/advisoryinstance"
	securityadvisory "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/tam/securityadvisory"
	collectioncontrolpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/techsupportmanagement/collectioncontrolpolicy"
	techsupportbundle "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/techsupportmanagement/techsupportbundle"
	executor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/terraform/executor"
	policythermal "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/thermal/policy"
	pooluuidpool "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/uuidpool/pool"
	ciscohypervisormanager "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/ciscohypervisormanager"
	esxiconsole "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/esxiconsole"
	iwedatacenter "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/iwedatacenter"
	virtualdisk "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/virtualdisk"
	virtualmachine "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/virtualmachine"
	virtualnetwork "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/virtualization/virtualnetwork"
	policyvmedia "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vmedia/policy"
	console "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vmrc/console"
	consolevnc "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnc/console"
	ethadapterpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/ethadapterpolicy"
	ethif "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/ethif"
	ethnetworkpolicyvnic "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/ethnetworkpolicy"
	ethqospolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/ethqospolicy"
	fcadapterpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/fcadapterpolicy"
	fcif "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/fcif"
	fcnetworkpolicyvnic "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/fcnetworkpolicy"
	fcqospolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/fcqospolicy"
	iscsiadapterpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/iscsiadapterpolicy"
	iscsibootpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/iscsibootpolicy"
	iscsistatictargetpolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/iscsistatictargetpolicy"
	lanconnectivitypolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/lanconnectivitypolicy"
	sanconnectivitypolicy "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vnic/sanconnectivitypolicy"
	vrf "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/vrf/vrf"
	ansiblebatchexecutor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/ansiblebatchexecutor"
	batchapiexecutor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/batchapiexecutor"
	customdatatypedefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/customdatatypedefinition"
	errorresponsehandler "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/errorresponsehandler"
	rollbackworkflow "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/rollbackworkflow"
	solutionactiondefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/solutionactiondefinition"
	solutionactioninstance "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/solutionactioninstance"
	solutiondefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/solutiondefinition"
	solutioninstance "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/solutioninstance"
	solutionoutput "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/solutionoutput"
	sshbatchexecutor "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/sshbatchexecutor"
	taskdefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/taskdefinition"
	workflowdefinition "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/workflowdefinition"
	workflowinfo "github.com/crossplane-contrib/provider-jet-intersight/internal/controller/workflow/workflowinfo"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		retentionpolicy.Setup,
		policy.Setup,
		configpolicy.Setup,
		autormapolicy.Setup,
		backup.Setup,
		backuppolicy.Setup,
		dataexportpolicy.Setup,
		deviceclaim.Setup,
		diagsetting.Setup,
		remotefileimport.Setup,
		restore.Setup,
		target.Setup,
		policybios.Setup,
		precisionpolicy.Setup,
		export.Setup,
		request.Setup,
		adapterunitdescriptor.Setup,
		chassisdescriptor.Setup,
		chassismanufacturingdef.Setup,
		cimcfirmwaredescriptor.Setup,
		equipmentphysicaldef.Setup,
		equipmentslotarray.Setup,
		fanmoduledescriptor.Setup,
		fanmodulemanufacturingdef.Setup,
		iocardcapabilitydef.Setup,
		iocarddescriptor.Setup,
		iocardmanufacturingdef.Setup,
		portgroupaggregationdef.Setup,
		psudescriptor.Setup,
		psumanufacturingdef.Setup,
		servermodelscapabilitydef.Setup,
		serverschemadescriptor.Setup,
		siocmodulecapabilitydef.Setup,
		siocmoduledescriptor.Setup,
		siocmodulemanufacturingdef.Setup,
		switchcapability.Setup,
		switchdescriptor.Setup,
		switchmanufacturingdef.Setup,
		policycertificatemanagement.Setup,
		configimport.Setup,
		profile.Setup,
		httpproxypolicy.Setup,
		connectorpackupgrade.Setup,
		healthcheckdefinition.Setup,
		customresource.Setup,
		policydeviceconnector.Setup,
		authorization.Setup,
		appliancepcrole.Setup,
		appliancerole.Setup,
		ethnetworkcontrolpolicy.Setup,
		ethnetworkgrouppolicy.Setup,
		ethnetworkpolicy.Setup,
		fcnetworkpolicy.Setup,
		fcoeuplinkpcrole.Setup,
		fcoeuplinkrole.Setup,
		fcstoragerole.Setup,
		fcuplinkpcrole.Setup,
		fcuplinkrole.Setup,
		flowcontrolpolicy.Setup,
		linkaggregationpolicy.Setup,
		linkcontrolpolicy.Setup,
		multicastpolicy.Setup,
		pcoperation.Setup,
		portmode.Setup,
		portoperation.Setup,
		portpolicy.Setup,
		serverrole.Setup,
		switchclusterprofile.Setup,
		switchcontrolpolicy.Setup,
		switchprofile.Setup,
		systemqospolicy.Setup,
		uplinkpcrole.Setup,
		uplinkrole.Setup,
		vlan.Setup,
		vsan.Setup,
		pool.Setup,
		biosdescriptor.Setup,
		boardcontrollerdescriptor.Setup,
		chassisupgrade.Setup,
		cimcdescriptor.Setup,
		dimmdescriptor.Setup,
		distributable.Setup,
		drivedescriptor.Setup,
		driverdistributable.Setup,
		eula.Setup,
		gpudescriptor.Setup,
		hbadescriptor.Setup,
		iomdescriptor.Setup,
		mswitchdescriptor.Setup,
		nxosdescriptor.Setup,
		pciedescriptor.Setup,
		psudescriptorfirmware.Setup,
		sasexpanderdescriptor.Setup,
		serverconfigurationutilitydistributable.Setup,
		storagecontrollerdescriptor.Setup,
		switchupgrade.Setup,
		unsupportedversionupgrade.Setup,
		upgrade.Setup,
		hyperflexsoftwarecompatibilityinfo.Setup,
		appcatalog.Setup,
		autosupportpolicy.Setup,
		capabilityinfo.Setup,
		clusterbackuppolicy.Setup,
		clusterbackuppolicydeployment.Setup,
		clusternetworkpolicy.Setup,
		clusterprofile.Setup,
		clusterreplicationnetworkpolicy.Setup,
		clusterreplicationnetworkpolicydeployment.Setup,
		clusterstoragepolicy.Setup,
		extfcstoragepolicy.Setup,
		extiscsistoragepolicy.Setup,
		featurelimitexternal.Setup,
		featurelimitinternal.Setup,
		healthcheckdefinitionhyperflex.Setup,
		healthcheckpackagechecksum.Setup,
		hxdpversion.Setup,
		localcredentialpolicy.Setup,
		nodeconfigpolicy.Setup,
		nodeprofile.Setup,
		proxysettingpolicy.Setup,
		serverfirmwareversion.Setup,
		serverfirmwareversionentry.Setup,
		servermodel.Setup,
		serviceauthtoken.Setup,
		softwaredistributioncomponent.Setup,
		softwaredistributionentry.Setup,
		softwaredistributionversion.Setup,
		softwareversionpolicy.Setup,
		sysconfigpolicy.Setup,
		ucsmconfigpolicy.Setup,
		vcenterconfigpolicy.Setup,
		vmimportoperation.Setup,
		vmrestoreoperation.Setup,
		account.Setup,
		accountexperience.Setup,
		apikey.Setup,
		appregistration.Setup,
		certificate.Setup,
		certificaterequest.Setup,
		endpointuser.Setup,
		endpointuserpolicy.Setup,
		endpointuserrole.Setup,
		idp.Setup,
		ipaccessmanagement.Setup,
		ipaddress.Setup,
		ldapgroup.Setup,
		ldappolicy.Setup,
		ldapprovider.Setup,
		permission.Setup,
		privatekeyspec.Setup,
		qualifier.Setup,
		resourceroles.Setup,
		sessionlimits.Setup,
		trustpoint.Setup,
		user.Setup,
		usergroup.Setup,
		policyipmioverlan.Setup,
		poolippool.Setup,
		pooliqnpool.Setup,
		acicniapic.Setup,
		acicniprofile.Setup,
		acicnitenantclusterallocation.Setup,
		addondefinition.Setup,
		addonpolicy.Setup,
		addonrepository.Setup,
		baremetalnodeprofile.Setup,
		cluster.Setup,
		clusteraddonprofile.Setup,
		clusterprofilekubernetes.Setup,
		containerruntimepolicy.Setup,
		networkpolicy.Setup,
		nodegroupprofile.Setup,
		sysconfigpolicykubernetes.Setup,
		trustedregistriespolicy.Setup,
		version.Setup,
		versionpolicy.Setup,
		virtualmachineinfraconfigpolicy.Setup,
		virtualmachineinfrastructureprovider.Setup,
		virtualmachineinstancetype.Setup,
		virtualmachinenodeprofile.Setup,
		policykvm.Setup,
		session.Setup,
		tunnel.Setup,
		ikslicensecount.Setup,
		iwolicensecount.Setup,
		licenseinfo.Setup,
		licensereservationop.Setup,
		poolmacpool.Setup,
		persistentmemorypolicy.Setup,
		policynetworkconfig.Setup,
		accountsubscription.Setup,
		policyntp.Setup,
		authorizationoauth.Setup,
		deployment.Setup,
		synctargetlistmessage.Setup,
		organization.Setup,
		bulkinstallinfo.Setup,
		configurationfile.Setup,
		install.Setup,
		policypower.Setup,
		providerconfig.Setup,
		backupconfigpolicy.Setup,
		backupprofile.Setup,
		ondemandbackup.Setup,
		restorerecovery.Setup,
		scheduleconfigpolicy.Setup,
		group.Setup,
		reservation.Setup,
		poolresourcepool.Setup,
		policysdcard.Setup,
		profilesdwan.Setup,
		routernode.Setup,
		routerpolicy.Setup,
		vmanageaccountpolicy.Setup,
		configimportserver.Setup,
		profileserver.Setup,
		profiletemplate.Setup,
		policysmtp.Setup,
		policysnmp.Setup,
		appliancedistributable.Setup,
		hclmeta.Setup,
		hyperflexbundledistributable.Setup,
		hyperflexdistributable.Setup,
		releasemeta.Setup,
		solutiondistributable.Setup,
		ucsdbundledistributable.Setup,
		ucsddistributable.Setup,
		authorizationsoftwarerepository.Setup,
		categorymapper.Setup,
		categorymappermodel.Setup,
		categorysupportconstraint.Setup,
		operatingsystemfile.Setup,
		release.Setup,
		policysol.Setup,
		policyssh.Setup,
		drivegroup.Setup,
		storagepolicy.Setup,
		policysyslog.Setup,
		advisorycount.Setup,
		advisorydefinition.Setup,
		advisoryinfo.Setup,
		advisoryinstance.Setup,
		securityadvisory.Setup,
		collectioncontrolpolicy.Setup,
		techsupportbundle.Setup,
		executor.Setup,
		policythermal.Setup,
		pooluuidpool.Setup,
		ciscohypervisormanager.Setup,
		esxiconsole.Setup,
		iwedatacenter.Setup,
		virtualdisk.Setup,
		virtualmachine.Setup,
		virtualnetwork.Setup,
		policyvmedia.Setup,
		console.Setup,
		consolevnc.Setup,
		ethadapterpolicy.Setup,
		ethif.Setup,
		ethnetworkpolicyvnic.Setup,
		ethqospolicy.Setup,
		fcadapterpolicy.Setup,
		fcif.Setup,
		fcnetworkpolicyvnic.Setup,
		fcqospolicy.Setup,
		iscsiadapterpolicy.Setup,
		iscsibootpolicy.Setup,
		iscsistatictargetpolicy.Setup,
		lanconnectivitypolicy.Setup,
		sanconnectivitypolicy.Setup,
		vrf.Setup,
		ansiblebatchexecutor.Setup,
		batchapiexecutor.Setup,
		customdatatypedefinition.Setup,
		errorresponsehandler.Setup,
		rollbackworkflow.Setup,
		solutionactiondefinition.Setup,
		solutionactioninstance.Setup,
		solutiondefinition.Setup,
		solutioninstance.Setup,
		solutionoutput.Setup,
		sshbatchexecutor.Setup,
		taskdefinition.Setup,
		workflowdefinition.Setup,
		workflowinfo.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
