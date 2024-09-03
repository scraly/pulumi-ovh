// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.ovh.ovh.CloudProject.outputs.GetLoadBalancerFloatingIp;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLoadBalancerResult {
    /**
     * @return Date of creation of the loadbalancer
     * 
     */
    private String createdAt;
    /**
     * @return ID of the flavor
     * 
     */
    private String flavorId;
    /**
     * @return Information about the floating IP
     * 
     */
    private GetLoadBalancerFloatingIp floatingIp;
    /**
     * @return ID of the floating IP
     * 
     */
    private String id;
    /**
     * @return Name of the loadbalancer
     * 
     */
    private String name;
    /**
     * @return Operating status of the loadbalancer
     * 
     */
    private String operatingStatus;
    /**
     * @return Provisioning status of the loadbalancer
     * 
     */
    private String provisioningStatus;
    /**
     * @return Region of the loadbalancer
     * 
     */
    private String regionName;
    /**
     * @return ID of the public cloud project
     * 
     */
    private String serviceName;
    /**
     * @return Last update date of the loadbalancer
     * 
     */
    private String updatedAt;
    /**
     * @return IP address of the Virtual IP
     * 
     */
    private String vipAddress;
    /**
     * @return Openstack ID of the network for the Virtual IP
     * 
     */
    private String vipNetworkId;
    /**
     * @return ID of the subnet for the Virtual IP
     * 
     */
    private String vipSubnetId;

    private GetLoadBalancerResult() {}
    /**
     * @return Date of creation of the loadbalancer
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return ID of the flavor
     * 
     */
    public String flavorId() {
        return this.flavorId;
    }
    /**
     * @return Information about the floating IP
     * 
     */
    public GetLoadBalancerFloatingIp floatingIp() {
        return this.floatingIp;
    }
    /**
     * @return ID of the floating IP
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the loadbalancer
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Operating status of the loadbalancer
     * 
     */
    public String operatingStatus() {
        return this.operatingStatus;
    }
    /**
     * @return Provisioning status of the loadbalancer
     * 
     */
    public String provisioningStatus() {
        return this.provisioningStatus;
    }
    /**
     * @return Region of the loadbalancer
     * 
     */
    public String regionName() {
        return this.regionName;
    }
    /**
     * @return ID of the public cloud project
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Last update date of the loadbalancer
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return IP address of the Virtual IP
     * 
     */
    public String vipAddress() {
        return this.vipAddress;
    }
    /**
     * @return Openstack ID of the network for the Virtual IP
     * 
     */
    public String vipNetworkId() {
        return this.vipNetworkId;
    }
    /**
     * @return ID of the subnet for the Virtual IP
     * 
     */
    public String vipSubnetId() {
        return this.vipSubnetId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLoadBalancerResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String flavorId;
        private GetLoadBalancerFloatingIp floatingIp;
        private String id;
        private String name;
        private String operatingStatus;
        private String provisioningStatus;
        private String regionName;
        private String serviceName;
        private String updatedAt;
        private String vipAddress;
        private String vipNetworkId;
        private String vipSubnetId;
        public Builder() {}
        public Builder(GetLoadBalancerResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.flavorId = defaults.flavorId;
    	      this.floatingIp = defaults.floatingIp;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.operatingStatus = defaults.operatingStatus;
    	      this.provisioningStatus = defaults.provisioningStatus;
    	      this.regionName = defaults.regionName;
    	      this.serviceName = defaults.serviceName;
    	      this.updatedAt = defaults.updatedAt;
    	      this.vipAddress = defaults.vipAddress;
    	      this.vipNetworkId = defaults.vipNetworkId;
    	      this.vipSubnetId = defaults.vipSubnetId;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder flavorId(String flavorId) {
            if (flavorId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "flavorId");
            }
            this.flavorId = flavorId;
            return this;
        }
        @CustomType.Setter
        public Builder floatingIp(GetLoadBalancerFloatingIp floatingIp) {
            if (floatingIp == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "floatingIp");
            }
            this.floatingIp = floatingIp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder operatingStatus(String operatingStatus) {
            if (operatingStatus == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "operatingStatus");
            }
            this.operatingStatus = operatingStatus;
            return this;
        }
        @CustomType.Setter
        public Builder provisioningStatus(String provisioningStatus) {
            if (provisioningStatus == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "provisioningStatus");
            }
            this.provisioningStatus = provisioningStatus;
            return this;
        }
        @CustomType.Setter
        public Builder regionName(String regionName) {
            if (regionName == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "regionName");
            }
            this.regionName = regionName;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder vipAddress(String vipAddress) {
            if (vipAddress == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "vipAddress");
            }
            this.vipAddress = vipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder vipNetworkId(String vipNetworkId) {
            if (vipNetworkId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "vipNetworkId");
            }
            this.vipNetworkId = vipNetworkId;
            return this;
        }
        @CustomType.Setter
        public Builder vipSubnetId(String vipSubnetId) {
            if (vipSubnetId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancerResult", "vipSubnetId");
            }
            this.vipSubnetId = vipSubnetId;
            return this;
        }
        public GetLoadBalancerResult build() {
            final var _resultValue = new GetLoadBalancerResult();
            _resultValue.createdAt = createdAt;
            _resultValue.flavorId = flavorId;
            _resultValue.floatingIp = floatingIp;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.operatingStatus = operatingStatus;
            _resultValue.provisioningStatus = provisioningStatus;
            _resultValue.regionName = regionName;
            _resultValue.serviceName = serviceName;
            _resultValue.updatedAt = updatedAt;
            _resultValue.vipAddress = vipAddress;
            _resultValue.vipNetworkId = vipNetworkId;
            _resultValue.vipSubnetId = vipSubnetId;
            return _resultValue;
        }
    }
}