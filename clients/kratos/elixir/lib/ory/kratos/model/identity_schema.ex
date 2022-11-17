# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.IdentitySchema do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :id,
    :schema
  ]

  @type t :: %__MODULE__{
    :id => String.t | nil,
    :schema => map() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.IdentitySchema do
  def decode(value, _options) do
    value
  end
end

