# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetVolumeResult',
    'AwaitableGetVolumeResult',
    'get_volume',
    'get_volume_output',
]

@pulumi.output_type
class GetVolumeResult:
    """
    A collection of values returned by getVolume.
    """
    def __init__(__self__, id=None, name=None, region_name=None, service_name=None, size=None, volume_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region_name and not isinstance(region_name, str):
            raise TypeError("Expected argument 'region_name' to be a str")
        pulumi.set(__self__, "region_name", region_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        pulumi.set(__self__, "size", size)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the volume (E.g.: "GRA", meaning Gravelines, for region "GRA1")
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        """
        The region name where volume is available
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The id of the public cloud project.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> float:
        """
        The size of the volume
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
        """
        The id of the volume
        """
        return pulumi.get(self, "volume_id")


class AwaitableGetVolumeResult(GetVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeResult(
            id=self.id,
            name=self.name,
            region_name=self.region_name,
            service_name=self.service_name,
            size=self.size,
            volume_id=self.volume_id)


def get_volume(region_name: Optional[str] = None,
               service_name: Optional[str] = None,
               volume_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeResult:
    """
    Get information about a volume in a public cloud project

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    volume = ovh.CloudProject.get_volume(region_name="xxx",
        service_name="yyy",
        volume_id="aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
    ```


    :param str region_name: A valid OVHcloud public cloud region name in which the volume is available. Ex.: "GRA11".
    :param str service_name: The id of the public cloud project.
    :param str volume_id: Volume id to get the informations
    """
    __args__ = dict()
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getVolume:getVolume', __args__, opts=opts, typ=GetVolumeResult).value

    return AwaitableGetVolumeResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region_name=pulumi.get(__ret__, 'region_name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        size=pulumi.get(__ret__, 'size'),
        volume_id=pulumi.get(__ret__, 'volume_id'))
def get_volume_output(region_name: Optional[pulumi.Input[str]] = None,
                      service_name: Optional[pulumi.Input[str]] = None,
                      volume_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVolumeResult]:
    """
    Get information about a volume in a public cloud project

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    volume = ovh.CloudProject.get_volume(region_name="xxx",
        service_name="yyy",
        volume_id="aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
    ```


    :param str region_name: A valid OVHcloud public cloud region name in which the volume is available. Ex.: "GRA11".
    :param str service_name: The id of the public cloud project.
    :param str volume_id: Volume id to get the informations
    """
    __args__ = dict()
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getVolume:getVolume', __args__, opts=opts, typ=GetVolumeResult)
    return __ret__.apply(lambda __response__: GetVolumeResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region_name=pulumi.get(__response__, 'region_name'),
        service_name=pulumi.get(__response__, 'service_name'),
        size=pulumi.get(__response__, 'size'),
        volume_id=pulumi.get(__response__, 'volume_id')))