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
    /// A singular authenticator used during authentication / login.
    /// </summary>
    [DataContract(Name = "sessionAuthenticationMethod")]
    public partial class ClientSessionAuthenticationMethod : IEquatable<ClientSessionAuthenticationMethod>, IValidatableObject
    {
        /// <summary>
        /// Defines Method
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum MethodEnum
        {
            /// <summary>
            /// Enum LinkRecovery for value: link_recovery
            /// </summary>
            [EnumMember(Value = "link_recovery")]
            LinkRecovery = 1,

            /// <summary>
            /// Enum CodeRecovery for value: code_recovery
            /// </summary>
            [EnumMember(Value = "code_recovery")]
            CodeRecovery = 2,

            /// <summary>
            /// Enum Password for value: password
            /// </summary>
            [EnumMember(Value = "password")]
            Password = 3,

            /// <summary>
            /// Enum Totp for value: totp
            /// </summary>
            [EnumMember(Value = "totp")]
            Totp = 4,

            /// <summary>
            /// Enum Oidc for value: oidc
            /// </summary>
            [EnumMember(Value = "oidc")]
            Oidc = 5,

            /// <summary>
            /// Enum Webauthn for value: webauthn
            /// </summary>
            [EnumMember(Value = "webauthn")]
            Webauthn = 6,

            /// <summary>
            /// Enum LookupSecret for value: lookup_secret
            /// </summary>
            [EnumMember(Value = "lookup_secret")]
            LookupSecret = 7,

            /// <summary>
            /// Enum V06LegacySession for value: v0.6_legacy_session
            /// </summary>
            [EnumMember(Value = "v0.6_legacy_session")]
            V06LegacySession = 8

        }


        /// <summary>
        /// Gets or Sets Method
        /// </summary>
        [DataMember(Name = "method", EmitDefaultValue = false)]
        public MethodEnum? Method { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientSessionAuthenticationMethod" /> class.
        /// </summary>
        /// <param name="aal">aal.</param>
        /// <param name="completedAt">When the authentication challenge was completed..</param>
        /// <param name="method">method.</param>
        public ClientSessionAuthenticationMethod(ClientAuthenticatorAssuranceLevel aal = default(ClientAuthenticatorAssuranceLevel), DateTime completedAt = default(DateTime), MethodEnum? method = default(MethodEnum?))
        {
            this.Aal = aal;
            this.CompletedAt = completedAt;
            this.Method = method;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Aal
        /// </summary>
        [DataMember(Name = "aal", EmitDefaultValue = false)]
        public ClientAuthenticatorAssuranceLevel Aal { get; set; }

        /// <summary>
        /// When the authentication challenge was completed.
        /// </summary>
        /// <value>When the authentication challenge was completed.</value>
        [DataMember(Name = "completed_at", EmitDefaultValue = false)]
        public DateTime CompletedAt { get; set; }

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
            sb.Append("class ClientSessionAuthenticationMethod {\n");
            sb.Append("  Aal: ").Append(Aal).Append("\n");
            sb.Append("  CompletedAt: ").Append(CompletedAt).Append("\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
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
            return this.Equals(input as ClientSessionAuthenticationMethod);
        }

        /// <summary>
        /// Returns true if ClientSessionAuthenticationMethod instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientSessionAuthenticationMethod to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientSessionAuthenticationMethod input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Aal == input.Aal ||
                    (this.Aal != null &&
                    this.Aal.Equals(input.Aal))
                ) && 
                (
                    this.CompletedAt == input.CompletedAt ||
                    (this.CompletedAt != null &&
                    this.CompletedAt.Equals(input.CompletedAt))
                ) && 
                (
                    this.Method == input.Method ||
                    this.Method.Equals(input.Method)
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
                if (this.Aal != null)
                {
                    hashCode = (hashCode * 59) + this.Aal.GetHashCode();
                }
                if (this.CompletedAt != null)
                {
                    hashCode = (hashCode * 59) + this.CompletedAt.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Method.GetHashCode();
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
