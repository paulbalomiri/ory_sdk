# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.Identity do
  @moduledoc """
  An identity can be a real human, a service, an IoT device - everything that can be described as an \"actor\" in a system.
  """

  @derive [Poison.Encoder]
  defstruct [
    :created_at,
    :credentials,
    :id,
    :metadata_admin,
    :metadata_public,
    :recovery_addresses,
    :schema_id,
    :schema_url,
    :state,
    :state_changed_at,
    :traits,
    :updated_at,
    :verifiable_addresses
  ]

  @type t :: %__MODULE__{
    :created_at => DateTime.t | nil,
    :credentials => %{optional(String.t) => Ory.Kratos.Model.IdentityCredentials.t} | nil,
    :id => String.t,
    :metadata_admin => AnyType | nil,
    :metadata_public => AnyType | nil,
    :recovery_addresses => [Ory.Kratos.Model.RecoveryAddress.t] | nil,
    :schema_id => String.t,
    :schema_url => String.t,
    :state => Ory.Kratos.Model.IdentityState.t | nil,
    :state_changed_at => DateTime.t | nil,
    :traits => AnyType | nil,
    :updated_at => DateTime.t | nil,
    :verifiable_addresses => [Ory.Kratos.Model.VerifiableIdentityAddress.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.Identity do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:credentials, :map, Ory.Kratos.Model.IdentityCredentials, options)
    |> deserialize(:metadata_admin, :struct, Ory.Kratos.Model.AnyType, options)
    |> deserialize(:metadata_public, :struct, Ory.Kratos.Model.AnyType, options)
    |> deserialize(:recovery_addresses, :list, Ory.Kratos.Model.RecoveryAddress, options)
    |> deserialize(:state, :struct, Ory.Kratos.Model.IdentityState, options)
    |> deserialize(:traits, :struct, Ory.Kratos.Model.AnyType, options)
    |> deserialize(:verifiable_addresses, :list, Ory.Kratos.Model.VerifiableIdentityAddress, options)
  end
end

