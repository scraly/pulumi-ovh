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
    'GetVRackResult',
    'AwaitableGetVRackResult',
    'get_v_rack',
    'get_v_rack_output',
]

@pulumi.output_type
class GetVRackResult:
    """
    A collection of values returned by getVRack.
    """
    def __init__(__self__, description=None, id=None, name=None, service_name=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the vrack
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the vrack
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The id of the public cloud project
        """
        return pulumi.get(self, "service_name")


class AwaitableGetVRackResult(GetVRackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVRackResult(
            description=self.description,
            id=self.id,
            name=self.name,
            service_name=self.service_name)


def get_v_rack(service_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVRackResult:
    """
    Use this data source to get the linked vrack on your public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    vrack_v_rack = ovh.CloudProject.get_v_rack(service_name="XXXXXX")
    pulumi.export("vrack", vrack_v_rack)
    ```


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getVRack:getVRack', __args__, opts=opts, typ=GetVRackResult).value

    return AwaitableGetVRackResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        service_name=pulumi.get(__ret__, 'service_name'))


@_utilities.lift_output_func(get_v_rack)
def get_v_rack_output(service_name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVRackResult]:
    """
    Use this data source to get the linked vrack on your public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    vrack_v_rack = ovh.CloudProject.get_v_rack(service_name="XXXXXX")
    pulumi.export("vrack", vrack_v_rack)
    ```


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...