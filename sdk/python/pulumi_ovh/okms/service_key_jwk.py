# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceKeyJWKArgs', 'ServiceKeyJWK']

@pulumi.input_type
class ServiceKeyJWKArgs:
    def __init__(__self__, *,
                 keys: pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]],
                 okms_id: pulumi.Input[str],
                 context: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceKeyJWK resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]] keys: Set of JSON Web Keys to import
        :param pulumi.Input[str] okms_id: Okms ID
        :param pulumi.Input[str] context: Context of the key
        :param pulumi.Input[str] name: Key name
        """
        pulumi.set(__self__, "keys", keys)
        pulumi.set(__self__, "okms_id", okms_id)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]]:
        """
        Set of JSON Web Keys to import
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter(name="okmsId")
    def okms_id(self) -> pulumi.Input[str]:
        """
        Okms ID
        """
        return pulumi.get(self, "okms_id")

    @okms_id.setter
    def okms_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "okms_id", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """
        Context of the key
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Key name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ServiceKeyJWKState:
    def __init__(__self__, *,
                 context: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 deactivation_reason: Optional[pulumi.Input[str]] = None,
                 iam: Optional[pulumi.Input['ServiceKeyJWKIamArgs']] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 okms_id: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[float]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceKeyJWK resources.
        :param pulumi.Input[str] context: Context of the key
        :param pulumi.Input[str] created_at: Creation time of the key
        :param pulumi.Input[str] deactivation_reason: Key deactivation reason
        :param pulumi.Input['ServiceKeyJWKIamArgs'] iam: IAM resource metadata
        :param pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]] keys: Set of JSON Web Keys to import
        :param pulumi.Input[str] name: Key name
        :param pulumi.Input[str] okms_id: Okms ID
        :param pulumi.Input[float] size: Size of the key to be created
        :param pulumi.Input[str] state: State of the key
        :param pulumi.Input[str] type: Type of the key to be created
        """
        if context is not None:
            pulumi.set(__self__, "context", context)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if deactivation_reason is not None:
            pulumi.set(__self__, "deactivation_reason", deactivation_reason)
        if iam is not None:
            pulumi.set(__self__, "iam", iam)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if okms_id is not None:
            pulumi.set(__self__, "okms_id", okms_id)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """
        Context of the key
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of the key
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deactivationReason")
    def deactivation_reason(self) -> Optional[pulumi.Input[str]]:
        """
        Key deactivation reason
        """
        return pulumi.get(self, "deactivation_reason")

    @deactivation_reason.setter
    def deactivation_reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deactivation_reason", value)

    @property
    @pulumi.getter
    def iam(self) -> Optional[pulumi.Input['ServiceKeyJWKIamArgs']]:
        """
        IAM resource metadata
        """
        return pulumi.get(self, "iam")

    @iam.setter
    def iam(self, value: Optional[pulumi.Input['ServiceKeyJWKIamArgs']]):
        pulumi.set(self, "iam", value)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]]]:
        """
        Set of JSON Web Keys to import
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceKeyJWKKeyArgs']]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Key name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="okmsId")
    def okms_id(self) -> Optional[pulumi.Input[str]]:
        """
        Okms ID
        """
        return pulumi.get(self, "okms_id")

    @okms_id.setter
    def okms_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "okms_id", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[float]]:
        """
        Size of the key to be created
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the key
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the key to be created
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ServiceKeyJWK(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceKeyJWKKeyArgs', 'ServiceKeyJWKKeyArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 okms_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Import an existing key in the JWK format in an OVHcloud KMS.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context: Context of the key
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServiceKeyJWKKeyArgs', 'ServiceKeyJWKKeyArgsDict']]]] keys: Set of JSON Web Keys to import
        :param pulumi.Input[str] name: Key name
        :param pulumi.Input[str] okms_id: Okms ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceKeyJWKArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Import an existing key in the JWK format in an OVHcloud KMS.

        :param str resource_name: The name of the resource.
        :param ServiceKeyJWKArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceKeyJWKArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceKeyJWKKeyArgs', 'ServiceKeyJWKKeyArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 okms_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceKeyJWKArgs.__new__(ServiceKeyJWKArgs)

            __props__.__dict__["context"] = context
            if keys is None and not opts.urn:
                raise TypeError("Missing required property 'keys'")
            __props__.__dict__["keys"] = keys
            __props__.__dict__["name"] = name
            if okms_id is None and not opts.urn:
                raise TypeError("Missing required property 'okms_id'")
            __props__.__dict__["okms_id"] = okms_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["deactivation_reason"] = None
            __props__.__dict__["iam"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
        super(ServiceKeyJWK, __self__).__init__(
            'ovh:Okms/serviceKeyJWK:ServiceKeyJWK',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            context: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            deactivation_reason: Optional[pulumi.Input[str]] = None,
            iam: Optional[pulumi.Input[Union['ServiceKeyJWKIamArgs', 'ServiceKeyJWKIamArgsDict']]] = None,
            keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceKeyJWKKeyArgs', 'ServiceKeyJWKKeyArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            okms_id: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[float]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ServiceKeyJWK':
        """
        Get an existing ServiceKeyJWK resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context: Context of the key
        :param pulumi.Input[str] created_at: Creation time of the key
        :param pulumi.Input[str] deactivation_reason: Key deactivation reason
        :param pulumi.Input[Union['ServiceKeyJWKIamArgs', 'ServiceKeyJWKIamArgsDict']] iam: IAM resource metadata
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServiceKeyJWKKeyArgs', 'ServiceKeyJWKKeyArgsDict']]]] keys: Set of JSON Web Keys to import
        :param pulumi.Input[str] name: Key name
        :param pulumi.Input[str] okms_id: Okms ID
        :param pulumi.Input[float] size: Size of the key to be created
        :param pulumi.Input[str] state: State of the key
        :param pulumi.Input[str] type: Type of the key to be created
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceKeyJWKState.__new__(_ServiceKeyJWKState)

        __props__.__dict__["context"] = context
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["deactivation_reason"] = deactivation_reason
        __props__.__dict__["iam"] = iam
        __props__.__dict__["keys"] = keys
        __props__.__dict__["name"] = name
        __props__.__dict__["okms_id"] = okms_id
        __props__.__dict__["size"] = size
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        return ServiceKeyJWK(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[str]]:
        """
        Context of the key
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation time of the key
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deactivationReason")
    def deactivation_reason(self) -> pulumi.Output[str]:
        """
        Key deactivation reason
        """
        return pulumi.get(self, "deactivation_reason")

    @property
    @pulumi.getter
    def iam(self) -> pulumi.Output['outputs.ServiceKeyJWKIam']:
        """
        IAM resource metadata
        """
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Output[Sequence['outputs.ServiceKeyJWKKey']]:
        """
        Set of JSON Web Keys to import
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Key name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="okmsId")
    def okms_id(self) -> pulumi.Output[str]:
        """
        Okms ID
        """
        return pulumi.get(self, "okms_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[float]:
        """
        Size of the key to be created
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the key
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the key to be created
        """
        return pulumi.get(self, "type")
