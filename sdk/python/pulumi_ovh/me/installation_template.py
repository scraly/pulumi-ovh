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
from ._inputs import *

__all__ = ['InstallationTemplateArgs', 'InstallationTemplate']

@pulumi.input_type
class InstallationTemplateArgs:
    def __init__(__self__, *,
                 base_template_name: pulumi.Input[str],
                 template_name: pulumi.Input[str],
                 customization: Optional[pulumi.Input['InstallationTemplateCustomizationArgs']] = None,
                 remove_default_partition_schemes: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a InstallationTemplate resource.
        :param pulumi.Input[str] base_template_name: The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        :param pulumi.Input[str] template_name: This template name.
        :param pulumi.Input[bool] remove_default_partition_schemes: Remove default partition schemes at creation.
        """
        pulumi.set(__self__, "base_template_name", base_template_name)
        pulumi.set(__self__, "template_name", template_name)
        if customization is not None:
            pulumi.set(__self__, "customization", customization)
        if remove_default_partition_schemes is not None:
            pulumi.set(__self__, "remove_default_partition_schemes", remove_default_partition_schemes)

    @property
    @pulumi.getter(name="baseTemplateName")
    def base_template_name(self) -> pulumi.Input[str]:
        """
        The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        """
        return pulumi.get(self, "base_template_name")

    @base_template_name.setter
    def base_template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "base_template_name", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        """
        This template name.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter
    def customization(self) -> Optional[pulumi.Input['InstallationTemplateCustomizationArgs']]:
        return pulumi.get(self, "customization")

    @customization.setter
    def customization(self, value: Optional[pulumi.Input['InstallationTemplateCustomizationArgs']]):
        pulumi.set(self, "customization", value)

    @property
    @pulumi.getter(name="removeDefaultPartitionSchemes")
    def remove_default_partition_schemes(self) -> Optional[pulumi.Input[bool]]:
        """
        Remove default partition schemes at creation.
        """
        return pulumi.get(self, "remove_default_partition_schemes")

    @remove_default_partition_schemes.setter
    def remove_default_partition_schemes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_default_partition_schemes", value)


@pulumi.input_type
class _InstallationTemplateState:
    def __init__(__self__, *,
                 base_template_name: Optional[pulumi.Input[str]] = None,
                 bit_format: Optional[pulumi.Input[int]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 customization: Optional[pulumi.Input['InstallationTemplateCustomizationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 distribution: Optional[pulumi.Input[str]] = None,
                 end_of_install: Optional[pulumi.Input[str]] = None,
                 family: Optional[pulumi.Input[str]] = None,
                 filesystems: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hard_raid_configuration: Optional[pulumi.Input[bool]] = None,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input['InstallationTemplateInputArgs']]]] = None,
                 lvm_ready: Optional[pulumi.Input[bool]] = None,
                 no_partitioning: Optional[pulumi.Input[bool]] = None,
                 remove_default_partition_schemes: Optional[pulumi.Input[bool]] = None,
                 soft_raid_only_mirroring: Optional[pulumi.Input[bool]] = None,
                 subfamily: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstallationTemplate resources.
        :param pulumi.Input[str] base_template_name: The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        :param pulumi.Input[int] bit_format: This template bit format (32 or 64).
        :param pulumi.Input[str] category: Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
        :param pulumi.Input[str] description: information about this template.
        :param pulumi.Input[str] distribution: the distribution this template is based on.
        :param pulumi.Input[str] end_of_install: after this date, install of this template will not be possible at OVH
        :param pulumi.Input[str] family: this template family type.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] filesystems: Filesystems available.
        :param pulumi.Input[bool] hard_raid_configuration: This distribution supports hardware raid configuration through the OVHcloud API. Deprecated, will be removed in next release.
        :param pulumi.Input[bool] lvm_ready: Whether this distribution supports Logical Volumes (Linux LVM)
        :param pulumi.Input[bool] no_partitioning: Partitioning customization is not available for this OS template
        :param pulumi.Input[bool] remove_default_partition_schemes: Remove default partition schemes at creation.
        :param pulumi.Input[bool] soft_raid_only_mirroring: Partitioning customization is available but limited to mirroring for this OS template
        :param pulumi.Input[str] subfamily: this template subfamily type
        :param pulumi.Input[str] template_name: This template name.
        """
        if base_template_name is not None:
            pulumi.set(__self__, "base_template_name", base_template_name)
        if bit_format is not None:
            pulumi.set(__self__, "bit_format", bit_format)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if customization is not None:
            pulumi.set(__self__, "customization", customization)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if distribution is not None:
            pulumi.set(__self__, "distribution", distribution)
        if end_of_install is not None:
            pulumi.set(__self__, "end_of_install", end_of_install)
        if family is not None:
            pulumi.set(__self__, "family", family)
        if filesystems is not None:
            pulumi.set(__self__, "filesystems", filesystems)
        if hard_raid_configuration is not None:
            warnings.warn("""This will be deprecated in the next release""", DeprecationWarning)
            pulumi.log.warn("""hard_raid_configuration is deprecated: This will be deprecated in the next release""")
        if hard_raid_configuration is not None:
            pulumi.set(__self__, "hard_raid_configuration", hard_raid_configuration)
        if inputs is not None:
            pulumi.set(__self__, "inputs", inputs)
        if lvm_ready is not None:
            pulumi.set(__self__, "lvm_ready", lvm_ready)
        if no_partitioning is not None:
            pulumi.set(__self__, "no_partitioning", no_partitioning)
        if remove_default_partition_schemes is not None:
            pulumi.set(__self__, "remove_default_partition_schemes", remove_default_partition_schemes)
        if soft_raid_only_mirroring is not None:
            pulumi.set(__self__, "soft_raid_only_mirroring", soft_raid_only_mirroring)
        if subfamily is not None:
            pulumi.set(__self__, "subfamily", subfamily)
        if template_name is not None:
            pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter(name="baseTemplateName")
    def base_template_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        """
        return pulumi.get(self, "base_template_name")

    @base_template_name.setter
    def base_template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_template_name", value)

    @property
    @pulumi.getter(name="bitFormat")
    def bit_format(self) -> Optional[pulumi.Input[int]]:
        """
        This template bit format (32 or 64).
        """
        return pulumi.get(self, "bit_format")

    @bit_format.setter
    def bit_format(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bit_format", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def customization(self) -> Optional[pulumi.Input['InstallationTemplateCustomizationArgs']]:
        return pulumi.get(self, "customization")

    @customization.setter
    def customization(self, value: Optional[pulumi.Input['InstallationTemplateCustomizationArgs']]):
        pulumi.set(self, "customization", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        information about this template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def distribution(self) -> Optional[pulumi.Input[str]]:
        """
        the distribution this template is based on.
        """
        return pulumi.get(self, "distribution")

    @distribution.setter
    def distribution(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distribution", value)

    @property
    @pulumi.getter(name="endOfInstall")
    def end_of_install(self) -> Optional[pulumi.Input[str]]:
        """
        after this date, install of this template will not be possible at OVH
        """
        return pulumi.get(self, "end_of_install")

    @end_of_install.setter
    def end_of_install(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_of_install", value)

    @property
    @pulumi.getter
    def family(self) -> Optional[pulumi.Input[str]]:
        """
        this template family type.
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def filesystems(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Filesystems available.
        """
        return pulumi.get(self, "filesystems")

    @filesystems.setter
    def filesystems(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "filesystems", value)

    @property
    @pulumi.getter(name="hardRaidConfiguration")
    def hard_raid_configuration(self) -> Optional[pulumi.Input[bool]]:
        """
        This distribution supports hardware raid configuration through the OVHcloud API. Deprecated, will be removed in next release.
        """
        warnings.warn("""This will be deprecated in the next release""", DeprecationWarning)
        pulumi.log.warn("""hard_raid_configuration is deprecated: This will be deprecated in the next release""")

        return pulumi.get(self, "hard_raid_configuration")

    @hard_raid_configuration.setter
    def hard_raid_configuration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hard_raid_configuration", value)

    @property
    @pulumi.getter
    def inputs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstallationTemplateInputArgs']]]]:
        return pulumi.get(self, "inputs")

    @inputs.setter
    def inputs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstallationTemplateInputArgs']]]]):
        pulumi.set(self, "inputs", value)

    @property
    @pulumi.getter(name="lvmReady")
    def lvm_ready(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this distribution supports Logical Volumes (Linux LVM)
        """
        return pulumi.get(self, "lvm_ready")

    @lvm_ready.setter
    def lvm_ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lvm_ready", value)

    @property
    @pulumi.getter(name="noPartitioning")
    def no_partitioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Partitioning customization is not available for this OS template
        """
        return pulumi.get(self, "no_partitioning")

    @no_partitioning.setter
    def no_partitioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_partitioning", value)

    @property
    @pulumi.getter(name="removeDefaultPartitionSchemes")
    def remove_default_partition_schemes(self) -> Optional[pulumi.Input[bool]]:
        """
        Remove default partition schemes at creation.
        """
        return pulumi.get(self, "remove_default_partition_schemes")

    @remove_default_partition_schemes.setter
    def remove_default_partition_schemes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_default_partition_schemes", value)

    @property
    @pulumi.getter(name="softRaidOnlyMirroring")
    def soft_raid_only_mirroring(self) -> Optional[pulumi.Input[bool]]:
        """
        Partitioning customization is available but limited to mirroring for this OS template
        """
        return pulumi.get(self, "soft_raid_only_mirroring")

    @soft_raid_only_mirroring.setter
    def soft_raid_only_mirroring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "soft_raid_only_mirroring", value)

    @property
    @pulumi.getter
    def subfamily(self) -> Optional[pulumi.Input[str]]:
        """
        this template subfamily type
        """
        return pulumi.get(self, "subfamily")

    @subfamily.setter
    def subfamily(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subfamily", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[pulumi.Input[str]]:
        """
        This template name.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_name", value)


class InstallationTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_template_name: Optional[pulumi.Input[str]] = None,
                 customization: Optional[pulumi.Input[pulumi.InputType['InstallationTemplateCustomizationArgs']]] = None,
                 remove_default_partition_schemes: Optional[pulumi.Input[bool]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create a custom installation template available for dedicated servers.

        ## Import

        Custom installation template available for dedicated servers can be imported using the `base_template_name`, `template_name` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Me/installationTemplate:InstallationTemplate mytemplate base_template_name/template_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_template_name: The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        :param pulumi.Input[bool] remove_default_partition_schemes: Remove default partition schemes at creation.
        :param pulumi.Input[str] template_name: This template name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstallationTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create a custom installation template available for dedicated servers.

        ## Import

        Custom installation template available for dedicated servers can be imported using the `base_template_name`, `template_name` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Me/installationTemplate:InstallationTemplate mytemplate base_template_name/template_name
        ```

        :param str resource_name: The name of the resource.
        :param InstallationTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstallationTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_template_name: Optional[pulumi.Input[str]] = None,
                 customization: Optional[pulumi.Input[pulumi.InputType['InstallationTemplateCustomizationArgs']]] = None,
                 remove_default_partition_schemes: Optional[pulumi.Input[bool]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstallationTemplateArgs.__new__(InstallationTemplateArgs)

            if base_template_name is None and not opts.urn:
                raise TypeError("Missing required property 'base_template_name'")
            __props__.__dict__["base_template_name"] = base_template_name
            __props__.__dict__["customization"] = customization
            __props__.__dict__["remove_default_partition_schemes"] = remove_default_partition_schemes
            if template_name is None and not opts.urn:
                raise TypeError("Missing required property 'template_name'")
            __props__.__dict__["template_name"] = template_name
            __props__.__dict__["bit_format"] = None
            __props__.__dict__["category"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["distribution"] = None
            __props__.__dict__["end_of_install"] = None
            __props__.__dict__["family"] = None
            __props__.__dict__["filesystems"] = None
            __props__.__dict__["hard_raid_configuration"] = None
            __props__.__dict__["inputs"] = None
            __props__.__dict__["lvm_ready"] = None
            __props__.__dict__["no_partitioning"] = None
            __props__.__dict__["soft_raid_only_mirroring"] = None
            __props__.__dict__["subfamily"] = None
        super(InstallationTemplate, __self__).__init__(
            'ovh:Me/installationTemplate:InstallationTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            base_template_name: Optional[pulumi.Input[str]] = None,
            bit_format: Optional[pulumi.Input[int]] = None,
            category: Optional[pulumi.Input[str]] = None,
            customization: Optional[pulumi.Input[pulumi.InputType['InstallationTemplateCustomizationArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            distribution: Optional[pulumi.Input[str]] = None,
            end_of_install: Optional[pulumi.Input[str]] = None,
            family: Optional[pulumi.Input[str]] = None,
            filesystems: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            hard_raid_configuration: Optional[pulumi.Input[bool]] = None,
            inputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstallationTemplateInputArgs']]]]] = None,
            lvm_ready: Optional[pulumi.Input[bool]] = None,
            no_partitioning: Optional[pulumi.Input[bool]] = None,
            remove_default_partition_schemes: Optional[pulumi.Input[bool]] = None,
            soft_raid_only_mirroring: Optional[pulumi.Input[bool]] = None,
            subfamily: Optional[pulumi.Input[str]] = None,
            template_name: Optional[pulumi.Input[str]] = None) -> 'InstallationTemplate':
        """
        Get an existing InstallationTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_template_name: The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        :param pulumi.Input[int] bit_format: This template bit format (32 or 64).
        :param pulumi.Input[str] category: Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
        :param pulumi.Input[str] description: information about this template.
        :param pulumi.Input[str] distribution: the distribution this template is based on.
        :param pulumi.Input[str] end_of_install: after this date, install of this template will not be possible at OVH
        :param pulumi.Input[str] family: this template family type.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] filesystems: Filesystems available.
        :param pulumi.Input[bool] hard_raid_configuration: This distribution supports hardware raid configuration through the OVHcloud API. Deprecated, will be removed in next release.
        :param pulumi.Input[bool] lvm_ready: Whether this distribution supports Logical Volumes (Linux LVM)
        :param pulumi.Input[bool] no_partitioning: Partitioning customization is not available for this OS template
        :param pulumi.Input[bool] remove_default_partition_schemes: Remove default partition schemes at creation.
        :param pulumi.Input[bool] soft_raid_only_mirroring: Partitioning customization is available but limited to mirroring for this OS template
        :param pulumi.Input[str] subfamily: this template subfamily type
        :param pulumi.Input[str] template_name: This template name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstallationTemplateState.__new__(_InstallationTemplateState)

        __props__.__dict__["base_template_name"] = base_template_name
        __props__.__dict__["bit_format"] = bit_format
        __props__.__dict__["category"] = category
        __props__.__dict__["customization"] = customization
        __props__.__dict__["description"] = description
        __props__.__dict__["distribution"] = distribution
        __props__.__dict__["end_of_install"] = end_of_install
        __props__.__dict__["family"] = family
        __props__.__dict__["filesystems"] = filesystems
        __props__.__dict__["hard_raid_configuration"] = hard_raid_configuration
        __props__.__dict__["inputs"] = inputs
        __props__.__dict__["lvm_ready"] = lvm_ready
        __props__.__dict__["no_partitioning"] = no_partitioning
        __props__.__dict__["remove_default_partition_schemes"] = remove_default_partition_schemes
        __props__.__dict__["soft_raid_only_mirroring"] = soft_raid_only_mirroring
        __props__.__dict__["subfamily"] = subfamily
        __props__.__dict__["template_name"] = template_name
        return InstallationTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="baseTemplateName")
    def base_template_name(self) -> pulumi.Output[str]:
        """
        The name of an existing installation template, choose one among the list given by `get_installation_templates` datasource.
        """
        return pulumi.get(self, "base_template_name")

    @property
    @pulumi.getter(name="bitFormat")
    def bit_format(self) -> pulumi.Output[int]:
        """
        This template bit format (32 or 64).
        """
        return pulumi.get(self, "bit_format")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def customization(self) -> pulumi.Output[Optional['outputs.InstallationTemplateCustomization']]:
        return pulumi.get(self, "customization")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        information about this template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distribution(self) -> pulumi.Output[str]:
        """
        the distribution this template is based on.
        """
        return pulumi.get(self, "distribution")

    @property
    @pulumi.getter(name="endOfInstall")
    def end_of_install(self) -> pulumi.Output[str]:
        """
        after this date, install of this template will not be possible at OVH
        """
        return pulumi.get(self, "end_of_install")

    @property
    @pulumi.getter
    def family(self) -> pulumi.Output[str]:
        """
        this template family type.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def filesystems(self) -> pulumi.Output[Sequence[str]]:
        """
        Filesystems available.
        """
        return pulumi.get(self, "filesystems")

    @property
    @pulumi.getter(name="hardRaidConfiguration")
    def hard_raid_configuration(self) -> pulumi.Output[bool]:
        """
        This distribution supports hardware raid configuration through the OVHcloud API. Deprecated, will be removed in next release.
        """
        warnings.warn("""This will be deprecated in the next release""", DeprecationWarning)
        pulumi.log.warn("""hard_raid_configuration is deprecated: This will be deprecated in the next release""")

        return pulumi.get(self, "hard_raid_configuration")

    @property
    @pulumi.getter
    def inputs(self) -> pulumi.Output[Sequence['outputs.InstallationTemplateInput']]:
        return pulumi.get(self, "inputs")

    @property
    @pulumi.getter(name="lvmReady")
    def lvm_ready(self) -> pulumi.Output[bool]:
        """
        Whether this distribution supports Logical Volumes (Linux LVM)
        """
        return pulumi.get(self, "lvm_ready")

    @property
    @pulumi.getter(name="noPartitioning")
    def no_partitioning(self) -> pulumi.Output[bool]:
        """
        Partitioning customization is not available for this OS template
        """
        return pulumi.get(self, "no_partitioning")

    @property
    @pulumi.getter(name="removeDefaultPartitionSchemes")
    def remove_default_partition_schemes(self) -> pulumi.Output[bool]:
        """
        Remove default partition schemes at creation.
        """
        return pulumi.get(self, "remove_default_partition_schemes")

    @property
    @pulumi.getter(name="softRaidOnlyMirroring")
    def soft_raid_only_mirroring(self) -> pulumi.Output[bool]:
        """
        Partitioning customization is available but limited to mirroring for this OS template
        """
        return pulumi.get(self, "soft_raid_only_mirroring")

    @property
    @pulumi.getter
    def subfamily(self) -> pulumi.Output[str]:
        """
        this template subfamily type
        """
        return pulumi.get(self, "subfamily")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Output[str]:
        """
        This template name.
        """
        return pulumi.get(self, "template_name")

