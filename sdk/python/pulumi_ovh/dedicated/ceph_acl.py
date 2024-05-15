# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CephAclArgs', 'CephAcl']

@pulumi.input_type
class CephAclArgs:
    def __init__(__self__, *,
                 netmask: pulumi.Input[str],
                 network: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a CephAcl resource.
        :param pulumi.Input[str] netmask: The network mask to apply
        :param pulumi.Input[str] network: The network IP to authorize
        :param pulumi.Input[str] service_name: The internal name of your dedicated CEPH
        """
        pulumi.set(__self__, "netmask", netmask)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Input[str]:
        """
        The network mask to apply
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: pulumi.Input[str]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        The network IP to authorize
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your dedicated CEPH
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _CephAclState:
    def __init__(__self__, *,
                 family: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CephAcl resources.
        :param pulumi.Input[str] family: IP family. `IPv4` or `IPv6`
        :param pulumi.Input[str] netmask: The network mask to apply
        :param pulumi.Input[str] network: The network IP to authorize
        :param pulumi.Input[str] service_name: The internal name of your dedicated CEPH
        """
        if family is not None:
            pulumi.set(__self__, "family", family)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def family(self) -> Optional[pulumi.Input[str]]:
        """
        IP family. `IPv4` or `IPv6`
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The network mask to apply
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The network IP to authorize
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your dedicated CEPH
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class CephAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Add a new access ACL for the given network/mask.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_ceph = ovh.Dedicated.get_ceph(service_name="94d423da-0e55-45f2-9812-836460a19939")
        my_acl = ovh.dedicated.CephAcl("my-acl",
            service_name=my_ceph.id,
            network="1.2.3.4",
            netmask="255.255.255.255")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] netmask: The network mask to apply
        :param pulumi.Input[str] network: The network IP to authorize
        :param pulumi.Input[str] service_name: The internal name of your dedicated CEPH
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CephAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Add a new access ACL for the given network/mask.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_ceph = ovh.Dedicated.get_ceph(service_name="94d423da-0e55-45f2-9812-836460a19939")
        my_acl = ovh.dedicated.CephAcl("my-acl",
            service_name=my_ceph.id,
            network="1.2.3.4",
            netmask="255.255.255.255")
        ```

        :param str resource_name: The name of the resource.
        :param CephAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CephAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CephAclArgs.__new__(CephAclArgs)

            if netmask is None and not opts.urn:
                raise TypeError("Missing required property 'netmask'")
            __props__.__dict__["netmask"] = netmask
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["family"] = None
        super(CephAcl, __self__).__init__(
            'ovh:Dedicated/cephAcl:CephAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            family: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'CephAcl':
        """
        Get an existing CephAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] family: IP family. `IPv4` or `IPv6`
        :param pulumi.Input[str] netmask: The network mask to apply
        :param pulumi.Input[str] network: The network IP to authorize
        :param pulumi.Input[str] service_name: The internal name of your dedicated CEPH
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CephAclState.__new__(_CephAclState)

        __props__.__dict__["family"] = family
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["network"] = network
        __props__.__dict__["service_name"] = service_name
        return CephAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def family(self) -> pulumi.Output[str]:
        """
        IP family. `IPv4` or `IPv6`
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[str]:
        """
        The network mask to apply
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The network IP to authorize
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your dedicated CEPH
        """
        return pulumi.get(self, "service_name")

