# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.SelfServiceRecoveryFlow do
  @moduledoc """
  This request is used when an identity wants to recover their account.  We recommend reading the [Account Recovery Documentation](../self-service/flows/password-reset-account-recovery)
  """

  @derive [Poison.Encoder]
  defstruct [
    :active,
    :expires_at,
    :id,
    :issued_at,
    :request_url,
    :return_to,
    :state,
    :type,
    :ui
  ]

  @type t :: %__MODULE__{
    :active => String.t | nil,
    :expires_at => DateTime.t,
    :id => String.t,
    :issued_at => DateTime.t,
    :request_url => String.t,
    :return_to => String.t | nil,
    :state => Ory.Kratos.Model.SelfServiceRecoveryFlowState.t,
    :type => String.t,
    :ui => Ory.Kratos.Model.UiContainer.t
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.SelfServiceRecoveryFlow do
  import Ory.Kratos.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:state, :struct, Ory.Kratos.Model.SelfServiceRecoveryFlowState, options)
    |> deserialize(:ui, :struct, Ory.Kratos.Model.UiContainer, options)
  end
end

