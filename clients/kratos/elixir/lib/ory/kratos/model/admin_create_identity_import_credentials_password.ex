# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.AdminCreateIdentityImportCredentialsPassword do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :config
  ]

  @type t :: %__MODULE__{
    :config => Ory.Kratos.Model.AdminCreateIdentityImportCredentialsPasswordConfig.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.AdminCreateIdentityImportCredentialsPassword do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:config, :struct, Ory.Kratos.Model.AdminCreateIdentityImportCredentialsPasswordConfig, options)
  end
end

