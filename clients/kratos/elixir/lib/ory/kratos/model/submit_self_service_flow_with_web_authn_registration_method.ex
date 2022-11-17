# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.SubmitSelfServiceFlowWithWebAuthnRegistrationMethod do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :webauthn_register,
    :webauthn_register_displayname
  ]

  @type t :: %__MODULE__{
    :webauthn_register => String.t | nil,
    :webauthn_register_displayname => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.SubmitSelfServiceFlowWithWebAuthnRegistrationMethod do
  def decode(value, _options) do
    value
  end
end

