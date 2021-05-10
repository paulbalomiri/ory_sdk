/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.4
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// LoginFlow : This object represents a login flow. A login flow is initiated at the \"Initiate Login API / Browser Flow\" endpoint by a client.  Once a login flow is completed successfully, a session cookie or session token will be issued.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct LoginFlow {
    /// and so on.
    #[serde(rename = "active", skip_serializing_if = "Option::is_none")]
    pub active: Option<String>,
    /// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
    #[serde(rename = "expires_at")]
    pub expires_at: String,
    /// Forced stores whether this login flow should enforce re-authentication.
    #[serde(rename = "forced", skip_serializing_if = "Option::is_none")]
    pub forced: Option<bool>,
    #[serde(rename = "id")]
    pub id: String,
    /// IssuedAt is the time (UTC) when the flow started.
    #[serde(rename = "issued_at")]
    pub issued_at: String,
    /// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
    #[serde(rename = "request_url")]
    pub request_url: String,
    /// The flow type can either be `api` or `browser`.
    #[serde(rename = "type")]
    pub _type: String,
    #[serde(rename = "ui")]
    pub ui: Box<crate::models::UiContainer>,
}

impl LoginFlow {
    /// This object represents a login flow. A login flow is initiated at the \"Initiate Login API / Browser Flow\" endpoint by a client.  Once a login flow is completed successfully, a session cookie or session token will be issued.
    pub fn new(expires_at: String, id: String, issued_at: String, request_url: String, _type: String, ui: crate::models::UiContainer) -> LoginFlow {
        LoginFlow {
            active: None,
            expires_at,
            forced: None,
            id,
            issued_at,
            request_url,
            _type,
            ui: Box::new(ui),
        }
    }
}


