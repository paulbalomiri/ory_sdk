# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.OidcConfiguration do
  @moduledoc """
  It includes links to several endpoints (for example `/oauth2/token`) and exposes information on supported signature algorithms among others.
  """

  @derive [Poison.Encoder]
  defstruct [
    :authorization_endpoint,
    :backchannel_logout_session_supported,
    :backchannel_logout_supported,
    :claims_parameter_supported,
    :claims_supported,
    :code_challenge_methods_supported,
    :end_session_endpoint,
    :frontchannel_logout_session_supported,
    :frontchannel_logout_supported,
    :grant_types_supported,
    :id_token_signed_response_alg,
    :id_token_signing_alg_values_supported,
    :issuer,
    :jwks_uri,
    :registration_endpoint,
    :request_object_signing_alg_values_supported,
    :request_parameter_supported,
    :request_uri_parameter_supported,
    :require_request_uri_registration,
    :response_modes_supported,
    :response_types_supported,
    :revocation_endpoint,
    :scopes_supported,
    :subject_types_supported,
    :token_endpoint,
    :token_endpoint_auth_methods_supported,
    :userinfo_endpoint,
    :userinfo_signed_response_alg,
    :userinfo_signing_alg_values_supported
  ]

  @type t :: %__MODULE__{
    :authorization_endpoint => String.t,
    :backchannel_logout_session_supported => boolean() | nil,
    :backchannel_logout_supported => boolean() | nil,
    :claims_parameter_supported => boolean() | nil,
    :claims_supported => [String.t] | nil,
    :code_challenge_methods_supported => [String.t] | nil,
    :end_session_endpoint => String.t | nil,
    :frontchannel_logout_session_supported => boolean() | nil,
    :frontchannel_logout_supported => boolean() | nil,
    :grant_types_supported => [String.t] | nil,
    :id_token_signed_response_alg => [String.t],
    :id_token_signing_alg_values_supported => [String.t],
    :issuer => String.t,
    :jwks_uri => String.t,
    :registration_endpoint => String.t | nil,
    :request_object_signing_alg_values_supported => [String.t] | nil,
    :request_parameter_supported => boolean() | nil,
    :request_uri_parameter_supported => boolean() | nil,
    :require_request_uri_registration => boolean() | nil,
    :response_modes_supported => [String.t] | nil,
    :response_types_supported => [String.t],
    :revocation_endpoint => String.t | nil,
    :scopes_supported => [String.t] | nil,
    :subject_types_supported => [String.t],
    :token_endpoint => String.t,
    :token_endpoint_auth_methods_supported => [String.t] | nil,
    :userinfo_endpoint => String.t | nil,
    :userinfo_signed_response_alg => [String.t],
    :userinfo_signing_alg_values_supported => [String.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.OidcConfiguration do
  def decode(value, _options) do
    value
  end
end

