# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.GetVersion200Response do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :version
  ]

  @type t :: %__MODULE__{
    :version => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.GetVersion200Response do
  def decode(value, _options) do
    value
  end
end

