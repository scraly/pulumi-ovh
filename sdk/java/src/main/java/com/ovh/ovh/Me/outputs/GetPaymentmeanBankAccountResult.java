// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPaymentmeanBankAccountResult {
    /**
     * @return a boolean which tells if the retrieved bank account
     * is marked as the default payment mean
     * 
     */
    private Boolean default_;
    /**
     * @return the description attribute of the bank account
     * 
     */
    private String description;
    private @Nullable String descriptionRegexp;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String state;
    private @Nullable Boolean useDefault;
    private @Nullable Boolean useOldest;

    private GetPaymentmeanBankAccountResult() {}
    /**
     * @return a boolean which tells if the retrieved bank account
     * is marked as the default payment mean
     * 
     */
    public Boolean default_() {
        return this.default_;
    }
    /**
     * @return the description attribute of the bank account
     * 
     */
    public String description() {
        return this.description;
    }
    public Optional<String> descriptionRegexp() {
        return Optional.ofNullable(this.descriptionRegexp);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String state() {
        return this.state;
    }
    public Optional<Boolean> useDefault() {
        return Optional.ofNullable(this.useDefault);
    }
    public Optional<Boolean> useOldest() {
        return Optional.ofNullable(this.useOldest);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPaymentmeanBankAccountResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean default_;
        private String description;
        private @Nullable String descriptionRegexp;
        private String id;
        private String state;
        private @Nullable Boolean useDefault;
        private @Nullable Boolean useOldest;
        public Builder() {}
        public Builder(GetPaymentmeanBankAccountResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.default_ = defaults.default_;
    	      this.description = defaults.description;
    	      this.descriptionRegexp = defaults.descriptionRegexp;
    	      this.id = defaults.id;
    	      this.state = defaults.state;
    	      this.useDefault = defaults.useDefault;
    	      this.useOldest = defaults.useOldest;
        }

        @CustomType.Setter("default")
        public Builder default_(Boolean default_) {
            if (default_ == null) {
              throw new MissingRequiredPropertyException("GetPaymentmeanBankAccountResult", "default_");
            }
            this.default_ = default_;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetPaymentmeanBankAccountResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder descriptionRegexp(@Nullable String descriptionRegexp) {

            this.descriptionRegexp = descriptionRegexp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetPaymentmeanBankAccountResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetPaymentmeanBankAccountResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder useDefault(@Nullable Boolean useDefault) {

            this.useDefault = useDefault;
            return this;
        }
        @CustomType.Setter
        public Builder useOldest(@Nullable Boolean useOldest) {

            this.useOldest = useOldest;
            return this;
        }
        public GetPaymentmeanBankAccountResult build() {
            final var _resultValue = new GetPaymentmeanBankAccountResult();
            _resultValue.default_ = default_;
            _resultValue.description = description;
            _resultValue.descriptionRegexp = descriptionRegexp;
            _resultValue.id = id;
            _resultValue.state = state;
            _resultValue.useDefault = useDefault;
            _resultValue.useOldest = useOldest;
            return _resultValue;
        }
    }
}
