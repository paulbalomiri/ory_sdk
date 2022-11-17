# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.IdentityCredentialsOidc do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :providers
  ]

  @type t :: %__MODULE__{
    :providers => [Ory.Kratos.Model.IdentityCredentialsOidcProvider.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.IdentityCredentialsOidc do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:providers, :list, Ory.Kratos.Model.IdentityCredentialsOidcProvider, options)
  end
end

