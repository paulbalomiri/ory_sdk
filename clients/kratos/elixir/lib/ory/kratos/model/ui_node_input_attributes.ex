# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.UiNodeInputAttributes do
  @moduledoc """
  InputAttributes represents the attributes of an input node
  """

  @derive [Poison.Encoder]
  defstruct [
    :disabled,
    :label,
    :name,
    :node_type,
    :onclick,
    :pattern,
    :required,
    :type,
    :value
  ]

  @type t :: %__MODULE__{
    :disabled => boolean(),
    :label => Ory.Kratos.Model.UiText.t | nil,
    :name => String.t,
    :node_type => String.t,
    :onclick => String.t | nil,
    :pattern => String.t | nil,
    :required => boolean() | nil,
    :type => String.t,
    :value => any() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.UiNodeInputAttributes do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:label, :struct, Ory.Kratos.Model.UiText, options)
    |> deserialize(:value, :any, :any, options)
  end
end

