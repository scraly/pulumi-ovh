// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServerDetails {
    /**
     * @return Personnal hostname to use in server reinstallation
     * 
     */
    private @Nullable String customHostname;
    /**
     * @return Disk group id to process install on (only available for some templates)
     * 
     */
    private @Nullable Double diskGroupId;
    /**
     * @return true if you want to install only on the first disk
     * 
     */
    private @Nullable Boolean noRaid;
    /**
     * @return Number of devices to use for system&#39;s software RAID
     * 
     */
    private @Nullable Double softRaidDevices;

    private ServerDetails() {}
    /**
     * @return Personnal hostname to use in server reinstallation
     * 
     */
    public Optional<String> customHostname() {
        return Optional.ofNullable(this.customHostname);
    }
    /**
     * @return Disk group id to process install on (only available for some templates)
     * 
     */
    public Optional<Double> diskGroupId() {
        return Optional.ofNullable(this.diskGroupId);
    }
    /**
     * @return true if you want to install only on the first disk
     * 
     */
    public Optional<Boolean> noRaid() {
        return Optional.ofNullable(this.noRaid);
    }
    /**
     * @return Number of devices to use for system&#39;s software RAID
     * 
     */
    public Optional<Double> softRaidDevices() {
        return Optional.ofNullable(this.softRaidDevices);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerDetails defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String customHostname;
        private @Nullable Double diskGroupId;
        private @Nullable Boolean noRaid;
        private @Nullable Double softRaidDevices;
        public Builder() {}
        public Builder(ServerDetails defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customHostname = defaults.customHostname;
    	      this.diskGroupId = defaults.diskGroupId;
    	      this.noRaid = defaults.noRaid;
    	      this.softRaidDevices = defaults.softRaidDevices;
        }

        @CustomType.Setter
        public Builder customHostname(@Nullable String customHostname) {

            this.customHostname = customHostname;
            return this;
        }
        @CustomType.Setter
        public Builder diskGroupId(@Nullable Double diskGroupId) {

            this.diskGroupId = diskGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder noRaid(@Nullable Boolean noRaid) {

            this.noRaid = noRaid;
            return this;
        }
        @CustomType.Setter
        public Builder softRaidDevices(@Nullable Double softRaidDevices) {

            this.softRaidDevices = softRaidDevices;
            return this;
        }
        public ServerDetails build() {
            final var _resultValue = new ServerDetails();
            _resultValue.customHostname = customHostname;
            _resultValue.diskGroupId = diskGroupId;
            _resultValue.noRaid = noRaid;
            _resultValue.softRaidDevices = softRaidDevices;
            return _resultValue;
        }
    }
}
