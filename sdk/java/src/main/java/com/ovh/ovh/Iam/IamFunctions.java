// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Iam;

import com.ovh.ovh.Iam.inputs.GetPolicyArgs;
import com.ovh.ovh.Iam.inputs.GetPolicyPlainArgs;
import com.ovh.ovh.Iam.inputs.GetReferenceActionsArgs;
import com.ovh.ovh.Iam.inputs.GetReferenceActionsPlainArgs;
import com.ovh.ovh.Iam.inputs.GetResourceGroupArgs;
import com.ovh.ovh.Iam.inputs.GetResourceGroupPlainArgs;
import com.ovh.ovh.Iam.outputs.GetPoliciesResult;
import com.ovh.ovh.Iam.outputs.GetPolicyResult;
import com.ovh.ovh.Iam.outputs.GetReferenceActionsResult;
import com.ovh.ovh.Iam.outputs.GetReferenceResourceTypeResult;
import com.ovh.ovh.Iam.outputs.GetResourceGroupResult;
import com.ovh.ovh.Iam.outputs.GetResourceGroupsResult;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class IamFunctions {
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetPoliciesResult> getPolicies() {
        return getPolicies(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetPoliciesResult> getPoliciesPlain() {
        return getPoliciesPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetPoliciesResult> getPolicies(InvokeArgs args) {
        return getPolicies(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetPoliciesResult> getPoliciesPlain(InvokeArgs args) {
        return getPoliciesPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetPoliciesResult> getPolicies(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getPolicies:getPolicies", TypeShape.of(GetPoliciesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myPolicies = IamFunctions.getPolicies();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetPoliciesResult> getPoliciesPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getPolicies:getPolicies", TypeShape.of(GetPoliciesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve am IAM policy.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetPolicyArgs;
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
     *         final var myPolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
     *             .id("my_policy_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args) {
        return getPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve am IAM policy.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetPolicyArgs;
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
     *         final var myPolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
     *             .id("my_policy_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args) {
        return getPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve am IAM policy.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetPolicyArgs;
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
     *         final var myPolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
     *             .id("my_policy_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve am IAM policy.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetPolicyArgs;
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
     *         final var myPolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
     *             .id("my_policy_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list the IAM action associated with a resource type.
     * 
     * ## Example Usage
     * 
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetReferenceActionsArgs;
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
     *         final var vpsActions = IamFunctions.getReferenceActions(GetReferenceActionsArgs.builder()
     *             .resourceType("vps")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetReferenceActionsResult> getReferenceActions(GetReferenceActionsArgs args) {
        return getReferenceActions(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the IAM action associated with a resource type.
     * 
     * ## Example Usage
     * 
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetReferenceActionsArgs;
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
     *         final var vpsActions = IamFunctions.getReferenceActions(GetReferenceActionsArgs.builder()
     *             .resourceType("vps")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetReferenceActionsResult> getReferenceActionsPlain(GetReferenceActionsPlainArgs args) {
        return getReferenceActionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the IAM action associated with a resource type.
     * 
     * ## Example Usage
     * 
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetReferenceActionsArgs;
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
     *         final var vpsActions = IamFunctions.getReferenceActions(GetReferenceActionsArgs.builder()
     *             .resourceType("vps")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetReferenceActionsResult> getReferenceActions(GetReferenceActionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getReferenceActions:getReferenceActions", TypeShape.of(GetReferenceActionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list the IAM action associated with a resource type.
     * 
     * ## Example Usage
     * 
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetReferenceActionsArgs;
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
     *         final var vpsActions = IamFunctions.getReferenceActions(GetReferenceActionsArgs.builder()
     *             .resourceType("vps")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetReferenceActionsResult> getReferenceActionsPlain(GetReferenceActionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getReferenceActions:getReferenceActions", TypeShape.of(GetReferenceActionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetReferenceResourceTypeResult> getReferenceResourceType() {
        return getReferenceResourceType(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetReferenceResourceTypeResult> getReferenceResourceTypePlain() {
        return getReferenceResourceTypePlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetReferenceResourceTypeResult> getReferenceResourceType(InvokeArgs args) {
        return getReferenceResourceType(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetReferenceResourceTypeResult> getReferenceResourceTypePlain(InvokeArgs args) {
        return getReferenceResourceTypePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetReferenceResourceTypeResult> getReferenceResourceType(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getReferenceResourceType:getReferenceResourceType", TypeShape.of(GetReferenceResourceTypeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list all the IAM resource types.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var types = IamFunctions.getReferenceResourceType();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetReferenceResourceTypeResult> getReferenceResourceTypePlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getReferenceResourceType:getReferenceResourceType", TypeShape.of(GetReferenceResourceTypeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source get details about a resource group.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetResourceGroupArgs;
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
     *         final var myResourceGroup = IamFunctions.getResourceGroup(GetResourceGroupArgs.builder()
     *             .id("my_resource_group_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetResourceGroupResult> getResourceGroup(GetResourceGroupArgs args) {
        return getResourceGroup(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source get details about a resource group.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetResourceGroupArgs;
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
     *         final var myResourceGroup = IamFunctions.getResourceGroup(GetResourceGroupArgs.builder()
     *             .id("my_resource_group_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetResourceGroupResult> getResourceGroupPlain(GetResourceGroupPlainArgs args) {
        return getResourceGroupPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source get details about a resource group.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetResourceGroupArgs;
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
     *         final var myResourceGroup = IamFunctions.getResourceGroup(GetResourceGroupArgs.builder()
     *             .id("my_resource_group_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetResourceGroupResult> getResourceGroup(GetResourceGroupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getResourceGroup:getResourceGroup", TypeShape.of(GetResourceGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source get details about a resource group.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
     * import com.pulumi.ovh.Iam.inputs.GetResourceGroupArgs;
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
     *         final var myResourceGroup = IamFunctions.getResourceGroup(GetResourceGroupArgs.builder()
     *             .id("my_resource_group_id")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetResourceGroupResult> getResourceGroupPlain(GetResourceGroupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getResourceGroup:getResourceGroup", TypeShape.of(GetResourceGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetResourceGroupsResult> getResourceGroups() {
        return getResourceGroups(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetResourceGroupsResult> getResourceGroupsPlain() {
        return getResourceGroupsPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetResourceGroupsResult> getResourceGroups(InvokeArgs args) {
        return getResourceGroups(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetResourceGroupsResult> getResourceGroupsPlain(InvokeArgs args) {
        return getResourceGroupsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static Output<GetResourceGroupsResult> getResourceGroups(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Iam/getResourceGroups:getResourceGroups", TypeShape.of(GetResourceGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to list the existing IAM policies of an account.
     * 
     * ## Example Usage
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Iam.IamFunctions;
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
     *         final var myGroups = IamFunctions.getResourceGroups();
     * 
     *     }
     * }
     * }
     * </pre>
     * 
     */
    public static CompletableFuture<GetResourceGroupsResult> getResourceGroupsPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Iam/getResourceGroups:getResourceGroups", TypeShape.of(GetResourceGroupsResult.class), args, Utilities.withVersion(options));
    }
}