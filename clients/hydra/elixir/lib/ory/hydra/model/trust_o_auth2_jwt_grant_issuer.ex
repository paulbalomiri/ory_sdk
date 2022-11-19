# NOTE: This file is auto generated by OpenAPI Generator 6.3.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.TrustOAuth2JwtGrantIssuer do
  @moduledoc """
  Trust OAuth2 JWT Bearer Grant Type Issuer Request Body
  """

  @derive [Poison.Encoder]
  defstruct [
    :allow_any_subject,
    :expires_at,
    :issuer,
    :jwk,
    :scope,
    :subject
  ]

  @type t :: %__MODULE__{
    :allow_any_subject => boolean() | nil,
    :expires_at => DateTime.t,
    :issuer => String.t,
    :jwk => Ory.Hydra.Model.JsonWebKey.t,
    :scope => [String.t],
    :subject => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.TrustOAuth2JwtGrantIssuer do
  import Ory.Hydra.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:jwk, :struct, Ory.Hydra.Model.JsonWebKey, options)
  end
end

