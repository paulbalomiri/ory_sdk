# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.AdminCreateIdentityImportCredentialsOidcProvider do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :provider,
    :subject
  ]

  @type t :: %__MODULE__{
    :provider => String.t,
    :subject => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.AdminCreateIdentityImportCredentialsOidcProvider do
  def decode(value, _options) do
    value
  end
end

