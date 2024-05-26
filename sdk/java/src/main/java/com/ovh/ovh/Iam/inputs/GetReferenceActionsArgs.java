// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Iam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetReferenceActionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReferenceActionsArgs Empty = new GetReferenceActionsArgs();

    /**
     * Kind of resource we want the actions for
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Kind of resource we want the actions for
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GetReferenceActionsArgs() {}

    private GetReferenceActionsArgs(GetReferenceActionsArgs $) {
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReferenceActionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReferenceActionsArgs $;

        public Builder() {
            $ = new GetReferenceActionsArgs();
        }

        public Builder(GetReferenceActionsArgs defaults) {
            $ = new GetReferenceActionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type Kind of resource we want the actions for
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Kind of resource we want the actions for
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetReferenceActionsArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("GetReferenceActionsArgs", "type");
            }
            return $;
        }
    }

}
