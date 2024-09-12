// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Domain;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ZoneImportArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneImportArgs Empty = new ZoneImportArgs();

    /**
     * Content of the zone file to import
     * 
     */
    @Import(name="zoneFile", required=true)
    private Output<String> zoneFile;

    /**
     * @return Content of the zone file to import
     * 
     */
    public Output<String> zoneFile() {
        return this.zoneFile;
    }

    /**
     * The name of the domain zone
     * 
     */
    @Import(name="zoneName", required=true)
    private Output<String> zoneName;

    /**
     * @return The name of the domain zone
     * 
     */
    public Output<String> zoneName() {
        return this.zoneName;
    }

    private ZoneImportArgs() {}

    private ZoneImportArgs(ZoneImportArgs $) {
        this.zoneFile = $.zoneFile;
        this.zoneName = $.zoneName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneImportArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneImportArgs $;

        public Builder() {
            $ = new ZoneImportArgs();
        }

        public Builder(ZoneImportArgs defaults) {
            $ = new ZoneImportArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param zoneFile Content of the zone file to import
         * 
         * @return builder
         * 
         */
        public Builder zoneFile(Output<String> zoneFile) {
            $.zoneFile = zoneFile;
            return this;
        }

        /**
         * @param zoneFile Content of the zone file to import
         * 
         * @return builder
         * 
         */
        public Builder zoneFile(String zoneFile) {
            return zoneFile(Output.of(zoneFile));
        }

        /**
         * @param zoneName The name of the domain zone
         * 
         * @return builder
         * 
         */
        public Builder zoneName(Output<String> zoneName) {
            $.zoneName = zoneName;
            return this;
        }

        /**
         * @param zoneName The name of the domain zone
         * 
         * @return builder
         * 
         */
        public Builder zoneName(String zoneName) {
            return zoneName(Output.of(zoneName));
        }

        public ZoneImportArgs build() {
            if ($.zoneFile == null) {
                throw new MissingRequiredPropertyException("ZoneImportArgs", "zoneFile");
            }
            if ($.zoneName == null) {
                throw new MissingRequiredPropertyException("ZoneImportArgs", "zoneName");
            }
            return $;
        }
    }

}
