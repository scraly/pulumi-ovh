// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsInputConfigurationLogstashArgs extends com.pulumi.resources.ResourceArgs {

    public static final LogsInputConfigurationLogstashArgs Empty = new LogsInputConfigurationLogstashArgs();

    /**
     * The filter section of logstash.conf
     * 
     */
    @Import(name="filterSection")
    private @Nullable Output<String> filterSection;

    /**
     * @return The filter section of logstash.conf
     * 
     */
    public Optional<Output<String>> filterSection() {
        return Optional.ofNullable(this.filterSection);
    }

    /**
     * The filter section of logstash.conf
     * 
     */
    @Import(name="inputSection", required=true)
    private Output<String> inputSection;

    /**
     * @return The filter section of logstash.conf
     * 
     */
    public Output<String> inputSection() {
        return this.inputSection;
    }

    /**
     * The list of customs Grok patterns
     * 
     */
    @Import(name="patternSection")
    private @Nullable Output<String> patternSection;

    /**
     * @return The list of customs Grok patterns
     * 
     */
    public Optional<Output<String>> patternSection() {
        return Optional.ofNullable(this.patternSection);
    }

    private LogsInputConfigurationLogstashArgs() {}

    private LogsInputConfigurationLogstashArgs(LogsInputConfigurationLogstashArgs $) {
        this.filterSection = $.filterSection;
        this.inputSection = $.inputSection;
        this.patternSection = $.patternSection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsInputConfigurationLogstashArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsInputConfigurationLogstashArgs $;

        public Builder() {
            $ = new LogsInputConfigurationLogstashArgs();
        }

        public Builder(LogsInputConfigurationLogstashArgs defaults) {
            $ = new LogsInputConfigurationLogstashArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filterSection The filter section of logstash.conf
         * 
         * @return builder
         * 
         */
        public Builder filterSection(@Nullable Output<String> filterSection) {
            $.filterSection = filterSection;
            return this;
        }

        /**
         * @param filterSection The filter section of logstash.conf
         * 
         * @return builder
         * 
         */
        public Builder filterSection(String filterSection) {
            return filterSection(Output.of(filterSection));
        }

        /**
         * @param inputSection The filter section of logstash.conf
         * 
         * @return builder
         * 
         */
        public Builder inputSection(Output<String> inputSection) {
            $.inputSection = inputSection;
            return this;
        }

        /**
         * @param inputSection The filter section of logstash.conf
         * 
         * @return builder
         * 
         */
        public Builder inputSection(String inputSection) {
            return inputSection(Output.of(inputSection));
        }

        /**
         * @param patternSection The list of customs Grok patterns
         * 
         * @return builder
         * 
         */
        public Builder patternSection(@Nullable Output<String> patternSection) {
            $.patternSection = patternSection;
            return this;
        }

        /**
         * @param patternSection The list of customs Grok patterns
         * 
         * @return builder
         * 
         */
        public Builder patternSection(String patternSection) {
            return patternSection(Output.of(patternSection));
        }

        public LogsInputConfigurationLogstashArgs build() {
            if ($.inputSection == null) {
                throw new MissingRequiredPropertyException("LogsInputConfigurationLogstashArgs", "inputSection");
            }
            return $;
        }
    }

}
