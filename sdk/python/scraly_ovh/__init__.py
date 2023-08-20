# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .get_installation_templates import *
from .get_server import *
from .get_servers import *
from .get_vrack_networks import *
from .provider import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import scraly_ovh.cloudproject as __cloudproject
    cloudproject = __cloudproject
    import scraly_ovh.cloudprojectdatabase as __cloudprojectdatabase
    cloudprojectdatabase = __cloudprojectdatabase
    import scraly_ovh.config as __config
    config = __config
    import scraly_ovh.dbaas as __dbaas
    dbaas = __dbaas
    import scraly_ovh.dedicated as __dedicated
    dedicated = __dedicated
    import scraly_ovh.domain as __domain
    domain = __domain
    import scraly_ovh.hosting as __hosting
    hosting = __hosting
    import scraly_ovh.iam as __iam
    iam = __iam
    import scraly_ovh.ip as __ip
    ip = __ip
    import scraly_ovh.iploadbalancing as __iploadbalancing
    iploadbalancing = __iploadbalancing
    import scraly_ovh.me as __me
    me = __me
    import scraly_ovh.order as __order
    order = __order
    import scraly_ovh.vps as __vps
    vps = __vps
    import scraly_ovh.vrack as __vrack
    vrack = __vrack
else:
    cloudproject = _utilities.lazy_import('scraly_ovh.cloudproject')
    cloudprojectdatabase = _utilities.lazy_import('scraly_ovh.cloudprojectdatabase')
    config = _utilities.lazy_import('scraly_ovh.config')
    dbaas = _utilities.lazy_import('scraly_ovh.dbaas')
    dedicated = _utilities.lazy_import('scraly_ovh.dedicated')
    domain = _utilities.lazy_import('scraly_ovh.domain')
    hosting = _utilities.lazy_import('scraly_ovh.hosting')
    iam = _utilities.lazy_import('scraly_ovh.iam')
    ip = _utilities.lazy_import('scraly_ovh.ip')
    iploadbalancing = _utilities.lazy_import('scraly_ovh.iploadbalancing')
    me = _utilities.lazy_import('scraly_ovh.me')
    order = _utilities.lazy_import('scraly_ovh.order')
    vps = _utilities.lazy_import('scraly_ovh.vps')
    vrack = _utilities.lazy_import('scraly_ovh.vrack')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "ovh",
  "mod": "CloudProject/containerRegistry",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/containerRegistry:ContainerRegistry": "ContainerRegistry"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/containerRegistryUser",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/containerRegistryUser:ContainerRegistryUser": "ContainerRegistryUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/database",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/database:Database": "Database"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/failoverIpAttach",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/failoverIpAttach:FailoverIpAttach": "FailoverIpAttach"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/kube",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/kube:Kube": "Kube"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/kubeIpRestrictions",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions": "KubeIpRestrictions"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/kubeNodePool",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/kubeNodePool:KubeNodePool": "KubeNodePool"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/kubeOidc",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/kubeOidc:KubeOidc": "KubeOidc"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/networkPrivate",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/networkPrivate:NetworkPrivate": "NetworkPrivate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/networkPrivateSubnet",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet": "NetworkPrivateSubnet"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/project",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/project:Project": "Project"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/regionStoragePresign",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/regionStoragePresign:RegionStoragePresign": "RegionStoragePresign"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/s3Credential",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/s3Credential:S3Credential": "S3Credential"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/s3Policy",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/s3Policy:S3Policy": "S3Policy"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/user",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/user:User": "User"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProject/workflowBackup",
  "fqn": "scraly_ovh.cloudproject",
  "classes": {
   "ovh:CloudProject/workflowBackup:WorkflowBackup": "WorkflowBackup"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/databaseInstance",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance": "DatabaseInstance"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/integration",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/ipRestriction",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/ipRestriction:IpRestriction": "IpRestriction"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/kafkaAcl",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl": "KafkaAcl"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/kafkaTopic",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic": "KafkaTopic"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/m3DbNamespace",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace": "M3DbNamespace"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/m3DbUser",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/m3DbUser:M3DbUser": "M3DbUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/mongoDbUser",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/mongoDbUser:MongoDbUser": "MongoDbUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/opensearchPattern",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern": "OpensearchPattern"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/opensearchUser",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/opensearchUser:OpensearchUser": "OpensearchUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/postgresSqlUser",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser": "PostgresSqlUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/redisUser",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/redisUser:RedisUser": "RedisUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "CloudProjectDatabase/user",
  "fqn": "scraly_ovh.cloudprojectdatabase",
  "classes": {
   "ovh:CloudProjectDatabase/user:User": "User"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dbaas/logsCluster",
  "fqn": "scraly_ovh.dbaas",
  "classes": {
   "ovh:Dbaas/logsCluster:LogsCluster": "LogsCluster"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dbaas/logsInput",
  "fqn": "scraly_ovh.dbaas",
  "classes": {
   "ovh:Dbaas/logsInput:LogsInput": "LogsInput"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dbaas/logsOutputGraylogStream",
  "fqn": "scraly_ovh.dbaas",
  "classes": {
   "ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream": "LogsOutputGraylogStream"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/cephAcl",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/cephAcl:CephAcl": "CephAcl"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/nasHAPartition",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/nasHAPartition:NasHAPartition": "NasHAPartition"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/nasHAPartitionAccess",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess": "NasHAPartitionAccess"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/nasHAPartitionSnapshot",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot": "NasHAPartitionSnapshot"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/serverInstallTask",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/serverInstallTask:ServerInstallTask": "ServerInstallTask"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/serverNetworking",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/serverNetworking:ServerNetworking": "ServerNetworking"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/serverRebootTask",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/serverRebootTask:ServerRebootTask": "ServerRebootTask"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Dedicated/serverUpdate",
  "fqn": "scraly_ovh.dedicated",
  "classes": {
   "ovh:Dedicated/serverUpdate:ServerUpdate": "ServerUpdate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Domain/zone",
  "fqn": "scraly_ovh.domain",
  "classes": {
   "ovh:Domain/zone:Zone": "Zone"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Domain/zoneRecord",
  "fqn": "scraly_ovh.domain",
  "classes": {
   "ovh:Domain/zoneRecord:ZoneRecord": "ZoneRecord"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Domain/zoneRedirection",
  "fqn": "scraly_ovh.domain",
  "classes": {
   "ovh:Domain/zoneRedirection:ZoneRedirection": "ZoneRedirection"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Hosting/privateDatabase",
  "fqn": "scraly_ovh.hosting",
  "classes": {
   "ovh:Hosting/privateDatabase:PrivateDatabase": "PrivateDatabase"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Hosting/privateDatabaseAllowlist",
  "fqn": "scraly_ovh.hosting",
  "classes": {
   "ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist": "PrivateDatabaseAllowlist"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Hosting/privateDatabaseDb",
  "fqn": "scraly_ovh.hosting",
  "classes": {
   "ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb": "PrivateDatabaseDb"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Hosting/privateDatabaseUser",
  "fqn": "scraly_ovh.hosting",
  "classes": {
   "ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser": "PrivateDatabaseUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Hosting/privateDatabaseUserGrant",
  "fqn": "scraly_ovh.hosting",
  "classes": {
   "ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant": "PrivateDatabaseUserGrant"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Iam/policy",
  "fqn": "scraly_ovh.iam",
  "classes": {
   "ovh:Iam/policy:Policy": "Policy"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Ip/ipService",
  "fqn": "scraly_ovh.ip",
  "classes": {
   "ovh:Ip/ipService:IpService": "IpService"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Ip/reverse",
  "fqn": "scraly_ovh.ip",
  "classes": {
   "ovh:Ip/reverse:Reverse": "Reverse"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/httpFarm",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/httpFarm:HttpFarm": "HttpFarm"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/httpFarmServer",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer": "HttpFarmServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/httpFrontend",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/httpFrontend:HttpFrontend": "HttpFrontend"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/httpRoute",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/httpRoute:HttpRoute": "HttpRoute"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/httpRouteRule",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule": "HttpRouteRule"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/loadBalancer",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/loadBalancer:LoadBalancer": "LoadBalancer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/refresh",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/refresh:Refresh": "Refresh"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/tcpFarm",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/tcpFarm:TcpFarm": "TcpFarm"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/tcpFarmServer",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/tcpFarmServer:TcpFarmServer": "TcpFarmServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/tcpFrontend",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/tcpFrontend:TcpFrontend": "TcpFrontend"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/tcpRoute",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/tcpRoute:TcpRoute": "TcpRoute"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/tcpRouteRule",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule": "TcpRouteRule"
  }
 },
 {
  "pkg": "ovh",
  "mod": "IpLoadBalancing/vrackNetwork",
  "fqn": "scraly_ovh.iploadbalancing",
  "classes": {
   "ovh:IpLoadBalancing/vrackNetwork:VrackNetwork": "VrackNetwork"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/identityGroup",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/identityGroup:IdentityGroup": "IdentityGroup"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/identityUser",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/identityUser:IdentityUser": "IdentityUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/installationTemplate",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/installationTemplate:InstallationTemplate": "InstallationTemplate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/installationTemplatePartitionScheme",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/installationTemplatePartitionScheme:InstallationTemplatePartitionScheme": "InstallationTemplatePartitionScheme"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/installationTemplatePartitionSchemeHardwareRaid",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid": "InstallationTemplatePartitionSchemeHardwareRaid"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/installationTemplatePartitionSchemePartition",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition": "InstallationTemplatePartitionSchemePartition"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/ipxeScript",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/ipxeScript:IpxeScript": "IpxeScript"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Me/sshKey",
  "fqn": "scraly_ovh.me",
  "classes": {
   "ovh:Me/sshKey:SshKey": "SshKey"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/cloudProject",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/cloudProject:CloudProject": "CloudProject"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/dedicatedServer",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/dedicatedServer:DedicatedServer": "DedicatedServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/dedicatedServerInterface",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface": "DedicatedServerInterface"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/ipAddress",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/ipAddress:IpAddress": "IpAddress"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/ipLoadbalancing",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/ipLoadbalancing:IpLoadbalancing": "IpLoadbalancing"
  }
 },
 {
  "pkg": "ovh",
  "mod": "Vrack/vrack",
  "fqn": "scraly_ovh.vrack",
  "classes": {
   "ovh:Vrack/vrack:Vrack": "Vrack"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "ovh",
  "token": "pulumi:providers:ovh",
  "fqn": "scraly_ovh",
  "class": "Provider"
 }
]
"""
)