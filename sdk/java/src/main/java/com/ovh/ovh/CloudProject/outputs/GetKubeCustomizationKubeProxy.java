// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.ovh.ovh.CloudProject.outputs.GetKubeCustomizationKubeProxyIptables;
import com.ovh.ovh.CloudProject.outputs.GetKubeCustomizationKubeProxyIpvs;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKubeCustomizationKubeProxy {
    /**
     * @return Kubernetes cluster kube-proxy customization of iptables specific config.
     * 
     */
    private @Nullable GetKubeCustomizationKubeProxyIptables iptables;
    /**
     * @return Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
     * 
     */
    private @Nullable GetKubeCustomizationKubeProxyIpvs ipvs;

    private GetKubeCustomizationKubeProxy() {}
    /**
     * @return Kubernetes cluster kube-proxy customization of iptables specific config.
     * 
     */
    public Optional<GetKubeCustomizationKubeProxyIptables> iptables() {
        return Optional.ofNullable(this.iptables);
    }
    /**
     * @return Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
     * 
     */
    public Optional<GetKubeCustomizationKubeProxyIpvs> ipvs() {
        return Optional.ofNullable(this.ipvs);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeCustomizationKubeProxy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable GetKubeCustomizationKubeProxyIptables iptables;
        private @Nullable GetKubeCustomizationKubeProxyIpvs ipvs;
        public Builder() {}
        public Builder(GetKubeCustomizationKubeProxy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.iptables = defaults.iptables;
    	      this.ipvs = defaults.ipvs;
        }

        @CustomType.Setter
        public Builder iptables(@Nullable GetKubeCustomizationKubeProxyIptables iptables) {

            this.iptables = iptables;
            return this;
        }
        @CustomType.Setter
        public Builder ipvs(@Nullable GetKubeCustomizationKubeProxyIpvs ipvs) {

            this.ipvs = ipvs;
            return this;
        }
        public GetKubeCustomizationKubeProxy build() {
            final var _resultValue = new GetKubeCustomizationKubeProxy();
            _resultValue.iptables = iptables;
            _resultValue.ipvs = ipvs;
            return _resultValue;
        }
    }
}
