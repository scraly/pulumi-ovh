// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 */
export class Vrack extends pulumi.CustomResource {
    /**
     * Get an existing Vrack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VrackState, opts?: pulumi.CustomResourceOptions): Vrack {
        return new Vrack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Vrack/vrack:Vrack';

    /**
     * Returns true if the given object is an instance of Vrack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vrack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vrack.__pulumiType;
    }

    /**
     * yourvrackdescription
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * yourvrackname
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.Vrack.VrackOrder[]>;
    /**
     * OVHcloud Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    public readonly paymentMean!: pulumi.Output<string | undefined>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.Vrack.VrackPlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.Vrack.VrackPlanOption[] | undefined>;
    /**
     * The internal name of your vrack
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * The URN of the vrack, used with IAM permissions
     */
    public /*out*/ readonly urn!: pulumi.Output<string>;

    /**
     * Create a Vrack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VrackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VrackArgs | VrackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VrackState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["urn"] = state ? state.urn : undefined;
        } else {
            const args = argsOrState as VrackArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["orders"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["urn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vrack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vrack resources.
 */
export interface VrackState {
    /**
     * yourvrackdescription
     */
    description?: pulumi.Input<string>;
    /**
     * yourvrackname
     */
    name?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackOrder>[]>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.Vrack.VrackPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackPlanOption>[]>;
    /**
     * The internal name of your vrack
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The URN of the vrack, used with IAM permissions
     */
    urn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vrack resource.
 */
export interface VrackArgs {
    /**
     * yourvrackdescription
     */
    description?: pulumi.Input<string>;
    /**
     * yourvrackname
     */
    name?: pulumi.Input<string>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.Vrack.VrackPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackPlanOption>[]>;
}