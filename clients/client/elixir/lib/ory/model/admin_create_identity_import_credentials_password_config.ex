# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.AdminCreateIdentityImportCredentialsPasswordConfig do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :hashed_password,
    :password
  ]

  @type t :: %__MODULE__{
    :hashed_password => String.t | nil,
    :password => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.AdminCreateIdentityImportCredentialsPasswordConfig do
  def decode(value, _options) do
    value
  end
end

