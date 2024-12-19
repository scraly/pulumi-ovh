// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * List your public cloud loadbalancers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lbsLoadBalancers = ovh.CloudProject.getLoadBalancers({
 *     serviceName: "XXXXXX",
 *     regionName: "XXX",
 * });
 * export const lbs = lbsLoadBalancers;
 * ```
 */
export function getLoadBalancers(args: GetLoadBalancersArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getLoadBalancers:getLoadBalancers", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancers.
 */
export interface GetLoadBalancersArgs {
    /**
     * Region of the loadbalancers.
     */
    regionName: string;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getLoadBalancers.
 */
export interface GetLoadBalancersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of loadbalancer
     */
    readonly loadbalancers: outputs.CloudProject.GetLoadBalancersLoadbalancer[];
    /**
     * Region of the loadbalancers
     */
    readonly regionName: string;
    /**
     * ID of the public cloud project
     */
    readonly serviceName: string;
}
/**
 * List your public cloud loadbalancers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lbsLoadBalancers = ovh.CloudProject.getLoadBalancers({
 *     serviceName: "XXXXXX",
 *     regionName: "XXX",
 * });
 * export const lbs = lbsLoadBalancers;
 * ```
 */
export function getLoadBalancersOutput(args: GetLoadBalancersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLoadBalancersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getLoadBalancers:getLoadBalancers", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancers.
 */
export interface GetLoadBalancersOutputArgs {
    /**
     * Region of the loadbalancers.
     */
    regionName: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
