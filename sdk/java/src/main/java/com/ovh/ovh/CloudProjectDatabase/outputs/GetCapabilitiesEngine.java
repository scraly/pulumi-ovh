// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesEngine {
    /**
     * @return Default version used for the engine.
     * 
     */
    private String defaultVersion;
    /**
     * @return Description of the plan.
     * 
     */
    private String description;
    /**
     * @return Name of the plan.
     * 
     */
    private String name;
    /**
     * @return SSL modes for this engine.
     * 
     */
    private List<String> sslModes;
    /**
     * @return Versions available for this engine.
     * 
     */
    private List<String> versions;

    private GetCapabilitiesEngine() {}
    /**
     * @return Default version used for the engine.
     * 
     */
    public String defaultVersion() {
        return this.defaultVersion;
    }
    /**
     * @return Description of the plan.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Name of the plan.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return SSL modes for this engine.
     * 
     */
    public List<String> sslModes() {
        return this.sslModes;
    }
    /**
     * @return Versions available for this engine.
     * 
     */
    public List<String> versions() {
        return this.versions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesEngine defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String defaultVersion;
        private String description;
        private String name;
        private List<String> sslModes;
        private List<String> versions;
        public Builder() {}
        public Builder(GetCapabilitiesEngine defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultVersion = defaults.defaultVersion;
    	      this.description = defaults.description;
    	      this.name = defaults.name;
    	      this.sslModes = defaults.sslModes;
    	      this.versions = defaults.versions;
        }

        @CustomType.Setter
        public Builder defaultVersion(String defaultVersion) {
            if (defaultVersion == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesEngine", "defaultVersion");
            }
            this.defaultVersion = defaultVersion;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesEngine", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesEngine", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sslModes(List<String> sslModes) {
            if (sslModes == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesEngine", "sslModes");
            }
            this.sslModes = sslModes;
            return this;
        }
        public Builder sslModes(String... sslModes) {
            return sslModes(List.of(sslModes));
        }
        @CustomType.Setter
        public Builder versions(List<String> versions) {
            if (versions == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesEngine", "versions");
            }
            this.versions = versions;
            return this;
        }
        public Builder versions(String... versions) {
            return versions(List.of(versions));
        }
        public GetCapabilitiesEngine build() {
            final var _resultValue = new GetCapabilitiesEngine();
            _resultValue.defaultVersion = defaultVersion;
            _resultValue.description = description;
            _resultValue.name = name;
            _resultValue.sslModes = sslModes;
            _resultValue.versions = versions;
            return _resultValue;
        }
    }
}
