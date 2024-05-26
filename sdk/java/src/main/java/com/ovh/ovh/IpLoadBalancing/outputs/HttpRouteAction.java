// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HttpRouteAction {
    /**
     * @return HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
     * 
     */
    private @Nullable Integer status;
    /**
     * @return Farm ID for &#34;farm&#34; action type or URL template for &#34;redirect&#34; action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target
     * 
     */
    private @Nullable String target;
    /**
     * @return Action to trigger if all the rules of this route matches
     * 
     */
    private String type;

    private HttpRouteAction() {}
    /**
     * @return HTTP status code for &#34;redirect&#34; and &#34;reject&#34; actions
     * 
     */
    public Optional<Integer> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return Farm ID for &#34;farm&#34; action type or URL template for &#34;redirect&#34; action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target
     * 
     */
    public Optional<String> target() {
        return Optional.ofNullable(this.target);
    }
    /**
     * @return Action to trigger if all the rules of this route matches
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HttpRouteAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer status;
        private @Nullable String target;
        private String type;
        public Builder() {}
        public Builder(HttpRouteAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
    	      this.target = defaults.target;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder status(@Nullable Integer status) {

            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder target(@Nullable String target) {

            this.target = target;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("HttpRouteAction", "type");
            }
            this.type = type;
            return this;
        }
        public HttpRouteAction build() {
            final var _resultValue = new HttpRouteAction();
            _resultValue.status = status;
            _resultValue.target = target;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
