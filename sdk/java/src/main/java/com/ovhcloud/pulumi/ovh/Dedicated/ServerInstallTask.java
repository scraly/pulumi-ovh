// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated;

import com.ovhcloud.pulumi.ovh.Dedicated.ServerInstallTaskArgs;
import com.ovhcloud.pulumi.ovh.Dedicated.inputs.ServerInstallTaskState;
import com.ovhcloud.pulumi.ovh.Dedicated.outputs.ServerInstallTaskDetails;
import com.ovhcloud.pulumi.ovh.Dedicated.outputs.ServerInstallTaskUserMetadata;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Using a custom template based on an OVHCloud template
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Dedicated.DedicatedFunctions;
 * import com.pulumi.ovh.Dedicated.inputs.GetServerBootsArgs;
 * import com.pulumi.ovh.Me.InstallationTemplate;
 * import com.pulumi.ovh.Me.InstallationTemplateArgs;
 * import com.pulumi.ovh.Me.inputs.InstallationTemplateCustomizationArgs;
 * import com.pulumi.ovh.Dedicated.ServerInstallTask;
 * import com.pulumi.ovh.Dedicated.ServerInstallTaskArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskDetailsArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskUserMetadataArgs;
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
 *         final var rescue = DedicatedFunctions.getServerBoots(GetServerBootsArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .bootType("rescue")
 *             .build());
 * 
 *         var debian = new InstallationTemplate("debian", InstallationTemplateArgs.builder()
 *             .baseTemplateName("debian12_64")
 *             .templateName("mydebian12")
 *             .customization(InstallationTemplateCustomizationArgs.builder()
 *                 .customHostname("mytest")
 *                 .build())
 *             .build());
 * 
 *         var serverInstall = new ServerInstallTask("serverInstall", ServerInstallTaskArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .templateName(debian.templateName())
 *             .bootidOnDestroy(rescue.applyValue(getServerBootsResult -> getServerBootsResult.results()[0]))
 *             .details(ServerInstallTaskDetailsArgs.builder()
 *                 .customHostname("mytest")
 *                 .build())
 *             .userMetadatas(            
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("sshKey")
 *                     .value("ssh-ed25519 AAAAC3...")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("postInstallationScript")
 *                     .value("""
 * #!/bin/bash
 *   echo "coucou postInstallationScript" > /opt/coucou
 *   cat /etc/machine-id  >> /opt/coucou
 *   date "+%Y-%m-%d %H:%M:%S" --utc >> /opt/coucou
 *                     """)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Using a BringYourOwnLinux (BYOLinux) template (with userMetadata)
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.OvhFunctions;
 * import com.pulumi.ovh.inputs.GetServerArgs;
 * import com.pulumi.ovh.Dedicated.DedicatedFunctions;
 * import com.pulumi.ovh.Dedicated.inputs.GetServerBootsArgs;
 * import com.pulumi.ovh.Dedicated.ServerInstallTask;
 * import com.pulumi.ovh.Dedicated.ServerInstallTaskArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskDetailsArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskUserMetadataArgs;
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
 *         final var server = OvhFunctions.getServer(GetServerArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .build());
 * 
 *         final var rescue = DedicatedFunctions.getServerBoots(GetServerBootsArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .bootType("rescue")
 *             .build());
 * 
 *         var serverInstall = new ServerInstallTask("serverInstall", ServerInstallTaskArgs.builder()
 *             .serviceName(server.applyValue(getServerResult -> getServerResult.serviceName()))
 *             .templateName("byolinux_64")
 *             .bootidOnDestroy(rescue.applyValue(getServerBootsResult -> getServerBootsResult.results()[0]))
 *             .details(ServerInstallTaskDetailsArgs.builder()
 *                 .customHostname("mytest")
 *                 .build())
 *             .userMetadatas(            
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("imageURL")
 *                     .value("https://myimage.qcow2")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("imageType")
 *                     .value("qcow2")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("httpHeaders0Key")
 *                     .value("Authorization")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("httpHeaders0Value")
 *                     .value("Basic bG9naW46xxxxxxx=")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("imageChecksumType")
 *                     .value("sha512")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("imageCheckSum")
 *                     .value("047122c9ff4d2a69512212104b06c678f5a9cdb22b75467353613ff87ccd03b57b38967e56d810e61366f9d22d6bd39ac0addf4e00a4c6445112a2416af8f225")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("configDriveUserData")
 *                     .value("""
 * #cloud-config
 * ssh_authorized_keys:
 *   - %s
 * 
 * users:
 *   - name: patient0
 *     sudo: ALL=(ALL) NOPASSWD:ALL
 *     groups: users, sudo
 *     shell: /bin/bash
 *     lock_passwd: false
 *     ssh_authorized_keys:
 *       - %s
 * disable_root: false
 * packages:
 *   - vim
 *   - tree
 * final_message: The system is finally up, after $UPTIME seconds
 * ", data.ovh_me_ssh_key().mykey().key(),data.ovh_me_ssh_key().mykey().key()))
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Using a Microsoft Windows server OVHcloud template with a specific language
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.OvhFunctions;
 * import com.pulumi.ovh.inputs.GetServerArgs;
 * import com.pulumi.ovh.Dedicated.DedicatedFunctions;
 * import com.pulumi.ovh.Dedicated.inputs.GetServerBootsArgs;
 * import com.pulumi.ovh.Dedicated.ServerInstallTask;
 * import com.pulumi.ovh.Dedicated.ServerInstallTaskArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskDetailsArgs;
 * import com.pulumi.ovh.Dedicated.inputs.ServerInstallTaskUserMetadataArgs;
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
 *         final var server = OvhFunctions.getServer(GetServerArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .build());
 * 
 *         final var rescue = DedicatedFunctions.getServerBoots(GetServerBootsArgs.builder()
 *             .serviceName("nsxxxxxxx.ip-xx-xx-xx.eu")
 *             .bootType("rescue")
 *             .build());
 * 
 *         var serverInstall = new ServerInstallTask("serverInstall", ServerInstallTaskArgs.builder()
 *             .serviceName(server.applyValue(getServerResult -> getServerResult.serviceName()))
 *             .templateName("win2019-std_64")
 *             .bootidOnDestroy(rescue.applyValue(getServerBootsResult -> getServerBootsResult.results()[0]))
 *             .details(ServerInstallTaskDetailsArgs.builder()
 *                 .customHostname("mytest")
 *                 .build())
 *             .userMetadatas(            
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("language")
 *                     .value("fr-fr")
 *                     .build(),
 *                 ServerInstallTaskUserMetadataArgs.builder()
 *                     .key("postInstallationScript")
 *                     .value("""
 * coucou postInstallationScriptPowerShell" | Out-File -FilePath "c:\ovhupd\script\coucou.txt"
 *       (Get-ItemProperty -LiteralPath "Registry::HKLM\SOFTWARE\Microsoft\Cryptography" -Name "MachineGuid").MachineGuid | Out-File -FilePath "c:\ovhupd\script\coucou.txt" -Append
 *       (Get-Date).ToUniversalTime().ToString("yyyy-MM-dd HH:mm:ss") | Out-File -FilePath "c:\ovhupd\script\coucou.txt" -Append
 *                     """)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Installation task can be imported using the `service_name` (`nsXXXX.ip...`) of the baremetal server, the `template_name` used  and ths `task_id`, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Dedicated/serverInstallTask:ServerInstallTask ovh_dedicated_server_install_task nsXXXX.ipXXXX/template_name/12345
 * ```
 * 
 */
@ResourceType(type="ovh:Dedicated/serverInstallTask:ServerInstallTask")
public class ServerInstallTask extends com.pulumi.resources.CustomResource {
    /**
     * If set, reboot the server on the specified boot id during destroy phase.
     * 
     */
    @Export(name="bootidOnDestroy", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> bootidOnDestroy;

    /**
     * @return If set, reboot the server on the specified boot id during destroy phase.
     * 
     */
    public Output<Optional<Integer>> bootidOnDestroy() {
        return Codegen.optional(this.bootidOnDestroy);
    }
    /**
     * Details of this task. (should be `Install asked`)
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output<String> comment;

    /**
     * @return Details of this task. (should be `Install asked`)
     * 
     */
    public Output<String> comment() {
        return this.comment;
    }
    /**
     * see `details` block below.
     * 
     */
    @Export(name="details", refs={ServerInstallTaskDetails.class}, tree="[0]")
    private Output</* @Nullable */ ServerInstallTaskDetails> details;

    /**
     * @return see `details` block below.
     * 
     */
    public Output<Optional<ServerInstallTaskDetails>> details() {
        return Codegen.optional(this.details);
    }
    /**
     * Completion date in RFC3339 format.
     * 
     */
    @Export(name="doneDate", refs={String.class}, tree="[0]")
    private Output<String> doneDate;

    /**
     * @return Completion date in RFC3339 format.
     * 
     */
    public Output<String> doneDate() {
        return this.doneDate;
    }
    /**
     * Function name (should be `hardInstall`).
     * 
     */
    @Export(name="function", refs={String.class}, tree="[0]")
    private Output<String> function;

    /**
     * @return Function name (should be `hardInstall`).
     * 
     */
    public Output<String> function() {
        return this.function;
    }
    /**
     * Last update in RFC3339 format.
     * 
     */
    @Export(name="lastUpdate", refs={String.class}, tree="[0]")
    private Output<String> lastUpdate;

    /**
     * @return Last update in RFC3339 format.
     * 
     */
    public Output<String> lastUpdate() {
        return this.lastUpdate;
    }
    /**
     * Partition scheme name.
     * 
     */
    @Export(name="partitionSchemeName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partitionSchemeName;

    /**
     * @return Partition scheme name.
     * 
     */
    public Output<Optional<String>> partitionSchemeName() {
        return Codegen.optional(this.partitionSchemeName);
    }
    /**
     * The service_name of your dedicated server.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service_name of your dedicated server.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Task creation date in RFC3339 format.
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output<String> startDate;

    /**
     * @return Task creation date in RFC3339 format.
     * 
     */
    public Output<String> startDate() {
        return this.startDate;
    }
    /**
     * Task status (should be `done`)
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Task status (should be `done`)
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Template name.
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output<String> templateName;

    /**
     * @return Template name.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }
    /**
     * see `user_metadata` block below.
     * 
     */
    @Export(name="userMetadatas", refs={List.class,ServerInstallTaskUserMetadata.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ServerInstallTaskUserMetadata>> userMetadatas;

    /**
     * @return see `user_metadata` block below.
     * 
     */
    public Output<Optional<List<ServerInstallTaskUserMetadata>>> userMetadatas() {
        return Codegen.optional(this.userMetadatas);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerInstallTask(java.lang.String name) {
        this(name, ServerInstallTaskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerInstallTask(java.lang.String name, ServerInstallTaskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerInstallTask(java.lang.String name, ServerInstallTaskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/serverInstallTask:ServerInstallTask", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServerInstallTask(java.lang.String name, Output<java.lang.String> id, @Nullable ServerInstallTaskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/serverInstallTask:ServerInstallTask", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerInstallTaskArgs makeArgs(ServerInstallTaskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerInstallTaskArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ServerInstallTask get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerInstallTaskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerInstallTask(name, id, state, options);
    }
}
