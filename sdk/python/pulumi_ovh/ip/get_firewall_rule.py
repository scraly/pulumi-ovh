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
    'GetFirewallRuleResult',
    'AwaitableGetFirewallRuleResult',
    'get_firewall_rule',
    'get_firewall_rule_output',
]

@pulumi.output_type
class GetFirewallRuleResult:
    """
    A collection of values returned by getFirewallRule.
    """
    def __init__(__self__, action=None, creation_date=None, destination=None, destination_port=None, fragments=None, id=None, ip=None, ip_on_firewall=None, protocol=None, rule=None, sequence=None, source=None, source_port=None, state=None, tcp_option=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if destination_port and not isinstance(destination_port, str):
            raise TypeError("Expected argument 'destination_port' to be a str")
        pulumi.set(__self__, "destination_port", destination_port)
        if fragments and not isinstance(fragments, bool):
            raise TypeError("Expected argument 'fragments' to be a bool")
        pulumi.set(__self__, "fragments", fragments)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if ip_on_firewall and not isinstance(ip_on_firewall, str):
            raise TypeError("Expected argument 'ip_on_firewall' to be a str")
        pulumi.set(__self__, "ip_on_firewall", ip_on_firewall)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if rule and not isinstance(rule, str):
            raise TypeError("Expected argument 'rule' to be a str")
        pulumi.set(__self__, "rule", rule)
        if sequence and not isinstance(sequence, float):
            raise TypeError("Expected argument 'sequence' to be a float")
        pulumi.set(__self__, "sequence", sequence)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if source_port and not isinstance(source_port, str):
            raise TypeError("Expected argument 'source_port' to be a str")
        pulumi.set(__self__, "source_port", source_port)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tcp_option and not isinstance(tcp_option, str):
            raise TypeError("Expected argument 'tcp_option' to be a str")
        pulumi.set(__self__, "tcp_option", tcp_option)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Possible values for action (deny|permit)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Creation date of the rule
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def destination(self) -> str:
        """
        Destination IP for your rule
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> str:
        """
        Destination port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "destination_port")

    @property
    @pulumi.getter
    def fragments(self) -> bool:
        """
        Fragments option
        """
        return pulumi.get(self, "fragments")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipOnFirewall")
    def ip_on_firewall(self) -> str:
        """
        IPv4 address
        """
        return pulumi.get(self, "ip_on_firewall")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def rule(self) -> str:
        """
        Description of the rule
        """
        return pulumi.get(self, "rule")

    @property
    @pulumi.getter
    def sequence(self) -> float:
        """
        Rule position in the rules array
        """
        return pulumi.get(self, "sequence")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        IPv4 CIDR notation (e.g., 192.0.2.0/24)
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> str:
        """
        Source port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "source_port")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of your rule
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="tcpOption")
    def tcp_option(self) -> str:
        """
        TCP option on your rule (syn|established)
        """
        return pulumi.get(self, "tcp_option")


class AwaitableGetFirewallRuleResult(GetFirewallRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallRuleResult(
            action=self.action,
            creation_date=self.creation_date,
            destination=self.destination,
            destination_port=self.destination_port,
            fragments=self.fragments,
            id=self.id,
            ip=self.ip,
            ip_on_firewall=self.ip_on_firewall,
            protocol=self.protocol,
            rule=self.rule,
            sequence=self.sequence,
            source=self.source,
            source_port=self.source_port,
            state=self.state,
            tcp_option=self.tcp_option)


def get_firewall_rule(ip: Optional[str] = None,
                      ip_on_firewall: Optional[str] = None,
                      sequence: Optional[float] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallRuleResult:
    """
    Use this data source to retrieve information about a rule on an IP firewall.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    myfirewallrule = ovh.Ip.get_firewall_rule(ip="XXXXXX",
        ip_on_firewall="XXXXXX",
        sequence=0)
    ```


    :param str ip: The IP or the CIDR
    :param str ip_on_firewall: IPv4 address
    :param float sequence: Rule position in the rules array
    """
    __args__ = dict()
    __args__['ip'] = ip
    __args__['ipOnFirewall'] = ip_on_firewall
    __args__['sequence'] = sequence
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Ip/getFirewallRule:getFirewallRule', __args__, opts=opts, typ=GetFirewallRuleResult).value

    return AwaitableGetFirewallRuleResult(
        action=pulumi.get(__ret__, 'action'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        destination=pulumi.get(__ret__, 'destination'),
        destination_port=pulumi.get(__ret__, 'destination_port'),
        fragments=pulumi.get(__ret__, 'fragments'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        ip_on_firewall=pulumi.get(__ret__, 'ip_on_firewall'),
        protocol=pulumi.get(__ret__, 'protocol'),
        rule=pulumi.get(__ret__, 'rule'),
        sequence=pulumi.get(__ret__, 'sequence'),
        source=pulumi.get(__ret__, 'source'),
        source_port=pulumi.get(__ret__, 'source_port'),
        state=pulumi.get(__ret__, 'state'),
        tcp_option=pulumi.get(__ret__, 'tcp_option'))


@_utilities.lift_output_func(get_firewall_rule)
def get_firewall_rule_output(ip: Optional[pulumi.Input[str]] = None,
                             ip_on_firewall: Optional[pulumi.Input[str]] = None,
                             sequence: Optional[pulumi.Input[float]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallRuleResult]:
    """
    Use this data source to retrieve information about a rule on an IP firewall.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    myfirewallrule = ovh.Ip.get_firewall_rule(ip="XXXXXX",
        ip_on_firewall="XXXXXX",
        sequence=0)
    ```


    :param str ip: The IP or the CIDR
    :param str ip_on_firewall: IPv4 address
    :param float sequence: Rule position in the rules array
    """
    ...
