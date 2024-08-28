// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Hosting.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetPrivateDatabasePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivateDatabasePlainArgs Empty = new GetPrivateDatabasePlainArgs();

    /**
     * The internal name of your private database
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The internal name of your private database
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetPrivateDatabasePlainArgs() {}

    private GetPrivateDatabasePlainArgs(GetPrivateDatabasePlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivateDatabasePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivateDatabasePlainArgs $;

        public Builder() {
            $ = new GetPrivateDatabasePlainArgs();
        }

        public Builder(GetPrivateDatabasePlainArgs defaults) {
            $ = new GetPrivateDatabasePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The internal name of your private database
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetPrivateDatabasePlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetPrivateDatabasePlainArgs", "serviceName");
            }
            return $;
        }
    }

}
