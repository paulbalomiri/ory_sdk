/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.9.0-alpha.3
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceLoginFlowWithWebAuthnMethodBody {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Identifier is the email or username of the user trying to log in. This field is only required when using WebAuthn for passwordless login. When using WebAuthn for multi-factor authentication, it is not needed.
    #[serde(rename = "identifier", skip_serializing_if = "Option::is_none")]
    pub identifier: Option<String>,
    /// Method should be set to \"webAuthn\" when logging in using the WebAuthn strategy.
    #[serde(rename = "method")]
    pub method: String,
    /// Login a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
    #[serde(rename = "webauthn_login", skip_serializing_if = "Option::is_none")]
    pub webauthn_login: Option<String>,
}

impl SubmitSelfServiceLoginFlowWithWebAuthnMethodBody {
    pub fn new(method: String) -> SubmitSelfServiceLoginFlowWithWebAuthnMethodBody {
        SubmitSelfServiceLoginFlowWithWebAuthnMethodBody {
            csrf_token: None,
            identifier: None,
            method,
            webauthn_login: None,
        }
    }
}


