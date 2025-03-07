# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.UiNodeAnchorAttributes do
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
    :title => Ory.Model.UiText.t
  }
end

defimpl Poison.Decoder, for: Ory.Model.UiNodeAnchorAttributes do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:title, :struct, Ory.Model.UiText, options)
  end
end

