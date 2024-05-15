# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateDatabaseUserGrantArgs', 'PrivateDatabaseUserGrant']

@pulumi.input_type
class PrivateDatabaseUserGrantArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 grant: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 user_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a PrivateDatabaseUserGrant resource.
        :param pulumi.Input[str] database_name: Database name where add grant.
        :param pulumi.Input[str] grant: Database name where add grant. Values can be: 
               - admin
               - none
               - ro
               - rw
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "grant", grant)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Database name where add grant.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def grant(self) -> pulumi.Input[str]:
        """
        Database name where add grant. Values can be: 
        - admin
        - none
        - ro
        - rw
        """
        return pulumi.get(self, "grant")

    @grant.setter
    def grant(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant", value)

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
        User name used to connect on your databases.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class _PrivateDatabaseUserGrantState:
    def __init__(__self__, *,
                 database_name: Optional[pulumi.Input[str]] = None,
                 grant: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateDatabaseUserGrant resources.
        :param pulumi.Input[str] database_name: Database name where add grant.
        :param pulumi.Input[str] grant: Database name where add grant. Values can be: 
               - admin
               - none
               - ro
               - rw
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases.
        """
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if grant is not None:
            pulumi.set(__self__, "grant", grant)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database name where add grant.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def grant(self) -> Optional[pulumi.Input[str]]:
        """
        Database name where add grant. Values can be: 
        - admin
        - none
        - ro
        - rw
        """
        return pulumi.get(self, "grant")

    @grant.setter
    def grant(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant", value)

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
        User name used to connect on your databases.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class PrivateDatabaseUserGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 grant: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Add grant on a database in your private cloud database instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        user_grant = ovh.hosting.PrivateDatabaseUserGrant("userGrant",
            database_name="ovhcloud",
            grant="admin",
            service_name="XXXXXX",
            user_name="terraform")
        ```

        ## Import

        OVHcloud database user's grant can be imported using the `service_name`, the `user_name`, the `database_name` and the `grant`, separated by "/" E.g.,

        ```sh
        $ pulumi import ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant user service_name/user_name/database_name/grant
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Database name where add grant.
        :param pulumi.Input[str] grant: Database name where add grant. Values can be: 
               - admin
               - none
               - ro
               - rw
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateDatabaseUserGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Add grant on a database in your private cloud database instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        user_grant = ovh.hosting.PrivateDatabaseUserGrant("userGrant",
            database_name="ovhcloud",
            grant="admin",
            service_name="XXXXXX",
            user_name="terraform")
        ```

        ## Import

        OVHcloud database user's grant can be imported using the `service_name`, the `user_name`, the `database_name` and the `grant`, separated by "/" E.g.,

        ```sh
        $ pulumi import ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant user service_name/user_name/database_name/grant
        ```

        :param str resource_name: The name of the resource.
        :param PrivateDatabaseUserGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateDatabaseUserGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 grant: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateDatabaseUserGrantArgs.__new__(PrivateDatabaseUserGrantArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if grant is None and not opts.urn:
                raise TypeError("Missing required property 'grant'")
            __props__.__dict__["grant"] = grant
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(PrivateDatabaseUserGrant, __self__).__init__(
            'ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            grant: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'PrivateDatabaseUserGrant':
        """
        Get an existing PrivateDatabaseUserGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Database name where add grant.
        :param pulumi.Input[str] grant: Database name where add grant. Values can be: 
               - admin
               - none
               - ro
               - rw
        :param pulumi.Input[str] service_name: The internal name of your private database.
        :param pulumi.Input[str] user_name: User name used to connect on your databases.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateDatabaseUserGrantState.__new__(_PrivateDatabaseUserGrantState)

        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["grant"] = grant
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["user_name"] = user_name
        return PrivateDatabaseUserGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Database name where add grant.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def grant(self) -> pulumi.Output[str]:
        """
        Database name where add grant. Values can be: 
        - admin
        - none
        - ro
        - rw
        """
        return pulumi.get(self, "grant")

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
        User name used to connect on your databases.
        """
        return pulumi.get(self, "user_name")

