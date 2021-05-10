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
    /// KratosPluginConfig
    /// </summary>
    [DataContract(Name = "PluginConfig")]
    public partial class KratosPluginConfig : IEquatable<KratosPluginConfig>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosPluginConfig" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosPluginConfig() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosPluginConfig" /> class.
        /// </summary>
        /// <param name="args">args (required).</param>
        /// <param name="description">description (required).</param>
        /// <param name="dockerVersion">Docker Version used to create the plugin.</param>
        /// <param name="documentation">documentation (required).</param>
        /// <param name="entrypoint">entrypoint (required).</param>
        /// <param name="env">env (required).</param>
        /// <param name="_interface">_interface (required).</param>
        /// <param name="ipcHost">ipc host (required).</param>
        /// <param name="linux">linux (required).</param>
        /// <param name="mounts">mounts (required).</param>
        /// <param name="network">network (required).</param>
        /// <param name="pidHost">pid host (required).</param>
        /// <param name="propagatedMount">propagated mount (required).</param>
        /// <param name="user">user.</param>
        /// <param name="workDir">work dir (required).</param>
        /// <param name="rootfs">rootfs.</param>
        public KratosPluginConfig(KratosPluginConfigArgs args = default(KratosPluginConfigArgs), string description = default(string), string dockerVersion = default(string), string documentation = default(string), List<string> entrypoint = default(List<string>), List<KratosPluginEnv> env = default(List<KratosPluginEnv>), KratosPluginConfigInterface _interface = default(KratosPluginConfigInterface), bool ipcHost = default(bool), KratosPluginConfigLinux linux = default(KratosPluginConfigLinux), List<KratosPluginMount> mounts = default(List<KratosPluginMount>), KratosPluginConfigNetwork network = default(KratosPluginConfigNetwork), bool pidHost = default(bool), string propagatedMount = default(string), KratosPluginConfigUser user = default(KratosPluginConfigUser), string workDir = default(string), KratosPluginConfigRootfs rootfs = default(KratosPluginConfigRootfs))
        {
            // to ensure "args" is required (not null)
            this.Args = args ?? throw new ArgumentNullException("args is a required property for KratosPluginConfig and cannot be null");
            // to ensure "description" is required (not null)
            this.Description = description ?? throw new ArgumentNullException("description is a required property for KratosPluginConfig and cannot be null");
            // to ensure "documentation" is required (not null)
            this.Documentation = documentation ?? throw new ArgumentNullException("documentation is a required property for KratosPluginConfig and cannot be null");
            // to ensure "entrypoint" is required (not null)
            this.Entrypoint = entrypoint ?? throw new ArgumentNullException("entrypoint is a required property for KratosPluginConfig and cannot be null");
            // to ensure "env" is required (not null)
            this.Env = env ?? throw new ArgumentNullException("env is a required property for KratosPluginConfig and cannot be null");
            // to ensure "_interface" is required (not null)
            this.Interface = _interface ?? throw new ArgumentNullException("_interface is a required property for KratosPluginConfig and cannot be null");
            this.IpcHost = ipcHost;
            // to ensure "linux" is required (not null)
            this.Linux = linux ?? throw new ArgumentNullException("linux is a required property for KratosPluginConfig and cannot be null");
            // to ensure "mounts" is required (not null)
            this.Mounts = mounts ?? throw new ArgumentNullException("mounts is a required property for KratosPluginConfig and cannot be null");
            // to ensure "network" is required (not null)
            this.Network = network ?? throw new ArgumentNullException("network is a required property for KratosPluginConfig and cannot be null");
            this.PidHost = pidHost;
            // to ensure "propagatedMount" is required (not null)
            this.PropagatedMount = propagatedMount ?? throw new ArgumentNullException("propagatedMount is a required property for KratosPluginConfig and cannot be null");
            // to ensure "workDir" is required (not null)
            this.WorkDir = workDir ?? throw new ArgumentNullException("workDir is a required property for KratosPluginConfig and cannot be null");
            this.DockerVersion = dockerVersion;
            this.User = user;
            this.Rootfs = rootfs;
        }

        /// <summary>
        /// Gets or Sets Args
        /// </summary>
        [DataMember(Name = "Args", IsRequired = true, EmitDefaultValue = false)]
        public KratosPluginConfigArgs Args { get; set; }

        /// <summary>
        /// description
        /// </summary>
        /// <value>description</value>
        [DataMember(Name = "Description", IsRequired = true, EmitDefaultValue = false)]
        public string Description { get; set; }

        /// <summary>
        /// Docker Version used to create the plugin
        /// </summary>
        /// <value>Docker Version used to create the plugin</value>
        [DataMember(Name = "DockerVersion", EmitDefaultValue = false)]
        public string DockerVersion { get; set; }

        /// <summary>
        /// documentation
        /// </summary>
        /// <value>documentation</value>
        [DataMember(Name = "Documentation", IsRequired = true, EmitDefaultValue = false)]
        public string Documentation { get; set; }

        /// <summary>
        /// entrypoint
        /// </summary>
        /// <value>entrypoint</value>
        [DataMember(Name = "Entrypoint", IsRequired = true, EmitDefaultValue = false)]
        public List<string> Entrypoint { get; set; }

        /// <summary>
        /// env
        /// </summary>
        /// <value>env</value>
        [DataMember(Name = "Env", IsRequired = true, EmitDefaultValue = false)]
        public List<KratosPluginEnv> Env { get; set; }

        /// <summary>
        /// Gets or Sets Interface
        /// </summary>
        [DataMember(Name = "Interface", IsRequired = true, EmitDefaultValue = false)]
        public KratosPluginConfigInterface Interface { get; set; }

        /// <summary>
        /// ipc host
        /// </summary>
        /// <value>ipc host</value>
        [DataMember(Name = "IpcHost", IsRequired = true, EmitDefaultValue = true)]
        public bool IpcHost { get; set; }

        /// <summary>
        /// Gets or Sets Linux
        /// </summary>
        [DataMember(Name = "Linux", IsRequired = true, EmitDefaultValue = false)]
        public KratosPluginConfigLinux Linux { get; set; }

        /// <summary>
        /// mounts
        /// </summary>
        /// <value>mounts</value>
        [DataMember(Name = "Mounts", IsRequired = true, EmitDefaultValue = false)]
        public List<KratosPluginMount> Mounts { get; set; }

        /// <summary>
        /// Gets or Sets Network
        /// </summary>
        [DataMember(Name = "Network", IsRequired = true, EmitDefaultValue = false)]
        public KratosPluginConfigNetwork Network { get; set; }

        /// <summary>
        /// pid host
        /// </summary>
        /// <value>pid host</value>
        [DataMember(Name = "PidHost", IsRequired = true, EmitDefaultValue = true)]
        public bool PidHost { get; set; }

        /// <summary>
        /// propagated mount
        /// </summary>
        /// <value>propagated mount</value>
        [DataMember(Name = "PropagatedMount", IsRequired = true, EmitDefaultValue = false)]
        public string PropagatedMount { get; set; }

        /// <summary>
        /// Gets or Sets User
        /// </summary>
        [DataMember(Name = "User", EmitDefaultValue = false)]
        public KratosPluginConfigUser User { get; set; }

        /// <summary>
        /// work dir
        /// </summary>
        /// <value>work dir</value>
        [DataMember(Name = "WorkDir", IsRequired = true, EmitDefaultValue = false)]
        public string WorkDir { get; set; }

        /// <summary>
        /// Gets or Sets Rootfs
        /// </summary>
        [DataMember(Name = "rootfs", EmitDefaultValue = false)]
        public KratosPluginConfigRootfs Rootfs { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class KratosPluginConfig {\n");
            sb.Append("  Args: ").Append(Args).Append("\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  DockerVersion: ").Append(DockerVersion).Append("\n");
            sb.Append("  Documentation: ").Append(Documentation).Append("\n");
            sb.Append("  Entrypoint: ").Append(Entrypoint).Append("\n");
            sb.Append("  Env: ").Append(Env).Append("\n");
            sb.Append("  Interface: ").Append(Interface).Append("\n");
            sb.Append("  IpcHost: ").Append(IpcHost).Append("\n");
            sb.Append("  Linux: ").Append(Linux).Append("\n");
            sb.Append("  Mounts: ").Append(Mounts).Append("\n");
            sb.Append("  Network: ").Append(Network).Append("\n");
            sb.Append("  PidHost: ").Append(PidHost).Append("\n");
            sb.Append("  PropagatedMount: ").Append(PropagatedMount).Append("\n");
            sb.Append("  User: ").Append(User).Append("\n");
            sb.Append("  WorkDir: ").Append(WorkDir).Append("\n");
            sb.Append("  Rootfs: ").Append(Rootfs).Append("\n");
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
            return this.Equals(input as KratosPluginConfig);
        }

        /// <summary>
        /// Returns true if KratosPluginConfig instances are equal
        /// </summary>
        /// <param name="input">Instance of KratosPluginConfig to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(KratosPluginConfig input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Args == input.Args ||
                    (this.Args != null &&
                    this.Args.Equals(input.Args))
                ) && 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && 
                (
                    this.DockerVersion == input.DockerVersion ||
                    (this.DockerVersion != null &&
                    this.DockerVersion.Equals(input.DockerVersion))
                ) && 
                (
                    this.Documentation == input.Documentation ||
                    (this.Documentation != null &&
                    this.Documentation.Equals(input.Documentation))
                ) && 
                (
                    this.Entrypoint == input.Entrypoint ||
                    this.Entrypoint != null &&
                    input.Entrypoint != null &&
                    this.Entrypoint.SequenceEqual(input.Entrypoint)
                ) && 
                (
                    this.Env == input.Env ||
                    this.Env != null &&
                    input.Env != null &&
                    this.Env.SequenceEqual(input.Env)
                ) && 
                (
                    this.Interface == input.Interface ||
                    (this.Interface != null &&
                    this.Interface.Equals(input.Interface))
                ) && 
                (
                    this.IpcHost == input.IpcHost ||
                    this.IpcHost.Equals(input.IpcHost)
                ) && 
                (
                    this.Linux == input.Linux ||
                    (this.Linux != null &&
                    this.Linux.Equals(input.Linux))
                ) && 
                (
                    this.Mounts == input.Mounts ||
                    this.Mounts != null &&
                    input.Mounts != null &&
                    this.Mounts.SequenceEqual(input.Mounts)
                ) && 
                (
                    this.Network == input.Network ||
                    (this.Network != null &&
                    this.Network.Equals(input.Network))
                ) && 
                (
                    this.PidHost == input.PidHost ||
                    this.PidHost.Equals(input.PidHost)
                ) && 
                (
                    this.PropagatedMount == input.PropagatedMount ||
                    (this.PropagatedMount != null &&
                    this.PropagatedMount.Equals(input.PropagatedMount))
                ) && 
                (
                    this.User == input.User ||
                    (this.User != null &&
                    this.User.Equals(input.User))
                ) && 
                (
                    this.WorkDir == input.WorkDir ||
                    (this.WorkDir != null &&
                    this.WorkDir.Equals(input.WorkDir))
                ) && 
                (
                    this.Rootfs == input.Rootfs ||
                    (this.Rootfs != null &&
                    this.Rootfs.Equals(input.Rootfs))
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
                if (this.Args != null)
                    hashCode = hashCode * 59 + this.Args.GetHashCode();
                if (this.Description != null)
                    hashCode = hashCode * 59 + this.Description.GetHashCode();
                if (this.DockerVersion != null)
                    hashCode = hashCode * 59 + this.DockerVersion.GetHashCode();
                if (this.Documentation != null)
                    hashCode = hashCode * 59 + this.Documentation.GetHashCode();
                if (this.Entrypoint != null)
                    hashCode = hashCode * 59 + this.Entrypoint.GetHashCode();
                if (this.Env != null)
                    hashCode = hashCode * 59 + this.Env.GetHashCode();
                if (this.Interface != null)
                    hashCode = hashCode * 59 + this.Interface.GetHashCode();
                hashCode = hashCode * 59 + this.IpcHost.GetHashCode();
                if (this.Linux != null)
                    hashCode = hashCode * 59 + this.Linux.GetHashCode();
                if (this.Mounts != null)
                    hashCode = hashCode * 59 + this.Mounts.GetHashCode();
                if (this.Network != null)
                    hashCode = hashCode * 59 + this.Network.GetHashCode();
                hashCode = hashCode * 59 + this.PidHost.GetHashCode();
                if (this.PropagatedMount != null)
                    hashCode = hashCode * 59 + this.PropagatedMount.GetHashCode();
                if (this.User != null)
                    hashCode = hashCode * 59 + this.User.GetHashCode();
                if (this.WorkDir != null)
                    hashCode = hashCode * 59 + this.WorkDir.GetHashCode();
                if (this.Rootfs != null)
                    hashCode = hashCode * 59 + this.Rootfs.GetHashCode();
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
