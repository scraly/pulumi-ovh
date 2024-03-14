// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const rescue = ovh.Dedicated.getServerBoots({
 *     serviceName: "nsxxxxxxx.ip-xx-xx-xx.eu",
 *     bootType: "rescue",
 *     kernel: "rescue64-pro",
 * });
 * const server = new ovh.dedicated.ServerUpdate("server", {
 *     serviceName: "nsxxxxxxx.ip-xx-xx-xx.eu",
 *     bootId: rescue.then(rescue => rescue.results?.[0]),
 *     monitoring: true,
 *     state: "ok",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ServerUpdate extends pulumi.CustomResource {
    /**
     * Get an existing ServerUpdate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerUpdateState, opts?: pulumi.CustomResourceOptions): ServerUpdate {
        return new ServerUpdate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/serverUpdate:ServerUpdate';

    /**
     * Returns true if the given object is an instance of ServerUpdate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerUpdate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerUpdate.__pulumiType;
    }

    /**
     * boot id of the server
     */
    public readonly bootId!: pulumi.Output<number>;
    /**
     * boot script of the server
     */
    public readonly bootScript!: pulumi.Output<string | undefined>;
    /**
     * Icmp monitoring state
     */
    public readonly monitoring!: pulumi.Output<boolean>;
    /**
     * The serviceName of your dedicated server.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * error, hacked, hackedBlocked, ok
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a ServerUpdate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerUpdateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerUpdateArgs | ServerUpdateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerUpdateState | undefined;
            resourceInputs["bootId"] = state ? state.bootId : undefined;
            resourceInputs["bootScript"] = state ? state.bootScript : undefined;
            resourceInputs["monitoring"] = state ? state.monitoring : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as ServerUpdateArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["bootId"] = args ? args.bootId : undefined;
            resourceInputs["bootScript"] = args ? args.bootScript : undefined;
            resourceInputs["monitoring"] = args ? args.monitoring : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerUpdate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerUpdate resources.
 */
export interface ServerUpdateState {
    /**
     * boot id of the server
     */
    bootId?: pulumi.Input<number>;
    /**
     * boot script of the server
     */
    bootScript?: pulumi.Input<string>;
    /**
     * Icmp monitoring state
     */
    monitoring?: pulumi.Input<boolean>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * error, hacked, hackedBlocked, ok
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerUpdate resource.
 */
export interface ServerUpdateArgs {
    /**
     * boot id of the server
     */
    bootId?: pulumi.Input<number>;
    /**
     * boot script of the server
     */
    bootScript?: pulumi.Input<string>;
    /**
     * Icmp monitoring state
     */
    monitoring?: pulumi.Input<boolean>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
    /**
     * error, hacked, hackedBlocked, ok
     */
    state?: pulumi.Input<string>;
}
