# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.UiNodeMeta do
  @moduledoc """
  This might include a label and other information that can optionally be used to render UIs.
  """

  @derive [Poison.Encoder]
  defstruct [
    :label
  ]

  @type t :: %__MODULE__{
    :label => Ory.Kratos.Model.UiText.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.UiNodeMeta do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:label, :struct, Ory.Kratos.Model.UiText, options)
  end
end

