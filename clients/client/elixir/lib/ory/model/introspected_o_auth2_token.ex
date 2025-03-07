# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.IntrospectedOAuth2Token do
  @moduledoc """
  Introspection contains an access token's session data as specified by [IETF RFC 7662](https://tools.ietf.org/html/rfc7662)
  """

  @derive [Poison.Encoder]
  defstruct [
    :active,
    :aud,
    :client_id,
    :exp,
    :ext,
    :iat,
    :iss,
    :nbf,
    :obfuscated_subject,
    :scope,
    :sub,
    :token_type,
    :token_use,
    :username
  ]

  @type t :: %__MODULE__{
    :active => boolean(),
    :aud => [String.t] | nil,
    :client_id => String.t | nil,
    :exp => integer() | nil,
    :ext => %{optional(String.t) => AnyType} | nil,
    :iat => integer() | nil,
    :iss => String.t | nil,
    :nbf => integer() | nil,
    :obfuscated_subject => String.t | nil,
    :scope => String.t | nil,
    :sub => String.t | nil,
    :token_type => String.t | nil,
    :token_use => String.t | nil,
    :username => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.IntrospectedOAuth2Token do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:ext, :map, Ory.Model.AnyType, options)
  end
end

