// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUserS3CredentialResult {
    private String accessKeyId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Sensitive) the Secret Access Key
     * 
     */
    private String secretAccessKey;
    private String serviceName;
    private String userId;

    private GetUserS3CredentialResult() {}
    public String accessKeyId() {
        return this.accessKeyId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Sensitive) the Secret Access Key
     * 
     */
    public String secretAccessKey() {
        return this.secretAccessKey;
    }
    public String serviceName() {
        return this.serviceName;
    }
    public String userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserS3CredentialResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessKeyId;
        private String id;
        private String secretAccessKey;
        private String serviceName;
        private String userId;
        public Builder() {}
        public Builder(GetUserS3CredentialResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessKeyId = defaults.accessKeyId;
    	      this.id = defaults.id;
    	      this.secretAccessKey = defaults.secretAccessKey;
    	      this.serviceName = defaults.serviceName;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder accessKeyId(String accessKeyId) {
            if (accessKeyId == null) {
              throw new MissingRequiredPropertyException("GetUserS3CredentialResult", "accessKeyId");
            }
            this.accessKeyId = accessKeyId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUserS3CredentialResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder secretAccessKey(String secretAccessKey) {
            if (secretAccessKey == null) {
              throw new MissingRequiredPropertyException("GetUserS3CredentialResult", "secretAccessKey");
            }
            this.secretAccessKey = secretAccessKey;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetUserS3CredentialResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetUserS3CredentialResult", "userId");
            }
            this.userId = userId;
            return this;
        }
        public GetUserS3CredentialResult build() {
            final var _resultValue = new GetUserS3CredentialResult();
            _resultValue.accessKeyId = accessKeyId;
            _resultValue.id = id;
            _resultValue.secretAccessKey = secretAccessKey;
            _resultValue.serviceName = serviceName;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
