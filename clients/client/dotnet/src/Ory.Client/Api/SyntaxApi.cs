/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.3.0
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net;
using System.Net.Mime;
using Ory.Client.Client;
using Ory.Client.Model;

namespace Ory.Client.Api
{

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface ISyntaxApiSync : IApiAccessor
    {
        #region Synchronous Operations
        /// <summary>
        /// Check the syntax of an OPL file
        /// </summary>
        /// <remarks>
        /// The OPL file is expected in the body of the request.
        /// </remarks>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <returns>ClientPostCheckOplSyntaxResponse</returns>
        ClientPostCheckOplSyntaxResponse CheckOplSyntax(string body);

        /// <summary>
        /// Check the syntax of an OPL file
        /// </summary>
        /// <remarks>
        /// The OPL file is expected in the body of the request.
        /// </remarks>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <returns>ApiResponse of ClientPostCheckOplSyntaxResponse</returns>
        ApiResponse<ClientPostCheckOplSyntaxResponse> CheckOplSyntaxWithHttpInfo(string body);
        #endregion Synchronous Operations
    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface ISyntaxApiAsync : IApiAccessor
    {
        #region Asynchronous Operations
        /// <summary>
        /// Check the syntax of an OPL file
        /// </summary>
        /// <remarks>
        /// The OPL file is expected in the body of the request.
        /// </remarks>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <param name="cancellationToken">Cancellation Token to cancel the request.</param>
        /// <returns>Task of ClientPostCheckOplSyntaxResponse</returns>
        System.Threading.Tasks.Task<ClientPostCheckOplSyntaxResponse> CheckOplSyntaxAsync(string body, System.Threading.CancellationToken cancellationToken = default(System.Threading.CancellationToken));

        /// <summary>
        /// Check the syntax of an OPL file
        /// </summary>
        /// <remarks>
        /// The OPL file is expected in the body of the request.
        /// </remarks>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <param name="cancellationToken">Cancellation Token to cancel the request.</param>
        /// <returns>Task of ApiResponse (ClientPostCheckOplSyntaxResponse)</returns>
        System.Threading.Tasks.Task<ApiResponse<ClientPostCheckOplSyntaxResponse>> CheckOplSyntaxWithHttpInfoAsync(string body, System.Threading.CancellationToken cancellationToken = default(System.Threading.CancellationToken));
        #endregion Asynchronous Operations
    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface ISyntaxApi : ISyntaxApiSync, ISyntaxApiAsync
    {

    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public partial class SyntaxApi : ISyntaxApi
    {
        private Ory.Client.Client.ExceptionFactory _exceptionFactory = (name, response) => null;

        /// <summary>
        /// Initializes a new instance of the <see cref="SyntaxApi"/> class.
        /// </summary>
        /// <returns></returns>
        public SyntaxApi() : this((string)null)
        {
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SyntaxApi"/> class.
        /// </summary>
        /// <returns></returns>
        public SyntaxApi(string basePath)
        {
            this.Configuration = Ory.Client.Client.Configuration.MergeConfigurations(
                Ory.Client.Client.GlobalConfiguration.Instance,
                new Ory.Client.Client.Configuration { BasePath = basePath }
            );
            this.Client = new Ory.Client.Client.ApiClient(this.Configuration.BasePath);
            this.AsynchronousClient = new Ory.Client.Client.ApiClient(this.Configuration.BasePath);
            this.ExceptionFactory = Ory.Client.Client.Configuration.DefaultExceptionFactory;
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SyntaxApi"/> class
        /// using Configuration object
        /// </summary>
        /// <param name="configuration">An instance of Configuration</param>
        /// <returns></returns>
        public SyntaxApi(Ory.Client.Client.Configuration configuration)
        {
            if (configuration == null) throw new ArgumentNullException("configuration");

            this.Configuration = Ory.Client.Client.Configuration.MergeConfigurations(
                Ory.Client.Client.GlobalConfiguration.Instance,
                configuration
            );
            this.Client = new Ory.Client.Client.ApiClient(this.Configuration.BasePath);
            this.AsynchronousClient = new Ory.Client.Client.ApiClient(this.Configuration.BasePath);
            ExceptionFactory = Ory.Client.Client.Configuration.DefaultExceptionFactory;
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="SyntaxApi"/> class
        /// using a Configuration object and client instance.
        /// </summary>
        /// <param name="client">The client interface for synchronous API access.</param>
        /// <param name="asyncClient">The client interface for asynchronous API access.</param>
        /// <param name="configuration">The configuration object.</param>
        public SyntaxApi(Ory.Client.Client.ISynchronousClient client, Ory.Client.Client.IAsynchronousClient asyncClient, Ory.Client.Client.IReadableConfiguration configuration)
        {
            if (client == null) throw new ArgumentNullException("client");
            if (asyncClient == null) throw new ArgumentNullException("asyncClient");
            if (configuration == null) throw new ArgumentNullException("configuration");

            this.Client = client;
            this.AsynchronousClient = asyncClient;
            this.Configuration = configuration;
            this.ExceptionFactory = Ory.Client.Client.Configuration.DefaultExceptionFactory;
        }

        /// <summary>
        /// The client for accessing this underlying API asynchronously.
        /// </summary>
        public Ory.Client.Client.IAsynchronousClient AsynchronousClient { get; set; }

        /// <summary>
        /// The client for accessing this underlying API synchronously.
        /// </summary>
        public Ory.Client.Client.ISynchronousClient Client { get; set; }

        /// <summary>
        /// Gets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        public string GetBasePath()
        {
            return this.Configuration.BasePath;
        }

        /// <summary>
        /// Gets or sets the configuration object
        /// </summary>
        /// <value>An instance of the Configuration</value>
        public Ory.Client.Client.IReadableConfiguration Configuration { get; set; }

        /// <summary>
        /// Provides a factory method hook for the creation of exceptions.
        /// </summary>
        public Ory.Client.Client.ExceptionFactory ExceptionFactory
        {
            get
            {
                if (_exceptionFactory != null && _exceptionFactory.GetInvocationList().Length > 1)
                {
                    throw new InvalidOperationException("Multicast delegate for ExceptionFactory is unsupported.");
                }
                return _exceptionFactory;
            }
            set { _exceptionFactory = value; }
        }

        /// <summary>
        /// Check the syntax of an OPL file The OPL file is expected in the body of the request.
        /// </summary>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <returns>ClientPostCheckOplSyntaxResponse</returns>
        public ClientPostCheckOplSyntaxResponse CheckOplSyntax(string body)
        {
            Ory.Client.Client.ApiResponse<ClientPostCheckOplSyntaxResponse> localVarResponse = CheckOplSyntaxWithHttpInfo(body);
            return localVarResponse.Data;
        }

        /// <summary>
        /// Check the syntax of an OPL file The OPL file is expected in the body of the request.
        /// </summary>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <returns>ApiResponse of ClientPostCheckOplSyntaxResponse</returns>
        public Ory.Client.Client.ApiResponse<ClientPostCheckOplSyntaxResponse> CheckOplSyntaxWithHttpInfo(string body)
        {
            // verify the required parameter 'body' is set
            if (body == null)
            {
                throw new Ory.Client.Client.ApiException(400, "Missing required parameter 'body' when calling SyntaxApi->CheckOplSyntax");
            }

            Ory.Client.Client.RequestOptions localVarRequestOptions = new Ory.Client.Client.RequestOptions();

            string[] _contentTypes = new string[] {
                "text/plain"
            };

            // to determine the Accept header
            string[] _accepts = new string[] {
                "application/json"
            };

            var localVarContentType = Ory.Client.Client.ClientUtils.SelectHeaderContentType(_contentTypes);
            if (localVarContentType != null)
            {
                localVarRequestOptions.HeaderParameters.Add("Content-Type", localVarContentType);
            }

            var localVarAccept = Ory.Client.Client.ClientUtils.SelectHeaderAccept(_accepts);
            if (localVarAccept != null)
            {
                localVarRequestOptions.HeaderParameters.Add("Accept", localVarAccept);
            }

            localVarRequestOptions.Data = body;

            // authentication (oryAccessToken) required
            // bearer authentication required
            if (!string.IsNullOrEmpty(this.Configuration.AccessToken) && !localVarRequestOptions.HeaderParameters.ContainsKey("Authorization"))
            {
                localVarRequestOptions.HeaderParameters.Add("Authorization", "Bearer " + this.Configuration.AccessToken);
            }

            // make the HTTP request
            var localVarResponse = this.Client.Post<ClientPostCheckOplSyntaxResponse>("/opl/syntax/check", localVarRequestOptions, this.Configuration);
            if (this.ExceptionFactory != null)
            {
                Exception _exception = this.ExceptionFactory("CheckOplSyntax", localVarResponse);
                if (_exception != null)
                {
                    throw _exception;
                }
            }

            return localVarResponse;
        }

        /// <summary>
        /// Check the syntax of an OPL file The OPL file is expected in the body of the request.
        /// </summary>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <param name="cancellationToken">Cancellation Token to cancel the request.</param>
        /// <returns>Task of ClientPostCheckOplSyntaxResponse</returns>
        public async System.Threading.Tasks.Task<ClientPostCheckOplSyntaxResponse> CheckOplSyntaxAsync(string body, System.Threading.CancellationToken cancellationToken = default(System.Threading.CancellationToken))
        {
            Ory.Client.Client.ApiResponse<ClientPostCheckOplSyntaxResponse> localVarResponse = await CheckOplSyntaxWithHttpInfoAsync(body, cancellationToken).ConfigureAwait(false);
            return localVarResponse.Data;
        }

        /// <summary>
        /// Check the syntax of an OPL file The OPL file is expected in the body of the request.
        /// </summary>
        /// <exception cref="Ory.Client.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="body">the OPL content to check</param>
        /// <param name="cancellationToken">Cancellation Token to cancel the request.</param>
        /// <returns>Task of ApiResponse (ClientPostCheckOplSyntaxResponse)</returns>
        public async System.Threading.Tasks.Task<Ory.Client.Client.ApiResponse<ClientPostCheckOplSyntaxResponse>> CheckOplSyntaxWithHttpInfoAsync(string body, System.Threading.CancellationToken cancellationToken = default(System.Threading.CancellationToken))
        {
            // verify the required parameter 'body' is set
            if (body == null)
            {
                throw new Ory.Client.Client.ApiException(400, "Missing required parameter 'body' when calling SyntaxApi->CheckOplSyntax");
            }


            Ory.Client.Client.RequestOptions localVarRequestOptions = new Ory.Client.Client.RequestOptions();

            string[] _contentTypes = new string[] {
                "text/plain"
            };

            // to determine the Accept header
            string[] _accepts = new string[] {
                "application/json"
            };

            var localVarContentType = Ory.Client.Client.ClientUtils.SelectHeaderContentType(_contentTypes);
            if (localVarContentType != null)
            {
                localVarRequestOptions.HeaderParameters.Add("Content-Type", localVarContentType);
            }

            var localVarAccept = Ory.Client.Client.ClientUtils.SelectHeaderAccept(_accepts);
            if (localVarAccept != null)
            {
                localVarRequestOptions.HeaderParameters.Add("Accept", localVarAccept);
            }

            localVarRequestOptions.Data = body;

            // authentication (oryAccessToken) required
            // bearer authentication required
            if (!string.IsNullOrEmpty(this.Configuration.AccessToken) && !localVarRequestOptions.HeaderParameters.ContainsKey("Authorization"))
            {
                localVarRequestOptions.HeaderParameters.Add("Authorization", "Bearer " + this.Configuration.AccessToken);
            }

            // make the HTTP request
            var localVarResponse = await this.AsynchronousClient.PostAsync<ClientPostCheckOplSyntaxResponse>("/opl/syntax/check", localVarRequestOptions, this.Configuration, cancellationToken).ConfigureAwait(false);

            if (this.ExceptionFactory != null)
            {
                Exception _exception = this.ExceptionFactory("CheckOplSyntax", localVarResponse);
                if (_exception != null)
                {
                    throw _exception;
                }
            }

            return localVarResponse;
        }

    }
}
