// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIdentityUsersResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of the user&#39;s logins of all the identity users.
     * 
     */
    private List<String> users;

    private GetIdentityUsersResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of the user&#39;s logins of all the identity users.
     * 
     */
    public List<String> users() {
        return this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityUsersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> users;
        public Builder() {}
        public Builder(GetIdentityUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder users(List<String> users) {
            if (users == null) {
              throw new MissingRequiredPropertyException("GetIdentityUsersResult", "users");
            }
            this.users = users;
            return this;
        }
        public Builder users(String... users) {
            return users(List.of(users));
        }
        public GetIdentityUsersResult build() {
            final var _resultValue = new GetIdentityUsersResult();
            _resultValue.id = id;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}
