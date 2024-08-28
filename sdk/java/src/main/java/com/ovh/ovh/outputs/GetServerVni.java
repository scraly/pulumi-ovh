// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetServerVni {
    /**
     * @return VirtualNetworkInterface activation state
     * 
     */
    private Boolean enabled;
    /**
     * @return VirtualNetworkInterface mode (public,vrack,vrack_aggregation)
     * 
     */
    private String mode;
    /**
     * @return User defined VirtualNetworkInterface name
     * 
     */
    private String name;
    /**
     * @return NetworkInterfaceControllers bound to this VirtualNetworkInterface
     * 
     */
    private List<String> nics;
    /**
     * @return Server bound to this VirtualNetworkInterface
     * 
     */
    private String serverName;
    /**
     * @return VirtualNetworkInterface unique id
     * 
     */
    private String uuid;
    /**
     * @return vRack name
     * 
     */
    private String vrack;

    private GetServerVni() {}
    /**
     * @return VirtualNetworkInterface activation state
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return VirtualNetworkInterface mode (public,vrack,vrack_aggregation)
     * 
     */
    public String mode() {
        return this.mode;
    }
    /**
     * @return User defined VirtualNetworkInterface name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return NetworkInterfaceControllers bound to this VirtualNetworkInterface
     * 
     */
    public List<String> nics() {
        return this.nics;
    }
    /**
     * @return Server bound to this VirtualNetworkInterface
     * 
     */
    public String serverName() {
        return this.serverName;
    }
    /**
     * @return VirtualNetworkInterface unique id
     * 
     */
    public String uuid() {
        return this.uuid;
    }
    /**
     * @return vRack name
     * 
     */
    public String vrack() {
        return this.vrack;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerVni defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private String mode;
        private String name;
        private List<String> nics;
        private String serverName;
        private String uuid;
        private String vrack;
        public Builder() {}
        public Builder(GetServerVni defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.nics = defaults.nics;
    	      this.serverName = defaults.serverName;
    	      this.uuid = defaults.uuid;
    	      this.vrack = defaults.vrack;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            if (mode == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "mode");
            }
            this.mode = mode;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nics(List<String> nics) {
            if (nics == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "nics");
            }
            this.nics = nics;
            return this;
        }
        public Builder nics(String... nics) {
            return nics(List.of(nics));
        }
        @CustomType.Setter
        public Builder serverName(String serverName) {
            if (serverName == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "serverName");
            }
            this.serverName = serverName;
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            if (uuid == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "uuid");
            }
            this.uuid = uuid;
            return this;
        }
        @CustomType.Setter
        public Builder vrack(String vrack) {
            if (vrack == null) {
              throw new MissingRequiredPropertyException("GetServerVni", "vrack");
            }
            this.vrack = vrack;
            return this;
        }
        public GetServerVni build() {
            final var _resultValue = new GetServerVni();
            _resultValue.enabled = enabled;
            _resultValue.mode = mode;
            _resultValue.name = name;
            _resultValue.nics = nics;
            _resultValue.serverName = serverName;
            _resultValue.uuid = uuid;
            _resultValue.vrack = vrack;
            return _resultValue;
        }
    }
}
