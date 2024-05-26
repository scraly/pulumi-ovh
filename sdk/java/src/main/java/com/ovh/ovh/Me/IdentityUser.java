// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me;

import com.ovh.ovh.Me.IdentityUserArgs;
import com.ovh.ovh.Me.inputs.IdentityUserState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates an identity user.
 * 
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.IdentityUser;
 * import com.pulumi.ovh.Me.IdentityUserArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var myUser = new IdentityUser("myUser", IdentityUserArgs.builder()        
 *             .description("Some custom description")
 *             .email("my_login{@literal @}example.com")
 *             .group("DEFAULT")
 *             .login("my_login")
 *             .password("super-s3cr3t!password")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 */
@ResourceType(type="ovh:Me/identityUser:IdentityUser")
public class IdentityUser extends com.pulumi.resources.CustomResource {
    /**
     * URN of the user, used when writing IAM policies
     * 
     */
    @Export(name="UserURN", refs={String.class}, tree="[0]")
    private Output<String> UserURN;

    /**
     * @return URN of the user, used when writing IAM policies
     * 
     */
    public Output<String> UserURN() {
        return this.UserURN;
    }
    /**
     * Creation date of this user.
     * 
     */
    @Export(name="creation", refs={String.class}, tree="[0]")
    private Output<String> creation;

    /**
     * @return Creation date of this user.
     * 
     */
    public Output<String> creation() {
        return this.creation;
    }
    /**
     * User description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User&#39;s email.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return User&#39;s email.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * User&#39;s group.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> group;

    /**
     * @return User&#39;s group.
     * 
     */
    public Output<Optional<String>> group() {
        return Codegen.optional(this.group);
    }
    /**
     * Last update of this user.
     * 
     */
    @Export(name="lastUpdate", refs={String.class}, tree="[0]")
    private Output<String> lastUpdate;

    /**
     * @return Last update of this user.
     * 
     */
    public Output<String> lastUpdate() {
        return this.lastUpdate;
    }
    /**
     * User&#39;s login suffix.
     * 
     */
    @Export(name="login", refs={String.class}, tree="[0]")
    private Output<String> login;

    /**
     * @return User&#39;s login suffix.
     * 
     */
    public Output<String> login() {
        return this.login;
    }
    /**
     * User&#39;s password.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return User&#39;s password.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * When the user changed his password for the last time.
     * 
     */
    @Export(name="passwordLastUpdate", refs={String.class}, tree="[0]")
    private Output<String> passwordLastUpdate;

    /**
     * @return When the user changed his password for the last time.
     * 
     */
    public Output<String> passwordLastUpdate() {
        return this.passwordLastUpdate;
    }
    /**
     * Current user&#39;s status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Current user&#39;s status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IdentityUser(String name) {
        this(name, IdentityUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IdentityUser(String name, IdentityUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IdentityUser(String name, IdentityUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/identityUser:IdentityUser", name, args == null ? IdentityUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IdentityUser(String name, Output<String> id, @Nullable IdentityUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/identityUser:IdentityUser", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static IdentityUser get(String name, Output<String> id, @Nullable IdentityUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IdentityUser(name, id, state, options);
    }
}
