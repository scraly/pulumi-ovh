# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
    'get_load_balancer_output',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, created_at=None, flavor_id=None, floating_ip=None, id=None, name=None, operating_status=None, provisioning_status=None, region_name=None, service_name=None, updated_at=None, vip_address=None, vip_network_id=None, vip_subnet_id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if floating_ip and not isinstance(floating_ip, dict):
            raise TypeError("Expected argument 'floating_ip' to be a dict")
        pulumi.set(__self__, "floating_ip", floating_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operating_status and not isinstance(operating_status, str):
            raise TypeError("Expected argument 'operating_status' to be a str")
        pulumi.set(__self__, "operating_status", operating_status)
        if provisioning_status and not isinstance(provisioning_status, str):
            raise TypeError("Expected argument 'provisioning_status' to be a str")
        pulumi.set(__self__, "provisioning_status", provisioning_status)
        if region_name and not isinstance(region_name, str):
            raise TypeError("Expected argument 'region_name' to be a str")
        pulumi.set(__self__, "region_name", region_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if vip_address and not isinstance(vip_address, str):
            raise TypeError("Expected argument 'vip_address' to be a str")
        pulumi.set(__self__, "vip_address", vip_address)
        if vip_network_id and not isinstance(vip_network_id, str):
            raise TypeError("Expected argument 'vip_network_id' to be a str")
        pulumi.set(__self__, "vip_network_id", vip_network_id)
        if vip_subnet_id and not isinstance(vip_subnet_id, str):
            raise TypeError("Expected argument 'vip_subnet_id' to be a str")
        pulumi.set(__self__, "vip_subnet_id", vip_subnet_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of creation of the loadbalancer
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> str:
        """
        ID of the flavor
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> 'outputs.GetLoadBalancerFloatingIpResult':
        """
        Information about the floating IP
        """
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the floating IP
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the loadbalancer
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operatingStatus")
    def operating_status(self) -> str:
        """
        Operating status of the loadbalancer
        """
        return pulumi.get(self, "operating_status")

    @property
    @pulumi.getter(name="provisioningStatus")
    def provisioning_status(self) -> str:
        """
        Provisioning status of the loadbalancer
        """
        return pulumi.get(self, "provisioning_status")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        """
        Region of the loadbalancer
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Last update date of the loadbalancer
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vipAddress")
    def vip_address(self) -> str:
        """
        IP address of the Virtual IP
        """
        return pulumi.get(self, "vip_address")

    @property
    @pulumi.getter(name="vipNetworkId")
    def vip_network_id(self) -> str:
        """
        Openstack ID of the network for the Virtual IP
        """
        return pulumi.get(self, "vip_network_id")

    @property
    @pulumi.getter(name="vipSubnetId")
    def vip_subnet_id(self) -> str:
        """
        ID of the subnet for the Virtual IP
        """
        return pulumi.get(self, "vip_subnet_id")


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            created_at=self.created_at,
            flavor_id=self.flavor_id,
            floating_ip=self.floating_ip,
            id=self.id,
            name=self.name,
            operating_status=self.operating_status,
            provisioning_status=self.provisioning_status,
            region_name=self.region_name,
            service_name=self.service_name,
            updated_at=self.updated_at,
            vip_address=self.vip_address,
            vip_network_id=self.vip_network_id,
            vip_subnet_id=self.vip_subnet_id)


def get_load_balancer(id: Optional[str] = None,
                      region_name: Optional[str] = None,
                      service_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    Get the details of a public cloud project loadbalancer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_load_balancer = ovh.CloudProject.get_load_balancer(service_name="XXXXXX",
        region_name="XXX",
        id="XXX")
    pulumi.export("lb", lb_load_balancer)
    ```


    :param str id: ID of the loadbalancer
    :param str region_name: Region of the loadbalancer.
    :param str service_name: The ID of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getLoadBalancer:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        floating_ip=pulumi.get(__ret__, 'floating_ip'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        operating_status=pulumi.get(__ret__, 'operating_status'),
        provisioning_status=pulumi.get(__ret__, 'provisioning_status'),
        region_name=pulumi.get(__ret__, 'region_name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        vip_address=pulumi.get(__ret__, 'vip_address'),
        vip_network_id=pulumi.get(__ret__, 'vip_network_id'),
        vip_subnet_id=pulumi.get(__ret__, 'vip_subnet_id'))


@_utilities.lift_output_func(get_load_balancer)
def get_load_balancer_output(id: Optional[pulumi.Input[str]] = None,
                             region_name: Optional[pulumi.Input[str]] = None,
                             service_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerResult]:
    """
    Get the details of a public cloud project loadbalancer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    lb_load_balancer = ovh.CloudProject.get_load_balancer(service_name="XXXXXX",
        region_name="XXX",
        id="XXX")
    pulumi.export("lb", lb_load_balancer)
    ```


    :param str id: ID of the loadbalancer
    :param str region_name: Region of the loadbalancer.
    :param str service_name: The ID of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
