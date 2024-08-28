// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas;

import com.ovh.ovh.Dbaas.LogsOutputGraylogStreamArgs;
import com.ovh.ovh.Dbaas.inputs.LogsOutputGraylogStreamState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a DBaaS Logs Graylog output stream.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Dbaas.LogsOutputGraylogStream;
 * import com.pulumi.ovh.Dbaas.LogsOutputGraylogStreamArgs;
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
 *         var stream = new LogsOutputGraylogStream("stream", LogsOutputGraylogStreamArgs.builder()
 *             .description("my graylog stream")
 *             .serviceName("....")
 *             .title("my stream")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * To define the retention of the stream, you can use the following configuration:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Dbaas.DbaasFunctions;
 * import com.pulumi.ovh.Dbaas.inputs.GetLogsClustersRetentionArgs;
 * import com.pulumi.ovh.Dbaas.LogsOutputGraylogStream;
 * import com.pulumi.ovh.Dbaas.LogsOutputGraylogStreamArgs;
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
 *         final var retention = DbaasFunctions.getLogsClustersRetention(GetLogsClustersRetentionArgs.builder()
 *             .serviceName("ldp-xx-xxxxx")
 *             .clusterId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .duration("P14D")
 *             .build());
 * 
 *         var stream = new LogsOutputGraylogStream("stream", LogsOutputGraylogStreamArgs.builder()
 *             .serviceName("....")
 *             .title("my stream")
 *             .description("my graylog stream")
 *             .retentionId(retention.applyValue(getLogsClustersRetentionResult -> getLogsClustersRetentionResult.retentionId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream")
public class LogsOutputGraylogStream extends com.pulumi.resources.CustomResource {
    /**
     * Indicates if the current user can create alert on the stream
     * 
     */
    @Export(name="canAlert", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> canAlert;

    /**
     * @return Indicates if the current user can create alert on the stream
     * 
     */
    public Output<Boolean> canAlert() {
        return this.canAlert;
    }
    /**
     * Cold storage compression method. One of &#34;LZMA&#34;, &#34;GZIP&#34;, &#34;DEFLATED&#34;, &#34;ZSTD&#34;
     * 
     */
    @Export(name="coldStorageCompression", refs={String.class}, tree="[0]")
    private Output<String> coldStorageCompression;

    /**
     * @return Cold storage compression method. One of &#34;LZMA&#34;, &#34;GZIP&#34;, &#34;DEFLATED&#34;, &#34;ZSTD&#34;
     * 
     */
    public Output<String> coldStorageCompression() {
        return this.coldStorageCompression;
    }
    /**
     * ColdStorage content. One of &#34;ALL&#34;, &#34;GLEF&#34;, &#34;PLAIN&#34;
     * 
     */
    @Export(name="coldStorageContent", refs={String.class}, tree="[0]")
    private Output<String> coldStorageContent;

    /**
     * @return ColdStorage content. One of &#34;ALL&#34;, &#34;GLEF&#34;, &#34;PLAIN&#34;
     * 
     */
    public Output<String> coldStorageContent() {
        return this.coldStorageContent;
    }
    /**
     * Is Cold storage enabled?
     * 
     */
    @Export(name="coldStorageEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> coldStorageEnabled;

    /**
     * @return Is Cold storage enabled?
     * 
     */
    public Output<Boolean> coldStorageEnabled() {
        return this.coldStorageEnabled;
    }
    /**
     * Notify on new Cold storage archive
     * 
     */
    @Export(name="coldStorageNotifyEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> coldStorageNotifyEnabled;

    /**
     * @return Notify on new Cold storage archive
     * 
     */
    public Output<Boolean> coldStorageNotifyEnabled() {
        return this.coldStorageNotifyEnabled;
    }
    /**
     * Cold storage retention in year
     * 
     */
    @Export(name="coldStorageRetention", refs={Integer.class}, tree="[0]")
    private Output<Integer> coldStorageRetention;

    /**
     * @return Cold storage retention in year
     * 
     */
    public Output<Integer> coldStorageRetention() {
        return this.coldStorageRetention;
    }
    /**
     * ColdStorage destination. One of &#34;PCA&#34;, &#34;PCS&#34;
     * 
     */
    @Export(name="coldStorageTarget", refs={String.class}, tree="[0]")
    private Output<String> coldStorageTarget;

    /**
     * @return ColdStorage destination. One of &#34;PCA&#34;, &#34;PCS&#34;
     * 
     */
    public Output<String> coldStorageTarget() {
        return this.coldStorageTarget;
    }
    /**
     * Stream creation
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Stream creation
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Stream description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Stream description
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Enable ES indexing
     * 
     */
    @Export(name="indexingEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> indexingEnabled;

    /**
     * @return Enable ES indexing
     * 
     */
    public Output<Boolean> indexingEnabled() {
        return this.indexingEnabled;
    }
    /**
     * Maximum indexing size (in GB)
     * 
     */
    @Export(name="indexingMaxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> indexingMaxSize;

    /**
     * @return Maximum indexing size (in GB)
     * 
     */
    public Output<Integer> indexingMaxSize() {
        return this.indexingMaxSize;
    }
    /**
     * If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    @Export(name="indexingNotifyEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> indexingNotifyEnabled;

    /**
     * @return If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    public Output<Boolean> indexingNotifyEnabled() {
        return this.indexingNotifyEnabled;
    }
    /**
     * Indicates if you are allowed to edit entry
     * 
     */
    @Export(name="isEditable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isEditable;

    /**
     * @return Indicates if you are allowed to edit entry
     * 
     */
    public Output<Boolean> isEditable() {
        return this.isEditable;
    }
    /**
     * Indicates if you are allowed to share entry
     * 
     */
    @Export(name="isShareable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isShareable;

    /**
     * @return Indicates if you are allowed to share entry
     * 
     */
    public Output<Boolean> isShareable() {
        return this.isShareable;
    }
    /**
     * Number of alert condition
     * 
     */
    @Export(name="nbAlertCondition", refs={Integer.class}, tree="[0]")
    private Output<Integer> nbAlertCondition;

    /**
     * @return Number of alert condition
     * 
     */
    public Output<Integer> nbAlertCondition() {
        return this.nbAlertCondition;
    }
    /**
     * Number of coldstored archivesr
     * 
     */
    @Export(name="nbArchive", refs={Integer.class}, tree="[0]")
    private Output<Integer> nbArchive;

    /**
     * @return Number of coldstored archivesr
     * 
     */
    public Output<Integer> nbArchive() {
        return this.nbArchive;
    }
    /**
     * Parent stream ID
     * 
     */
    @Export(name="parentStreamId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parentStreamId;

    /**
     * @return Parent stream ID
     * 
     */
    public Output<Optional<String>> parentStreamId() {
        return Codegen.optional(this.parentStreamId);
    }
    /**
     * If set, pause indexing when maximum size is reach
     * 
     */
    @Export(name="pauseIndexingOnMaxSize", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pauseIndexingOnMaxSize;

    /**
     * @return If set, pause indexing when maximum size is reach
     * 
     */
    public Output<Boolean> pauseIndexingOnMaxSize() {
        return this.pauseIndexingOnMaxSize;
    }
    /**
     * Retention ID
     * 
     */
    @Export(name="retentionId", refs={String.class}, tree="[0]")
    private Output<String> retentionId;

    /**
     * @return Retention ID
     * 
     */
    public Output<String> retentionId() {
        return this.retentionId;
    }
    /**
     * The service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Stream ID
     * 
     */
    @Export(name="streamId", refs={String.class}, tree="[0]")
    private Output<String> streamId;

    /**
     * @return Stream ID
     * 
     */
    public Output<String> streamId() {
        return this.streamId;
    }
    /**
     * Stream description
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Stream description
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * Stream last updater
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Stream last updater
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Enable Websocket
     * 
     */
    @Export(name="webSocketEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> webSocketEnabled;

    /**
     * @return Enable Websocket
     * 
     */
    public Output<Boolean> webSocketEnabled() {
        return this.webSocketEnabled;
    }
    /**
     * Write token of the stream (empty if the caller is not the owner of the stream)
     * 
     */
    @Export(name="writeToken", refs={String.class}, tree="[0]")
    private Output<String> writeToken;

    /**
     * @return Write token of the stream (empty if the caller is not the owner of the stream)
     * 
     */
    public Output<String> writeToken() {
        return this.writeToken;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogsOutputGraylogStream(String name) {
        this(name, LogsOutputGraylogStreamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogsOutputGraylogStream(String name, LogsOutputGraylogStreamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogsOutputGraylogStream(String name, LogsOutputGraylogStreamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private LogsOutputGraylogStream(String name, Output<String> id, @Nullable LogsOutputGraylogStreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream", name, state, makeResourceOptions(options, id));
    }

    private static LogsOutputGraylogStreamArgs makeArgs(LogsOutputGraylogStreamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LogsOutputGraylogStreamArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "writeToken"
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
    public static LogsOutputGraylogStream get(String name, Output<String> id, @Nullable LogsOutputGraylogStreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogsOutputGraylogStream(name, id, state, options);
    }
}
