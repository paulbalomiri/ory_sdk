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
    /// ClientUiNodeScriptAttributes
    /// </summary>
    [DataContract(Name = "uiNodeScriptAttributes")]
    public partial class ClientUiNodeScriptAttributes : IEquatable<ClientUiNodeScriptAttributes>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiNodeScriptAttributes" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientUiNodeScriptAttributes()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiNodeScriptAttributes" /> class.
        /// </summary>
        /// <param name="async">The script async type (required).</param>
        /// <param name="crossorigin">The script cross origin policy (required).</param>
        /// <param name="id">A unique identifier (required).</param>
        /// <param name="integrity">The script&#39;s integrity hash (required).</param>
        /// <param name="nodeType">NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. (required).</param>
        /// <param name="nonce">Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value! (required).</param>
        /// <param name="referrerpolicy">The script referrer policy (required).</param>
        /// <param name="src">The script source (required).</param>
        /// <param name="type">The script MIME type (required).</param>
        public ClientUiNodeScriptAttributes(bool async = default(bool), string crossorigin = default(string), string id = default(string), string integrity = default(string), string nodeType = default(string), string nonce = default(string), string referrerpolicy = default(string), string src = default(string), string type = default(string))
        {
            this.Async = async;
            // to ensure "crossorigin" is required (not null)
            if (crossorigin == null) {
                throw new ArgumentNullException("crossorigin is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Crossorigin = crossorigin;
            // to ensure "id" is required (not null)
            if (id == null) {
                throw new ArgumentNullException("id is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Id = id;
            // to ensure "integrity" is required (not null)
            if (integrity == null) {
                throw new ArgumentNullException("integrity is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Integrity = integrity;
            // to ensure "nodeType" is required (not null)
            if (nodeType == null) {
                throw new ArgumentNullException("nodeType is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.NodeType = nodeType;
            // to ensure "nonce" is required (not null)
            if (nonce == null) {
                throw new ArgumentNullException("nonce is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Nonce = nonce;
            // to ensure "referrerpolicy" is required (not null)
            if (referrerpolicy == null) {
                throw new ArgumentNullException("referrerpolicy is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Referrerpolicy = referrerpolicy;
            // to ensure "src" is required (not null)
            if (src == null) {
                throw new ArgumentNullException("src is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Src = src;
            // to ensure "type" is required (not null)
            if (type == null) {
                throw new ArgumentNullException("type is a required property for ClientUiNodeScriptAttributes and cannot be null");
            }
            this.Type = type;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The script async type
        /// </summary>
        /// <value>The script async type</value>
        [DataMember(Name = "async", IsRequired = true, EmitDefaultValue = true)]
        public bool Async { get; set; }

        /// <summary>
        /// The script cross origin policy
        /// </summary>
        /// <value>The script cross origin policy</value>
        [DataMember(Name = "crossorigin", IsRequired = true, EmitDefaultValue = false)]
        public string Crossorigin { get; set; }

        /// <summary>
        /// A unique identifier
        /// </summary>
        /// <value>A unique identifier</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// The script&#39;s integrity hash
        /// </summary>
        /// <value>The script&#39;s integrity hash</value>
        [DataMember(Name = "integrity", IsRequired = true, EmitDefaultValue = false)]
        public string Integrity { get; set; }

        /// <summary>
        /// NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;.
        /// </summary>
        /// <value>NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;.</value>
        [DataMember(Name = "node_type", IsRequired = true, EmitDefaultValue = false)]
        public string NodeType { get; set; }

        /// <summary>
        /// Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!
        /// </summary>
        /// <value>Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!</value>
        [DataMember(Name = "nonce", IsRequired = true, EmitDefaultValue = false)]
        public string Nonce { get; set; }

        /// <summary>
        /// The script referrer policy
        /// </summary>
        /// <value>The script referrer policy</value>
        [DataMember(Name = "referrerpolicy", IsRequired = true, EmitDefaultValue = false)]
        public string Referrerpolicy { get; set; }

        /// <summary>
        /// The script source
        /// </summary>
        /// <value>The script source</value>
        [DataMember(Name = "src", IsRequired = true, EmitDefaultValue = false)]
        public string Src { get; set; }

        /// <summary>
        /// The script MIME type
        /// </summary>
        /// <value>The script MIME type</value>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = false)]
        public string Type { get; set; }

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
            sb.Append("class ClientUiNodeScriptAttributes {\n");
            sb.Append("  Async: ").Append(Async).Append("\n");
            sb.Append("  Crossorigin: ").Append(Crossorigin).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Integrity: ").Append(Integrity).Append("\n");
            sb.Append("  NodeType: ").Append(NodeType).Append("\n");
            sb.Append("  Nonce: ").Append(Nonce).Append("\n");
            sb.Append("  Referrerpolicy: ").Append(Referrerpolicy).Append("\n");
            sb.Append("  Src: ").Append(Src).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
            return this.Equals(input as ClientUiNodeScriptAttributes);
        }

        /// <summary>
        /// Returns true if ClientUiNodeScriptAttributes instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientUiNodeScriptAttributes to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientUiNodeScriptAttributes input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Async == input.Async ||
                    this.Async.Equals(input.Async)
                ) && 
                (
                    this.Crossorigin == input.Crossorigin ||
                    (this.Crossorigin != null &&
                    this.Crossorigin.Equals(input.Crossorigin))
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.Integrity == input.Integrity ||
                    (this.Integrity != null &&
                    this.Integrity.Equals(input.Integrity))
                ) && 
                (
                    this.NodeType == input.NodeType ||
                    (this.NodeType != null &&
                    this.NodeType.Equals(input.NodeType))
                ) && 
                (
                    this.Nonce == input.Nonce ||
                    (this.Nonce != null &&
                    this.Nonce.Equals(input.Nonce))
                ) && 
                (
                    this.Referrerpolicy == input.Referrerpolicy ||
                    (this.Referrerpolicy != null &&
                    this.Referrerpolicy.Equals(input.Referrerpolicy))
                ) && 
                (
                    this.Src == input.Src ||
                    (this.Src != null &&
                    this.Src.Equals(input.Src))
                ) && 
                (
                    this.Type == input.Type ||
                    (this.Type != null &&
                    this.Type.Equals(input.Type))
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
                hashCode = (hashCode * 59) + this.Async.GetHashCode();
                if (this.Crossorigin != null)
                {
                    hashCode = (hashCode * 59) + this.Crossorigin.GetHashCode();
                }
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.Integrity != null)
                {
                    hashCode = (hashCode * 59) + this.Integrity.GetHashCode();
                }
                if (this.NodeType != null)
                {
                    hashCode = (hashCode * 59) + this.NodeType.GetHashCode();
                }
                if (this.Nonce != null)
                {
                    hashCode = (hashCode * 59) + this.Nonce.GetHashCode();
                }
                if (this.Referrerpolicy != null)
                {
                    hashCode = (hashCode * 59) + this.Referrerpolicy.GetHashCode();
                }
                if (this.Src != null)
                {
                    hashCode = (hashCode * 59) + this.Src.GetHashCode();
                }
                if (this.Type != null)
                {
                    hashCode = (hashCode * 59) + this.Type.GetHashCode();
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
