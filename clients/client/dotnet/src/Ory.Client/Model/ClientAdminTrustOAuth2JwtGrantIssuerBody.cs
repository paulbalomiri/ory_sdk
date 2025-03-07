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
    /// ClientAdminTrustOAuth2JwtGrantIssuerBody
    /// </summary>
    [DataContract(Name = "adminTrustOAuth2JwtGrantIssuerBody")]
    public partial class ClientAdminTrustOAuth2JwtGrantIssuerBody : IEquatable<ClientAdminTrustOAuth2JwtGrantIssuerBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientAdminTrustOAuth2JwtGrantIssuerBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientAdminTrustOAuth2JwtGrantIssuerBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientAdminTrustOAuth2JwtGrantIssuerBody" /> class.
        /// </summary>
        /// <param name="allowAnySubject">The \&quot;allow_any_subject\&quot; indicates that the issuer is allowed to have any principal as the subject of the JWT..</param>
        /// <param name="expiresAt">The \&quot;expires_at\&quot; indicates, when grant will expire, so we will reject assertion from \&quot;issuer\&quot; targeting \&quot;subject\&quot;. (required).</param>
        /// <param name="issuer">The \&quot;issuer\&quot; identifies the principal that issued the JWT assertion (same as \&quot;iss\&quot; claim in JWT). (required).</param>
        /// <param name="jwk">jwk (required).</param>
        /// <param name="scope">The \&quot;scope\&quot; contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) (required).</param>
        /// <param name="subject">The \&quot;subject\&quot; identifies the principal that is the subject of the JWT..</param>
        public ClientAdminTrustOAuth2JwtGrantIssuerBody(bool allowAnySubject = default(bool), DateTime expiresAt = default(DateTime), string issuer = default(string), ClientJsonWebKey jwk = default(ClientJsonWebKey), List<string> scope = default(List<string>), string subject = default(string))
        {
            this.ExpiresAt = expiresAt;
            // to ensure "issuer" is required (not null)
            if (issuer == null) {
                throw new ArgumentNullException("issuer is a required property for ClientAdminTrustOAuth2JwtGrantIssuerBody and cannot be null");
            }
            this.Issuer = issuer;
            // to ensure "jwk" is required (not null)
            if (jwk == null) {
                throw new ArgumentNullException("jwk is a required property for ClientAdminTrustOAuth2JwtGrantIssuerBody and cannot be null");
            }
            this.Jwk = jwk;
            // to ensure "scope" is required (not null)
            if (scope == null) {
                throw new ArgumentNullException("scope is a required property for ClientAdminTrustOAuth2JwtGrantIssuerBody and cannot be null");
            }
            this.Scope = scope;
            this.AllowAnySubject = allowAnySubject;
            this.Subject = subject;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The \&quot;allow_any_subject\&quot; indicates that the issuer is allowed to have any principal as the subject of the JWT.
        /// </summary>
        /// <value>The \&quot;allow_any_subject\&quot; indicates that the issuer is allowed to have any principal as the subject of the JWT.</value>
        [DataMember(Name = "allow_any_subject", EmitDefaultValue = true)]
        public bool AllowAnySubject { get; set; }

        /// <summary>
        /// The \&quot;expires_at\&quot; indicates, when grant will expire, so we will reject assertion from \&quot;issuer\&quot; targeting \&quot;subject\&quot;.
        /// </summary>
        /// <value>The \&quot;expires_at\&quot; indicates, when grant will expire, so we will reject assertion from \&quot;issuer\&quot; targeting \&quot;subject\&quot;.</value>
        [DataMember(Name = "expires_at", IsRequired = true, EmitDefaultValue = false)]
        public DateTime ExpiresAt { get; set; }

        /// <summary>
        /// The \&quot;issuer\&quot; identifies the principal that issued the JWT assertion (same as \&quot;iss\&quot; claim in JWT).
        /// </summary>
        /// <value>The \&quot;issuer\&quot; identifies the principal that issued the JWT assertion (same as \&quot;iss\&quot; claim in JWT).</value>
        [DataMember(Name = "issuer", IsRequired = true, EmitDefaultValue = false)]
        public string Issuer { get; set; }

        /// <summary>
        /// Gets or Sets Jwk
        /// </summary>
        [DataMember(Name = "jwk", IsRequired = true, EmitDefaultValue = false)]
        public ClientJsonWebKey Jwk { get; set; }

        /// <summary>
        /// The \&quot;scope\&quot; contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
        /// </summary>
        /// <value>The \&quot;scope\&quot; contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])</value>
        [DataMember(Name = "scope", IsRequired = true, EmitDefaultValue = false)]
        public List<string> Scope { get; set; }

        /// <summary>
        /// The \&quot;subject\&quot; identifies the principal that is the subject of the JWT.
        /// </summary>
        /// <value>The \&quot;subject\&quot; identifies the principal that is the subject of the JWT.</value>
        [DataMember(Name = "subject", EmitDefaultValue = false)]
        public string Subject { get; set; }

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
            sb.Append("class ClientAdminTrustOAuth2JwtGrantIssuerBody {\n");
            sb.Append("  AllowAnySubject: ").Append(AllowAnySubject).Append("\n");
            sb.Append("  ExpiresAt: ").Append(ExpiresAt).Append("\n");
            sb.Append("  Issuer: ").Append(Issuer).Append("\n");
            sb.Append("  Jwk: ").Append(Jwk).Append("\n");
            sb.Append("  Scope: ").Append(Scope).Append("\n");
            sb.Append("  Subject: ").Append(Subject).Append("\n");
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
            return this.Equals(input as ClientAdminTrustOAuth2JwtGrantIssuerBody);
        }

        /// <summary>
        /// Returns true if ClientAdminTrustOAuth2JwtGrantIssuerBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientAdminTrustOAuth2JwtGrantIssuerBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientAdminTrustOAuth2JwtGrantIssuerBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.AllowAnySubject == input.AllowAnySubject ||
                    this.AllowAnySubject.Equals(input.AllowAnySubject)
                ) && 
                (
                    this.ExpiresAt == input.ExpiresAt ||
                    (this.ExpiresAt != null &&
                    this.ExpiresAt.Equals(input.ExpiresAt))
                ) && 
                (
                    this.Issuer == input.Issuer ||
                    (this.Issuer != null &&
                    this.Issuer.Equals(input.Issuer))
                ) && 
                (
                    this.Jwk == input.Jwk ||
                    (this.Jwk != null &&
                    this.Jwk.Equals(input.Jwk))
                ) && 
                (
                    this.Scope == input.Scope ||
                    this.Scope != null &&
                    input.Scope != null &&
                    this.Scope.SequenceEqual(input.Scope)
                ) && 
                (
                    this.Subject == input.Subject ||
                    (this.Subject != null &&
                    this.Subject.Equals(input.Subject))
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
                hashCode = (hashCode * 59) + this.AllowAnySubject.GetHashCode();
                if (this.ExpiresAt != null)
                {
                    hashCode = (hashCode * 59) + this.ExpiresAt.GetHashCode();
                }
                if (this.Issuer != null)
                {
                    hashCode = (hashCode * 59) + this.Issuer.GetHashCode();
                }
                if (this.Jwk != null)
                {
                    hashCode = (hashCode * 59) + this.Jwk.GetHashCode();
                }
                if (this.Scope != null)
                {
                    hashCode = (hashCode * 59) + this.Scope.GetHashCode();
                }
                if (this.Subject != null)
                {
                    hashCode = (hashCode * 59) + this.Subject.GetHashCode();
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
