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
from ._inputs import *

__all__ = ['NetworkPrivateSubnetArgs', 'NetworkPrivateSubnet']

@pulumi.input_type
class NetworkPrivateSubnetArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 network: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 region: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 start: pulumi.Input[str],
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 no_gateway: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NetworkPrivateSubnet resource.
        :param pulumi.Input[str] end: Last ip for this region.
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network: Global network in CIDR format.
               Changing this value recreates the subnet
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] start: First ip for this region.
               Changing this value recreates the subnet.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to false.
               _
        :param pulumi.Input[bool] no_gateway: Set to true if you don't want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to false.
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "start", start)
        if dhcp is not None:
            pulumi.set(__self__, "dhcp", dhcp)
        if no_gateway is not None:
            pulumi.set(__self__, "no_gateway", no_gateway)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        """
        Last ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        Global network in CIDR format.
        Changing this value recreates the subnet
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        """
        First ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def dhcp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to false.
        _
        """
        return pulumi.get(self, "dhcp")

    @dhcp.setter
    def dhcp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp", value)

    @property
    @pulumi.getter(name="noGateway")
    def no_gateway(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true if you don't want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to false.
        """
        return pulumi.get(self, "no_gateway")

    @no_gateway.setter
    def no_gateway(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_gateway", value)


@pulumi.input_type
class _NetworkPrivateSubnetState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetIpPoolArgs']]]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 no_gateway: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkPrivateSubnet resources.
        :param pulumi.Input[str] cidr: Ip Block representing the subnet cidr.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to false.
               _
        :param pulumi.Input[str] end: Last ip for this region.
               Changing this value recreates the subnet.
        :param pulumi.Input[str] gateway_ip: The IP of the gateway
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetIpPoolArgs']]] ip_pools: List of ip pools allocated in the subnet.
               * `ip_pools/network` - Global network with cidr.
               * `ip_pools/region` - Region where this subnet is created.
               * `ip_pools/dhcp` - DHCP enabled.
               * `ip_pools/end` - Last ip for this region.
               * `ip_pools/start` - First ip for this region.
        :param pulumi.Input[str] network: Global network in CIDR format.
               Changing this value recreates the subnet
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[bool] no_gateway: Set to true if you don't want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to false.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] start: First ip for this region.
               Changing this value recreates the subnet.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if dhcp is not None:
            pulumi.set(__self__, "dhcp", dhcp)
        if end is not None:
            pulumi.set(__self__, "end", end)
        if gateway_ip is not None:
            pulumi.set(__self__, "gateway_ip", gateway_ip)
        if ip_pools is not None:
            pulumi.set(__self__, "ip_pools", ip_pools)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if no_gateway is not None:
            pulumi.set(__self__, "no_gateway", no_gateway)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Ip Block representing the subnet cidr.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def dhcp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to false.
        _
        """
        return pulumi.get(self, "dhcp")

    @dhcp.setter
    def dhcp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp", value)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[str]]:
        """
        Last ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP of the gateway
        """
        return pulumi.get(self, "gateway_ip")

    @gateway_ip.setter
    def gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_ip", value)

    @property
    @pulumi.getter(name="ipPools")
    def ip_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetIpPoolArgs']]]]:
        """
        List of ip pools allocated in the subnet.
        * `ip_pools/network` - Global network with cidr.
        * `ip_pools/region` - Region where this subnet is created.
        * `ip_pools/dhcp` - DHCP enabled.
        * `ip_pools/end` - Last ip for this region.
        * `ip_pools/start` - First ip for this region.
        """
        return pulumi.get(self, "ip_pools")

    @ip_pools.setter
    def ip_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetIpPoolArgs']]]]):
        pulumi.set(self, "ip_pools", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Global network in CIDR format.
        Changing this value recreates the subnet
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="noGateway")
    def no_gateway(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true if you don't want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to false.
        """
        return pulumi.get(self, "no_gateway")

    @no_gateway.setter
    def no_gateway(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_gateway", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        """
        First ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)


class NetworkPrivateSubnet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 no_gateway: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a subnet in a private network of a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        subnet = ovh.cloud_project.NetworkPrivateSubnet("subnet",
            dhcp=True,
            end="192.168.168.200",
            network="192.168.168.0/24",
            network_id="0234543",
            no_gateway=False,
            region="GRA1",
            service_name="xxxxx",
            start="192.168.168.100")
        ```

        ## Import

        Subnet in a private network of a public cloud project can be imported using the `service_name` , the `network_id` and the `subnet_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to false.
               _
        :param pulumi.Input[str] end: Last ip for this region.
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network: Global network in CIDR format.
               Changing this value recreates the subnet
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[bool] no_gateway: Set to true if you don't want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to false.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] start: First ip for this region.
               Changing this value recreates the subnet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkPrivateSubnetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a subnet in a private network of a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        subnet = ovh.cloud_project.NetworkPrivateSubnet("subnet",
            dhcp=True,
            end="192.168.168.200",
            network="192.168.168.0/24",
            network_id="0234543",
            no_gateway=False,
            region="GRA1",
            service_name="xxxxx",
            start="192.168.168.100")
        ```

        ## Import

        Subnet in a private network of a public cloud project can be imported using the `service_name` , the `network_id` and the `subnet_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909
        ```

        :param str resource_name: The name of the resource.
        :param NetworkPrivateSubnetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkPrivateSubnetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 no_gateway: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkPrivateSubnetArgs.__new__(NetworkPrivateSubnetArgs)

            __props__.__dict__["dhcp"] = dhcp
            if end is None and not opts.urn:
                raise TypeError("Missing required property 'end'")
            __props__.__dict__["end"] = end
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["no_gateway"] = no_gateway
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if start is None and not opts.urn:
                raise TypeError("Missing required property 'start'")
            __props__.__dict__["start"] = start
            __props__.__dict__["cidr"] = None
            __props__.__dict__["gateway_ip"] = None
            __props__.__dict__["ip_pools"] = None
        super(NetworkPrivateSubnet, __self__).__init__(
            'ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            dhcp: Optional[pulumi.Input[bool]] = None,
            end: Optional[pulumi.Input[str]] = None,
            gateway_ip: Optional[pulumi.Input[str]] = None,
            ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPrivateSubnetIpPoolArgs']]]]] = None,
            network: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            no_gateway: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            start: Optional[pulumi.Input[str]] = None) -> 'NetworkPrivateSubnet':
        """
        Get an existing NetworkPrivateSubnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: Ip Block representing the subnet cidr.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to false.
               _
        :param pulumi.Input[str] end: Last ip for this region.
               Changing this value recreates the subnet.
        :param pulumi.Input[str] gateway_ip: The IP of the gateway
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkPrivateSubnetIpPoolArgs']]]] ip_pools: List of ip pools allocated in the subnet.
               * `ip_pools/network` - Global network with cidr.
               * `ip_pools/region` - Region where this subnet is created.
               * `ip_pools/dhcp` - DHCP enabled.
               * `ip_pools/end` - Last ip for this region.
               * `ip_pools/start` - First ip for this region.
        :param pulumi.Input[str] network: Global network in CIDR format.
               Changing this value recreates the subnet
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[bool] no_gateway: Set to true if you don't want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to false.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] start: First ip for this region.
               Changing this value recreates the subnet.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkPrivateSubnetState.__new__(_NetworkPrivateSubnetState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["dhcp"] = dhcp
        __props__.__dict__["end"] = end
        __props__.__dict__["gateway_ip"] = gateway_ip
        __props__.__dict__["ip_pools"] = ip_pools
        __props__.__dict__["network"] = network
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["no_gateway"] = no_gateway
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["start"] = start
        return NetworkPrivateSubnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        Ip Block representing the subnet cidr.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def dhcp(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to false.
        _
        """
        return pulumi.get(self, "dhcp")

    @property
    @pulumi.getter
    def end(self) -> pulumi.Output[str]:
        """
        Last ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> pulumi.Output[str]:
        """
        The IP of the gateway
        """
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter(name="ipPools")
    def ip_pools(self) -> pulumi.Output[Sequence['outputs.NetworkPrivateSubnetIpPool']]:
        """
        List of ip pools allocated in the subnet.
        * `ip_pools/network` - Global network with cidr.
        * `ip_pools/region` - Region where this subnet is created.
        * `ip_pools/dhcp` - DHCP enabled.
        * `ip_pools/end` - Last ip for this region.
        * `ip_pools/start` - First ip for this region.
        """
        return pulumi.get(self, "ip_pools")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        Global network in CIDR format.
        Changing this value recreates the subnet
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="noGateway")
    def no_gateway(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to true if you don't want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to false.
        """
        return pulumi.get(self, "no_gateway")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[str]:
        """
        First ip for this region.
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "start")

