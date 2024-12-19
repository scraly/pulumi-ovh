// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a dedicated server associated with your OVHcloud Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const server = ovh.getServer({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:index/getServer:getServer", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerArgs {
    /**
     * The serviceName of your dedicated server.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getServer.
 */
export interface GetServerResult {
    /**
     * URN of the dedicated server instance
     */
    readonly ServerURN: string;
    /**
     * Dedicated AZ localisation
     */
    readonly availabilityZone: string;
    /**
     * Boot id of the server
     */
    readonly bootId: number;
    /**
     * Boot script of the server
     */
    readonly bootScript: string;
    /**
     * Dedicated server commercial range
     */
    readonly commercialRange: string;
    /**
     * Dedicated datacenter localisation (bhs1,bhs2,...)
     */
    readonly datacenter: string;
    /**
     * Dedicated server display name
     */
    readonly displayName: string;
    /**
     * List of enabled public VNI uuids
     */
    readonly enabledPublicVnis: string[];
    /**
     * List of enabled vrackAggregation VNI uuids
     */
    readonly enabledVrackAggregationVnis: string[];
    /**
     * List of enabled vrack VNI uuids
     */
    readonly enabledVrackVnis: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Dedicated server ip (IPv4)
     */
    readonly ip: string;
    /**
     * Dedicated server ip blocks
     */
    readonly ips: string[];
    /**
     * Link speed of the server
     */
    readonly linkSpeed: number;
    /**
     * Icmp monitoring state
     */
    readonly monitoring: boolean;
    /**
     * User defined VirtualNetworkInterface name
     */
    readonly name: string;
    readonly newUpgradeSystem: boolean;
    /**
     * Prevent datacenter intervention
     */
    readonly noIntervention: boolean;
    /**
     * Operating system
     */
    readonly os: string;
    /**
     * Power state of the server (poweroff, poweron)
     */
    readonly powerState: string;
    /**
     * Does this server have professional use option
     */
    readonly professionalUse: boolean;
    /**
     * Rack id of the server
     */
    readonly rack: string;
    /**
     * Dedicated region localisation
     */
    readonly region: string;
    /**
     * Rescue mail of the server
     */
    readonly rescueMail: string;
    /**
     * Public SSH Key used in the rescue mode
     */
    readonly rescueSshKey: string;
    /**
     * Dedicated server reverse
     */
    readonly reverse: string;
    /**
     * Root device of the server
     */
    readonly rootDevice: string;
    /**
     * Server id
     */
    readonly serverId: number;
    readonly serviceName: string;
    /**
     * Error, hacked, hackedBlocked, ok
     */
    readonly state: string;
    /**
     * Dedicated server support level (critical, fastpath, gs, pro)
     */
    readonly supportLevel: string;
    /**
     * The list of Virtualnetworkinterface associated with this server
     */
    readonly vnis: outputs.GetServerVni[];
}
/**
 * Use this data source to retrieve information about a dedicated server associated with your OVHcloud Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const server = ovh.getServer({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:index/getServer:getServer", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerOutputArgs {
    /**
     * The serviceName of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
}
