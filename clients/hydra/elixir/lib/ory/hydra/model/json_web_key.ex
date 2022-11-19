# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.JsonWebKey do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :alg,
    :crv,
    :d,
    :dp,
    :dq,
    :e,
    :k,
    :kid,
    :kty,
    :n,
    :p,
    :q,
    :qi,
    :use,
    :x,
    :x5c,
    :y
  ]

  @type t :: %__MODULE__{
    :alg => String.t,
    :crv => String.t | nil,
    :d => String.t | nil,
    :dp => String.t | nil,
    :dq => String.t | nil,
    :e => String.t | nil,
    :k => String.t | nil,
    :kid => String.t,
    :kty => String.t,
    :n => String.t | nil,
    :p => String.t | nil,
    :q => String.t | nil,
    :qi => String.t | nil,
    :use => String.t,
    :x => String.t | nil,
    :x5c => [String.t] | nil,
    :y => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.JsonWebKey do
  def decode(value, _options) do
    value
  end
end

