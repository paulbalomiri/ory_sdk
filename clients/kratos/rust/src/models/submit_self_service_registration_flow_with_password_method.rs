/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.17
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SubmitSelfServiceRegistrationFlowWithPasswordMethod : SubmitSelfServiceRegistrationFlowWithPasswordMethod is used to decode the registration form payload when using the password method.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceRegistrationFlowWithPasswordMethod {
    /// The CSRF Token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method to use  This field must be set to `password` when using the password method.
    #[serde(rename = "method")]
    pub method: String,
    /// Password to sign the user up with
    #[serde(rename = "password", skip_serializing_if = "Option::is_none")]
    pub password: Option<String>,
    /// The identity's traits
    #[serde(rename = "traits", skip_serializing_if = "Option::is_none")]
    pub traits: Option<serde_json::Value>,
}

impl SubmitSelfServiceRegistrationFlowWithPasswordMethod {
    /// SubmitSelfServiceRegistrationFlowWithPasswordMethod is used to decode the registration form payload when using the password method.
    pub fn new(method: String) -> SubmitSelfServiceRegistrationFlowWithPasswordMethod {
        SubmitSelfServiceRegistrationFlowWithPasswordMethod {
            csrf_token: None,
            method,
            password: None,
            traits: None,
        }
    }
}


