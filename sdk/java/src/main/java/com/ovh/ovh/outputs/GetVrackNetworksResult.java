// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVrackNetworksResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of vrack network ids.
     * 
     */
    private List<Integer> results;
    private String serviceName;
    private @Nullable String subnet;
    private @Nullable Integer vlanId;

    private GetVrackNetworksResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of vrack network ids.
     * 
     */
    public List<Integer> results() {
        return this.results;
    }
    public String serviceName() {
        return this.serviceName;
    }
    public Optional<String> subnet() {
        return Optional.ofNullable(this.subnet);
    }
    public Optional<Integer> vlanId() {
        return Optional.ofNullable(this.vlanId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVrackNetworksResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<Integer> results;
        private String serviceName;
        private @Nullable String subnet;
        private @Nullable Integer vlanId;
        public Builder() {}
        public Builder(GetVrackNetworksResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.results = defaults.results;
    	      this.serviceName = defaults.serviceName;
    	      this.subnet = defaults.subnet;
    	      this.vlanId = defaults.vlanId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworksResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder results(List<Integer> results) {
            if (results == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworksResult", "results");
            }
            this.results = results;
            return this;
        }
        public Builder results(Integer... results) {
            return results(List.of(results));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworksResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder subnet(@Nullable String subnet) {

            this.subnet = subnet;
            return this;
        }
        @CustomType.Setter
        public Builder vlanId(@Nullable Integer vlanId) {

            this.vlanId = vlanId;
            return this;
        }
        public GetVrackNetworksResult build() {
            final var _resultValue = new GetVrackNetworksResult();
            _resultValue.id = id;
            _resultValue.results = results;
            _resultValue.serviceName = serviceName;
            _resultValue.subnet = subnet;
            _resultValue.vlanId = vlanId;
            return _resultValue;
        }
    }
}
