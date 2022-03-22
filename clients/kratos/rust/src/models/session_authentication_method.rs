/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.9.0-alpha.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SessionAuthenticationMethod : A singular authenticator used during authentication / login.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SessionAuthenticationMethod {
    #[serde(rename = "aal", skip_serializing_if = "Option::is_none")]
    pub aal: Option<crate::models::AuthenticatorAssuranceLevel>,
    /// When the authentication challenge was completed.
    #[serde(rename = "completed_at", skip_serializing_if = "Option::is_none")]
    pub completed_at: Option<String>,
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<Method>,
}

impl SessionAuthenticationMethod {
    /// A singular authenticator used during authentication / login.
    pub fn new() -> SessionAuthenticationMethod {
        SessionAuthenticationMethod {
            aal: None,
            completed_at: None,
            method: None,
        }
    }
}

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Method {
    #[serde(rename = "link_recovery")]
    LinkRecovery,
    #[serde(rename = "password")]
    Password,
    #[serde(rename = "totp")]
    Totp,
    #[serde(rename = "oidc")]
    Oidc,
    #[serde(rename = "webauthn")]
    Webauthn,
    #[serde(rename = "lookup_secret")]
    LookupSecret,
}

