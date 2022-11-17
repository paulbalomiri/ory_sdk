/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.3.0
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SelfServiceRecoveryCode {
    /// Expires At is the timestamp of when the recovery flow expires  The timestamp when the recovery link expires.
    #[serde(rename = "expires_at", skip_serializing_if = "Option::is_none")]
    pub expires_at: Option<String>,
    /// RecoveryCode is the code that can be used to recover the account
    #[serde(rename = "recovery_code")]
    pub recovery_code: String,
    /// RecoveryLink with flow  This link opens the recovery UI with an empty `code` field.
    #[serde(rename = "recovery_link")]
    pub recovery_link: String,
}


impl SelfServiceRecoveryCode {
    pub fn new(recovery_code: String, recovery_link: String) -> SelfServiceRecoveryCode {
        SelfServiceRecoveryCode {
                expires_at: None,
                recovery_code,
                recovery_link,
        }
    }
}


