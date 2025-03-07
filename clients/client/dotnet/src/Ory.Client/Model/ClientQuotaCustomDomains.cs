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
    /// ClientQuotaCustomDomains
    /// </summary>
    [DataContract(Name = "QuotaCustomDomains")]
    public partial class ClientQuotaCustomDomains : IEquatable<ClientQuotaCustomDomains>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientQuotaCustomDomains" /> class.
        /// </summary>
        /// <param name="availableDomains">availableDomains.</param>
        /// <param name="canUse">canUse.</param>
        /// <param name="usedDomains">usedDomains.</param>
        public ClientQuotaCustomDomains(long availableDomains = default(long), bool canUse = default(bool), long usedDomains = default(long))
        {
            this.AvailableDomains = availableDomains;
            this.CanUse = canUse;
            this.UsedDomains = usedDomains;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets AvailableDomains
        /// </summary>
        [DataMember(Name = "available_domains", EmitDefaultValue = false)]
        public long AvailableDomains { get; set; }

        /// <summary>
        /// Gets or Sets CanUse
        /// </summary>
        [DataMember(Name = "can_use", EmitDefaultValue = true)]
        public bool CanUse { get; set; }

        /// <summary>
        /// Gets or Sets UsedDomains
        /// </summary>
        [DataMember(Name = "used_domains", EmitDefaultValue = false)]
        public long UsedDomains { get; set; }

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
            sb.Append("class ClientQuotaCustomDomains {\n");
            sb.Append("  AvailableDomains: ").Append(AvailableDomains).Append("\n");
            sb.Append("  CanUse: ").Append(CanUse).Append("\n");
            sb.Append("  UsedDomains: ").Append(UsedDomains).Append("\n");
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
            return this.Equals(input as ClientQuotaCustomDomains);
        }

        /// <summary>
        /// Returns true if ClientQuotaCustomDomains instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientQuotaCustomDomains to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientQuotaCustomDomains input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.AvailableDomains == input.AvailableDomains ||
                    this.AvailableDomains.Equals(input.AvailableDomains)
                ) && 
                (
                    this.CanUse == input.CanUse ||
                    this.CanUse.Equals(input.CanUse)
                ) && 
                (
                    this.UsedDomains == input.UsedDomains ||
                    this.UsedDomains.Equals(input.UsedDomains)
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
                hashCode = (hashCode * 59) + this.AvailableDomains.GetHashCode();
                hashCode = (hashCode * 59) + this.CanUse.GetHashCode();
                hashCode = (hashCode * 59) + this.UsedDomains.GetHashCode();
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
