# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RefreshArgs', 'Refresh']

@pulumi.input_type
class RefreshArgs:
    def __init__(__self__, *,
                 keepers: pulumi.Input[Sequence[pulumi.Input[str]]],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a Refresh resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keepers: List of values tracked to trigger refresh, used also to form implicit dependencies
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        """
        pulumi.set(__self__, "keepers", keepers)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of values tracked to trigger refresh, used also to form implicit dependencies
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _RefreshState:
    def __init__(__self__, *,
                 keepers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Refresh resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keepers: List of values tracked to trigger refresh, used also to form implicit dependencies
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        """
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values tracked to trigger refresh, used also to form implicit dependencies
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class Refresh(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Applies changes from other `ovh_iploadbalancing_*` resources to the production configuration of loadbalancers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.TcpFarm("farmname",
            port=8080,
            service_name=lb.service_name,
            zone="all")
        backend = ovh.ip_load_balancing.TcpFarmServer("backend",
            address="4.5.6.7",
            backup=True,
            display_name="mybackend",
            farm_id=farmname.id,
            port=80,
            probe=True,
            proxy_protocol_version="v2",
            service_name=lb.service_name,
            ssl=False,
            status="active",
            weight=2)
        mylb = ovh.ip_load_balancing.Refresh("mylb",
            keepers=[[__item.address for __item in [backend]]],
            service_name=lb.service_name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keepers: List of values tracked to trigger refresh, used also to form implicit dependencies
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RefreshArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Applies changes from other `ovh_iploadbalancing_*` resources to the production configuration of loadbalancers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.TcpFarm("farmname",
            port=8080,
            service_name=lb.service_name,
            zone="all")
        backend = ovh.ip_load_balancing.TcpFarmServer("backend",
            address="4.5.6.7",
            backup=True,
            display_name="mybackend",
            farm_id=farmname.id,
            port=80,
            probe=True,
            proxy_protocol_version="v2",
            service_name=lb.service_name,
            ssl=False,
            status="active",
            weight=2)
        mylb = ovh.ip_load_balancing.Refresh("mylb",
            keepers=[[__item.address for __item in [backend]]],
            service_name=lb.service_name)
        ```

        :param str resource_name: The name of the resource.
        :param RefreshArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RefreshArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RefreshArgs.__new__(RefreshArgs)

            if keepers is None and not opts.urn:
                raise TypeError("Missing required property 'keepers'")
            __props__.__dict__["keepers"] = keepers
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(Refresh, __self__).__init__(
            'ovh:IpLoadBalancing/refresh:Refresh',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            keepers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'Refresh':
        """
        Get an existing Refresh resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keepers: List of values tracked to trigger refresh, used also to form implicit dependencies
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RefreshState.__new__(_RefreshState)

        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["service_name"] = service_name
        return Refresh(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Sequence[str]]:
        """
        List of values tracked to trigger refresh, used also to form implicit dependencies
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

