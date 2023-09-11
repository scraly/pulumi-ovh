# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateDatabaseUserArgs', 'PrivateDatabaseUser']

@pulumi.input_type
class PrivateDatabaseUserArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 user_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a PrivateDatabaseUser resource.
        :param pulumi.Input[str] password: Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        User name used to connect on your databases
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class _PrivateDatabaseUserState:
    def __init__(__self__, *,
                 password: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateDatabaseUser resources.
        :param pulumi.Input[str] password: Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases
        """
        if password is not None:
            pulumi.set(__self__, "password", password)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        User name used to connect on your databases
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class PrivateDatabaseUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a new user on your private cloud database instance.

        ## Example Usage

        ## Import

        OVHcloud database user can be imported using the `service_name` and the `user_name`, separated by "/" E.g.,

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] password: Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateDatabaseUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new user on your private cloud database instance.

        ## Example Usage

        ## Import

        OVHcloud database user can be imported using the `service_name` and the `user_name`, separated by "/" E.g.,

        :param str resource_name: The name of the resource.
        :param PrivateDatabaseUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateDatabaseUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateDatabaseUserArgs.__new__(PrivateDatabaseUserArgs)

            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PrivateDatabaseUser, __self__).__init__(
            'ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            password: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'PrivateDatabaseUser':
        """
        Get an existing PrivateDatabaseUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] password: Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateDatabaseUserState.__new__(_PrivateDatabaseUserState)

        __props__.__dict__["password"] = password
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["user_name"] = user_name
        return PrivateDatabaseUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        User name used to connect on your databases
        """
        return pulumi.get(self, "user_name")

