# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.AdminCreateSelfServiceRecoveryCodeBody do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :expires_in,
    :identity_id
  ]

  @type t :: %__MODULE__{
    :expires_in => String.t | nil,
    :identity_id => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Model.AdminCreateSelfServiceRecoveryCodeBody do
  def decode(value, _options) do
    value
  end
end

