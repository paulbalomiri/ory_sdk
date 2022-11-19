# NOTE: This file is auto generated by OpenAPI Generator 6.3.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.AcceptOAuth2ConsentRequestSession do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :access_token,
    :id_token
  ]

  @type t :: %__MODULE__{
    :access_token => any | nil,
    :id_token => any | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.AcceptOAuth2ConsentRequestSession do
  import Ory.Hydra.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:access_token, :any, :any, options)
    |> deserialize(:id_token, :any, :any, options)
  end
end

