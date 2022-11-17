# NOTE: This file is auto generated by OpenAPI Generator 6.0.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Kratos.Api.Metadata do
  @moduledoc """
  API calls for all endpoints tagged `Metadata`.
  """

  alias Ory.Kratos.Connection
  import Ory.Kratos.RequestBuilder

  @doc """
  Return Running Software Version.
  This endpoint returns the version of Ory Kratos.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of this service, the version will never refer to the cluster state, only to a single instance.

  ### Parameters

  - `connection` (Ory.Kratos.Connection): Connection to server
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, Ory.Kratos.Model.GetVersion200Response.t}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec get_version(Tesla.Env.client, keyword()) :: {:ok, Ory.Kratos.Model.GetVersion200Response.t} | {:error, Tesla.Env.t}
  def get_version(connection, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/version")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {200, %Ory.Kratos.Model.GetVersion200Response{}}
    ])
  end

  @doc """
  Check HTTP Server Status
  This endpoint returns a HTTP 200 status code when Ory Kratos is accepting incoming HTTP requests. This status does currently not include checks whether the database connection is working.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of this service, the health status will never refer to the cluster state, only to a single instance.

  ### Parameters

  - `connection` (Ory.Kratos.Connection): Connection to server
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, Ory.Kratos.Model.IsAlive200Response.t}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec is_alive(Tesla.Env.client, keyword()) :: {:ok, Ory.Kratos.Model.IsAlive200Response.t} | {:ok, Ory.Kratos.Model.GenericError.t} | {:error, Tesla.Env.t}
  def is_alive(connection, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/health/alive")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {200, %Ory.Kratos.Model.IsAlive200Response{}},
      {500, %Ory.Kratos.Model.GenericError{}}
    ])
  end

  @doc """
  Check HTTP Server and Database Status
  This endpoint returns a HTTP 200 status code when Ory Kratos is up running and the environment dependencies (e.g. the database) are responsive as well.  If the service supports TLS Edge Termination, this endpoint does not require the `X-Forwarded-Proto` header to be set.  Be aware that if you are running multiple nodes of Ory Kratos, the health status will never refer to the cluster state, only to a single instance.

  ### Parameters

  - `connection` (Ory.Kratos.Connection): Connection to server
  - `opts` (keyword): Optional parameters

  ### Returns

  - `{:ok, Ory.Kratos.Model.IsAlive200Response.t}` on success
  - `{:error, Tesla.Env.t}` on failure
  """
  @spec is_ready(Tesla.Env.client, keyword()) :: {:ok, Ory.Kratos.Model.IsAlive200Response.t} | {:ok, Ory.Kratos.Model.IsReady503Response.t} | {:error, Tesla.Env.t}
  def is_ready(connection, _opts \\ []) do
    request =
      %{}
      |> method(:get)
      |> url("/health/ready")
      |> Enum.into([])

    connection
    |> Connection.request(request)
    |> evaluate_response([
      {200, %Ory.Kratos.Model.IsAlive200Response{}},
      {503, %Ory.Kratos.Model.IsReady503Response{}}
    ])
  end
end
