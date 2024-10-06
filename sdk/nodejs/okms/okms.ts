// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const newkms = new ovh.okms.Okms("newkms", {
 *     displayName: "terraformed KMS",
 *     ovhSubsidiary: "FR",
 *     region: "EU_WEST_RBX",
 * });
 * ```
 */
export class Okms extends pulumi.CustomResource {
    /**
     * Get an existing Okms resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OkmsState, opts?: pulumi.CustomResourceOptions): Okms {
        return new Okms(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Okms/okms:Okms';

    /**
     * Returns true if the given object is an instance of Okms.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Okms {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Okms.__pulumiType;
    }

    /**
     * (String) Resource display name
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * (Attributes) IAM resource metadata (see below for nested schema)
     */
    public /*out*/ readonly iam!: pulumi.Output<outputs.Okms.OkmsIam>;
    /**
     * (String) KMS kmip API endpoint
     */
    public /*out*/ readonly kmipEndpoint!: pulumi.Output<string>;
    /**
     * OVH subsidiaries
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * (String) KMS public CA (Certificate Authority)
     */
    public /*out*/ readonly publicCa!: pulumi.Output<string>;
    /**
     * KMS region
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * (String) KMS rest API endpoint
     */
    public /*out*/ readonly restEndpoint!: pulumi.Output<string>;
    /**
     * (String) KMS rest API swagger UI
     */
    public /*out*/ readonly swaggerEndpoint!: pulumi.Output<string>;

    /**
     * Create a Okms resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OkmsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OkmsArgs | OkmsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OkmsState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["iam"] = state ? state.iam : undefined;
            resourceInputs["kmipEndpoint"] = state ? state.kmipEndpoint : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["publicCa"] = state ? state.publicCa : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["restEndpoint"] = state ? state.restEndpoint : undefined;
            resourceInputs["swaggerEndpoint"] = state ? state.swaggerEndpoint : undefined;
        } else {
            const args = argsOrState as OkmsArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["iam"] = undefined /*out*/;
            resourceInputs["kmipEndpoint"] = undefined /*out*/;
            resourceInputs["publicCa"] = undefined /*out*/;
            resourceInputs["restEndpoint"] = undefined /*out*/;
            resourceInputs["swaggerEndpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Okms.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Okms resources.
 */
export interface OkmsState {
    /**
     * (String) Resource display name
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Attributes) IAM resource metadata (see below for nested schema)
     */
    iam?: pulumi.Input<inputs.Okms.OkmsIam>;
    /**
     * (String) KMS kmip API endpoint
     */
    kmipEndpoint?: pulumi.Input<string>;
    /**
     * OVH subsidiaries
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * (String) KMS public CA (Certificate Authority)
     */
    publicCa?: pulumi.Input<string>;
    /**
     * KMS region
     */
    region?: pulumi.Input<string>;
    /**
     * (String) KMS rest API endpoint
     */
    restEndpoint?: pulumi.Input<string>;
    /**
     * (String) KMS rest API swagger UI
     */
    swaggerEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Okms resource.
 */
export interface OkmsArgs {
    /**
     * (String) Resource display name
     */
    displayName?: pulumi.Input<string>;
    /**
     * OVH subsidiaries
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * KMS region
     */
    region: pulumi.Input<string>;
}