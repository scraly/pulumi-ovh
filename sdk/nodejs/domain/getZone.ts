// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a domain zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const rootZone = ovh.Domain.getZone({
 *     name: "mysite.ovh",
 * });
 * ```
 */
export function getZone(args: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Domain/getZone:getZone", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    /**
     * The name of the domain zone.
     */
    name: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * URN of the DNS zone
     */
    readonly ZoneURN: string;
    /**
     * Is DNSSEC supported by this zone
     */
    readonly dnssecSupported: boolean;
    /**
     * hasDnsAnycast flag of the DNS zone
     */
    readonly hasDnsAnycast: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Last update date of the DNS zone
     */
    readonly lastUpdate: string;
    readonly name: string;
    /**
     * Name servers that host the DNS zone
     */
    readonly nameServers: string[];
}
/**
 * Use this data source to retrieve information about a domain zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const rootZone = ovh.Domain.getZone({
 *     name: "mysite.ovh",
 * });
 * ```
 */
export function getZoneOutput(args: GetZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Domain/getZone:getZone", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    /**
     * The name of the domain zone.
     */
    name: pulumi.Input<string>;
}
