# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Hydra.Model.TokenPaginationHeaders do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :link,
    :"x-total-count"
  ]

  @type t :: %__MODULE__{
    :link => String.t | nil,
    :"x-total-count" => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.TokenPaginationHeaders do
  def decode(value, _options) do
    value
  end
end

