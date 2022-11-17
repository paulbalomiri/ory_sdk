# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Kratos.Model.Session do
  @moduledoc """
  A Session
  """

  @derive [Poison.Encoder]
  defstruct [
    :active,
    :authenticated_at,
    :authentication_methods,
    :authenticator_assurance_level,
    :expires_at,
    :id,
    :identity,
    :issued_at
  ]

  @type t :: %__MODULE__{
    :active => boolean() | nil,
    :authenticated_at => DateTime.t | nil,
    :authentication_methods => [Ory.Kratos.Model.SessionAuthenticationMethod.t] | nil,
    :authenticator_assurance_level => Ory.Kratos.Model.AuthenticatorAssuranceLevel.t | nil,
    :expires_at => DateTime.t | nil,
    :id => String.t,
    :identity => Ory.Kratos.Model.Identity.t,
    :issued_at => DateTime.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.Session do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:authentication_methods, :list, Ory.Kratos.Model.SessionAuthenticationMethod, options)
    |> deserialize(:authenticator_assurance_level, :struct, Ory.Kratos.Model.AuthenticatorAssuranceLevel, options)
    |> deserialize(:identity, :struct, Ory.Kratos.Model.Identity, options)
  end
end

