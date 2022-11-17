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
pub struct AdminCreateIdentityImportCredentialsOidc {
    #[serde(rename = "config", skip_serializing_if = "Option::is_none")]
    pub config: Option<Box<crate::models::AdminCreateIdentityImportCredentialsOidcConfig>>,
}

impl Default for AdminCreateIdentityImportCredentialsOidc {
    fn default() -> Self {
        Self::new()
    }
}

impl AdminCreateIdentityImportCredentialsOidc {
    pub fn new() -> AdminCreateIdentityImportCredentialsOidc {
        AdminCreateIdentityImportCredentialsOidc {
                config: None,
        }
    }
}


