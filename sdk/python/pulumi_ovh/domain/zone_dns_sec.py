# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ZoneDNSSecArgs', 'ZoneDNSSec']

@pulumi.input_type
class ZoneDNSSecArgs:
    def __init__(__self__, *,
                 zone_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ZoneDNSSec resource.
        :param pulumi.Input[str] zone_name: The name of the domain zone
        """
        pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Input[str]:
        """
        The name of the domain zone
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_name", value)


@pulumi.input_type
class _ZoneDNSSecState:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZoneDNSSec resources.
        :param pulumi.Input[str] status: DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
        :param pulumi.Input[str] zone_name: The name of the domain zone
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if zone_name is not None:
            pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the domain zone
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_name", value)


class ZoneDNSSec(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Enable / disable DNSSEC on a domain zone.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        dnssec = ovh.domain.ZoneDNSSec("dnssec", zone_name="mysite.ovh")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] zone_name: The name of the domain zone
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneDNSSecArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enable / disable DNSSEC on a domain zone.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        dnssec = ovh.domain.ZoneDNSSec("dnssec", zone_name="mysite.ovh")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ZoneDNSSecArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneDNSSecArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneDNSSecArgs.__new__(ZoneDNSSecArgs)

            if zone_name is None and not opts.urn:
                raise TypeError("Missing required property 'zone_name'")
            __props__.__dict__["zone_name"] = zone_name
            __props__.__dict__["status"] = None
        super(ZoneDNSSec, __self__).__init__(
            'ovh:Domain/zoneDNSSec:ZoneDNSSec',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            status: Optional[pulumi.Input[str]] = None,
            zone_name: Optional[pulumi.Input[str]] = None) -> 'ZoneDNSSec':
        """
        Get an existing ZoneDNSSec resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
        :param pulumi.Input[str] zone_name: The name of the domain zone
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneDNSSecState.__new__(_ZoneDNSSecState)

        __props__.__dict__["status"] = status
        __props__.__dict__["zone_name"] = zone_name
        return ZoneDNSSec(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Output[str]:
        """
        The name of the domain zone
        """
        return pulumi.get(self, "zone_name")

