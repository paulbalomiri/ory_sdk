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
    /// ClientUiNodeImageAttributes
    /// </summary>
    [DataContract(Name = "uiNodeImageAttributes")]
    public partial class ClientUiNodeImageAttributes : IEquatable<ClientUiNodeImageAttributes>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiNodeImageAttributes" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientUiNodeImageAttributes()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiNodeImageAttributes" /> class.
        /// </summary>
        /// <param name="height">Height of the image (required).</param>
        /// <param name="id">A unique identifier (required).</param>
        /// <param name="nodeType">NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;. (required).</param>
        /// <param name="src">The image&#39;s source URL.  format: uri (required).</param>
        /// <param name="width">Width of the image (required).</param>
        public ClientUiNodeImageAttributes(long height = default(long), string id = default(string), string nodeType = default(string), string src = default(string), long width = default(long))
        {
            this.Height = height;
            // to ensure "id" is required (not null)
            if (id == null) {
                throw new ArgumentNullException("id is a required property for ClientUiNodeImageAttributes and cannot be null");
            }
            this.Id = id;
            // to ensure "nodeType" is required (not null)
            if (nodeType == null) {
                throw new ArgumentNullException("nodeType is a required property for ClientUiNodeImageAttributes and cannot be null");
            }
            this.NodeType = nodeType;
            // to ensure "src" is required (not null)
            if (src == null) {
                throw new ArgumentNullException("src is a required property for ClientUiNodeImageAttributes and cannot be null");
            }
            this.Src = src;
            this.Width = width;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Height of the image
        /// </summary>
        /// <value>Height of the image</value>
        [DataMember(Name = "height", IsRequired = true, EmitDefaultValue = false)]
        public long Height { get; set; }

        /// <summary>
        /// A unique identifier
        /// </summary>
        /// <value>A unique identifier</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;.
        /// </summary>
        /// <value>NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;.</value>
        [DataMember(Name = "node_type", IsRequired = true, EmitDefaultValue = false)]
        public string NodeType { get; set; }

        /// <summary>
        /// The image&#39;s source URL.  format: uri
        /// </summary>
        /// <value>The image&#39;s source URL.  format: uri</value>
        [DataMember(Name = "src", IsRequired = true, EmitDefaultValue = false)]
        public string Src { get; set; }

        /// <summary>
        /// Width of the image
        /// </summary>
        /// <value>Width of the image</value>
        [DataMember(Name = "width", IsRequired = true, EmitDefaultValue = false)]
        public long Width { get; set; }

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
            sb.Append("class ClientUiNodeImageAttributes {\n");
            sb.Append("  Height: ").Append(Height).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  NodeType: ").Append(NodeType).Append("\n");
            sb.Append("  Src: ").Append(Src).Append("\n");
            sb.Append("  Width: ").Append(Width).Append("\n");
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
            return this.Equals(input as ClientUiNodeImageAttributes);
        }

        /// <summary>
        /// Returns true if ClientUiNodeImageAttributes instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientUiNodeImageAttributes to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientUiNodeImageAttributes input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Height == input.Height ||
                    this.Height.Equals(input.Height)
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.NodeType == input.NodeType ||
                    (this.NodeType != null &&
                    this.NodeType.Equals(input.NodeType))
                ) && 
                (
                    this.Src == input.Src ||
                    (this.Src != null &&
                    this.Src.Equals(input.Src))
                ) && 
                (
                    this.Width == input.Width ||
                    this.Width.Equals(input.Width)
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
                hashCode = (hashCode * 59) + this.Height.GetHashCode();
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.NodeType != null)
                {
                    hashCode = (hashCode * 59) + this.NodeType.GetHashCode();
                }
                if (this.Src != null)
                {
                    hashCode = (hashCode * 59) + this.Src.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Width.GetHashCode();
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
