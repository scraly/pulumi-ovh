// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Ip;

import com.ovh.ovh.Ip.inputs.GetServiceArgs;
import com.ovh.ovh.Ip.inputs.GetServicePlainArgs;
import com.ovh.ovh.Ip.outputs.GetServiceResult;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class IpFunctions {
    /**
     * Use this data source to retrieve information about an IP service.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Ip.IpFunctions;
     * import com.pulumi.ovh.Ip.inputs.GetServiceArgs;
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
     *         final var myip = IpFunctions.getService(GetServiceArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args) {
        return getService(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an IP service.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Ip.IpFunctions;
     * import com.pulumi.ovh.Ip.inputs.GetServiceArgs;
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
     *         final var myip = IpFunctions.getService(GetServiceArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args) {
        return getServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an IP service.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Ip.IpFunctions;
     * import com.pulumi.ovh.Ip.inputs.GetServiceArgs;
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
     *         final var myip = IpFunctions.getService(GetServiceArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Ip/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an IP service.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Ip.IpFunctions;
     * import com.pulumi.ovh.Ip.inputs.GetServiceArgs;
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
     *         final var myip = IpFunctions.getService(GetServiceArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Ip/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
}