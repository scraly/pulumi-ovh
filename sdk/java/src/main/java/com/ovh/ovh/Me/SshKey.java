// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me;

import com.ovh.ovh.Me.SshKeyArgs;
import com.ovh.ovh.Me.inputs.SshKeyState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.SshKey;
 * import com.pulumi.ovh.Me.SshKeyArgs;
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
 *         var mykey = new SshKey("mykey", SshKeyArgs.builder()        
 *             .key("ssh-ed25519 AAAAC3...")
 *             .keyName("mykey")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 */
@ResourceType(type="ovh:Me/sshKey:SshKey")
public class SshKey extends com.pulumi.resources.CustomResource {
    /**
     * True when this public SSH key is used for rescue mode and reinstallations.
     * 
     */
    @Export(name="default", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> default_;

    /**
     * @return True when this public SSH key is used for rescue mode and reinstallations.
     * 
     */
    public Output<Boolean> default_() {
        return this.default_;
    }
    /**
     * The content of the public key in the form &#34;ssh-algo content&#34;, e.g. &#34;ssh-ed25519 AAAAC3...&#34;.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The content of the public key in the form &#34;ssh-algo content&#34;, e.g. &#34;ssh-ed25519 AAAAC3...&#34;.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The friendly name of this SSH key.
     * 
     */
    @Export(name="keyName", refs={String.class}, tree="[0]")
    private Output<String> keyName;

    /**
     * @return The friendly name of this SSH key.
     * 
     */
    public Output<String> keyName() {
        return this.keyName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SshKey(String name) {
        this(name, SshKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SshKey(String name, SshKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SshKey(String name, SshKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/sshKey:SshKey", name, args == null ? SshKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SshKey(String name, Output<String> id, @Nullable SshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/sshKey:SshKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static SshKey get(String name, Output<String> id, @Nullable SshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SshKey(name, id, state, options);
    }
}
