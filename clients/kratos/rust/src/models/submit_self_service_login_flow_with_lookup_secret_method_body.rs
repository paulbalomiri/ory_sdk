/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.10.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// The lookup secret.
    #[serde(rename = "lookup_secret")]
    pub lookup_secret: String,
    /// Method should be set to \"lookup_secret\" when logging in using the lookup_secret strategy.
    #[serde(rename = "method")]
    pub method: String,
}


impl SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
    pub fn new(lookup_secret: String, method: String) -> SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
        SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
                csrf_token: None,
                lookup_secret,
                method,
        }
    }
}


