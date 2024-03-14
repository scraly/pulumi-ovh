// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information of order cart product options.
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
 * const options = mycart.then(mycart => ovh.Order.getCartProductOptions({
 *     cartId: mycart.id,
 *     product: "cloud",
 *     planCode: "project",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCartProductOptions(args: GetCartProductOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetCartProductOptionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Order/getCartProductOptions:getCartProductOptions", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "planCode": args.planCode,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getCartProductOptions.
 */
export interface GetCartProductOptionsArgs {
    /**
     * Cart identifier
     */
    cartId: string;
    /**
     * Catalog name
     */
    catalogName?: string;
    /**
     * Product offer identifier
     */
    planCode: string;
    /**
     * Product
     */
    product: string;
}

/**
 * A collection of values returned by getCartProductOptions.
 */
export interface GetCartProductOptionsResult {
    readonly cartId: string;
    readonly catalogName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Product offer identifier
     */
    readonly planCode: string;
    readonly product: string;
    /**
     * products results
     */
    readonly results: outputs.Order.GetCartProductOptionsResult[];
}
/**
 * Use this data source to retrieve information of order cart product options.
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
 * const options = mycart.then(mycart => ovh.Order.getCartProductOptions({
 *     cartId: mycart.id,
 *     product: "cloud",
 *     planCode: "project",
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCartProductOptionsOutput(args: GetCartProductOptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCartProductOptionsResult> {
    return pulumi.output(args).apply((a: any) => getCartProductOptions(a, opts))
}

/**
 * A collection of arguments for invoking getCartProductOptions.
 */
export interface GetCartProductOptionsOutputArgs {
    /**
     * Cart identifier
     */
    cartId: pulumi.Input<string>;
    /**
     * Catalog name
     */
    catalogName?: pulumi.Input<string>;
    /**
     * Product offer identifier
     */
    planCode: pulumi.Input<string>;
    /**
     * Product
     */
    product: pulumi.Input<string>;
}
