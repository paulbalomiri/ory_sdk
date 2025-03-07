# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.CreateCustomHostnameBody do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :cookie_domain,
    :cors_allowed_origins,
    :cors_enabled,
    :hostname
  ]

  @type t :: %__MODULE__{
    :cookie_domain => String.t | nil,
    :cors_allowed_origins => [String.t] | nil,
    :cors_enabled => boolean() | nil,
    :hostname => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.CreateCustomHostnameBody do
  def decode(value, _options) do
    value
  end
end

