// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * OVHcloud Managed Redis clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g., bash
 *
 * ```sh
 *  $ pulumi import ovh:CloudProjectDatabase/redisUser:RedisUser my_user service_name/cluster_id/id
 * ```
 */
export class RedisUser extends pulumi.CustomResource {
    /**
     * Get an existing RedisUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RedisUserState, opts?: pulumi.CustomResourceOptions): RedisUser {
        return new RedisUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/redisUser:RedisUser';

    /**
     * Returns true if the given object is an instance of RedisUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisUser.__pulumiType;
    }

    /**
     * Categories of the user.
     */
    public readonly categories!: pulumi.Output<string[] | undefined>;
    /**
     * Channels of the user.
     */
    public readonly channels!: pulumi.Output<string[]>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Commands of the user.
     */
    public readonly commands!: pulumi.Output<string[] | undefined>;
    /**
     * Date of the creation of the user.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Keys of the user.
     */
    public readonly keys!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Sensitive) Password of the user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    public readonly passwordReset!: pulumi.Output<string | undefined>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current status of the user.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a RedisUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RedisUserArgs | RedisUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RedisUserState | undefined;
            resourceInputs["categories"] = state ? state.categories : undefined;
            resourceInputs["channels"] = state ? state.channels : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["commands"] = state ? state.commands : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["keys"] = state ? state.keys : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordReset"] = state ? state.passwordReset : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RedisUserArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["channels"] = args ? args.channels : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["commands"] = args ? args.commands : undefined;
            resourceInputs["keys"] = args ? args.keys : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passwordReset"] = args ? args.passwordReset : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RedisUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RedisUser resources.
 */
export interface RedisUserState {
    /**
     * Categories of the user.
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Channels of the user.
     */
    channels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Commands of the user.
     */
    commands?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date of the creation of the user.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Keys of the user.
     */
    keys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
     */
    name?: pulumi.Input<string>;
    /**
     * (Sensitive) Password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    passwordReset?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current status of the user.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RedisUser resource.
 */
export interface RedisUserArgs {
    /**
     * Categories of the user.
     */
    categories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Channels of the user.
     */
    channels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Commands of the user.
     */
    commands?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Keys of the user.
     */
    keys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
     */
    name?: pulumi.Input<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    passwordReset?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
