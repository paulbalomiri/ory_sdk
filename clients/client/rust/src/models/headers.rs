/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.60
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// Headers : Headers is the jwt headers



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Headers {
    #[serde(rename = "extra", skip_serializing_if = "Option::is_none")]
    pub extra: Option<::std::collections::HashMap<String, serde_json::Value>>,
}

impl Default for Headers {
    fn default() -> Self {
        Self::new()
    }
}

impl Headers {
    /// Headers is the jwt headers
    pub fn new() -> Headers {
        Headers {
                extra: None,
        }
    }
}


