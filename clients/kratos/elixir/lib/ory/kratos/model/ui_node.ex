# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.UiNode do
  @moduledoc """
  Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.
  """

  @derive [Poison.Encoder]
  defstruct [
    :attributes,
    :group,
    :messages,
    :meta,
    :type
  ]

  @type t :: %__MODULE__{
    :attributes => Ory.Kratos.Model.UiNodeAttributes.t,
    :group => String.t,
    :messages => [Ory.Kratos.Model.UiText.t],
    :meta => Ory.Kratos.Model.UiNodeMeta.t,
    :type => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.UiNode do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:attributes, :struct, Ory.Kratos.Model.UiNodeAttributes, options)
    |> deserialize(:messages, :list, Ory.Kratos.Model.UiText, options)
    |> deserialize(:meta, :struct, Ory.Kratos.Model.UiNodeMeta, options)
  end
end

