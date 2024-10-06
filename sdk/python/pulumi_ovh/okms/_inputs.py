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
    'OkmsIamArgs',
]

@pulumi.input_type
class OkmsIamArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 urn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] display_name: (String) Resource display name
        :param pulumi.Input[str] id: (String) Unique identifier of the resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        :param pulumi.Input[str] urn: (String) Unique resource name used in policies
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        (String) Resource display name
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        (String) Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def urn(self) -> Optional[pulumi.Input[str]]:
        """
        (String) Unique resource name used in policies
        """
        return pulumi.get(self, "urn")

    @urn.setter
    def urn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "urn", value)

