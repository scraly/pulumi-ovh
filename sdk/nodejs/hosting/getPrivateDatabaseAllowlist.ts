// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about an hosting privatedatabase whitelist.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const whitelist = ovh.Hosting.getPrivateDatabaseAllowlist({
 *     ip: "XXXXXX",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getPrivateDatabaseAllowlist(args: GetPrivateDatabaseAllowlistArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateDatabaseAllowlistResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", {
        "ip": args.ip,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateDatabaseAllowlist.
 */
export interface GetPrivateDatabaseAllowlistArgs {
    /**
     * The whitelisted IP in your instance
     */
    ip?: string;
    /**
     * The internal name of your private database
     */
    serviceName: string;
}

/**
 * A collection of values returned by getPrivateDatabaseAllowlist.
 */
export interface GetPrivateDatabaseAllowlistResult {
    /**
     * Creation date of the database
     */
    readonly creationDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ip?: string;
    /**
     * The last update date of this whitelist
     */
    readonly lastUpdate: string;
    /**
     * Custom name for your Whitelisted IP
     */
    readonly name: string;
    /**
     * Authorize this IP to access service port
     */
    readonly service: boolean;
    readonly serviceName: string;
    /**
     * Authorize this IP to access SFTP port
     */
    readonly sftp: boolean;
    /**
     * Whitelist status
     */
    readonly status: string;
}
/**
 * Use this data source to retrieve information about an hosting privatedatabase whitelist.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const whitelist = ovh.Hosting.getPrivateDatabaseAllowlist({
 *     ip: "XXXXXX",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getPrivateDatabaseAllowlistOutput(args: GetPrivateDatabaseAllowlistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateDatabaseAllowlistResult> {
    return pulumi.output(args).apply((a: any) => getPrivateDatabaseAllowlist(a, opts))
}

/**
 * A collection of arguments for invoking getPrivateDatabaseAllowlist.
 */
export interface GetPrivateDatabaseAllowlistOutputArgs {
    /**
     * The whitelisted IP in your instance
     */
    ip?: pulumi.Input<string>;
    /**
     * The internal name of your private database
     */
    serviceName: pulumi.Input<string>;
}
