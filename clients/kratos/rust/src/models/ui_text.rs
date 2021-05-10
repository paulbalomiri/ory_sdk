/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.17
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UiText {
    /// The message's context. Useful when customizing messages.
    #[serde(rename = "context", skip_serializing_if = "Option::is_none")]
    pub context: Option<serde_json::Value>,
    #[serde(rename = "id")]
    pub id: i64,
    /// The message text. Written in american english.
    #[serde(rename = "text")]
    pub text: String,
    #[serde(rename = "type")]
    pub _type: String,
}

impl UiText {
    pub fn new(id: i64, text: String, _type: String) -> UiText {
        UiText {
            context: None,
            id,
            text,
            _type,
        }
    }
}


