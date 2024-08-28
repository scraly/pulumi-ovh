// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInstallationTemplateProjectO {
    /**
     * @return OS template project item governance
     * 
     */
    private List<String> governances;
    /**
     * @return OS template project item name
     * 
     */
    private String name;
    /**
     * @return OS template project item release notes
     * 
     */
    private String releaseNotes;
    /**
     * @return OS template project item url
     * 
     */
    private String url;
    /**
     * @return OS template project item version
     * 
     */
    private String version;

    private GetInstallationTemplateProjectO() {}
    /**
     * @return OS template project item governance
     * 
     */
    public List<String> governances() {
        return this.governances;
    }
    /**
     * @return OS template project item name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return OS template project item release notes
     * 
     */
    public String releaseNotes() {
        return this.releaseNotes;
    }
    /**
     * @return OS template project item url
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return OS template project item version
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstallationTemplateProjectO defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> governances;
        private String name;
        private String releaseNotes;
        private String url;
        private String version;
        public Builder() {}
        public Builder(GetInstallationTemplateProjectO defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.governances = defaults.governances;
    	      this.name = defaults.name;
    	      this.releaseNotes = defaults.releaseNotes;
    	      this.url = defaults.url;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder governances(List<String> governances) {
            if (governances == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplateProjectO", "governances");
            }
            this.governances = governances;
            return this;
        }
        public Builder governances(String... governances) {
            return governances(List.of(governances));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplateProjectO", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder releaseNotes(String releaseNotes) {
            if (releaseNotes == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplateProjectO", "releaseNotes");
            }
            this.releaseNotes = releaseNotes;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplateProjectO", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplateProjectO", "version");
            }
            this.version = version;
            return this;
        }
        public GetInstallationTemplateProjectO build() {
            final var _resultValue = new GetInstallationTemplateProjectO();
            _resultValue.governances = governances;
            _resultValue.name = name;
            _resultValue.releaseNotes = releaseNotes;
            _resultValue.url = url;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
