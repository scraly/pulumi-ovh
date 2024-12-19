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

__all__ = [
    'GetLogsInputEngineResult',
    'AwaitableGetLogsInputEngineResult',
    'get_logs_input_engine',
    'get_logs_input_engine_output',
]

@pulumi.output_type
class GetLogsInputEngineResult:
    """
    A collection of values returned by getLogsInputEngine.
    """
    def __init__(__self__, id=None, is_deprecated=None, name=None, service_name=None, version=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_deprecated and not isinstance(is_deprecated, bool):
            raise TypeError("Expected argument 'is_deprecated' to be a bool")
        pulumi.set(__self__, "is_deprecated", is_deprecated)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDeprecated")
    def is_deprecated(self) -> bool:
        return pulumi.get(self, "is_deprecated")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


class AwaitableGetLogsInputEngineResult(GetLogsInputEngineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogsInputEngineResult(
            id=self.id,
            is_deprecated=self.is_deprecated,
            name=self.name,
            service_name=self.service_name,
            version=self.version)


def get_logs_input_engine(is_deprecated: Optional[bool] = None,
                          name: Optional[str] = None,
                          service_name: Optional[str] = None,
                          version: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLogsInputEngineResult:
    """
    Use this data source to retrieve information about a DBaas logs input engine.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    logstash = ovh.Dbaas.get_logs_input_engine(is_deprecated=True,
        name="logstash",
        service_name="ldp-xx-xxxxx",
        version="6.8")
    ```


    :param bool is_deprecated: Indicates if engine will soon not be supported.
    :param str name: The name of the logs input engine.
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    :param str version: Software version
    """
    __args__ = dict()
    __args__['isDeprecated'] = is_deprecated
    __args__['name'] = name
    __args__['serviceName'] = service_name
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dbaas/getLogsInputEngine:getLogsInputEngine', __args__, opts=opts, typ=GetLogsInputEngineResult).value

    return AwaitableGetLogsInputEngineResult(
        id=pulumi.get(__ret__, 'id'),
        is_deprecated=pulumi.get(__ret__, 'is_deprecated'),
        name=pulumi.get(__ret__, 'name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        version=pulumi.get(__ret__, 'version'))
def get_logs_input_engine_output(is_deprecated: Optional[pulumi.Input[Optional[bool]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 service_name: Optional[pulumi.Input[str]] = None,
                                 version: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLogsInputEngineResult]:
    """
    Use this data source to retrieve information about a DBaas logs input engine.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    logstash = ovh.Dbaas.get_logs_input_engine(is_deprecated=True,
        name="logstash",
        service_name="ldp-xx-xxxxx",
        version="6.8")
    ```


    :param bool is_deprecated: Indicates if engine will soon not be supported.
    :param str name: The name of the logs input engine.
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    :param str version: Software version
    """
    __args__ = dict()
    __args__['isDeprecated'] = is_deprecated
    __args__['name'] = name
    __args__['serviceName'] = service_name
    __args__['version'] = version
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Dbaas/getLogsInputEngine:getLogsInputEngine', __args__, opts=opts, typ=GetLogsInputEngineResult)
    return __ret__.apply(lambda __response__: GetLogsInputEngineResult(
        id=pulumi.get(__response__, 'id'),
        is_deprecated=pulumi.get(__response__, 'is_deprecated'),
        name=pulumi.get(__response__, 'name'),
        service_name=pulumi.get(__response__, 'service_name'),
        version=pulumi.get(__response__, 'version')))
