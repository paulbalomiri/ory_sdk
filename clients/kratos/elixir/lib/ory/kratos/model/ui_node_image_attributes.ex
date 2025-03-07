# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.UiNodeImageAttributes do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :height,
    :id,
    :node_type,
    :src,
    :width
  ]

  @type t :: %__MODULE__{
    :height => integer(),
    :id => String.t,
    :node_type => String.t,
    :src => String.t,
    :width => integer()
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.UiNodeImageAttributes do
  def decode(value, _options) do
    value
  end
end

