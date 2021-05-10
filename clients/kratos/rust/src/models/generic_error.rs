/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.4
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// GenericError : Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GenericError {
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<Box<crate::models::GenericErrorPayload>>,
}

impl GenericError {
    /// Error responses are sent when an error (e.g. unauthorized, bad request, ...) occurred.
    pub fn new() -> GenericError {
        GenericError {
            error: None,
        }
    }
}


