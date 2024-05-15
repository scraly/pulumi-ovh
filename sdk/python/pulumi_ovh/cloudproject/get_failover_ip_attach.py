# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFailoverIpAttachResult',
    'AwaitableGetFailoverIpAttachResult',
    'get_failover_ip_attach',
    'get_failover_ip_attach_output',
]

@pulumi.output_type
class GetFailoverIpAttachResult:
    """
    A collection of values returned by getFailoverIpAttach.
    """
    def __init__(__self__, block=None, continent_code=None, geo_loc=None, id=None, ip=None, progress=None, routed_to=None, service_name=None, status=None, sub_type=None):
        if block and not isinstance(block, str):
            raise TypeError("Expected argument 'block' to be a str")
        pulumi.set(__self__, "block", block)
        if continent_code and not isinstance(continent_code, str):
            raise TypeError("Expected argument 'continent_code' to be a str")
        pulumi.set(__self__, "continent_code", continent_code)
        if geo_loc and not isinstance(geo_loc, str):
            raise TypeError("Expected argument 'geo_loc' to be a str")
        pulumi.set(__self__, "geo_loc", geo_loc)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if progress and not isinstance(progress, int):
            raise TypeError("Expected argument 'progress' to be a int")
        pulumi.set(__self__, "progress", progress)
        if routed_to and not isinstance(routed_to, str):
            raise TypeError("Expected argument 'routed_to' to be a str")
        pulumi.set(__self__, "routed_to", routed_to)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if sub_type and not isinstance(sub_type, str):
            raise TypeError("Expected argument 'sub_type' to be a str")
        pulumi.set(__self__, "sub_type", sub_type)

    @property
    @pulumi.getter
    def block(self) -> str:
        """
        The IP block
        """
        return pulumi.get(self, "block")

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> str:
        return pulumi.get(self, "continent_code")

    @property
    @pulumi.getter(name="geoLoc")
    def geo_loc(self) -> str:
        return pulumi.get(self, "geo_loc")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The Ip id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The Ip Address
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def progress(self) -> int:
        """
        Current operation progress in percent
        """
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter(name="routedTo")
    def routed_to(self) -> str:
        return pulumi.get(self, "routed_to")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Ip status, can be `ok` or `operationPending`
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subType")
    def sub_type(self) -> str:
        return pulumi.get(self, "sub_type")


class AwaitableGetFailoverIpAttachResult(GetFailoverIpAttachResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFailoverIpAttachResult(
            block=self.block,
            continent_code=self.continent_code,
            geo_loc=self.geo_loc,
            id=self.id,
            ip=self.ip,
            progress=self.progress,
            routed_to=self.routed_to,
            service_name=self.service_name,
            status=self.status,
            sub_type=self.sub_type)


def get_failover_ip_attach(block: Optional[str] = None,
                           continent_code: Optional[str] = None,
                           geo_loc: Optional[str] = None,
                           ip: Optional[str] = None,
                           routed_to: Optional[str] = None,
                           service_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFailoverIpAttachResult:
    """
    Use this data source to get the details of a failover IP address of a service in a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    myfailoverip = ovh.CloudProject.get_failover_ip_attach(ip="XXXXXX",
        service_name="XXXXXX")
    ```


    :param str block: The IP block
    :param str ip: The failover ip address to query
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['block'] = block
    __args__['continentCode'] = continent_code
    __args__['geoLoc'] = geo_loc
    __args__['ip'] = ip
    __args__['routedTo'] = routed_to
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach', __args__, opts=opts, typ=GetFailoverIpAttachResult).value

    return AwaitableGetFailoverIpAttachResult(
        block=pulumi.get(__ret__, 'block'),
        continent_code=pulumi.get(__ret__, 'continent_code'),
        geo_loc=pulumi.get(__ret__, 'geo_loc'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        progress=pulumi.get(__ret__, 'progress'),
        routed_to=pulumi.get(__ret__, 'routed_to'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'),
        sub_type=pulumi.get(__ret__, 'sub_type'))


@_utilities.lift_output_func(get_failover_ip_attach)
def get_failover_ip_attach_output(block: Optional[pulumi.Input[Optional[str]]] = None,
                                  continent_code: Optional[pulumi.Input[Optional[str]]] = None,
                                  geo_loc: Optional[pulumi.Input[Optional[str]]] = None,
                                  ip: Optional[pulumi.Input[Optional[str]]] = None,
                                  routed_to: Optional[pulumi.Input[Optional[str]]] = None,
                                  service_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFailoverIpAttachResult]:
    """
    Use this data source to get the details of a failover IP address of a service in a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    myfailoverip = ovh.CloudProject.get_failover_ip_attach(ip="XXXXXX",
        service_name="XXXXXX")
    ```


    :param str block: The IP block
    :param str ip: The failover ip address to query
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
