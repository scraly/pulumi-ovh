// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Iam.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetReferenceActionsAction {
    /**
     * @return Name of the action
     * 
     */
    private String action;
    /**
     * @return List of the categories of the action
     * 
     */
    private List<String> categories;
    /**
     * @return Description of the action
     * 
     */
    private String description;
    /**
     * @return Resource type the action is related to
     * 
     */
    private String resourceType;

    private GetReferenceActionsAction() {}
    /**
     * @return Name of the action
     * 
     */
    public String action() {
        return this.action;
    }
    /**
     * @return List of the categories of the action
     * 
     */
    public List<String> categories() {
        return this.categories;
    }
    /**
     * @return Description of the action
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Resource type the action is related to
     * 
     */
    public String resourceType() {
        return this.resourceType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetReferenceActionsAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String action;
        private List<String> categories;
        private String description;
        private String resourceType;
        public Builder() {}
        public Builder(GetReferenceActionsAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.categories = defaults.categories;
    	      this.description = defaults.description;
    	      this.resourceType = defaults.resourceType;
        }

        @CustomType.Setter
        public Builder action(String action) {
            if (action == null) {
              throw new MissingRequiredPropertyException("GetReferenceActionsAction", "action");
            }
            this.action = action;
            return this;
        }
        @CustomType.Setter
        public Builder categories(List<String> categories) {
            if (categories == null) {
              throw new MissingRequiredPropertyException("GetReferenceActionsAction", "categories");
            }
            this.categories = categories;
            return this;
        }
        public Builder categories(String... categories) {
            return categories(List.of(categories));
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetReferenceActionsAction", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder resourceType(String resourceType) {
            if (resourceType == null) {
              throw new MissingRequiredPropertyException("GetReferenceActionsAction", "resourceType");
            }
            this.resourceType = resourceType;
            return this;
        }
        public GetReferenceActionsAction build() {
            final var _resultValue = new GetReferenceActionsAction();
            _resultValue.action = action;
            _resultValue.categories = categories;
            _resultValue.description = description;
            _resultValue.resourceType = resourceType;
            return _resultValue;
        }
    }
}