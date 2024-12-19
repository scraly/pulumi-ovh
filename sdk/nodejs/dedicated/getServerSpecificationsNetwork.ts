// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the network information about a dedicated server associated with your OVHcloud Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const spec = ovh.Dedicated.getServerSpecificationsNetwork({
 *     serviceName: "myserver",
 * });
 * ```
 */
export function getServerSpecificationsNetwork(args: GetServerSpecificationsNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetServerSpecificationsNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dedicated/getServerSpecificationsNetwork:getServerSpecificationsNetwork", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerSpecificationsNetwork.
 */
export interface GetServerSpecificationsNetworkArgs {
    /**
     * The internal name of your dedicated server.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getServerSpecificationsNetwork.
 */
export interface GetServerSpecificationsNetworkResult {
    /**
     * vrack bandwidth limitation
     */
    readonly bandwidth: outputs.Dedicated.GetServerSpecificationsNetworkBandwidth;
    /**
     * Network connection flow rate
     */
    readonly connectionVal: outputs.Dedicated.GetServerSpecificationsNetworkConnectionVal;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * OLA details
     */
    readonly ola: outputs.Dedicated.GetServerSpecificationsNetworkOla;
    /**
     * Routing details
     */
    readonly routing: outputs.Dedicated.GetServerSpecificationsNetworkRouting;
    readonly serviceName: string;
    /**
     * Switching details
     */
    readonly switching: outputs.Dedicated.GetServerSpecificationsNetworkSwitching;
    /**
     * Traffic details
     */
    readonly traffic: outputs.Dedicated.GetServerSpecificationsNetworkTraffic;
    /**
     * VMAC information for this dedicated server
     */
    readonly vmac: outputs.Dedicated.GetServerSpecificationsNetworkVmac;
    /**
     * vRack details
     */
    readonly vrack: outputs.Dedicated.GetServerSpecificationsNetworkVrack;
}
/**
 * Use this data source to get the network information about a dedicated server associated with your OVHcloud Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const spec = ovh.Dedicated.getServerSpecificationsNetwork({
 *     serviceName: "myserver",
 * });
 * ```
 */
export function getServerSpecificationsNetworkOutput(args: GetServerSpecificationsNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServerSpecificationsNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Dedicated/getServerSpecificationsNetwork:getServerSpecificationsNetwork", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerSpecificationsNetwork.
 */
export interface GetServerSpecificationsNetworkOutputArgs {
    /**
     * The internal name of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
}
