// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSshKeyResult {
    /**
     * @return True when this public SSH key is used for rescue mode and reinstallations.
     * 
     */
    private Boolean default_;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The content of the public key.
     * E.g.: &#34;ssh-ed25519 AAAAC3...&#34;
     * 
     */
    private String key;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String keyName;

    private GetSshKeyResult() {}
    /**
     * @return True when this public SSH key is used for rescue mode and reinstallations.
     * 
     */
    public Boolean default_() {
        return this.default_;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The content of the public key.
     * E.g.: &#34;ssh-ed25519 AAAAC3...&#34;
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String keyName() {
        return this.keyName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSshKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean default_;
        private String id;
        private String key;
        private String keyName;
        public Builder() {}
        public Builder(GetSshKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.default_ = defaults.default_;
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.keyName = defaults.keyName;
        }

        @CustomType.Setter("default")
        public Builder default_(Boolean default_) {
            if (default_ == null) {
              throw new MissingRequiredPropertyException("GetSshKeyResult", "default_");
            }
            this.default_ = default_;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSshKeyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetSshKeyResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder keyName(String keyName) {
            if (keyName == null) {
              throw new MissingRequiredPropertyException("GetSshKeyResult", "keyName");
            }
            this.keyName = keyName;
            return this;
        }
        public GetSshKeyResult build() {
            final var _resultValue = new GetSshKeyResult();
            _resultValue.default_ = default_;
            _resultValue.id = id;
            _resultValue.key = key;
            _resultValue.keyName = keyName;
            return _resultValue;
        }
    }
}