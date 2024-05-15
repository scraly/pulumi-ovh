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

__all__ = [
    'GetPrivateDatabaseDbResult',
    'AwaitableGetPrivateDatabaseDbResult',
    'get_private_database_db',
    'get_private_database_db_output',
]

@pulumi.output_type
class GetPrivateDatabaseDbResult:
    """
    A collection of values returned by getPrivateDatabaseDb.
    """
    def __init__(__self__, backup_time=None, creation_date=None, database_name=None, id=None, quota_used=None, service_name=None, users=None):
        if backup_time and not isinstance(backup_time, str):
            raise TypeError("Expected argument 'backup_time' to be a str")
        pulumi.set(__self__, "backup_time", backup_time)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if quota_used and not isinstance(quota_used, int):
            raise TypeError("Expected argument 'quota_used' to be a int")
        pulumi.set(__self__, "quota_used", quota_used)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> str:
        """
        Time of the next backup (every day)
        """
        return pulumi.get(self, "backup_time")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Creation date of the database
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="quotaUsed")
    def quota_used(self) -> int:
        """
        Space used by the database (in MB)
        """
        return pulumi.get(self, "quota_used")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetPrivateDatabaseDbUserResult']:
        """
        Users granted to this database
        """
        return pulumi.get(self, "users")


class AwaitableGetPrivateDatabaseDbResult(GetPrivateDatabaseDbResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateDatabaseDbResult(
            backup_time=self.backup_time,
            creation_date=self.creation_date,
            database_name=self.database_name,
            id=self.id,
            quota_used=self.quota_used,
            service_name=self.service_name,
            users=self.users)


def get_private_database_db(database_name: Optional[str] = None,
                            service_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateDatabaseDbResult:
    """
    Use this data source to retrieve information about an hosting privatedatabase.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    mydatabase = ovh.Hosting.get_private_database_db(database_name="XXXXXX",
        service_name="XXXXXX")
    ```


    :param str database_name: Database name
    :param str service_name: The internal name of your private database
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Hosting/getPrivateDatabaseDb:getPrivateDatabaseDb', __args__, opts=opts, typ=GetPrivateDatabaseDbResult).value

    return AwaitableGetPrivateDatabaseDbResult(
        backup_time=pulumi.get(__ret__, 'backup_time'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        database_name=pulumi.get(__ret__, 'database_name'),
        id=pulumi.get(__ret__, 'id'),
        quota_used=pulumi.get(__ret__, 'quota_used'),
        service_name=pulumi.get(__ret__, 'service_name'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_private_database_db)
def get_private_database_db_output(database_name: Optional[pulumi.Input[str]] = None,
                                   service_name: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateDatabaseDbResult]:
    """
    Use this data source to retrieve information about an hosting privatedatabase.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    mydatabase = ovh.Hosting.get_private_database_db(database_name="XXXXXX",
        service_name="XXXXXX")
    ```


    :param str database_name: Database name
    :param str service_name: The internal name of your private database
    """
    ...
