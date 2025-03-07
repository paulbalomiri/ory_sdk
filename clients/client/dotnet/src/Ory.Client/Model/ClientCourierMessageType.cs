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
    /// It can either be &#x60;email&#x60; or &#x60;phone&#x60;
    /// </summary>
    /// <value>It can either be &#x60;email&#x60; or &#x60;phone&#x60;</value>
    [JsonConverter(typeof(StringEnumConverter))]
    public enum ClientCourierMessageType
    {
        /// <summary>
        /// Enum Email for value: email
        /// </summary>
        [EnumMember(Value = "email")]
        Email = 1,

        /// <summary>
        /// Enum Phone for value: phone
        /// </summary>
        [EnumMember(Value = "phone")]
        Phone = 2

    }

}
