/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.9.0-alpha.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
    /// CSRFToken is the anti-CSRF token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.
    #[serde(rename = "method")]
    pub method: String,
    /// The identity's traits
    #[serde(rename = "traits")]
    pub traits: serde_json::Value,
    /// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
    #[serde(rename = "webauthn_register", skip_serializing_if = "Option::is_none")]
    pub webauthn_register: Option<String>,
    /// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
    #[serde(rename = "webauthn_register_displayname", skip_serializing_if = "Option::is_none")]
    pub webauthn_register_displayname: Option<String>,
}

impl SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
    pub fn new(method: String, traits: serde_json::Value) -> SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
        SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
            csrf_token: None,
            method,
            traits,
            webauthn_register: None,
            webauthn_register_displayname: None,
        }
    }
}


