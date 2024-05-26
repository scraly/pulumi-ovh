// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Me.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetIpxeScriptArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpxeScriptArgs Empty = new GetIpxeScriptArgs();

    /**
     * The name of the IPXE Script.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the IPXE Script.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetIpxeScriptArgs() {}

    private GetIpxeScriptArgs(GetIpxeScriptArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpxeScriptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpxeScriptArgs $;

        public Builder() {
            $ = new GetIpxeScriptArgs();
        }

        public Builder(GetIpxeScriptArgs defaults) {
            $ = new GetIpxeScriptArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the IPXE Script.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the IPXE Script.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetIpxeScriptArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetIpxeScriptArgs", "name");
            }
            return $;
        }
    }

}