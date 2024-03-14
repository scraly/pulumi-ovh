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
    'GetResourceGroupResult',
    'AwaitableGetResourceGroupResult',
    'get_resource_group',
    'get_resource_group_output',
]

@pulumi.output_type
class GetResourceGroupResult:
    """
    A collection of values returned by getResourceGroup.
    """
    def __init__(__self__, group_urn=None, created_at=None, id=None, name=None, owner=None, read_only=None, resources=None, updated_at=None):
        if group_urn and not isinstance(group_urn, str):
            raise TypeError("Expected argument 'group_urn' to be a str")
        pulumi.set(__self__, "group_urn", group_urn)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if read_only and not isinstance(read_only, bool):
            raise TypeError("Expected argument 'read_only' to be a bool")
        pulumi.set(__self__, "read_only", read_only)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="GroupURN")
    def group_urn(self) -> str:
        """
        URN of the resource group, used when writing policies
        """
        return pulumi.get(self, "group_urn")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of the creation of the resource group
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Name of the account owning the resource group
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> bool:
        """
        Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
        """
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter
    def resources(self) -> Sequence[str]:
        """
        Set of the URNs of the resources contained in the resource group
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date of the last modification of the resource group
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetResourceGroupResult(GetResourceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceGroupResult(
            group_urn=self.group_urn,
            created_at=self.created_at,
            id=self.id,
            name=self.name,
            owner=self.owner,
            read_only=self.read_only,
            resources=self.resources,
            updated_at=self.updated_at)


def get_resource_group(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceGroupResult:
    """
    Use this data source get details about a resource group.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_resource_group = ovh.Iam.get_resource_group(id="my_resource_group_id")
    ```
    <!--End PulumiCodeChooser -->


    :param str id: Id of the resource group
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Iam/getResourceGroup:getResourceGroup', __args__, opts=opts, typ=GetResourceGroupResult).value

    return AwaitableGetResourceGroupResult(
        group_urn=pulumi.get(__ret__, 'group_urn'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        read_only=pulumi.get(__ret__, 'read_only'),
        resources=pulumi.get(__ret__, 'resources'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_resource_group)
def get_resource_group_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceGroupResult]:
    """
    Use this data source get details about a resource group.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_resource_group = ovh.Iam.get_resource_group(id="my_resource_group_id")
    ```
    <!--End PulumiCodeChooser -->


    :param str id: Id of the resource group
    """
    ...
