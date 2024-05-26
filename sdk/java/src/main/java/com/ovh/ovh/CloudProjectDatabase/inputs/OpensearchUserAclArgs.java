// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class OpensearchUserAclArgs extends com.pulumi.resources.ResourceArgs {

    public static final OpensearchUserAclArgs Empty = new OpensearchUserAclArgs();

    /**
     * Pattern of the ACL.
     * 
     */
    @Import(name="pattern", required=true)
    private Output<String> pattern;

    /**
     * @return Pattern of the ACL.
     * 
     */
    public Output<String> pattern() {
        return this.pattern;
    }

    /**
     * Permission of the ACL
     * Available permission:
     * 
     */
    @Import(name="permission", required=true)
    private Output<String> permission;

    /**
     * @return Permission of the ACL
     * Available permission:
     * 
     */
    public Output<String> permission() {
        return this.permission;
    }

    private OpensearchUserAclArgs() {}

    private OpensearchUserAclArgs(OpensearchUserAclArgs $) {
        this.pattern = $.pattern;
        this.permission = $.permission;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OpensearchUserAclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OpensearchUserAclArgs $;

        public Builder() {
            $ = new OpensearchUserAclArgs();
        }

        public Builder(OpensearchUserAclArgs defaults) {
            $ = new OpensearchUserAclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param pattern Pattern of the ACL.
         * 
         * @return builder
         * 
         */
        public Builder pattern(Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        /**
         * @param pattern Pattern of the ACL.
         * 
         * @return builder
         * 
         */
        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        /**
         * @param permission Permission of the ACL
         * Available permission:
         * 
         * @return builder
         * 
         */
        public Builder permission(Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Permission of the ACL
         * Available permission:
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        public OpensearchUserAclArgs build() {
            if ($.pattern == null) {
                throw new MissingRequiredPropertyException("OpensearchUserAclArgs", "pattern");
            }
            if ($.permission == null) {
                throw new MissingRequiredPropertyException("OpensearchUserAclArgs", "permission");
            }
            return $;
        }
    }

}
