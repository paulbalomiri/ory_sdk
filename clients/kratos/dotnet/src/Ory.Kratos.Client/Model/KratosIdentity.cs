/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.4
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
    /// KratosIdentity
    /// </summary>
    [DataContract(Name = "Identity")]
    public partial class KratosIdentity : IEquatable<KratosIdentity>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosIdentity" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosIdentity() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosIdentity" /> class.
        /// </summary>
        /// <param name="id">id (required).</param>
        /// <param name="recoveryAddresses">RecoveryAddresses contains all the addresses that can be used to recover an identity..</param>
        /// <param name="schemaId">SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits. (required).</param>
        /// <param name="schemaUrl">SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url (required).</param>
        /// <param name="traits">traits (required).</param>
        /// <param name="verifiableAddresses">VerifiableAddresses contains all the addresses that can be verified by the user..</param>
        public KratosIdentity(string id = default(string), List<KratosRecoveryAddress> recoveryAddresses = default(List<KratosRecoveryAddress>), string schemaId = default(string), string schemaUrl = default(string), Object traits = default(Object), List<KratosVerifiableAddress> verifiableAddresses = default(List<KratosVerifiableAddress>))
        {
            // to ensure "id" is required (not null)
            this.Id = id ?? throw new ArgumentNullException("id is a required property for KratosIdentity and cannot be null");
            // to ensure "schemaId" is required (not null)
            this.SchemaId = schemaId ?? throw new ArgumentNullException("schemaId is a required property for KratosIdentity and cannot be null");
            // to ensure "schemaUrl" is required (not null)
            this.SchemaUrl = schemaUrl ?? throw new ArgumentNullException("schemaUrl is a required property for KratosIdentity and cannot be null");
            // to ensure "traits" is required (not null)
            this.Traits = traits ?? throw new ArgumentNullException("traits is a required property for KratosIdentity and cannot be null");
            this.RecoveryAddresses = recoveryAddresses;
            this.VerifiableAddresses = verifiableAddresses;
        }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// RecoveryAddresses contains all the addresses that can be used to recover an identity.
        /// </summary>
        /// <value>RecoveryAddresses contains all the addresses that can be used to recover an identity.</value>
        [DataMember(Name = "recovery_addresses", EmitDefaultValue = false)]
        public List<KratosRecoveryAddress> RecoveryAddresses { get; set; }

        /// <summary>
        /// SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.
        /// </summary>
        /// <value>SchemaID is the ID of the JSON Schema to be used for validating the identity&#39;s traits.</value>
        [DataMember(Name = "schema_id", IsRequired = true, EmitDefaultValue = false)]
        public string SchemaId { get; set; }

        /// <summary>
        /// SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url
        /// </summary>
        /// <value>SchemaURL is the URL of the endpoint where the identity&#39;s traits schema can be fetched from.  format: url</value>
        [DataMember(Name = "schema_url", IsRequired = true, EmitDefaultValue = false)]
        public string SchemaUrl { get; set; }

        /// <summary>
        /// Gets or Sets Traits
        /// </summary>
        [DataMember(Name = "traits", IsRequired = true, EmitDefaultValue = false)]
        public Object Traits { get; set; }

        /// <summary>
        /// VerifiableAddresses contains all the addresses that can be verified by the user.
        /// </summary>
        /// <value>VerifiableAddresses contains all the addresses that can be verified by the user.</value>
        [DataMember(Name = "verifiable_addresses", EmitDefaultValue = false)]
        public List<KratosVerifiableAddress> VerifiableAddresses { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class KratosIdentity {\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  RecoveryAddresses: ").Append(RecoveryAddresses).Append("\n");
            sb.Append("  SchemaId: ").Append(SchemaId).Append("\n");
            sb.Append("  SchemaUrl: ").Append(SchemaUrl).Append("\n");
            sb.Append("  Traits: ").Append(Traits).Append("\n");
            sb.Append("  VerifiableAddresses: ").Append(VerifiableAddresses).Append("\n");
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
            return this.Equals(input as KratosIdentity);
        }

        /// <summary>
        /// Returns true if KratosIdentity instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosIdentity to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosIdentity input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.RecoveryAddresses == input.RecoveryAddresses ||
                    this.RecoveryAddresses != null &&
                    input.RecoveryAddresses != null &&
                    this.RecoveryAddresses.SequenceEqual(input.RecoveryAddresses)
                ) && 
                (
                    this.SchemaId == input.SchemaId ||
                    (this.SchemaId != null &&
                    this.SchemaId.Equals(input.SchemaId))
                ) && 
                (
                    this.SchemaUrl == input.SchemaUrl ||
                    (this.SchemaUrl != null &&
                    this.SchemaUrl.Equals(input.SchemaUrl))
                ) && 
                (
                    this.Traits == input.Traits ||
                    (this.Traits != null &&
                    this.Traits.Equals(input.Traits))
                ) && 
                (
                    this.VerifiableAddresses == input.VerifiableAddresses ||
                    this.VerifiableAddresses != null &&
                    input.VerifiableAddresses != null &&
                    this.VerifiableAddresses.SequenceEqual(input.VerifiableAddresses)
                );
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
                if (this.Id != null)
                    hashCode = hashCode * 59 + this.Id.GetHashCode();
                if (this.RecoveryAddresses != null)
                    hashCode = hashCode * 59 + this.RecoveryAddresses.GetHashCode();
                if (this.SchemaId != null)
                    hashCode = hashCode * 59 + this.SchemaId.GetHashCode();
                if (this.SchemaUrl != null)
                    hashCode = hashCode * 59 + this.SchemaUrl.GetHashCode();
                if (this.Traits != null)
                    hashCode = hashCode * 59 + this.Traits.GetHashCode();
                if (this.VerifiableAddresses != null)
                    hashCode = hashCode * 59 + this.VerifiableAddresses.GetHashCode();
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
