# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Hydra.Model.AcceptOAuth2LoginRequest do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :acr,
    :amr,
    :context,
    :force_subject_identifier,
    :remember,
    :remember_for,
    :subject
  ]

  @type t :: %__MODULE__{
    :acr => String.t | nil,
    :amr => [String.t] | nil,
    :context => AnyType | nil,
    :force_subject_identifier => String.t | nil,
    :remember => boolean() | nil,
    :remember_for => integer() | nil,
    :subject => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.AcceptOAuth2LoginRequest do
  import Ory.Hydra.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:context, :struct, Ory.Hydra.Model.AnyType, options)
  end
end

