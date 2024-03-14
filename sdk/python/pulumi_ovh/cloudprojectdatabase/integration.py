# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IntegrationArgs', 'Integration']

@pulumi.input_type
class IntegrationArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 destination_service_id: pulumi.Input[str],
                 engine: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 source_service_id: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Integration resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] destination_service_id: ID of the destination service.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               All engines available exept `mongodb`.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] source_service_id: ID of the source service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Parameters for the integration.
        :param pulumi.Input[str] type: Type of the integration.
               Available types:
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "destination_service_id", destination_service_id)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "source_service_id", source_service_id)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="destinationServiceId")
    def destination_service_id(self) -> pulumi.Input[str]:
        """
        ID of the destination service.
        """
        return pulumi.get(self, "destination_service_id")

    @destination_service_id.setter
    def destination_service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_service_id", value)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        All engines available exept `mongodb`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

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
    @pulumi.getter(name="sourceServiceId")
    def source_service_id(self) -> pulumi.Input[str]:
        """
        ID of the source service.
        """
        return pulumi.get(self, "source_service_id")

    @source_service_id.setter
    def source_service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_service_id", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Parameters for the integration.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the integration.
        Available types:
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _IntegrationState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 destination_service_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_service_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Integration resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] destination_service_id: ID of the destination service.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               All engines available exept `mongodb`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Parameters for the integration.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] source_service_id: ID of the source service.
        :param pulumi.Input[str] status: Current status of the integration.
        :param pulumi.Input[str] type: Type of the integration.
               Available types:
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if destination_service_id is not None:
            pulumi.set(__self__, "destination_service_id", destination_service_id)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if source_service_id is not None:
            pulumi.set(__self__, "source_service_id", source_service_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="destinationServiceId")
    def destination_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the destination service.
        """
        return pulumi.get(self, "destination_service_id")

    @destination_service_id.setter
    def destination_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_service_id", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        All engines available exept `mongodb`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Parameters for the integration.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

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
    @pulumi.getter(name="sourceServiceId")
    def source_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the source service.
        """
        return pulumi.get(self, "source_service_id")

    @source_service_id.setter
    def source_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_service_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current status of the integration.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the integration.
        Available types:
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Integration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 destination_service_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_service_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an integration for a database cluster associated with a public cloud project.

        With this resource you can create an integration for all engine exept `mongodb`.

        Please take a look at the list of available `types` in the `Argument references` section in order to know the list of avaulable integrations. For example, thanks to the integration feature you can have your PostgreSQL logs in your OpenSearch Database.

        ## Example Usage

        Push PostgreSQL logs in an OpenSearch DB:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        dbpostgresql = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="postgresql",
            id="ZZZZ")
        dbopensearch = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="opensearch",
            id="ZZZZ")
        integration = ovh.cloud_project_database.Integration("integration",
            service_name=dbpostgresql.service_name,
            engine=dbpostgresql.engine,
            cluster_id=dbpostgresql.id,
            source_service_id=dbpostgresql.id,
            destination_service_id=dbopensearch.id,
            type="opensearchLogs")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/integration:Integration my_user service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] destination_service_id: ID of the destination service.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               All engines available exept `mongodb`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Parameters for the integration.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] source_service_id: ID of the source service.
        :param pulumi.Input[str] type: Type of the integration.
               Available types:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an integration for a database cluster associated with a public cloud project.

        With this resource you can create an integration for all engine exept `mongodb`.

        Please take a look at the list of available `types` in the `Argument references` section in order to know the list of avaulable integrations. For example, thanks to the integration feature you can have your PostgreSQL logs in your OpenSearch Database.

        ## Example Usage

        Push PostgreSQL logs in an OpenSearch DB:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        dbpostgresql = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="postgresql",
            id="ZZZZ")
        dbopensearch = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="opensearch",
            id="ZZZZ")
        integration = ovh.cloud_project_database.Integration("integration",
            service_name=dbpostgresql.service_name,
            engine=dbpostgresql.engine,
            cluster_id=dbpostgresql.id,
            source_service_id=dbpostgresql.id,
            destination_service_id=dbopensearch.id,
            type="opensearchLogs")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/integration:Integration my_user service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 destination_service_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_service_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationArgs.__new__(IntegrationArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if destination_service_id is None and not opts.urn:
                raise TypeError("Missing required property 'destination_service_id'")
            __props__.__dict__["destination_service_id"] = destination_service_id
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["parameters"] = parameters
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if source_service_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_service_id'")
            __props__.__dict__["source_service_id"] = source_service_id
            __props__.__dict__["type"] = type
            __props__.__dict__["status"] = None
        super(Integration, __self__).__init__(
            'ovh:CloudProjectDatabase/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            destination_service_id: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            source_service_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] destination_service_id: ID of the destination service.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               All engines available exept `mongodb`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Parameters for the integration.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] source_service_id: ID of the source service.
        :param pulumi.Input[str] status: Current status of the integration.
        :param pulumi.Input[str] type: Type of the integration.
               Available types:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationState.__new__(_IntegrationState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["destination_service_id"] = destination_service_id
        __props__.__dict__["engine"] = engine
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["source_service_id"] = source_service_id
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="destinationServiceId")
    def destination_service_id(self) -> pulumi.Output[str]:
        """
        ID of the destination service.
        """
        return pulumi.get(self, "destination_service_id")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        All engines available exept `mongodb`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Parameters for the integration.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="sourceServiceId")
    def source_service_id(self) -> pulumi.Output[str]:
        """
        ID of the source service.
        """
        return pulumi.get(self, "source_service_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current status of the integration.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the integration.
        Available types:
        """
        return pulumi.get(self, "type")

