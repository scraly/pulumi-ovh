// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a DBaas logs input engine.
 *
 * ## Example Usage
 */
export function getLogsInputEngine(args: GetLogsInputEngineArgs, opts?: pulumi.InvokeOptions): Promise<GetLogsInputEngineResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dbaas/getLogsInputEngine:getLogsInputEngine", {
        "isDeprecated": args.isDeprecated,
        "name": args.name,
        "serviceName": args.serviceName,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getLogsInputEngine.
 */
export interface GetLogsInputEngineArgs {
    /**
     * Indicates if engine will soon not be supported.
     */
    isDeprecated?: boolean;
    /**
     * The name of the logs input engine.
     */
    name?: string;
    /**
     * The service name. It's the ID of your Logs Data Platform instance.
     */
    serviceName: string;
    /**
     * Software version
     */
    version?: string;
}

/**
 * A collection of values returned by getLogsInputEngine.
 */
export interface GetLogsInputEngineResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isDeprecated: boolean;
    readonly name: string;
    readonly serviceName: string;
    readonly version: string;
}
/**
 * Use this data source to retrieve information about a DBaas logs input engine.
 *
 * ## Example Usage
 */
export function getLogsInputEngineOutput(args: GetLogsInputEngineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogsInputEngineResult> {
    return pulumi.output(args).apply((a: any) => getLogsInputEngine(a, opts))
}

/**
 * A collection of arguments for invoking getLogsInputEngine.
 */
export interface GetLogsInputEngineOutputArgs {
    /**
     * Indicates if engine will soon not be supported.
     */
    isDeprecated?: pulumi.Input<boolean>;
    /**
     * The name of the logs input engine.
     */
    name?: pulumi.Input<string>;
    /**
     * The service name. It's the ID of your Logs Data Platform instance.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Software version
     */
    version?: pulumi.Input<string>;
}