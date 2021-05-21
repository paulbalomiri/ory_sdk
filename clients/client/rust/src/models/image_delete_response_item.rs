/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.9
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ImageDeleteResponseItem : ImageDeleteResponseItem image delete response item



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ImageDeleteResponseItem {
    /// The image ID of an image that was deleted
    #[serde(rename = "Deleted", skip_serializing_if = "Option::is_none")]
    pub deleted: Option<String>,
    /// The image ID of an image that was untagged
    #[serde(rename = "Untagged", skip_serializing_if = "Option::is_none")]
    pub untagged: Option<String>,
}

impl ImageDeleteResponseItem {
    /// ImageDeleteResponseItem image delete response item
    pub fn new() -> ImageDeleteResponseItem {
        ImageDeleteResponseItem {
            deleted: None,
            untagged: None,
        }
    }
}


