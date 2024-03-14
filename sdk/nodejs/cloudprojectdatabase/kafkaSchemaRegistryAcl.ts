// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const kafka = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXX",
 *     engine: "kafka",
 *     id: "ZZZ",
 * });
 * const schemaRegistryAcl = new ovh.cloudprojectdatabase.KafkaSchemaRegistryAcl("schemaRegistryAcl", {
 *     serviceName: kafka.then(kafka => kafka.serviceName),
 *     clusterId: kafka.then(kafka => kafka.id),
 *     permission: "schema_registry_read",
 *     resource: "Subject:myResource",
 *     username: "johndoe",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
 * ```
 */
export class KafkaSchemaRegistryAcl extends pulumi.CustomResource {
    /**
     * Get an existing KafkaSchemaRegistryAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaSchemaRegistryAclState, opts?: pulumi.CustomResourceOptions): KafkaSchemaRegistryAcl {
        return new KafkaSchemaRegistryAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl';

    /**
     * Returns true if the given object is an instance of KafkaSchemaRegistryAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaSchemaRegistryAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaSchemaRegistryAcl.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Permission to give to this username on this resource.
     * Available permissions:
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * Resource affected by this schema registry ACL.
     */
    public readonly resource!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Username affected by this schema registry ACL.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a KafkaSchemaRegistryAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaSchemaRegistryAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaSchemaRegistryAclArgs | KafkaSchemaRegistryAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaSchemaRegistryAclState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["resource"] = state ? state.resource : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as KafkaSchemaRegistryAclArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KafkaSchemaRegistryAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaSchemaRegistryAcl resources.
 */
export interface KafkaSchemaRegistryAclState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Permission to give to this username on this resource.
     * Available permissions:
     */
    permission?: pulumi.Input<string>;
    /**
     * Resource affected by this schema registry ACL.
     */
    resource?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Username affected by this schema registry ACL.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KafkaSchemaRegistryAcl resource.
 */
export interface KafkaSchemaRegistryAclArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Permission to give to this username on this resource.
     * Available permissions:
     */
    permission: pulumi.Input<string>;
    /**
     * Resource affected by this schema registry ACL.
     */
    resource: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Username affected by this schema registry ACL.
     */
    username: pulumi.Input<string>;
}
