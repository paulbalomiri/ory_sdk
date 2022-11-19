# NOTE: This file is auto generated by OpenAPI Generator 6.3.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Hydra.Model.RejectOAuth2Request do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :error,
    :error_debug,
    :error_description,
    :error_hint,
    :status_code
  ]

  @type t :: %__MODULE__{
    :error => String.t | nil,
    :error_debug => String.t | nil,
    :error_description => String.t | nil,
    :error_hint => String.t | nil,
    :status_code => integer() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Hydra.Model.RejectOAuth2Request do
  def decode(value, _options) do
    value
  end
end

