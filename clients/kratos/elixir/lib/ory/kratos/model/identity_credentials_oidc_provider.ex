# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Model.IdentityCredentialsOidcProvider do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :initial_access_token,
    :initial_id_token,
    :initial_refresh_token,
    :provider,
    :subject
  ]

  @type t :: %__MODULE__{
    :initial_access_token => String.t | nil,
    :initial_id_token => String.t | nil,
    :initial_refresh_token => String.t | nil,
    :provider => String.t | nil,
    :subject => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Kratos.Model.IdentityCredentialsOidcProvider do
  def decode(value, _options) do
    value
  end
end

