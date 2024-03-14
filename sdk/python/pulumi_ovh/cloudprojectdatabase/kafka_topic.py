# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KafkaTopicArgs', 'KafkaTopic']

@pulumi.input_type
class KafkaTopicArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 min_insync_replicas: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partitions: Optional[pulumi.Input[int]] = None,
                 replication: Optional[pulumi.Input[int]] = None,
                 retention_bytes: Optional[pulumi.Input[int]] = None,
                 retention_hours: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a KafkaTopic resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[int] min_insync_replicas: Minimum insync replica accepted for this topic. Should be superior to 0
        :param pulumi.Input[str] name: Name of the topic. No spaces allowed.
        :param pulumi.Input[int] partitions: Number of partitions for this topic. Should be superior to 0
        :param pulumi.Input[int] replication: Number of replication for this topic. Should be superior to 1
        :param pulumi.Input[int] retention_bytes: Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        :param pulumi.Input[int] retention_hours: Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "service_name", service_name)
        if min_insync_replicas is not None:
            pulumi.set(__self__, "min_insync_replicas", min_insync_replicas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partitions is not None:
            pulumi.set(__self__, "partitions", partitions)
        if replication is not None:
            pulumi.set(__self__, "replication", replication)
        if retention_bytes is not None:
            pulumi.set(__self__, "retention_bytes", retention_bytes)
        if retention_hours is not None:
            pulumi.set(__self__, "retention_hours", retention_hours)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

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
    @pulumi.getter(name="minInsyncReplicas")
    def min_insync_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum insync replica accepted for this topic. Should be superior to 0
        """
        return pulumi.get(self, "min_insync_replicas")

    @min_insync_replicas.setter
    def min_insync_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_insync_replicas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the topic. No spaces allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def partitions(self) -> Optional[pulumi.Input[int]]:
        """
        Number of partitions for this topic. Should be superior to 0
        """
        return pulumi.get(self, "partitions")

    @partitions.setter
    def partitions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partitions", value)

    @property
    @pulumi.getter
    def replication(self) -> Optional[pulumi.Input[int]]:
        """
        Number of replication for this topic. Should be superior to 1
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter(name="retentionBytes")
    def retention_bytes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_bytes")

    @retention_bytes.setter
    def retention_bytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_bytes", value)

    @property
    @pulumi.getter(name="retentionHours")
    def retention_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_hours")

    @retention_hours.setter
    def retention_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_hours", value)


@pulumi.input_type
class _KafkaTopicState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 min_insync_replicas: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partitions: Optional[pulumi.Input[int]] = None,
                 replication: Optional[pulumi.Input[int]] = None,
                 retention_bytes: Optional[pulumi.Input[int]] = None,
                 retention_hours: Optional[pulumi.Input[int]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KafkaTopic resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] min_insync_replicas: Minimum insync replica accepted for this topic. Should be superior to 0
        :param pulumi.Input[str] name: Name of the topic. No spaces allowed.
        :param pulumi.Input[int] partitions: Number of partitions for this topic. Should be superior to 0
        :param pulumi.Input[int] replication: Number of replication for this topic. Should be superior to 1
        :param pulumi.Input[int] retention_bytes: Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        :param pulumi.Input[int] retention_hours: Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if min_insync_replicas is not None:
            pulumi.set(__self__, "min_insync_replicas", min_insync_replicas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partitions is not None:
            pulumi.set(__self__, "partitions", partitions)
        if replication is not None:
            pulumi.set(__self__, "replication", replication)
        if retention_bytes is not None:
            pulumi.set(__self__, "retention_bytes", retention_bytes)
        if retention_hours is not None:
            pulumi.set(__self__, "retention_hours", retention_hours)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="minInsyncReplicas")
    def min_insync_replicas(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum insync replica accepted for this topic. Should be superior to 0
        """
        return pulumi.get(self, "min_insync_replicas")

    @min_insync_replicas.setter
    def min_insync_replicas(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_insync_replicas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the topic. No spaces allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def partitions(self) -> Optional[pulumi.Input[int]]:
        """
        Number of partitions for this topic. Should be superior to 0
        """
        return pulumi.get(self, "partitions")

    @partitions.setter
    def partitions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partitions", value)

    @property
    @pulumi.getter
    def replication(self) -> Optional[pulumi.Input[int]]:
        """
        Number of replication for this topic. Should be superior to 1
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter(name="retentionBytes")
    def retention_bytes(self) -> Optional[pulumi.Input[int]]:
        """
        Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_bytes")

    @retention_bytes.setter
    def retention_bytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_bytes", value)

    @property
    @pulumi.getter(name="retentionHours")
    def retention_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_hours")

    @retention_hours.setter
    def retention_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_hours", value)

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


class KafkaTopic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 min_insync_replicas: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partitions: Optional[pulumi.Input[int]] = None,
                 replication: Optional[pulumi.Input[int]] = None,
                 retention_bytes: Optional[pulumi.Input[int]] = None,
                 retention_hours: Optional[pulumi.Input[int]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a topic for a kafka cluster associated with a public cloud project.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        topic = ovh.cloud_project_database.KafkaTopic("topic",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            min_insync_replicas=1,
            partitions=3,
            replication=2,
            retention_bytes=4,
            retention_hours=5)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OVHcloud Managed kafka clusters topics can be imported using the `service_name`, `cluster_id` and `id` of the topic, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic my_topic service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] min_insync_replicas: Minimum insync replica accepted for this topic. Should be superior to 0
        :param pulumi.Input[str] name: Name of the topic. No spaces allowed.
        :param pulumi.Input[int] partitions: Number of partitions for this topic. Should be superior to 0
        :param pulumi.Input[int] replication: Number of replication for this topic. Should be superior to 1
        :param pulumi.Input[int] retention_bytes: Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        :param pulumi.Input[int] retention_hours: Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KafkaTopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a topic for a kafka cluster associated with a public cloud project.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        topic = ovh.cloud_project_database.KafkaTopic("topic",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            min_insync_replicas=1,
            partitions=3,
            replication=2,
            retention_bytes=4,
            retention_hours=5)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OVHcloud Managed kafka clusters topics can be imported using the `service_name`, `cluster_id` and `id` of the topic, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic my_topic service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param KafkaTopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KafkaTopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 min_insync_replicas: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partitions: Optional[pulumi.Input[int]] = None,
                 replication: Optional[pulumi.Input[int]] = None,
                 retention_bytes: Optional[pulumi.Input[int]] = None,
                 retention_hours: Optional[pulumi.Input[int]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KafkaTopicArgs.__new__(KafkaTopicArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["min_insync_replicas"] = min_insync_replicas
            __props__.__dict__["name"] = name
            __props__.__dict__["partitions"] = partitions
            __props__.__dict__["replication"] = replication
            __props__.__dict__["retention_bytes"] = retention_bytes
            __props__.__dict__["retention_hours"] = retention_hours
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(KafkaTopic, __self__).__init__(
            'ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            min_insync_replicas: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partitions: Optional[pulumi.Input[int]] = None,
            replication: Optional[pulumi.Input[int]] = None,
            retention_bytes: Optional[pulumi.Input[int]] = None,
            retention_hours: Optional[pulumi.Input[int]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'KafkaTopic':
        """
        Get an existing KafkaTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] min_insync_replicas: Minimum insync replica accepted for this topic. Should be superior to 0
        :param pulumi.Input[str] name: Name of the topic. No spaces allowed.
        :param pulumi.Input[int] partitions: Number of partitions for this topic. Should be superior to 0
        :param pulumi.Input[int] replication: Number of replication for this topic. Should be superior to 1
        :param pulumi.Input[int] retention_bytes: Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        :param pulumi.Input[int] retention_hours: Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KafkaTopicState.__new__(_KafkaTopicState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["min_insync_replicas"] = min_insync_replicas
        __props__.__dict__["name"] = name
        __props__.__dict__["partitions"] = partitions
        __props__.__dict__["replication"] = replication
        __props__.__dict__["retention_bytes"] = retention_bytes
        __props__.__dict__["retention_hours"] = retention_hours
        __props__.__dict__["service_name"] = service_name
        return KafkaTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="minInsyncReplicas")
    def min_insync_replicas(self) -> pulumi.Output[int]:
        """
        Minimum insync replica accepted for this topic. Should be superior to 0
        """
        return pulumi.get(self, "min_insync_replicas")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the topic. No spaces allowed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partitions(self) -> pulumi.Output[int]:
        """
        Number of partitions for this topic. Should be superior to 0
        """
        return pulumi.get(self, "partitions")

    @property
    @pulumi.getter
    def replication(self) -> pulumi.Output[int]:
        """
        Number of replication for this topic. Should be superior to 1
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter(name="retentionBytes")
    def retention_bytes(self) -> pulumi.Output[int]:
        """
        Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_bytes")

    @property
    @pulumi.getter(name="retentionHours")
    def retention_hours(self) -> pulumi.Output[int]:
        """
        Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
        """
        return pulumi.get(self, "retention_hours")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

