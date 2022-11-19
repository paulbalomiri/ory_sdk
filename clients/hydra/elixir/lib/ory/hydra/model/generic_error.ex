# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

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
    |> deserialize(:details, :any, :any, options)
  end
end

