# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule OryAPIs.Model.AdminCreateIdentityImportCredentialsOidcConfig do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :config,
    :providers
  ]

  @type t :: %__MODULE__{
    :config => OryAPIs.Model.AdminCreateIdentityImportCredentialsPasswordConfig.t | nil,
    :providers => [OryAPIs.Model.AdminCreateIdentityImportCredentialsOidcProvider.t] | nil
  }
end

defimpl Poison.Decoder, for: OryAPIs.Model.AdminCreateIdentityImportCredentialsOidcConfig do
  import OryAPIs.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:config, :struct, OryAPIs.Model.AdminCreateIdentityImportCredentialsPasswordConfig, options)
    |> deserialize(:providers, :list, OryAPIs.Model.AdminCreateIdentityImportCredentialsOidcProvider, options)
  end
end

