# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.IsReady200Response do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :status
  ]

  @type t :: %__MODULE__{
    :status => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.IsReady200Response do
  def decode(value, _options) do
    value
  end
end

