/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.60
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientOAuth2AccessRequest
    /// </summary>
    [DataContract(Name = "oAuth2AccessRequest")]
    public partial class ClientOAuth2AccessRequest : IEquatable<ClientOAuth2AccessRequest>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientOAuth2AccessRequest" /> class.
        /// </summary>
        /// <param name="clientId">ClientID is the identifier of the OAuth 2.0 client..</param>
        /// <param name="grantTypes">GrantTypes is the requests grant types..</param>
        /// <param name="grantedAudience">GrantedAudience is the list of audiences granted to the OAuth 2.0 client..</param>
        /// <param name="grantedScopes">GrantedScopes is the list of scopes granted to the OAuth 2.0 client..</param>
        public ClientOAuth2AccessRequest(string clientId = default(string), List<string> grantTypes = default(List<string>), List<string> grantedAudience = default(List<string>), List<string> grantedScopes = default(List<string>))
        {
            this.ClientId = clientId;
            this.GrantTypes = grantTypes;
            this.GrantedAudience = grantedAudience;
            this.GrantedScopes = grantedScopes;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// ClientID is the identifier of the OAuth 2.0 client.
        /// </summary>
        /// <value>ClientID is the identifier of the OAuth 2.0 client.</value>
        [DataMember(Name = "client_id", EmitDefaultValue = false)]
        public string ClientId { get; set; }

        /// <summary>
        /// GrantTypes is the requests grant types.
        /// </summary>
        /// <value>GrantTypes is the requests grant types.</value>
        [DataMember(Name = "grant_types", EmitDefaultValue = false)]
        public List<string> GrantTypes { get; set; }

        /// <summary>
        /// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
        /// </summary>
        /// <value>GrantedAudience is the list of audiences granted to the OAuth 2.0 client.</value>
        [DataMember(Name = "granted_audience", EmitDefaultValue = false)]
        public List<string> GrantedAudience { get; set; }

        /// <summary>
        /// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
        /// </summary>
        /// <value>GrantedScopes is the list of scopes granted to the OAuth 2.0 client.</value>
        [DataMember(Name = "granted_scopes", EmitDefaultValue = false)]
        public List<string> GrantedScopes { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientOAuth2AccessRequest {\n");
            sb.Append("  ClientId: ").Append(ClientId).Append("\n");
            sb.Append("  GrantTypes: ").Append(GrantTypes).Append("\n");
            sb.Append("  GrantedAudience: ").Append(GrantedAudience).Append("\n");
            sb.Append("  GrantedScopes: ").Append(GrantedScopes).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientOAuth2AccessRequest);
        }

        /// <summary>
        /// Returns true if ClientOAuth2AccessRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientOAuth2AccessRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientOAuth2AccessRequest input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.ClientId == input.ClientId ||
                    (this.ClientId != null &&
                    this.ClientId.Equals(input.ClientId))
                ) && 
                (
                    this.GrantTypes == input.GrantTypes ||
                    this.GrantTypes != null &&
                    input.GrantTypes != null &&
                    this.GrantTypes.SequenceEqual(input.GrantTypes)
                ) && 
                (
                    this.GrantedAudience == input.GrantedAudience ||
                    this.GrantedAudience != null &&
                    input.GrantedAudience != null &&
                    this.GrantedAudience.SequenceEqual(input.GrantedAudience)
                ) && 
                (
                    this.GrantedScopes == input.GrantedScopes ||
                    this.GrantedScopes != null &&
                    input.GrantedScopes != null &&
                    this.GrantedScopes.SequenceEqual(input.GrantedScopes)
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.ClientId != null)
                {
                    hashCode = (hashCode * 59) + this.ClientId.GetHashCode();
                }
                if (this.GrantTypes != null)
                {
                    hashCode = (hashCode * 59) + this.GrantTypes.GetHashCode();
                }
                if (this.GrantedAudience != null)
                {
                    hashCode = (hashCode * 59) + this.GrantedAudience.GetHashCode();
                }
                if (this.GrantedScopes != null)
                {
                    hashCode = (hashCode * 59) + this.GrantedScopes.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
