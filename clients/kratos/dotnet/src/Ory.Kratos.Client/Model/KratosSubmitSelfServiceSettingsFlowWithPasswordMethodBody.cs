/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.6-alpha.3
 * Contact: hi@ory.sh
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
using OpenAPIDateConverter = Ory.Kratos.Client.Client.OpenAPIDateConverter;

namespace Ory.Kratos.Client.Model
{
    /// <summary>
    /// KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody
    /// </summary>
    [DataContract(Name = "submitSelfServiceSettingsFlowWithPasswordMethodBody")]
    public partial class KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody : IEquatable<KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody" /> class.
        /// </summary>
        /// <param name="csrfToken">CSRFToken is the anti-CSRF token.</param>
        /// <param name="method">Method  Should be set to password when trying to update a password. (required).</param>
        /// <param name="password">Password is the updated password (required).</param>
        public KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody(string csrfToken = default(string), string method = default(string), string password = default(string))
        {
            // to ensure "method" is required (not null)
            if (method == null) {
                throw new ArgumentNullException("method is a required property for KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody and cannot be null");
            }
            this.Method = method;
            // to ensure "password" is required (not null)
            if (password == null) {
                throw new ArgumentNullException("password is a required property for KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody and cannot be null");
            }
            this.Password = password;
            this.CsrfToken = csrfToken;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// CSRFToken is the anti-CSRF token
        /// </summary>
        /// <value>CSRFToken is the anti-CSRF token</value>
        [DataMember(Name = "csrf_token", EmitDefaultValue = false)]
        public string CsrfToken { get; set; }

        /// <summary>
        /// Method  Should be set to password when trying to update a password.
        /// </summary>
        /// <value>Method  Should be set to password when trying to update a password.</value>
        [DataMember(Name = "method", IsRequired = true, EmitDefaultValue = false)]
        public string Method { get; set; }

        /// <summary>
        /// Password is the updated password
        /// </summary>
        /// <value>Password is the updated password</value>
        [DataMember(Name = "password", IsRequired = true, EmitDefaultValue = false)]
        public string Password { get; set; }

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
            var sb = new StringBuilder();
            sb.Append("class KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody {\n");
            sb.Append("  CsrfToken: ").Append(CsrfToken).Append("\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
            sb.Append("  Password: ").Append(Password).Append("\n");
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
            return this.Equals(input as KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody);
        }

        /// <summary>
        /// Returns true if KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosSubmitSelfServiceSettingsFlowWithPasswordMethodBody input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.CsrfToken == input.CsrfToken ||
                    (this.CsrfToken != null &&
                    this.CsrfToken.Equals(input.CsrfToken))
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
                    hashCode = hashCode * 59 + this.CsrfToken.GetHashCode();
                if (this.Method != null)
                    hashCode = hashCode * 59 + this.Method.GetHashCode();
                if (this.Password != null)
                    hashCode = hashCode * 59 + this.Password.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
