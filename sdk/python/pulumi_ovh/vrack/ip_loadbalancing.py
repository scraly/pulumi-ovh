# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IpLoadbalancingArgs', 'IpLoadbalancing']

@pulumi.input_type
class IpLoadbalancingArgs:
    def __init__(__self__, *,
                 loadbalancing_id: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IpLoadbalancing resource.
        :param pulumi.Input[str] loadbalancing_id: The id of the IP Load Balancing.
        :param pulumi.Input[str] service_name: The id of the vrack.
        """
        pulumi.set(__self__, "loadbalancing_id", loadbalancing_id)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="LoadbalancingId")
    def loadbalancing_id(self) -> pulumi.Input[str]:
        """
        The id of the IP Load Balancing.
        """
        return pulumi.get(self, "loadbalancing_id")

    @loadbalancing_id.setter
    def loadbalancing_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "loadbalancing_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the vrack.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _IpLoadbalancingState:
    def __init__(__self__, *,
                 loadbalancing_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpLoadbalancing resources.
        :param pulumi.Input[str] loadbalancing_id: The id of the IP Load Balancing.
        :param pulumi.Input[str] service_name: The id of the vrack.
        """
        if loadbalancing_id is not None:
            pulumi.set(__self__, "loadbalancing_id", loadbalancing_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="LoadbalancingId")
    def loadbalancing_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the IP Load Balancing.
        """
        return pulumi.get(self, "loadbalancing_id")

    @loadbalancing_id.setter
    def loadbalancing_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loadbalancing_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vrack.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class IpLoadbalancing(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 loadbalancing_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach an IP Load Balancing to a VRack.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] loadbalancing_id: The id of the IP Load Balancing.
        :param pulumi.Input[str] service_name: The id of the vrack.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpLoadbalancingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach an IP Load Balancing to a VRack.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param IpLoadbalancingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpLoadbalancingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 loadbalancing_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpLoadbalancingArgs.__new__(IpLoadbalancingArgs)

            if loadbalancing_id is None and not opts.urn:
                raise TypeError("Missing required property 'loadbalancing_id'")
            __props__.__dict__["loadbalancing_id"] = loadbalancing_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(IpLoadbalancing, __self__).__init__(
            'ovh:Vrack/ipLoadbalancing:IpLoadbalancing',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            loadbalancing_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'IpLoadbalancing':
        """
        Get an existing IpLoadbalancing resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] loadbalancing_id: The id of the IP Load Balancing.
        :param pulumi.Input[str] service_name: The id of the vrack.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpLoadbalancingState.__new__(_IpLoadbalancingState)

        __props__.__dict__["loadbalancing_id"] = loadbalancing_id
        __props__.__dict__["service_name"] = service_name
        return IpLoadbalancing(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="LoadbalancingId")
    def loadbalancing_id(self) -> pulumi.Output[str]:
        """
        The id of the IP Load Balancing.
        """
        return pulumi.get(self, "loadbalancing_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the vrack.
        """
        return pulumi.get(self, "service_name")
