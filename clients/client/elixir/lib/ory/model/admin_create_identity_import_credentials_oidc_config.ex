# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.AdminCreateIdentityImportCredentialsOidcConfig do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :config,
    :providers
  ]

  @type t :: %__MODULE__{
    :config => Ory.Model.AdminCreateIdentityImportCredentialsPasswordConfig.t | nil,
    :providers => [Ory.Model.AdminCreateIdentityImportCredentialsOidcProvider.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.AdminCreateIdentityImportCredentialsOidcConfig do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:config, :struct, Ory.Model.AdminCreateIdentityImportCredentialsPasswordConfig, options)
    |> deserialize(:providers, :list, Ory.Model.AdminCreateIdentityImportCredentialsOidcProvider, options)
  end
end

