/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.17
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
    /// PluginConfigArgs plugin config args
    /// </summary>
    [DataContract(Name = "PluginConfigArgs")]
    public partial class KratosPluginConfigArgs : IEquatable<KratosPluginConfigArgs>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosPluginConfigArgs" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosPluginConfigArgs()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosPluginConfigArgs" /> class.
        /// </summary>
        /// <param name="description">description (required).</param>
        /// <param name="name">name (required).</param>
        /// <param name="settable">settable (required).</param>
        /// <param name="value">value (required).</param>
        public KratosPluginConfigArgs(string description = default(string), string name = default(string), List<string> settable = default(List<string>), List<string> value = default(List<string>))
        {
            // to ensure "description" is required (not null)
            this.Description = description ?? throw new ArgumentNullException("description is a required property for KratosPluginConfigArgs and cannot be null");
            // to ensure "name" is required (not null)
            this.Name = name ?? throw new ArgumentNullException("name is a required property for KratosPluginConfigArgs and cannot be null");
            // to ensure "settable" is required (not null)
            this.Settable = settable ?? throw new ArgumentNullException("settable is a required property for KratosPluginConfigArgs and cannot be null");
            // to ensure "value" is required (not null)
            this.Value = value ?? throw new ArgumentNullException("value is a required property for KratosPluginConfigArgs and cannot be null");
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// description
        /// </summary>
        /// <value>description</value>
        [DataMember(Name = "Description", IsRequired = true, EmitDefaultValue = false)]
        public string Description { get; set; }

        /// <summary>
        /// name
        /// </summary>
        /// <value>name</value>
        [DataMember(Name = "Name", IsRequired = true, EmitDefaultValue = false)]
        public string Name { get; set; }

        /// <summary>
        /// settable
        /// </summary>
        /// <value>settable</value>
        [DataMember(Name = "Settable", IsRequired = true, EmitDefaultValue = false)]
        public List<string> Settable { get; set; }

        /// <summary>
        /// value
        /// </summary>
        /// <value>value</value>
        [DataMember(Name = "Value", IsRequired = true, EmitDefaultValue = false)]
        public List<string> Value { get; set; }

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
            sb.Append("class KratosPluginConfigArgs {\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  Name: ").Append(Name).Append("\n");
            sb.Append("  Settable: ").Append(Settable).Append("\n");
            sb.Append("  Value: ").Append(Value).Append("\n");
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
            return this.Equals(input as KratosPluginConfigArgs);
        }

        /// <summary>
        /// Returns true if KratosPluginConfigArgs instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosPluginConfigArgs to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosPluginConfigArgs input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && 
                (
                    this.Name == input.Name ||
                    (this.Name != null &&
                    this.Name.Equals(input.Name))
                ) && 
                (
                    this.Settable == input.Settable ||
                    this.Settable != null &&
                    input.Settable != null &&
                    this.Settable.SequenceEqual(input.Settable)
                ) && 
                (
                    this.Value == input.Value ||
                    this.Value != null &&
                    input.Value != null &&
                    this.Value.SequenceEqual(input.Value)
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
                if (this.Description != null)
                    hashCode = hashCode * 59 + this.Description.GetHashCode();
                if (this.Name != null)
                    hashCode = hashCode * 59 + this.Name.GetHashCode();
                if (this.Settable != null)
                    hashCode = hashCode * 59 + this.Settable.GetHashCode();
                if (this.Value != null)
                    hashCode = hashCode * 59 + this.Value.GetHashCode();
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
