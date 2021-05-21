/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.9
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginConfigUser : PluginConfigUser plugin config user



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginConfigUser {
    /// g ID
    #[serde(rename = "GID", skip_serializing_if = "Option::is_none")]
    pub GID: Option<i32>,
    /// UID
    #[serde(rename = "UID", skip_serializing_if = "Option::is_none")]
    pub UID: Option<i32>,
}

impl PluginConfigUser {
    /// PluginConfigUser plugin config user
    pub fn new() -> PluginConfigUser {
        PluginConfigUser {
            GID: None,
            UID: None,
        }
    }
}


