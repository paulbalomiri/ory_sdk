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
    /// ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody
    /// </summary>
    [DataContract(Name = "submitSelfServiceLoginFlowWithPasswordMethodBody")]
    public partial class ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody : IEquatable<ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody" /> class.
        /// </summary>
        /// <param name="csrfToken">Sending the anti-csrf token is only required for browser login flows..</param>
        /// <param name="identifier">Identifier is the email or username of the user trying to log in. (required).</param>
        /// <param name="method">Method should be set to \&quot;password\&quot; when logging in using the identifier and password strategy. (required).</param>
        /// <param name="password">The user&#39;s password. (required).</param>
        /// <param name="passwordIdentifier">Identifier is the email or username of the user trying to log in. This field is deprecated!.</param>
        public ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody(string csrfToken = default(string), string identifier = default(string), string method = default(string), string password = default(string), string passwordIdentifier = default(string))
        {
            // to ensure "identifier" is required (not null)
            if (identifier == null) {
                throw new ArgumentNullException("identifier is a required property for ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody and cannot be null");
            }
            this.Identifier = identifier;
            // to ensure "method" is required (not null)
            if (method == null) {
                throw new ArgumentNullException("method is a required property for ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody and cannot be null");
            }
            this.Method = method;
            // to ensure "password" is required (not null)
            if (password == null) {
                throw new ArgumentNullException("password is a required property for ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody and cannot be null");
            }
            this.Password = password;
            this.CsrfToken = csrfToken;
            this.PasswordIdentifier = passwordIdentifier;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Sending the anti-csrf token is only required for browser login flows.
        /// </summary>
        /// <value>Sending the anti-csrf token is only required for browser login flows.</value>
        [DataMember(Name = "csrf_token", EmitDefaultValue = false)]
        public string CsrfToken { get; set; }

        /// <summary>
        /// Identifier is the email or username of the user trying to log in.
        /// </summary>
        /// <value>Identifier is the email or username of the user trying to log in.</value>
        [DataMember(Name = "identifier", IsRequired = true, EmitDefaultValue = false)]
        public string Identifier { get; set; }

        /// <summary>
        /// Method should be set to \&quot;password\&quot; when logging in using the identifier and password strategy.
        /// </summary>
        /// <value>Method should be set to \&quot;password\&quot; when logging in using the identifier and password strategy.</value>
        [DataMember(Name = "method", IsRequired = true, EmitDefaultValue = false)]
        public string Method { get; set; }

        /// <summary>
        /// The user&#39;s password.
        /// </summary>
        /// <value>The user&#39;s password.</value>
        [DataMember(Name = "password", IsRequired = true, EmitDefaultValue = false)]
        public string Password { get; set; }

        /// <summary>
        /// Identifier is the email or username of the user trying to log in. This field is deprecated!
        /// </summary>
        /// <value>Identifier is the email or username of the user trying to log in. This field is deprecated!</value>
        [DataMember(Name = "password_identifier", EmitDefaultValue = false)]
        public string PasswordIdentifier { get; set; }

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
            sb.Append("class ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody {\n");
            sb.Append("  CsrfToken: ").Append(CsrfToken).Append("\n");
            sb.Append("  Identifier: ").Append(Identifier).Append("\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
            sb.Append("  Password: ").Append(Password).Append("\n");
            sb.Append("  PasswordIdentifier: ").Append(PasswordIdentifier).Append("\n");
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
            return this.Equals(input as ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody);
        }

        /// <summary>
        /// Returns true if ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientSubmitSelfServiceLoginFlowWithPasswordMethodBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.CsrfToken == input.CsrfToken ||
                    (this.CsrfToken != null &&
                    this.CsrfToken.Equals(input.CsrfToken))
                ) && 
                (
                    this.Identifier == input.Identifier ||
                    (this.Identifier != null &&
                    this.Identifier.Equals(input.Identifier))
                ) && 
                (
                    this.Method == input.Method ||
                    (this.Method != null &&
                    this.Method.Equals(input.Method))
                ) && 
                (
                    this.Password == input.Password ||
                    (this.Password != null &&
                    this.Password.Equals(input.Password))
                ) && 
                (
                    this.PasswordIdentifier == input.PasswordIdentifier ||
                    (this.PasswordIdentifier != null &&
                    this.PasswordIdentifier.Equals(input.PasswordIdentifier))
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
                if (this.CsrfToken != null)
                {
                    hashCode = (hashCode * 59) + this.CsrfToken.GetHashCode();
                }
                if (this.Identifier != null)
                {
                    hashCode = (hashCode * 59) + this.Identifier.GetHashCode();
                }
                if (this.Method != null)
                {
                    hashCode = (hashCode * 59) + this.Method.GetHashCode();
                }
                if (this.Password != null)
                {
                    hashCode = (hashCode * 59) + this.Password.GetHashCode();
                }
                if (this.PasswordIdentifier != null)
                {
                    hashCode = (hashCode * 59) + this.PasswordIdentifier.GetHashCode();
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
