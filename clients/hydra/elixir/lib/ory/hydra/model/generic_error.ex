# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Hydra.Model.GenericError do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :code,
    :debug,
    :details,
    :id,
    :message,
    :reason,
    :request,
    :status
  ]

  @type t :: %__MODULE__{
    :code => integer() | nil,
    :debug => String.t | nil,
    :details => AnyType | nil,
    :id => String.t | nil,
    :message => String.t,
    :reason => String.t | nil,
    :request => String.t | nil,
    :status => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.GenericError do
  import Ory.Hydra.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:details, :struct, Ory.Hydra.Model.AnyType, options)
  end
end

