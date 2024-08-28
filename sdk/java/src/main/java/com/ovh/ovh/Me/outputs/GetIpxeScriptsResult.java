// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIpxeScriptsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of the names of all the IPXE Scripts.
     * 
     */
    private List<String> results;

    private GetIpxeScriptsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of the names of all the IPXE Scripts.
     * 
     */
    public List<String> results() {
        return this.results;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpxeScriptsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> results;
        public Builder() {}
        public Builder(GetIpxeScriptsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.results = defaults.results;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIpxeScriptsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder results(List<String> results) {
            if (results == null) {
              throw new MissingRequiredPropertyException("GetIpxeScriptsResult", "results");
            }
            this.results = results;
            return this;
        }
        public Builder results(String... results) {
            return results(List.of(results));
        }
        public GetIpxeScriptsResult build() {
            final var _resultValue = new GetIpxeScriptsResult();
            _resultValue.id = id;
            _resultValue.results = results;
            return _resultValue;
        }
    }
}
