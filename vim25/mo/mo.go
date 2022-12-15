/*
Copyright (c) 2014-2022 VMware, Inc. All Rights Reserved.

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

package mo

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type Alarm struct {
	ExtensibleManagedObject

	Info types.AlarmInfo `mo:"info" json:"info"`
}

func init() {
	t["Alarm"] = reflect.TypeOf((*Alarm)(nil)).Elem()
}

type AlarmManager struct {
	Self types.ManagedObjectReference

	DefaultExpression []types.BaseAlarmExpression `mo:"defaultExpression" json:"defaultExpression,omitempty"`
	Description       types.AlarmDescription      `mo:"description" json:"description"`
}

func (m AlarmManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["AlarmManager"] = reflect.TypeOf((*AlarmManager)(nil)).Elem()
}

type AuthorizationManager struct {
	Self types.ManagedObjectReference

	PrivilegeList []types.AuthorizationPrivilege `mo:"privilegeList" json:"privilegeList,omitempty"`
	RoleList      []types.AuthorizationRole      `mo:"roleList" json:"roleList,omitempty"`
	Description   types.AuthorizationDescription `mo:"description" json:"description"`
}

func (m AuthorizationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["AuthorizationManager"] = reflect.TypeOf((*AuthorizationManager)(nil)).Elem()
}

type CertificateManager struct {
	Self types.ManagedObjectReference
}

func (m CertificateManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["CertificateManager"] = reflect.TypeOf((*CertificateManager)(nil)).Elem()
}

type ClusterComputeResource struct {
	ComputeResource

	Configuration     types.ClusterConfigInfo                    `mo:"configuration" json:"configuration"`
	Recommendation    []types.ClusterRecommendation              `mo:"recommendation" json:"recommendation,omitempty"`
	DrsRecommendation []types.ClusterDrsRecommendation           `mo:"drsRecommendation" json:"drsRecommendation,omitempty"`
	HciConfig         *types.ClusterComputeResourceHCIConfigInfo `mo:"hciConfig" json:"hciConfig,omitempty"`
	MigrationHistory  []types.ClusterDrsMigration                `mo:"migrationHistory" json:"migrationHistory,omitempty"`
	ActionHistory     []types.ClusterActionHistory               `mo:"actionHistory" json:"actionHistory,omitempty"`
	DrsFault          []types.ClusterDrsFaults                   `mo:"drsFault" json:"drsFault,omitempty"`
}

func init() {
	t["ClusterComputeResource"] = reflect.TypeOf((*ClusterComputeResource)(nil)).Elem()
}

type ClusterEVCManager struct {
	ExtensibleManagedObject

	ManagedCluster types.ManagedObjectReference    `mo:"managedCluster" json:"managedCluster"`
	EvcState       types.ClusterEVCManagerEVCState `mo:"evcState" json:"evcState"`
}

func init() {
	t["ClusterEVCManager"] = reflect.TypeOf((*ClusterEVCManager)(nil)).Elem()
}

type ClusterProfile struct {
	Profile
}

func init() {
	t["ClusterProfile"] = reflect.TypeOf((*ClusterProfile)(nil)).Elem()
}

type ClusterProfileManager struct {
	ProfileManager
}

func init() {
	t["ClusterProfileManager"] = reflect.TypeOf((*ClusterProfileManager)(nil)).Elem()
}

type ComputeResource struct {
	ManagedEntity

	ResourcePool       *types.ManagedObjectReference       `mo:"resourcePool" json:"resourcePool,omitempty"`
	Host               []types.ManagedObjectReference      `mo:"host" json:"host,omitempty"`
	Datastore          []types.ManagedObjectReference      `mo:"datastore" json:"datastore,omitempty"`
	Network            []types.ManagedObjectReference      `mo:"network" json:"network,omitempty"`
	Summary            types.BaseComputeResourceSummary    `mo:"summary" json:"summary"`
	EnvironmentBrowser *types.ManagedObjectReference       `mo:"environmentBrowser" json:"environmentBrowser,omitempty"`
	ConfigurationEx    types.BaseComputeResourceConfigInfo `mo:"configurationEx" json:"configurationEx"`
	LifecycleManaged   *bool                               `mo:"lifecycleManaged" json:"lifecycleManaged,omitempty"`
}

func (m *ComputeResource) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["ComputeResource"] = reflect.TypeOf((*ComputeResource)(nil)).Elem()
}

type ContainerView struct {
	ManagedObjectView

	Container types.ManagedObjectReference `mo:"container" json:"container"`
	Type      []string                     `mo:"type" json:"type,omitempty"`
	Recursive bool                         `mo:"recursive" json:"recursive"`
}

func init() {
	t["ContainerView"] = reflect.TypeOf((*ContainerView)(nil)).Elem()
}

type CryptoManager struct {
	Self types.ManagedObjectReference

	Enabled bool `mo:"enabled" json:"enabled"`
}

func (m CryptoManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["CryptoManager"] = reflect.TypeOf((*CryptoManager)(nil)).Elem()
}

type CryptoManagerHost struct {
	CryptoManager
}

func init() {
	t["CryptoManagerHost"] = reflect.TypeOf((*CryptoManagerHost)(nil)).Elem()
}

type CryptoManagerHostKMS struct {
	CryptoManagerHost
}

func init() {
	t["CryptoManagerHostKMS"] = reflect.TypeOf((*CryptoManagerHostKMS)(nil)).Elem()
}

type CryptoManagerKmip struct {
	CryptoManager

	KmipServers []types.KmipClusterInfo `mo:"kmipServers" json:"kmipServers,omitempty"`
}

func init() {
	t["CryptoManagerKmip"] = reflect.TypeOf((*CryptoManagerKmip)(nil)).Elem()
}

type CustomFieldsManager struct {
	Self types.ManagedObjectReference

	Field []types.CustomFieldDef `mo:"field" json:"field,omitempty"`
}

func (m CustomFieldsManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["CustomFieldsManager"] = reflect.TypeOf((*CustomFieldsManager)(nil)).Elem()
}

type CustomizationSpecManager struct {
	Self types.ManagedObjectReference

	Info          []types.CustomizationSpecInfo `mo:"info" json:"info,omitempty"`
	EncryptionKey []byte                        `mo:"encryptionKey" json:"encryptionKey,omitempty"`
}

func (m CustomizationSpecManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["CustomizationSpecManager"] = reflect.TypeOf((*CustomizationSpecManager)(nil)).Elem()
}

type Datacenter struct {
	ManagedEntity

	VmFolder        types.ManagedObjectReference   `mo:"vmFolder" json:"vmFolder"`
	HostFolder      types.ManagedObjectReference   `mo:"hostFolder" json:"hostFolder"`
	DatastoreFolder types.ManagedObjectReference   `mo:"datastoreFolder" json:"datastoreFolder"`
	NetworkFolder   types.ManagedObjectReference   `mo:"networkFolder" json:"networkFolder"`
	Datastore       []types.ManagedObjectReference `mo:"datastore" json:"datastore,omitempty"`
	Network         []types.ManagedObjectReference `mo:"network" json:"network,omitempty"`
	Configuration   types.DatacenterConfigInfo     `mo:"configuration" json:"configuration"`
}

func (m *Datacenter) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["Datacenter"] = reflect.TypeOf((*Datacenter)(nil)).Elem()
}

type Datastore struct {
	ManagedEntity

	Info              types.BaseDatastoreInfo        `mo:"info" json:"info"`
	Summary           types.DatastoreSummary         `mo:"summary" json:"summary"`
	Host              []types.DatastoreHostMount     `mo:"host" json:"host,omitempty"`
	Vm                []types.ManagedObjectReference `mo:"vm" json:"vm,omitempty"`
	Browser           types.ManagedObjectReference   `mo:"browser" json:"browser"`
	Capability        types.DatastoreCapability      `mo:"capability" json:"capability"`
	IormConfiguration *types.StorageIORMInfo         `mo:"iormConfiguration" json:"iormConfiguration,omitempty"`
}

func (m *Datastore) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["Datastore"] = reflect.TypeOf((*Datastore)(nil)).Elem()
}

type DatastoreNamespaceManager struct {
	Self types.ManagedObjectReference
}

func (m DatastoreNamespaceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["DatastoreNamespaceManager"] = reflect.TypeOf((*DatastoreNamespaceManager)(nil)).Elem()
}

type DiagnosticManager struct {
	Self types.ManagedObjectReference
}

func (m DiagnosticManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["DiagnosticManager"] = reflect.TypeOf((*DiagnosticManager)(nil)).Elem()
}

type DistributedVirtualPortgroup struct {
	Network

	Key      string                      `mo:"key" json:"key"`
	Config   types.DVPortgroupConfigInfo `mo:"config" json:"config"`
	PortKeys []string                    `mo:"portKeys" json:"portKeys,omitempty"`
}

func init() {
	t["DistributedVirtualPortgroup"] = reflect.TypeOf((*DistributedVirtualPortgroup)(nil)).Elem()
}

type DistributedVirtualSwitch struct {
	ManagedEntity

	Uuid                string                         `mo:"uuid" json:"uuid"`
	Capability          types.DVSCapability            `mo:"capability" json:"capability"`
	Summary             types.DVSSummary               `mo:"summary" json:"summary"`
	Config              types.BaseDVSConfigInfo        `mo:"config" json:"config"`
	NetworkResourcePool []types.DVSNetworkResourcePool `mo:"networkResourcePool" json:"networkResourcePool,omitempty"`
	Portgroup           []types.ManagedObjectReference `mo:"portgroup" json:"portgroup,omitempty"`
	Runtime             *types.DVSRuntimeInfo          `mo:"runtime" json:"runtime,omitempty"`
}

func (m *DistributedVirtualSwitch) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["DistributedVirtualSwitch"] = reflect.TypeOf((*DistributedVirtualSwitch)(nil)).Elem()
}

type DistributedVirtualSwitchManager struct {
	Self types.ManagedObjectReference
}

func (m DistributedVirtualSwitchManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["DistributedVirtualSwitchManager"] = reflect.TypeOf((*DistributedVirtualSwitchManager)(nil)).Elem()
}

type EnvironmentBrowser struct {
	Self types.ManagedObjectReference

	DatastoreBrowser *types.ManagedObjectReference `mo:"datastoreBrowser" json:"datastoreBrowser,omitempty"`
}

func (m EnvironmentBrowser) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["EnvironmentBrowser"] = reflect.TypeOf((*EnvironmentBrowser)(nil)).Elem()
}

type EventHistoryCollector struct {
	HistoryCollector

	LatestPage []types.BaseEvent `mo:"latestPage" json:"latestPage,omitempty"`
}

func init() {
	t["EventHistoryCollector"] = reflect.TypeOf((*EventHistoryCollector)(nil)).Elem()
}

type EventManager struct {
	Self types.ManagedObjectReference

	Description  types.EventDescription `mo:"description" json:"description"`
	LatestEvent  types.BaseEvent        `mo:"latestEvent" json:"latestEvent,omitempty"`
	MaxCollector int32                  `mo:"maxCollector" json:"maxCollector"`
}

func (m EventManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["EventManager"] = reflect.TypeOf((*EventManager)(nil)).Elem()
}

type ExtensibleManagedObject struct {
	Self types.ManagedObjectReference

	Value          []types.BaseCustomFieldValue `mo:"value" json:"value,omitempty"`
	AvailableField []types.CustomFieldDef       `mo:"availableField" json:"availableField,omitempty"`
}

func (m ExtensibleManagedObject) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ExtensibleManagedObject"] = reflect.TypeOf((*ExtensibleManagedObject)(nil)).Elem()
}

type ExtensionManager struct {
	Self types.ManagedObjectReference

	ExtensionList []types.Extension `mo:"extensionList" json:"extensionList,omitempty"`
}

func (m ExtensionManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ExtensionManager"] = reflect.TypeOf((*ExtensionManager)(nil)).Elem()
}

type FailoverClusterConfigurator struct {
	Self types.ManagedObjectReference

	DisabledConfigureMethod []string `mo:"disabledConfigureMethod" json:"disabledConfigureMethod,omitempty"`
}

func (m FailoverClusterConfigurator) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["FailoverClusterConfigurator"] = reflect.TypeOf((*FailoverClusterConfigurator)(nil)).Elem()
}

type FailoverClusterManager struct {
	Self types.ManagedObjectReference

	DisabledClusterMethod []string `mo:"disabledClusterMethod" json:"disabledClusterMethod,omitempty"`
}

func (m FailoverClusterManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["FailoverClusterManager"] = reflect.TypeOf((*FailoverClusterManager)(nil)).Elem()
}

type FileManager struct {
	Self types.ManagedObjectReference
}

func (m FileManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["FileManager"] = reflect.TypeOf((*FileManager)(nil)).Elem()
}

type Folder struct {
	ManagedEntity

	ChildType   []string                       `mo:"childType" json:"childType,omitempty"`
	ChildEntity []types.ManagedObjectReference `mo:"childEntity" json:"childEntity,omitempty"`
	Namespace   *string                        `mo:"namespace" json:"namespace,omitempty"`
}

func (m *Folder) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["Folder"] = reflect.TypeOf((*Folder)(nil)).Elem()
}

type GuestAliasManager struct {
	Self types.ManagedObjectReference
}

func (m GuestAliasManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestAliasManager"] = reflect.TypeOf((*GuestAliasManager)(nil)).Elem()
}

type GuestAuthManager struct {
	Self types.ManagedObjectReference
}

func (m GuestAuthManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestAuthManager"] = reflect.TypeOf((*GuestAuthManager)(nil)).Elem()
}

type GuestFileManager struct {
	Self types.ManagedObjectReference
}

func (m GuestFileManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestFileManager"] = reflect.TypeOf((*GuestFileManager)(nil)).Elem()
}

type GuestOperationsManager struct {
	Self types.ManagedObjectReference

	AuthManager                 *types.ManagedObjectReference `mo:"authManager" json:"authManager,omitempty"`
	FileManager                 *types.ManagedObjectReference `mo:"fileManager" json:"fileManager,omitempty"`
	ProcessManager              *types.ManagedObjectReference `mo:"processManager" json:"processManager,omitempty"`
	GuestWindowsRegistryManager *types.ManagedObjectReference `mo:"guestWindowsRegistryManager" json:"guestWindowsRegistryManager,omitempty"`
	AliasManager                *types.ManagedObjectReference `mo:"aliasManager" json:"aliasManager,omitempty"`
}

func (m GuestOperationsManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestOperationsManager"] = reflect.TypeOf((*GuestOperationsManager)(nil)).Elem()
}

type GuestProcessManager struct {
	Self types.ManagedObjectReference
}

func (m GuestProcessManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestProcessManager"] = reflect.TypeOf((*GuestProcessManager)(nil)).Elem()
}

type GuestWindowsRegistryManager struct {
	Self types.ManagedObjectReference
}

func (m GuestWindowsRegistryManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["GuestWindowsRegistryManager"] = reflect.TypeOf((*GuestWindowsRegistryManager)(nil)).Elem()
}

type HealthUpdateManager struct {
	Self types.ManagedObjectReference
}

func (m HealthUpdateManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HealthUpdateManager"] = reflect.TypeOf((*HealthUpdateManager)(nil)).Elem()
}

type HistoryCollector struct {
	Self types.ManagedObjectReference

	Filter types.AnyType `mo:"filter" json:"filter"`
}

func (m HistoryCollector) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HistoryCollector"] = reflect.TypeOf((*HistoryCollector)(nil)).Elem()
}

type HostAccessManager struct {
	Self types.ManagedObjectReference

	LockdownMode types.HostLockdownMode `mo:"lockdownMode" json:"lockdownMode"`
}

func (m HostAccessManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostAccessManager"] = reflect.TypeOf((*HostAccessManager)(nil)).Elem()
}

type HostActiveDirectoryAuthentication struct {
	HostDirectoryStore
}

func init() {
	t["HostActiveDirectoryAuthentication"] = reflect.TypeOf((*HostActiveDirectoryAuthentication)(nil)).Elem()
}

type HostAssignableHardwareManager struct {
	Self types.ManagedObjectReference

	Binding []types.HostAssignableHardwareBinding `mo:"binding" json:"binding,omitempty"`
	Config  types.HostAssignableHardwareConfig    `mo:"config" json:"config"`
}

func (m HostAssignableHardwareManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostAssignableHardwareManager"] = reflect.TypeOf((*HostAssignableHardwareManager)(nil)).Elem()
}

type HostAuthenticationManager struct {
	Self types.ManagedObjectReference

	Info           types.HostAuthenticationManagerInfo `mo:"info" json:"info"`
	SupportedStore []types.ManagedObjectReference      `mo:"supportedStore" json:"supportedStore"`
}

func (m HostAuthenticationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostAuthenticationManager"] = reflect.TypeOf((*HostAuthenticationManager)(nil)).Elem()
}

type HostAuthenticationStore struct {
	Self types.ManagedObjectReference

	Info types.BaseHostAuthenticationStoreInfo `mo:"info" json:"info"`
}

func (m HostAuthenticationStore) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostAuthenticationStore"] = reflect.TypeOf((*HostAuthenticationStore)(nil)).Elem()
}

type HostAutoStartManager struct {
	Self types.ManagedObjectReference

	Config types.HostAutoStartManagerConfig `mo:"config" json:"config"`
}

func (m HostAutoStartManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostAutoStartManager"] = reflect.TypeOf((*HostAutoStartManager)(nil)).Elem()
}

type HostBootDeviceSystem struct {
	Self types.ManagedObjectReference
}

func (m HostBootDeviceSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostBootDeviceSystem"] = reflect.TypeOf((*HostBootDeviceSystem)(nil)).Elem()
}

type HostCacheConfigurationManager struct {
	Self types.ManagedObjectReference

	CacheConfigurationInfo []types.HostCacheConfigurationInfo `mo:"cacheConfigurationInfo" json:"cacheConfigurationInfo,omitempty"`
}

func (m HostCacheConfigurationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostCacheConfigurationManager"] = reflect.TypeOf((*HostCacheConfigurationManager)(nil)).Elem()
}

type HostCertificateManager struct {
	Self types.ManagedObjectReference

	CertificateInfo types.HostCertificateManagerCertificateInfo `mo:"certificateInfo" json:"certificateInfo"`
}

func (m HostCertificateManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostCertificateManager"] = reflect.TypeOf((*HostCertificateManager)(nil)).Elem()
}

type HostCpuSchedulerSystem struct {
	ExtensibleManagedObject

	HyperthreadInfo *types.HostHyperThreadScheduleInfo `mo:"hyperthreadInfo" json:"hyperthreadInfo,omitempty"`
}

func init() {
	t["HostCpuSchedulerSystem"] = reflect.TypeOf((*HostCpuSchedulerSystem)(nil)).Elem()
}

type HostDatastoreBrowser struct {
	Self types.ManagedObjectReference

	Datastore     []types.ManagedObjectReference `mo:"datastore" json:"datastore,omitempty"`
	SupportedType []types.BaseFileQuery          `mo:"supportedType" json:"supportedType,omitempty"`
}

func (m HostDatastoreBrowser) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostDatastoreBrowser"] = reflect.TypeOf((*HostDatastoreBrowser)(nil)).Elem()
}

type HostDatastoreSystem struct {
	Self types.ManagedObjectReference

	Datastore    []types.ManagedObjectReference        `mo:"datastore" json:"datastore,omitempty"`
	Capabilities types.HostDatastoreSystemCapabilities `mo:"capabilities" json:"capabilities"`
}

func (m HostDatastoreSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostDatastoreSystem"] = reflect.TypeOf((*HostDatastoreSystem)(nil)).Elem()
}

type HostDateTimeSystem struct {
	Self types.ManagedObjectReference

	DateTimeInfo types.HostDateTimeInfo `mo:"dateTimeInfo" json:"dateTimeInfo"`
}

func (m HostDateTimeSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostDateTimeSystem"] = reflect.TypeOf((*HostDateTimeSystem)(nil)).Elem()
}

type HostDiagnosticSystem struct {
	Self types.ManagedObjectReference

	ActivePartition *types.HostDiagnosticPartition `mo:"activePartition" json:"activePartition,omitempty"`
}

func (m HostDiagnosticSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostDiagnosticSystem"] = reflect.TypeOf((*HostDiagnosticSystem)(nil)).Elem()
}

type HostDirectoryStore struct {
	HostAuthenticationStore
}

func init() {
	t["HostDirectoryStore"] = reflect.TypeOf((*HostDirectoryStore)(nil)).Elem()
}

type HostEsxAgentHostManager struct {
	Self types.ManagedObjectReference

	ConfigInfo types.HostEsxAgentHostManagerConfigInfo `mo:"configInfo" json:"configInfo"`
}

func (m HostEsxAgentHostManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostEsxAgentHostManager"] = reflect.TypeOf((*HostEsxAgentHostManager)(nil)).Elem()
}

type HostFirewallSystem struct {
	ExtensibleManagedObject

	FirewallInfo *types.HostFirewallInfo `mo:"firewallInfo" json:"firewallInfo,omitempty"`
}

func init() {
	t["HostFirewallSystem"] = reflect.TypeOf((*HostFirewallSystem)(nil)).Elem()
}

type HostFirmwareSystem struct {
	Self types.ManagedObjectReference
}

func (m HostFirmwareSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostFirmwareSystem"] = reflect.TypeOf((*HostFirmwareSystem)(nil)).Elem()
}

type HostGraphicsManager struct {
	ExtensibleManagedObject

	GraphicsInfo           []types.HostGraphicsInfo          `mo:"graphicsInfo" json:"graphicsInfo,omitempty"`
	GraphicsConfig         *types.HostGraphicsConfig         `mo:"graphicsConfig" json:"graphicsConfig,omitempty"`
	SharedPassthruGpuTypes []string                          `mo:"sharedPassthruGpuTypes" json:"sharedPassthruGpuTypes,omitempty"`
	SharedGpuCapabilities  []types.HostSharedGpuCapabilities `mo:"sharedGpuCapabilities" json:"sharedGpuCapabilities,omitempty"`
}

func init() {
	t["HostGraphicsManager"] = reflect.TypeOf((*HostGraphicsManager)(nil)).Elem()
}

type HostHealthStatusSystem struct {
	Self types.ManagedObjectReference

	Runtime types.HealthSystemRuntime `mo:"runtime" json:"runtime"`
}

func (m HostHealthStatusSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostHealthStatusSystem"] = reflect.TypeOf((*HostHealthStatusSystem)(nil)).Elem()
}

type HostImageConfigManager struct {
	Self types.ManagedObjectReference
}

func (m HostImageConfigManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostImageConfigManager"] = reflect.TypeOf((*HostImageConfigManager)(nil)).Elem()
}

type HostKernelModuleSystem struct {
	Self types.ManagedObjectReference
}

func (m HostKernelModuleSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostKernelModuleSystem"] = reflect.TypeOf((*HostKernelModuleSystem)(nil)).Elem()
}

type HostLocalAccountManager struct {
	Self types.ManagedObjectReference
}

func (m HostLocalAccountManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostLocalAccountManager"] = reflect.TypeOf((*HostLocalAccountManager)(nil)).Elem()
}

type HostLocalAuthentication struct {
	HostAuthenticationStore
}

func init() {
	t["HostLocalAuthentication"] = reflect.TypeOf((*HostLocalAuthentication)(nil)).Elem()
}

type HostMemorySystem struct {
	ExtensibleManagedObject

	ConsoleReservationInfo        *types.ServiceConsoleReservationInfo       `mo:"consoleReservationInfo" json:"consoleReservationInfo,omitempty"`
	VirtualMachineReservationInfo *types.VirtualMachineMemoryReservationInfo `mo:"virtualMachineReservationInfo" json:"virtualMachineReservationInfo,omitempty"`
}

func init() {
	t["HostMemorySystem"] = reflect.TypeOf((*HostMemorySystem)(nil)).Elem()
}

type HostNetworkSystem struct {
	ExtensibleManagedObject

	Capabilities         *types.HostNetCapabilities        `mo:"capabilities" json:"capabilities,omitempty"`
	NetworkInfo          *types.HostNetworkInfo            `mo:"networkInfo" json:"networkInfo,omitempty"`
	OffloadCapabilities  *types.HostNetOffloadCapabilities `mo:"offloadCapabilities" json:"offloadCapabilities,omitempty"`
	NetworkConfig        *types.HostNetworkConfig          `mo:"networkConfig" json:"networkConfig,omitempty"`
	DnsConfig            types.BaseHostDnsConfig           `mo:"dnsConfig" json:"dnsConfig,omitempty"`
	IpRouteConfig        types.BaseHostIpRouteConfig       `mo:"ipRouteConfig" json:"ipRouteConfig,omitempty"`
	ConsoleIpRouteConfig types.BaseHostIpRouteConfig       `mo:"consoleIpRouteConfig" json:"consoleIpRouteConfig,omitempty"`
}

func init() {
	t["HostNetworkSystem"] = reflect.TypeOf((*HostNetworkSystem)(nil)).Elem()
}

type HostNvdimmSystem struct {
	Self types.ManagedObjectReference

	NvdimmSystemInfo types.NvdimmSystemInfo `mo:"nvdimmSystemInfo" json:"nvdimmSystemInfo"`
}

func (m HostNvdimmSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostNvdimmSystem"] = reflect.TypeOf((*HostNvdimmSystem)(nil)).Elem()
}

type HostPatchManager struct {
	Self types.ManagedObjectReference
}

func (m HostPatchManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostPatchManager"] = reflect.TypeOf((*HostPatchManager)(nil)).Elem()
}

type HostPciPassthruSystem struct {
	ExtensibleManagedObject

	PciPassthruInfo     []types.BaseHostPciPassthruInfo     `mo:"pciPassthruInfo" json:"pciPassthruInfo"`
	SriovDevicePoolInfo []types.BaseHostSriovDevicePoolInfo `mo:"sriovDevicePoolInfo" json:"sriovDevicePoolInfo,omitempty"`
}

func init() {
	t["HostPciPassthruSystem"] = reflect.TypeOf((*HostPciPassthruSystem)(nil)).Elem()
}

type HostPowerSystem struct {
	Self types.ManagedObjectReference

	Capability types.PowerSystemCapability `mo:"capability" json:"capability"`
	Info       types.PowerSystemInfo       `mo:"info" json:"info"`
}

func (m HostPowerSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostPowerSystem"] = reflect.TypeOf((*HostPowerSystem)(nil)).Elem()
}

type HostProfile struct {
	Profile

	ValidationState           *string                                 `mo:"validationState" json:"validationState,omitempty"`
	ValidationStateUpdateTime *time.Time                              `mo:"validationStateUpdateTime" json:"validationStateUpdateTime,omitempty"`
	ValidationFailureInfo     *types.HostProfileValidationFailureInfo `mo:"validationFailureInfo" json:"validationFailureInfo,omitempty"`
	ReferenceHost             *types.ManagedObjectReference           `mo:"referenceHost" json:"referenceHost,omitempty"`
}

func init() {
	t["HostProfile"] = reflect.TypeOf((*HostProfile)(nil)).Elem()
}

type HostProfileManager struct {
	ProfileManager
}

func init() {
	t["HostProfileManager"] = reflect.TypeOf((*HostProfileManager)(nil)).Elem()
}

type HostServiceSystem struct {
	ExtensibleManagedObject

	ServiceInfo types.HostServiceInfo `mo:"serviceInfo" json:"serviceInfo"`
}

func init() {
	t["HostServiceSystem"] = reflect.TypeOf((*HostServiceSystem)(nil)).Elem()
}

type HostSnmpSystem struct {
	Self types.ManagedObjectReference

	Configuration types.HostSnmpConfigSpec        `mo:"configuration" json:"configuration"`
	Limits        types.HostSnmpSystemAgentLimits `mo:"limits" json:"limits"`
}

func (m HostSnmpSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostSnmpSystem"] = reflect.TypeOf((*HostSnmpSystem)(nil)).Elem()
}

type HostSpecificationManager struct {
	Self types.ManagedObjectReference
}

func (m HostSpecificationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostSpecificationManager"] = reflect.TypeOf((*HostSpecificationManager)(nil)).Elem()
}

type HostStorageSystem struct {
	ExtensibleManagedObject

	StorageDeviceInfo    *types.HostStorageDeviceInfo   `mo:"storageDeviceInfo" json:"storageDeviceInfo,omitempty"`
	FileSystemVolumeInfo types.HostFileSystemVolumeInfo `mo:"fileSystemVolumeInfo" json:"fileSystemVolumeInfo"`
	SystemFile           []string                       `mo:"systemFile" json:"systemFile,omitempty"`
	MultipathStateInfo   *types.HostMultipathStateInfo  `mo:"multipathStateInfo" json:"multipathStateInfo,omitempty"`
}

func init() {
	t["HostStorageSystem"] = reflect.TypeOf((*HostStorageSystem)(nil)).Elem()
}

type HostSystem struct {
	ManagedEntity

	Runtime                    types.HostRuntimeInfo                      `mo:"runtime" json:"runtime"`
	Summary                    types.HostListSummary                      `mo:"summary" json:"summary"`
	Hardware                   *types.HostHardwareInfo                    `mo:"hardware" json:"hardware,omitempty"`
	Capability                 *types.HostCapability                      `mo:"capability" json:"capability,omitempty"`
	LicensableResource         types.HostLicensableResourceInfo           `mo:"licensableResource" json:"licensableResource"`
	RemediationState           *types.HostSystemRemediationState          `mo:"remediationState" json:"remediationState,omitempty"`
	PrecheckRemediationResult  *types.ApplyHostProfileConfigurationSpec   `mo:"precheckRemediationResult" json:"precheckRemediationResult,omitempty"`
	RemediationResult          *types.ApplyHostProfileConfigurationResult `mo:"remediationResult" json:"remediationResult,omitempty"`
	ComplianceCheckState       *types.HostSystemComplianceCheckState      `mo:"complianceCheckState" json:"complianceCheckState,omitempty"`
	ComplianceCheckResult      *types.ComplianceResult                    `mo:"complianceCheckResult" json:"complianceCheckResult,omitempty"`
	ConfigManager              types.HostConfigManager                    `mo:"configManager" json:"configManager"`
	Config                     *types.HostConfigInfo                      `mo:"config" json:"config,omitempty"`
	Vm                         []types.ManagedObjectReference             `mo:"vm" json:"vm,omitempty"`
	Datastore                  []types.ManagedObjectReference             `mo:"datastore" json:"datastore,omitempty"`
	Network                    []types.ManagedObjectReference             `mo:"network" json:"network,omitempty"`
	DatastoreBrowser           types.ManagedObjectReference               `mo:"datastoreBrowser" json:"datastoreBrowser"`
	SystemResources            *types.HostSystemResourceInfo              `mo:"systemResources" json:"systemResources,omitempty"`
	AnswerFileValidationState  *types.AnswerFileStatusResult              `mo:"answerFileValidationState" json:"answerFileValidationState,omitempty"`
	AnswerFileValidationResult *types.AnswerFileStatusResult              `mo:"answerFileValidationResult" json:"answerFileValidationResult,omitempty"`
}

func (m *HostSystem) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["HostSystem"] = reflect.TypeOf((*HostSystem)(nil)).Elem()
}

type HostVFlashManager struct {
	Self types.ManagedObjectReference

	VFlashConfigInfo *types.HostVFlashManagerVFlashConfigInfo `mo:"vFlashConfigInfo" json:"vFlashConfigInfo,omitempty"`
}

func (m HostVFlashManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostVFlashManager"] = reflect.TypeOf((*HostVFlashManager)(nil)).Elem()
}

type HostVMotionSystem struct {
	ExtensibleManagedObject

	NetConfig *types.HostVMotionNetConfig `mo:"netConfig" json:"netConfig,omitempty"`
	IpConfig  *types.HostIpConfig         `mo:"ipConfig" json:"ipConfig,omitempty"`
}

func init() {
	t["HostVMotionSystem"] = reflect.TypeOf((*HostVMotionSystem)(nil)).Elem()
}

type HostVStorageObjectManager struct {
	VStorageObjectManagerBase
}

func init() {
	t["HostVStorageObjectManager"] = reflect.TypeOf((*HostVStorageObjectManager)(nil)).Elem()
}

type HostVirtualNicManager struct {
	ExtensibleManagedObject

	Info types.HostVirtualNicManagerInfo `mo:"info" json:"info"`
}

func init() {
	t["HostVirtualNicManager"] = reflect.TypeOf((*HostVirtualNicManager)(nil)).Elem()
}

type HostVsanInternalSystem struct {
	Self types.ManagedObjectReference
}

func (m HostVsanInternalSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostVsanInternalSystem"] = reflect.TypeOf((*HostVsanInternalSystem)(nil)).Elem()
}

type HostVsanSystem struct {
	Self types.ManagedObjectReference

	Config types.VsanHostConfigInfo `mo:"config" json:"config"`
}

func (m HostVsanSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HostVsanSystem"] = reflect.TypeOf((*HostVsanSystem)(nil)).Elem()
}

type HttpNfcLease struct {
	Self types.ManagedObjectReference

	InitializeProgress int32                          `mo:"initializeProgress" json:"initializeProgress"`
	TransferProgress   int32                          `mo:"transferProgress" json:"transferProgress"`
	Mode               string                         `mo:"mode" json:"mode"`
	Capabilities       types.HttpNfcLeaseCapabilities `mo:"capabilities" json:"capabilities"`
	Info               *types.HttpNfcLeaseInfo        `mo:"info" json:"info,omitempty"`
	State              types.HttpNfcLeaseState        `mo:"state" json:"state"`
	Error              *types.LocalizedMethodFault    `mo:"error" json:"error,omitempty"`
}

func (m HttpNfcLease) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["HttpNfcLease"] = reflect.TypeOf((*HttpNfcLease)(nil)).Elem()
}

type InventoryView struct {
	ManagedObjectView
}

func init() {
	t["InventoryView"] = reflect.TypeOf((*InventoryView)(nil)).Elem()
}

type IoFilterManager struct {
	Self types.ManagedObjectReference
}

func (m IoFilterManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["IoFilterManager"] = reflect.TypeOf((*IoFilterManager)(nil)).Elem()
}

type IpPoolManager struct {
	Self types.ManagedObjectReference
}

func (m IpPoolManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["IpPoolManager"] = reflect.TypeOf((*IpPoolManager)(nil)).Elem()
}

type IscsiManager struct {
	Self types.ManagedObjectReference
}

func (m IscsiManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["IscsiManager"] = reflect.TypeOf((*IscsiManager)(nil)).Elem()
}

type LicenseAssignmentManager struct {
	Self types.ManagedObjectReference
}

func (m LicenseAssignmentManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["LicenseAssignmentManager"] = reflect.TypeOf((*LicenseAssignmentManager)(nil)).Elem()
}

type LicenseManager struct {
	Self types.ManagedObjectReference

	Source                   types.BaseLicenseSource            `mo:"source" json:"source"`
	SourceAvailable          bool                               `mo:"sourceAvailable" json:"sourceAvailable"`
	Diagnostics              *types.LicenseDiagnostics          `mo:"diagnostics" json:"diagnostics,omitempty"`
	FeatureInfo              []types.LicenseFeatureInfo         `mo:"featureInfo" json:"featureInfo,omitempty"`
	LicensedEdition          string                             `mo:"licensedEdition" json:"licensedEdition"`
	Licenses                 []types.LicenseManagerLicenseInfo  `mo:"licenses" json:"licenses"`
	LicenseAssignmentManager *types.ManagedObjectReference      `mo:"licenseAssignmentManager" json:"licenseAssignmentManager,omitempty"`
	Evaluation               types.LicenseManagerEvaluationInfo `mo:"evaluation" json:"evaluation"`
}

func (m LicenseManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["LicenseManager"] = reflect.TypeOf((*LicenseManager)(nil)).Elem()
}

type ListView struct {
	ManagedObjectView
}

func init() {
	t["ListView"] = reflect.TypeOf((*ListView)(nil)).Elem()
}

type LocalizationManager struct {
	Self types.ManagedObjectReference

	Catalog []types.LocalizationManagerMessageCatalog `mo:"catalog" json:"catalog,omitempty"`
}

func (m LocalizationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["LocalizationManager"] = reflect.TypeOf((*LocalizationManager)(nil)).Elem()
}

type ManagedEntity struct {
	ExtensibleManagedObject

	Parent              *types.ManagedObjectReference  `mo:"parent" json:"parent,omitempty"`
	CustomValue         []types.BaseCustomFieldValue   `mo:"customValue" json:"customValue,omitempty"`
	OverallStatus       types.ManagedEntityStatus      `mo:"overallStatus" json:"overallStatus"`
	ConfigStatus        types.ManagedEntityStatus      `mo:"configStatus" json:"configStatus"`
	ConfigIssue         []types.BaseEvent              `mo:"configIssue" json:"configIssue,omitempty"`
	EffectiveRole       []int32                        `mo:"effectiveRole" json:"effectiveRole,omitempty"`
	Permission          []types.Permission             `mo:"permission" json:"permission,omitempty"`
	Name                string                         `mo:"name" json:"name"`
	DisabledMethod      []string                       `mo:"disabledMethod" json:"disabledMethod,omitempty"`
	RecentTask          []types.ManagedObjectReference `mo:"recentTask" json:"recentTask,omitempty"`
	DeclaredAlarmState  []types.AlarmState             `mo:"declaredAlarmState" json:"declaredAlarmState,omitempty"`
	TriggeredAlarmState []types.AlarmState             `mo:"triggeredAlarmState" json:"triggeredAlarmState,omitempty"`
	AlarmActionsEnabled *bool                          `mo:"alarmActionsEnabled" json:"alarmActionsEnabled,omitempty"`
	Tag                 []types.Tag                    `mo:"tag" json:"tag,omitempty"`
}

func init() {
	t["ManagedEntity"] = reflect.TypeOf((*ManagedEntity)(nil)).Elem()
}

type ManagedObjectView struct {
	Self types.ManagedObjectReference

	View []types.ManagedObjectReference `mo:"view" json:"view,omitempty"`
}

func (m ManagedObjectView) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ManagedObjectView"] = reflect.TypeOf((*ManagedObjectView)(nil)).Elem()
}

type MessageBusProxy struct {
	Self types.ManagedObjectReference
}

func (m MessageBusProxy) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["MessageBusProxy"] = reflect.TypeOf((*MessageBusProxy)(nil)).Elem()
}

type Network struct {
	ManagedEntity

	Summary types.BaseNetworkSummary       `mo:"summary" json:"summary"`
	Host    []types.ManagedObjectReference `mo:"host" json:"host,omitempty"`
	Vm      []types.ManagedObjectReference `mo:"vm" json:"vm,omitempty"`
	Name    string                         `mo:"name" json:"name"`
}

func (m *Network) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["Network"] = reflect.TypeOf((*Network)(nil)).Elem()
}

type OpaqueNetwork struct {
	Network

	Capability  *types.OpaqueNetworkCapability `mo:"capability" json:"capability,omitempty"`
	ExtraConfig []types.BaseOptionValue        `mo:"extraConfig" json:"extraConfig,omitempty"`
}

func init() {
	t["OpaqueNetwork"] = reflect.TypeOf((*OpaqueNetwork)(nil)).Elem()
}

type OptionManager struct {
	Self types.ManagedObjectReference

	SupportedOption []types.OptionDef       `mo:"supportedOption" json:"supportedOption,omitempty"`
	Setting         []types.BaseOptionValue `mo:"setting" json:"setting,omitempty"`
}

func (m OptionManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["OptionManager"] = reflect.TypeOf((*OptionManager)(nil)).Elem()
}

type OverheadMemoryManager struct {
	Self types.ManagedObjectReference
}

func (m OverheadMemoryManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["OverheadMemoryManager"] = reflect.TypeOf((*OverheadMemoryManager)(nil)).Elem()
}

type OvfManager struct {
	Self types.ManagedObjectReference

	OvfImportOption []types.OvfOptionInfo `mo:"ovfImportOption" json:"ovfImportOption,omitempty"`
	OvfExportOption []types.OvfOptionInfo `mo:"ovfExportOption" json:"ovfExportOption,omitempty"`
}

func (m OvfManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["OvfManager"] = reflect.TypeOf((*OvfManager)(nil)).Elem()
}

type PerformanceManager struct {
	Self types.ManagedObjectReference

	Description        types.PerformanceDescription `mo:"description" json:"description"`
	HistoricalInterval []types.PerfInterval         `mo:"historicalInterval" json:"historicalInterval,omitempty"`
	PerfCounter        []types.PerfCounterInfo      `mo:"perfCounter" json:"perfCounter,omitempty"`
}

func (m PerformanceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["PerformanceManager"] = reflect.TypeOf((*PerformanceManager)(nil)).Elem()
}

type Profile struct {
	Self types.ManagedObjectReference

	Config           types.BaseProfileConfigInfo    `mo:"config" json:"config"`
	Description      *types.ProfileDescription      `mo:"description" json:"description,omitempty"`
	Name             string                         `mo:"name" json:"name"`
	CreatedTime      time.Time                      `mo:"createdTime" json:"createdTime"`
	ModifiedTime     time.Time                      `mo:"modifiedTime" json:"modifiedTime"`
	Entity           []types.ManagedObjectReference `mo:"entity" json:"entity,omitempty"`
	ComplianceStatus string                         `mo:"complianceStatus" json:"complianceStatus"`
}

func (m Profile) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["Profile"] = reflect.TypeOf((*Profile)(nil)).Elem()
}

type ProfileComplianceManager struct {
	Self types.ManagedObjectReference
}

func (m ProfileComplianceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ProfileComplianceManager"] = reflect.TypeOf((*ProfileComplianceManager)(nil)).Elem()
}

type ProfileManager struct {
	Self types.ManagedObjectReference

	Profile []types.ManagedObjectReference `mo:"profile" json:"profile,omitempty"`
}

func (m ProfileManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ProfileManager"] = reflect.TypeOf((*ProfileManager)(nil)).Elem()
}

type PropertyCollector struct {
	Self types.ManagedObjectReference

	Filter []types.ManagedObjectReference `mo:"filter" json:"filter,omitempty"`
}

func (m PropertyCollector) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["PropertyCollector"] = reflect.TypeOf((*PropertyCollector)(nil)).Elem()
}

type PropertyFilter struct {
	Self types.ManagedObjectReference

	Spec           types.PropertyFilterSpec `mo:"spec" json:"spec"`
	PartialUpdates bool                     `mo:"partialUpdates" json:"partialUpdates"`
}

func (m PropertyFilter) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["PropertyFilter"] = reflect.TypeOf((*PropertyFilter)(nil)).Elem()
}

type ResourcePlanningManager struct {
	Self types.ManagedObjectReference
}

func (m ResourcePlanningManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ResourcePlanningManager"] = reflect.TypeOf((*ResourcePlanningManager)(nil)).Elem()
}

type ResourcePool struct {
	ManagedEntity

	Summary            types.BaseResourcePoolSummary  `mo:"summary" json:"summary"`
	Runtime            types.ResourcePoolRuntimeInfo  `mo:"runtime" json:"runtime"`
	Owner              types.ManagedObjectReference   `mo:"owner" json:"owner"`
	ResourcePool       []types.ManagedObjectReference `mo:"resourcePool" json:"resourcePool,omitempty"`
	Vm                 []types.ManagedObjectReference `mo:"vm" json:"vm,omitempty"`
	Config             types.ResourceConfigSpec       `mo:"config" json:"config"`
	Namespace          *string                        `mo:"namespace" json:"namespace,omitempty"`
	ChildConfiguration []types.ResourceConfigSpec     `mo:"childConfiguration" json:"childConfiguration,omitempty"`
}

func (m *ResourcePool) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["ResourcePool"] = reflect.TypeOf((*ResourcePool)(nil)).Elem()
}

type ScheduledTask struct {
	ExtensibleManagedObject

	Info types.ScheduledTaskInfo `mo:"info" json:"info"`
}

func init() {
	t["ScheduledTask"] = reflect.TypeOf((*ScheduledTask)(nil)).Elem()
}

type ScheduledTaskManager struct {
	Self types.ManagedObjectReference

	ScheduledTask []types.ManagedObjectReference `mo:"scheduledTask" json:"scheduledTask,omitempty"`
	Description   types.ScheduledTaskDescription `mo:"description" json:"description"`
}

func (m ScheduledTaskManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ScheduledTaskManager"] = reflect.TypeOf((*ScheduledTaskManager)(nil)).Elem()
}

type SearchIndex struct {
	Self types.ManagedObjectReference
}

func (m SearchIndex) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["SearchIndex"] = reflect.TypeOf((*SearchIndex)(nil)).Elem()
}

type ServiceInstance struct {
	Self types.ManagedObjectReference

	ServerClock time.Time            `mo:"serverClock" json:"serverClock"`
	Capability  types.Capability     `mo:"capability" json:"capability"`
	Content     types.ServiceContent `mo:"content" json:"content"`
}

func (m ServiceInstance) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ServiceInstance"] = reflect.TypeOf((*ServiceInstance)(nil)).Elem()
}

type ServiceManager struct {
	Self types.ManagedObjectReference

	Service []types.ServiceManagerServiceInfo `mo:"service" json:"service,omitempty"`
}

func (m ServiceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ServiceManager"] = reflect.TypeOf((*ServiceManager)(nil)).Elem()
}

type SessionManager struct {
	Self types.ManagedObjectReference

	SessionList         []types.UserSession `mo:"sessionList" json:"sessionList,omitempty"`
	CurrentSession      *types.UserSession  `mo:"currentSession" json:"currentSession,omitempty"`
	Message             *string             `mo:"message" json:"message,omitempty"`
	MessageLocaleList   []string            `mo:"messageLocaleList" json:"messageLocaleList,omitempty"`
	SupportedLocaleList []string            `mo:"supportedLocaleList" json:"supportedLocaleList,omitempty"`
	DefaultLocale       string              `mo:"defaultLocale" json:"defaultLocale"`
}

func (m SessionManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["SessionManager"] = reflect.TypeOf((*SessionManager)(nil)).Elem()
}

type SimpleCommand struct {
	Self types.ManagedObjectReference

	EncodingType types.SimpleCommandEncoding     `mo:"encodingType" json:"encodingType"`
	Entity       types.ServiceManagerServiceInfo `mo:"entity" json:"entity"`
}

func (m SimpleCommand) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["SimpleCommand"] = reflect.TypeOf((*SimpleCommand)(nil)).Elem()
}

type SiteInfoManager struct {
	Self types.ManagedObjectReference
}

func (m SiteInfoManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["SiteInfoManager"] = reflect.TypeOf((*SiteInfoManager)(nil)).Elem()
}

type StoragePod struct {
	Folder

	Summary            *types.StoragePodSummary  `mo:"summary" json:"summary,omitempty"`
	PodStorageDrsEntry *types.PodStorageDrsEntry `mo:"podStorageDrsEntry" json:"podStorageDrsEntry,omitempty"`
}

func init() {
	t["StoragePod"] = reflect.TypeOf((*StoragePod)(nil)).Elem()
}

type StorageQueryManager struct {
	Self types.ManagedObjectReference
}

func (m StorageQueryManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["StorageQueryManager"] = reflect.TypeOf((*StorageQueryManager)(nil)).Elem()
}

type StorageResourceManager struct {
	Self types.ManagedObjectReference
}

func (m StorageResourceManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["StorageResourceManager"] = reflect.TypeOf((*StorageResourceManager)(nil)).Elem()
}

type Task struct {
	ExtensibleManagedObject

	Info types.TaskInfo `mo:"info" json:"info"`
}

func init() {
	t["Task"] = reflect.TypeOf((*Task)(nil)).Elem()
}

type TaskHistoryCollector struct {
	HistoryCollector

	LatestPage []types.TaskInfo `mo:"latestPage" json:"latestPage,omitempty"`
}

func init() {
	t["TaskHistoryCollector"] = reflect.TypeOf((*TaskHistoryCollector)(nil)).Elem()
}

type TaskManager struct {
	Self types.ManagedObjectReference

	RecentTask   []types.ManagedObjectReference `mo:"recentTask" json:"recentTask,omitempty"`
	Description  types.TaskDescription          `mo:"description" json:"description"`
	MaxCollector int32                          `mo:"maxCollector" json:"maxCollector"`
}

func (m TaskManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["TaskManager"] = reflect.TypeOf((*TaskManager)(nil)).Elem()
}

type TenantTenantManager struct {
	Self types.ManagedObjectReference
}

func (m TenantTenantManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["TenantTenantManager"] = reflect.TypeOf((*TenantTenantManager)(nil)).Elem()
}

type UserDirectory struct {
	Self types.ManagedObjectReference

	DomainList []string `mo:"domainList" json:"domainList,omitempty"`
}

func (m UserDirectory) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["UserDirectory"] = reflect.TypeOf((*UserDirectory)(nil)).Elem()
}

type VStorageObjectManagerBase struct {
	Self types.ManagedObjectReference
}

func (m VStorageObjectManagerBase) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VStorageObjectManagerBase"] = reflect.TypeOf((*VStorageObjectManagerBase)(nil)).Elem()
}

type VcenterVStorageObjectManager struct {
	VStorageObjectManagerBase
}

func init() {
	t["VcenterVStorageObjectManager"] = reflect.TypeOf((*VcenterVStorageObjectManager)(nil)).Elem()
}

type View struct {
	Self types.ManagedObjectReference
}

func (m View) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["View"] = reflect.TypeOf((*View)(nil)).Elem()
}

type ViewManager struct {
	Self types.ManagedObjectReference

	ViewList []types.ManagedObjectReference `mo:"viewList" json:"viewList,omitempty"`
}

func (m ViewManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["ViewManager"] = reflect.TypeOf((*ViewManager)(nil)).Elem()
}

type VirtualApp struct {
	ResourcePool

	ParentFolder *types.ManagedObjectReference  `mo:"parentFolder" json:"parentFolder,omitempty"`
	Datastore    []types.ManagedObjectReference `mo:"datastore" json:"datastore,omitempty"`
	Network      []types.ManagedObjectReference `mo:"network" json:"network,omitempty"`
	VAppConfig   *types.VAppConfigInfo          `mo:"vAppConfig" json:"vAppConfig,omitempty"`
	ParentVApp   *types.ManagedObjectReference  `mo:"parentVApp" json:"parentVApp,omitempty"`
	ChildLink    []types.VirtualAppLinkInfo     `mo:"childLink" json:"childLink,omitempty"`
}

func init() {
	t["VirtualApp"] = reflect.TypeOf((*VirtualApp)(nil)).Elem()
}

type VirtualDiskManager struct {
	Self types.ManagedObjectReference
}

func (m VirtualDiskManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VirtualDiskManager"] = reflect.TypeOf((*VirtualDiskManager)(nil)).Elem()
}

type VirtualMachine struct {
	ManagedEntity

	Capability           types.VirtualMachineCapability    `mo:"capability" json:"capability"`
	Config               *types.VirtualMachineConfigInfo   `mo:"config" json:"config,omitempty"`
	Layout               *types.VirtualMachineFileLayout   `mo:"layout" json:"layout,omitempty"`
	LayoutEx             *types.VirtualMachineFileLayoutEx `mo:"layoutEx" json:"layoutEx,omitempty"`
	Storage              *types.VirtualMachineStorageInfo  `mo:"storage" json:"storage,omitempty"`
	EnvironmentBrowser   types.ManagedObjectReference      `mo:"environmentBrowser" json:"environmentBrowser"`
	ResourcePool         *types.ManagedObjectReference     `mo:"resourcePool" json:"resourcePool,omitempty"`
	ParentVApp           *types.ManagedObjectReference     `mo:"parentVApp" json:"parentVApp,omitempty"`
	ResourceConfig       *types.ResourceConfigSpec         `mo:"resourceConfig" json:"resourceConfig,omitempty"`
	Runtime              types.VirtualMachineRuntimeInfo   `mo:"runtime" json:"runtime"`
	Guest                *types.GuestInfo                  `mo:"guest" json:"guest,omitempty"`
	Summary              types.VirtualMachineSummary       `mo:"summary" json:"summary"`
	Datastore            []types.ManagedObjectReference    `mo:"datastore" json:"datastore,omitempty"`
	Network              []types.ManagedObjectReference    `mo:"network" json:"network,omitempty"`
	Snapshot             *types.VirtualMachineSnapshotInfo `mo:"snapshot" json:"snapshot,omitempty"`
	RootSnapshot         []types.ManagedObjectReference    `mo:"rootSnapshot" json:"rootSnapshot,omitempty"`
	GuestHeartbeatStatus types.ManagedEntityStatus         `mo:"guestHeartbeatStatus" json:"guestHeartbeatStatus"`
}

func (m *VirtualMachine) Entity() *ManagedEntity {
	return &m.ManagedEntity
}

func init() {
	t["VirtualMachine"] = reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

type VirtualMachineCompatibilityChecker struct {
	Self types.ManagedObjectReference
}

func (m VirtualMachineCompatibilityChecker) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VirtualMachineCompatibilityChecker"] = reflect.TypeOf((*VirtualMachineCompatibilityChecker)(nil)).Elem()
}

type VirtualMachineGuestCustomizationManager struct {
	Self types.ManagedObjectReference
}

func (m VirtualMachineGuestCustomizationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VirtualMachineGuestCustomizationManager"] = reflect.TypeOf((*VirtualMachineGuestCustomizationManager)(nil)).Elem()
}

type VirtualMachineProvisioningChecker struct {
	Self types.ManagedObjectReference
}

func (m VirtualMachineProvisioningChecker) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VirtualMachineProvisioningChecker"] = reflect.TypeOf((*VirtualMachineProvisioningChecker)(nil)).Elem()
}

type VirtualMachineSnapshot struct {
	ExtensibleManagedObject

	Config        types.VirtualMachineConfigInfo `mo:"config" json:"config"`
	ChildSnapshot []types.ManagedObjectReference `mo:"childSnapshot" json:"childSnapshot,omitempty"`
	Vm            types.ManagedObjectReference   `mo:"vm" json:"vm"`
}

func init() {
	t["VirtualMachineSnapshot"] = reflect.TypeOf((*VirtualMachineSnapshot)(nil)).Elem()
}

type VirtualizationManager struct {
	Self types.ManagedObjectReference
}

func (m VirtualizationManager) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VirtualizationManager"] = reflect.TypeOf((*VirtualizationManager)(nil)).Elem()
}

type VmwareDistributedVirtualSwitch struct {
	DistributedVirtualSwitch
}

func init() {
	t["VmwareDistributedVirtualSwitch"] = reflect.TypeOf((*VmwareDistributedVirtualSwitch)(nil)).Elem()
}

type VsanUpgradeSystem struct {
	Self types.ManagedObjectReference
}

func (m VsanUpgradeSystem) Reference() types.ManagedObjectReference {
	return m.Self
}

func init() {
	t["VsanUpgradeSystem"] = reflect.TypeOf((*VsanUpgradeSystem)(nil)).Elem()
}
