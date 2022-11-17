# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Kratos.Model.UiNodeAnchorAttributes do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :href,
    :id,
    :node_type,
    :title
  ]

  @type t :: %__MODULE__{
    :href => String.t,
    :id => String.t,
    :node_type => String.t,
    :title => Ory.Kratos.Model.UiText.t
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.UiNodeAnchorAttributes do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:title, :struct, Ory.Kratos.Model.UiText, options)
  end
end

