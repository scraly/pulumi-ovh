// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOpenSearchUserAcl {
    /**
     * @return Pattern of the ACL.
     * 
     */
    private String pattern;
    /**
     * @return Permission of the ACL.
     * 
     */
    private String permission;

    private GetOpenSearchUserAcl() {}
    /**
     * @return Pattern of the ACL.
     * 
     */
    public String pattern() {
        return this.pattern;
    }
    /**
     * @return Permission of the ACL.
     * 
     */
    public String permission() {
        return this.permission;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOpenSearchUserAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String pattern;
        private String permission;
        public Builder() {}
        public Builder(GetOpenSearchUserAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pattern = defaults.pattern;
    	      this.permission = defaults.permission;
        }

        @CustomType.Setter
        public Builder pattern(String pattern) {
            if (pattern == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchUserAcl", "pattern");
            }
            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            if (permission == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchUserAcl", "permission");
            }
            this.permission = permission;
            return this;
        }
        public GetOpenSearchUserAcl build() {
            final var _resultValue = new GetOpenSearchUserAcl();
            _resultValue.pattern = pattern;
            _resultValue.permission = permission;
            return _resultValue;
        }
    }
}
