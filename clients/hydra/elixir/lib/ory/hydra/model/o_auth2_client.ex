# NOTE: This file is auto generated by OpenAPI Generator 6.3.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.OAuth2Client do
  @moduledoc """
  OAuth 2.0 Clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
  """

  @derive [Poison.Encoder]
  defstruct [
    :allowed_cors_origins,
    :audience,
    :authorization_code_grant_access_token_lifespan,
    :authorization_code_grant_id_token_lifespan,
    :authorization_code_grant_refresh_token_lifespan,
    :backchannel_logout_session_required,
    :backchannel_logout_uri,
    :client_credentials_grant_access_token_lifespan,
    :client_id,
    :client_name,
    :client_secret,
    :client_secret_expires_at,
    :client_uri,
    :contacts,
    :created_at,
    :frontchannel_logout_session_required,
    :frontchannel_logout_uri,
    :grant_types,
    :implicit_grant_access_token_lifespan,
    :implicit_grant_id_token_lifespan,
    :jwks,
    :jwks_uri,
    :jwt_bearer_grant_access_token_lifespan,
    :logo_uri,
    :metadata,
    :owner,
    :policy_uri,
    :post_logout_redirect_uris,
    :redirect_uris,
    :refresh_token_grant_access_token_lifespan,
    :refresh_token_grant_id_token_lifespan,
    :refresh_token_grant_refresh_token_lifespan,
    :registration_access_token,
    :registration_client_uri,
    :request_object_signing_alg,
    :request_uris,
    :response_types,
    :scope,
    :sector_identifier_uri,
    :subject_type,
    :token_endpoint_auth_method,
    :token_endpoint_auth_signing_alg,
    :tos_uri,
    :updated_at,
    :userinfo_signed_response_alg
  ]

  @type t :: %__MODULE__{
    :allowed_cors_origins => [String.t] | nil,
    :audience => [String.t] | nil,
    :authorization_code_grant_access_token_lifespan => String.t | nil,
    :authorization_code_grant_id_token_lifespan => String.t | nil,
    :authorization_code_grant_refresh_token_lifespan => String.t | nil,
    :backchannel_logout_session_required => boolean() | nil,
    :backchannel_logout_uri => String.t | nil,
    :client_credentials_grant_access_token_lifespan => String.t | nil,
    :client_id => String.t | nil,
    :client_name => String.t | nil,
    :client_secret => String.t | nil,
    :client_secret_expires_at => integer() | nil,
    :client_uri => String.t | nil,
    :contacts => [String.t] | nil,
    :created_at => DateTime.t | nil,
    :frontchannel_logout_session_required => boolean() | nil,
    :frontchannel_logout_uri => String.t | nil,
    :grant_types => [String.t] | nil,
    :implicit_grant_access_token_lifespan => String.t | nil,
    :implicit_grant_id_token_lifespan => String.t | nil,
    :jwks => any | nil,
    :jwks_uri => String.t | nil,
    :jwt_bearer_grant_access_token_lifespan => String.t | nil,
    :logo_uri => String.t | nil,
    :metadata => any | nil,
    :owner => String.t | nil,
    :policy_uri => String.t | nil,
    :post_logout_redirect_uris => [String.t] | nil,
    :redirect_uris => [String.t] | nil,
    :refresh_token_grant_access_token_lifespan => String.t | nil,
    :refresh_token_grant_id_token_lifespan => String.t | nil,
    :refresh_token_grant_refresh_token_lifespan => String.t | nil,
    :registration_access_token => String.t | nil,
    :registration_client_uri => String.t | nil,
    :request_object_signing_alg => String.t | nil,
    :request_uris => [String.t] | nil,
    :response_types => [String.t] | nil,
    :scope => String.t | nil,
    :sector_identifier_uri => String.t | nil,
    :subject_type => String.t | nil,
    :token_endpoint_auth_method => String.t | nil,
    :token_endpoint_auth_signing_alg => String.t | nil,
    :tos_uri => String.t | nil,
    :updated_at => DateTime.t | nil,
    :userinfo_signed_response_alg => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.OAuth2Client do
  import Ory.Hydra.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:jwks, :any, :any, options)
    |> deserialize(:metadata, :any, :any, options)
  end
end

