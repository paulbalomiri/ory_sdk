/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.9
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UiNodeAnchorAttributes {
    /// The link's href (destination) URL.  format: uri
    #[serde(rename = "href")]
    pub href: String,
    #[serde(rename = "title")]
    pub title: Box<crate::models::UiText>,
}

impl UiNodeAnchorAttributes {
    pub fn new(href: String, title: crate::models::UiText) -> UiNodeAnchorAttributes {
        UiNodeAnchorAttributes {
            href,
            title: Box::new(title),
        }
    }
}


