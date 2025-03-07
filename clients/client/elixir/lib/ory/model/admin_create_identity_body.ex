# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.AdminCreateIdentityBody do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :credentials,
    :metadata_admin,
    :metadata_public,
    :recovery_addresses,
    :schema_id,
    :state,
    :traits,
    :verifiable_addresses
  ]

  @type t :: %__MODULE__{
    :credentials => Ory.Model.AdminIdentityImportCredentials.t | nil,
    :metadata_admin => AnyType | nil,
    :metadata_public => AnyType | nil,
    :recovery_addresses => [Ory.Model.RecoveryIdentityAddress.t] | nil,
    :schema_id => String.t,
    :state => Ory.Model.IdentityState.t | nil,
    :traits => map(),
    :verifiable_addresses => [Ory.Model.VerifiableIdentityAddress.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.AdminCreateIdentityBody do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:credentials, :struct, Ory.Model.AdminIdentityImportCredentials, options)
    |> deserialize(:metadata_admin, :struct, Ory.Model.AnyType, options)
    |> deserialize(:metadata_public, :struct, Ory.Model.AnyType, options)
    |> deserialize(:recovery_addresses, :list, Ory.Model.RecoveryIdentityAddress, options)
    |> deserialize(:state, :struct, Ory.Model.IdentityState, options)
    |> deserialize(:verifiable_addresses, :list, Ory.Model.VerifiableIdentityAddress, options)
  end
end

