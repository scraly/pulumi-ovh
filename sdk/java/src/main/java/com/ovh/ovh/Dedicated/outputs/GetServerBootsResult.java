// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetServerBootsResult {
    private @Nullable String bootType;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String kernel;
    /**
     * @return The list of dedicated server netboots.
     * 
     */
    private List<Integer> results;
    private String serviceName;

    private GetServerBootsResult() {}
    public Optional<String> bootType() {
        return Optional.ofNullable(this.bootType);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> kernel() {
        return Optional.ofNullable(this.kernel);
    }
    /**
     * @return The list of dedicated server netboots.
     * 
     */
    public List<Integer> results() {
        return this.results;
    }
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerBootsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String bootType;
        private String id;
        private @Nullable String kernel;
        private List<Integer> results;
        private String serviceName;
        public Builder() {}
        public Builder(GetServerBootsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bootType = defaults.bootType;
    	      this.id = defaults.id;
    	      this.kernel = defaults.kernel;
    	      this.results = defaults.results;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder bootType(@Nullable String bootType) {

            this.bootType = bootType;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServerBootsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kernel(@Nullable String kernel) {

            this.kernel = kernel;
            return this;
        }
        @CustomType.Setter
        public Builder results(List<Integer> results) {
            if (results == null) {
              throw new MissingRequiredPropertyException("GetServerBootsResult", "results");
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
              throw new MissingRequiredPropertyException("GetServerBootsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetServerBootsResult build() {
            final var _resultValue = new GetServerBootsResult();
            _resultValue.bootType = bootType;
            _resultValue.id = id;
            _resultValue.kernel = kernel;
            _resultValue.results = results;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
