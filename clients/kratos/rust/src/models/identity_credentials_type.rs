/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.9.0-alpha.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentityCredentialsType : and so on.

/// and so on.
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum IdentityCredentialsType {
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

impl ToString for IdentityCredentialsType {
    fn to_string(&self) -> String {
        match self {
            Self::Password => String::from("password"),
            Self::Totp => String::from("totp"),
            Self::Oidc => String::from("oidc"),
            Self::Webauthn => String::from("webauthn"),
            Self::LookupSecret => String::from("lookup_secret"),
        }
    }
}




