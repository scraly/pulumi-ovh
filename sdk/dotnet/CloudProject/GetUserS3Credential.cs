// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Ovh.CloudProject
{
    public static class GetUserS3Credential
    {
        /// <summary>
        /// Use this data source to retrieve the Secret Access Key of an Access Key ID associated with a public cloud project's user.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_users" "project_users" {
        ///   service_name = "XXX"
        /// }
        /// 
        /// locals {
        ///   # Get the user ID of a previously created user with the description "S3-User"
        ///   users      = [for user in data.ovh_cloud_project_users.project_users.users : user.user_id if user.description == "S3-User"]
        ///   s3_user_id = local.users[0]
        /// }
        /// 
        /// data "ovh_cloud_project_user_s3_credentials" "my_s3_credentials" {
        ///   service_name = data.ovh_cloud_project_users.project_users.service_name
        ///   user_id      = local.s3_user_id
        /// }
        /// 
        /// data "ovh_cloud_project_user_s3_credential" "my_s3_credential" {
        ///   service_name  = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.service_name
        ///   user_id       = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.user_id
        ///   access_key_id = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.access_key_ids[0]
        /// }
        /// 
        /// output "my_access_key_id" {
        ///   value = data.ovh_cloud_project_user_s3_credential.my_s3_credential.access_key_id
        /// }
        /// 
        /// output "my_secret_access_key" {
        ///   value     = data.ovh_cloud_project_user_s3_credential.my_s3_credential.secret_access_key
        ///   sensitive = true
        /// }
        /// ```
        /// </summary>
        public static Task<GetUserS3CredentialResult> InvokeAsync(GetUserS3CredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserS3CredentialResult>("ovh:CloudProject/getUserS3Credential:getUserS3Credential", args ?? new GetUserS3CredentialArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the Secret Access Key of an Access Key ID associated with a public cloud project's user.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_users" "project_users" {
        ///   service_name = "XXX"
        /// }
        /// 
        /// locals {
        ///   # Get the user ID of a previously created user with the description "S3-User"
        ///   users      = [for user in data.ovh_cloud_project_users.project_users.users : user.user_id if user.description == "S3-User"]
        ///   s3_user_id = local.users[0]
        /// }
        /// 
        /// data "ovh_cloud_project_user_s3_credentials" "my_s3_credentials" {
        ///   service_name = data.ovh_cloud_project_users.project_users.service_name
        ///   user_id      = local.s3_user_id
        /// }
        /// 
        /// data "ovh_cloud_project_user_s3_credential" "my_s3_credential" {
        ///   service_name  = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.service_name
        ///   user_id       = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.user_id
        ///   access_key_id = data.ovh_cloud_project_user_s3_credentials.my_s3_credentials.access_key_ids[0]
        /// }
        /// 
        /// output "my_access_key_id" {
        ///   value = data.ovh_cloud_project_user_s3_credential.my_s3_credential.access_key_id
        /// }
        /// 
        /// output "my_secret_access_key" {
        ///   value     = data.ovh_cloud_project_user_s3_credential.my_s3_credential.secret_access_key
        ///   sensitive = true
        /// }
        /// ```
        /// </summary>
        public static Output<GetUserS3CredentialResult> Invoke(GetUserS3CredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserS3CredentialResult>("ovh:CloudProject/getUserS3Credential:getUserS3Credential", args ?? new GetUserS3CredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserS3CredentialArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// the Access Key ID
        /// </summary>
        [Input("accessKeyId", required: true)]
        public string AccessKeyId { get; set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUserS3CredentialArgs()
        {
        }
        public static new GetUserS3CredentialArgs Empty => new GetUserS3CredentialArgs();
    }

    public sealed class GetUserS3CredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// the Access Key ID
        /// </summary>
        [Input("accessKeyId", required: true)]
        public Input<string> AccessKeyId { get; set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetUserS3CredentialInvokeArgs()
        {
        }
        public static new GetUserS3CredentialInvokeArgs Empty => new GetUserS3CredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserS3CredentialResult
    {
        public readonly string AccessKeyId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Sensitive) the Secret Access Key
        /// </summary>
        public readonly string SecretAccessKey;
        public readonly string ServiceName;
        public readonly string UserId;

        [OutputConstructor]
        private GetUserS3CredentialResult(
            string accessKeyId,

            string id,

            string secretAccessKey,

            string serviceName,

            string userId)
        {
            AccessKeyId = accessKeyId;
            Id = id;
            SecretAccessKey = secretAccessKey;
            ServiceName = serviceName;
            UserId = userId;
        }
    }
}