// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a connection pool of a postgresql cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const testPoolPostgresSqlConnectionPool = ovh.CloudProjectDatabase.getPostgresSqlConnectionPool({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const testPool = {
 *     service_name: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.serviceName),
 *     cluster_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.clusterId),
 *     name: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.name),
 *     database_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.databaseId),
 *     mode: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.mode),
 *     size: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.size),
 *     port: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.port),
 *     ssl_mode: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.sslMode),
 *     uri: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.uri),
 *     user_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.userId),
 * };
 * ```
 */
export function getPostgresSqlConnectionPool(args: GetPostgresSqlConnectionPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPostgresSqlConnectionPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getPostgresSqlConnectionPool:getPostgresSqlConnectionPool", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPostgresSqlConnectionPool.
 */
export interface GetPostgresSqlConnectionPoolArgs {
    /**
     * Cluster ID.
     */
    clusterId: string;
    /**
     * Name of the Connection pool.
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getPostgresSqlConnectionPool.
 */
export interface GetPostgresSqlConnectionPoolResult {
    /**
     * See Argument Reference above
     */
    readonly clusterId: string;
    /**
     * Database ID for a database that belongs to the Database cluster given above.
     */
    readonly databaseId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Connection mode to the connection pool
     * Available modes:
     */
    readonly mode: string;
    /**
     * See Argument Reference above
     */
    readonly name: string;
    /**
     * Port of the connection pool.
     */
    readonly port: number;
    /**
     * See Argument Reference above
     */
    readonly serviceName: string;
    /**
     * Size of the connection pool.
     */
    readonly size: number;
    /**
     * Ssl connection mode for the pool.
     */
    readonly sslMode: string;
    /**
     * Connection URI to the pool.
     */
    readonly uri: string;
    /**
     * Database user authorized to connect to the pool, if none all the users are allowed.
     */
    readonly userId: string;
}
/**
 * Use this data source to get information about a connection pool of a postgresql cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const testPoolPostgresSqlConnectionPool = ovh.CloudProjectDatabase.getPostgresSqlConnectionPool({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const testPool = {
 *     service_name: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.serviceName),
 *     cluster_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.clusterId),
 *     name: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.name),
 *     database_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.databaseId),
 *     mode: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.mode),
 *     size: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.size),
 *     port: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.port),
 *     ssl_mode: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.sslMode),
 *     uri: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.uri),
 *     user_id: testPoolPostgresSqlConnectionPool.then(testPoolPostgresSqlConnectionPool => testPoolPostgresSqlConnectionPool.userId),
 * };
 * ```
 */
export function getPostgresSqlConnectionPoolOutput(args: GetPostgresSqlConnectionPoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPostgresSqlConnectionPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getPostgresSqlConnectionPool:getPostgresSqlConnectionPool", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPostgresSqlConnectionPool.
 */
export interface GetPostgresSqlConnectionPoolOutputArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the Connection pool.
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
