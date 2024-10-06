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
    'OkmsIam',
    'GetOkmsResourceIamResult',
    'GetOkmsServiceKeyJwkKeyResult',
]

@pulumi.output_type
class OkmsIam(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayName":
            suggest = "display_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OkmsIam. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OkmsIam.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OkmsIam.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 display_name: Optional[str] = None,
                 id: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 urn: Optional[str] = None):
        """
        :param str display_name: (String) Resource display name
        :param str id: (String) Unique identifier of the resource
        :param Mapping[str, str] tags: (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: (String) Unique resource name used in policies
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
    def display_name(self) -> Optional[str]:
        """
        (String) Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        (String) Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        """
        (String) Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsResourceIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: (String) Resource display name
        :param str id: Should be set to the ID of your KMS
        :param Mapping[str, str] tags: (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: (String) Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        (String) Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Should be set to the ID of your KMS
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        (Map of String) Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        (String) Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsServiceKeyJwkKeyResult(dict):
    def __init__(__self__, *,
                 alg: str,
                 crv: str,
                 e: str,
                 key_ops: Sequence[str],
                 kid: str,
                 kty: str,
                 n: str,
                 use: str,
                 x: str,
                 y: str):
        """
        :param str alg: The algorithm intended to be used with the key
        :param str crv: The cryptographic curve used with the key
        :param str e: The exponent value for the RSA public key
        :param Sequence[str] key_ops: The operation for which the key is intended to be used
        :param str kid: key ID parameter used to match a specific key
        :param str kty: Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        :param str n: The modulus value for the RSA public key
        :param str use: The intended use of the public key
        :param str x: The x coordinate for the Elliptic Curve point
        :param str y: The y coordinate for the Elliptic Curve point
        """
        pulumi.set(__self__, "alg", alg)
        pulumi.set(__self__, "crv", crv)
        pulumi.set(__self__, "e", e)
        pulumi.set(__self__, "key_ops", key_ops)
        pulumi.set(__self__, "kid", kid)
        pulumi.set(__self__, "kty", kty)
        pulumi.set(__self__, "n", n)
        pulumi.set(__self__, "use", use)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def alg(self) -> str:
        """
        The algorithm intended to be used with the key
        """
        return pulumi.get(self, "alg")

    @property
    @pulumi.getter
    def crv(self) -> str:
        """
        The cryptographic curve used with the key
        """
        return pulumi.get(self, "crv")

    @property
    @pulumi.getter
    def e(self) -> str:
        """
        The exponent value for the RSA public key
        """
        return pulumi.get(self, "e")

    @property
    @pulumi.getter(name="keyOps")
    def key_ops(self) -> Sequence[str]:
        """
        The operation for which the key is intended to be used
        """
        return pulumi.get(self, "key_ops")

    @property
    @pulumi.getter
    def kid(self) -> str:
        """
        key ID parameter used to match a specific key
        """
        return pulumi.get(self, "kid")

    @property
    @pulumi.getter
    def kty(self) -> str:
        """
        Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        """
        return pulumi.get(self, "kty")

    @property
    @pulumi.getter
    def n(self) -> str:
        """
        The modulus value for the RSA public key
        """
        return pulumi.get(self, "n")

    @property
    @pulumi.getter
    def use(self) -> str:
        """
        The intended use of the public key
        """
        return pulumi.get(self, "use")

    @property
    @pulumi.getter
    def x(self) -> str:
        """
        The x coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> str:
        """
        The y coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "y")

