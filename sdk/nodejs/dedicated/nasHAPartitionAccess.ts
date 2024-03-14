// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing access rights to partitions on HA-NAS services
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const my_partition = new ovh.dedicated.NasHAPartitionAccess("my-partition", {
 *     ip: "123.123.123.123/32",
 *     partitionName: "my-partition",
 *     serviceName: "zpool-12345",
 *     type: "readwrite",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * HA-NAS partition access can be imported using the `{service_name}/{partition_name}/{ip}`, e.g.
 *
 * ```sh
 * $ pulumi import ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess my-partition zpool-12345/my-partition/123.123.123.123%!F(MISSING)32`
 * ```
 */
export class NasHAPartitionAccess extends pulumi.CustomResource {
    /**
     * Get an existing NasHAPartitionAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasHAPartitionAccessState, opts?: pulumi.CustomResourceOptions): NasHAPartitionAccess {
        return new NasHAPartitionAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess';

    /**
     * Returns true if the given object is an instance of NasHAPartitionAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasHAPartitionAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasHAPartitionAccess.__pulumiType;
    }

    /**
     * ip block in x.x.x.x/x format
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * name of the partition
     */
    public readonly partitionName!: pulumi.Output<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * one of "readwrite", "readonly"
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a NasHAPartitionAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasHAPartitionAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasHAPartitionAccessArgs | NasHAPartitionAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasHAPartitionAccessState | undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["partitionName"] = state ? state.partitionName : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NasHAPartitionAccessArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.partitionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["partitionName"] = args ? args.partitionName : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NasHAPartitionAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasHAPartitionAccess resources.
 */
export interface NasHAPartitionAccessState {
    /**
     * ip block in x.x.x.x/x format
     */
    ip?: pulumi.Input<string>;
    /**
     * name of the partition
     */
    partitionName?: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName?: pulumi.Input<string>;
    /**
     * one of "readwrite", "readonly"
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NasHAPartitionAccess resource.
 */
export interface NasHAPartitionAccessArgs {
    /**
     * ip block in x.x.x.x/x format
     */
    ip: pulumi.Input<string>;
    /**
     * name of the partition
     */
    partitionName: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName: pulumi.Input<string>;
    /**
     * one of "readwrite", "readonly"
     */
    type?: pulumi.Input<string>;
}
