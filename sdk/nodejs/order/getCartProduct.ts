// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information of order cart product products.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * const mycart = myaccount.then(myaccount => ovh.Order.getCart({
 *     ovhSubsidiary: myaccount.ovhSubsidiary,
 * }));
 * const plans = mycart.then(mycart => ovh.Order.getCartProduct({
 *     cartId: mycart.id,
 *     product: "...",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCartProduct(args: GetCartProductArgs, opts?: pulumi.InvokeOptions): Promise<GetCartProductResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Order/getCartProduct:getCartProduct", {
        "cartId": args.cartId,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getCartProduct.
 */
export interface GetCartProductArgs {
    /**
     * Cart identifier
     */
    cartId: string;
    /**
     * product
     */
    product: string;
}

/**
 * A collection of values returned by getCartProduct.
 */
export interface GetCartProductResult {
    readonly cartId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    /**
     * products results
     */
    readonly results: outputs.Order.GetCartProductResult[];
}
/**
 * Use this data source to retrieve information of order cart product products.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * const mycart = myaccount.then(myaccount => ovh.Order.getCart({
 *     ovhSubsidiary: myaccount.ovhSubsidiary,
 * }));
 * const plans = mycart.then(mycart => ovh.Order.getCartProduct({
 *     cartId: mycart.id,
 *     product: "...",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCartProductOutput(args: GetCartProductOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCartProductResult> {
    return pulumi.output(args).apply((a: any) => getCartProduct(a, opts))
}

/**
 * A collection of arguments for invoking getCartProduct.
 */
export interface GetCartProductOutputArgs {
    /**
     * Cart identifier
     */
    cartId: pulumi.Input<string>;
    /**
     * product
     */
    product: pulumi.Input<string>;
}
