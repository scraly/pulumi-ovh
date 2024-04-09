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
    'GetDatabaseLogSubscriptionsResult',
    'AwaitableGetDatabaseLogSubscriptionsResult',
    'get_database_log_subscriptions',
    'get_database_log_subscriptions_output',
]

@pulumi.output_type
class GetDatabaseLogSubscriptionsResult:
    """
    A collection of values returned by getDatabaseLogSubscriptions.
    """
    def __init__(__self__, cluster_id=None, engine=None, id=None, service_name=None, subscription_ids=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if subscription_ids and not isinstance(subscription_ids, list):
            raise TypeError("Expected argument 'subscription_ids' to be a list")
        pulumi.set(__self__, "subscription_ids", subscription_ids)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="subscriptionIds")
    def subscription_ids(self) -> Sequence[str]:
        """
        The list of log subscription ids of the cluster associated with the project.
        """
        return pulumi.get(self, "subscription_ids")


class AwaitableGetDatabaseLogSubscriptionsResult(GetDatabaseLogSubscriptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseLogSubscriptionsResult(
            cluster_id=self.cluster_id,
            engine=self.engine,
            id=self.id,
            service_name=self.service_name,
            subscription_ids=self.subscription_ids)


def get_database_log_subscriptions(cluster_id: Optional[str] = None,
                                   engine: Optional[str] = None,
                                   service_name: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseLogSubscriptionsResult:
    """
    Use this data source to get the list of log subscrition for a cluster associated with a public cloud project.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    subscriptions = ovh.CloudProjectDatabase.get_database_log_subscriptions(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ")
    pulumi.export("subscriptionIds", subscriptions.subscription_ids)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID.
    :param str engine: The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabaseLogSubscriptions:getDatabaseLogSubscriptions', __args__, opts=opts, typ=GetDatabaseLogSubscriptionsResult).value

    return AwaitableGetDatabaseLogSubscriptionsResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        subscription_ids=pulumi.get(__ret__, 'subscription_ids'))


@_utilities.lift_output_func(get_database_log_subscriptions)
def get_database_log_subscriptions_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                          engine: Optional[pulumi.Input[str]] = None,
                                          service_name: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseLogSubscriptionsResult]:
    """
    Use this data source to get the list of log subscrition for a cluster associated with a public cloud project.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    subscriptions = ovh.CloudProjectDatabase.get_database_log_subscriptions(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ")
    pulumi.export("subscriptionIds", subscriptions.subscription_ids)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID.
    :param str engine: The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
