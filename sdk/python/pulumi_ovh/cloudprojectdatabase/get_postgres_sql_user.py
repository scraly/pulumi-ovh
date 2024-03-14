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
    'GetPostgresSqlUserResult',
    'AwaitableGetPostgresSqlUserResult',
    'get_postgres_sql_user',
    'get_postgres_sql_user_output',
]

@pulumi.output_type
class GetPostgresSqlUserResult:
    """
    A collection of values returned by getPostgresSqlUser.
    """
    def __init__(__self__, cluster_id=None, created_at=None, id=None, name=None, roles=None, service_name=None, status=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of the creation of the user.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        """
        Roles the user belongs to.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Current status of the user.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status of the user.
        """
        return pulumi.get(self, "status")


class AwaitableGetPostgresSqlUserResult(GetPostgresSqlUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPostgresSqlUserResult(
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            id=self.id,
            name=self.name,
            roles=self.roles,
            service_name=self.service_name,
            status=self.status)


def get_postgres_sql_user(cluster_id: Optional[str] = None,
                          name: Optional[str] = None,
                          service_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPostgresSqlUserResult:
    """
    Use this data source to get information about a user of a postgresql cluster associated with a public cloud project.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    pguser = ovh.CloudProjectDatabase.get_postgres_sql_user(service_name="XXX",
        cluster_id="YYY",
        name="ZZZ")
    pulumi.export("pguserRoles", pguser.roles)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID
    :param str name: Name of the user.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getPostgresSqlUser:getPostgresSqlUser', __args__, opts=opts, typ=GetPostgresSqlUserResult).value

    return AwaitableGetPostgresSqlUserResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        roles=pulumi.get(__ret__, 'roles'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_postgres_sql_user)
def get_postgres_sql_user_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 service_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPostgresSqlUserResult]:
    """
    Use this data source to get information about a user of a postgresql cluster associated with a public cloud project.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    pguser = ovh.CloudProjectDatabase.get_postgres_sql_user(service_name="XXX",
        cluster_id="YYY",
        name="ZZZ")
    pulumi.export("pguserRoles", pguser.roles)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID
    :param str name: Name of the user.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
