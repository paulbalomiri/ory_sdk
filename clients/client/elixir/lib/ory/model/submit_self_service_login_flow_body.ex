# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SubmitSelfServiceLoginFlowBody do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :csrf_token,
    :identifier,
    :method,
    :password,
    :password_identifier,
    :provider,
    :traits,
    :totp_code,
    :webauthn_login,
    :lookup_secret
  ]

  @type t :: %__MODULE__{
    :csrf_token => String.t | nil,
    :identifier => String.t,
    :method => String.t,
    :password => String.t,
    :password_identifier => String.t | nil,
    :provider => String.t,
    :traits => map() | nil,
    :totp_code => String.t,
    :webauthn_login => String.t | nil,
    :lookup_secret => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Model.SubmitSelfServiceLoginFlowBody do
  def decode(value, _options) do
    value
  end
end

