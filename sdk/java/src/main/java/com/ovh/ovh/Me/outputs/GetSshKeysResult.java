// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSshKeysResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of the names of all the SSH keys.
     * 
     */
    private List<String> names;

    private GetSshKeysResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of the names of all the SSH keys.
     * 
     */
    public List<String> names() {
        return this.names;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSshKeysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> names;
        public Builder() {}
        public Builder(GetSshKeysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.names = defaults.names;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSshKeysResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            if (names == null) {
              throw new MissingRequiredPropertyException("GetSshKeysResult", "names");
            }
            this.names = names;
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        public GetSshKeysResult build() {
            final var _resultValue = new GetSshKeysResult();
            _resultValue.id = id;
            _resultValue.names = names;
            return _resultValue;
        }
    }
}
