# OryClient::IdentityApi

All URIs are relative to *https://playground.projects.oryapis.com*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**disable_session**](IdentityApi.md#disable_session) | **DELETE** /admin/sessions/{id} | Deactivate a Session |


## disable_session

> disable_session(id)

Deactivate a Session

Calling this endpoint deactivates the specified session. Session data is not deleted.

### Examples

```ruby
require 'time'
require 'ory-client'
# setup authorization
OryClient.configure do |config|
  # Configure Bearer authorization: oryAccessToken
  config.access_token = 'YOUR_BEARER_TOKEN'
end

api_instance = OryClient::IdentityApi.new
id = 'id_example' # String | ID is the session's ID.

begin
  # Deactivate a Session
  api_instance.disable_session(id)
rescue OryClient::ApiError => e
  puts "Error when calling IdentityApi->disable_session: #{e}"
end
```

#### Using the disable_session_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> disable_session_with_http_info(id)

```ruby
begin
  # Deactivate a Session
  data, status_code, headers = api_instance.disable_session_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OryClient::ApiError => e
  puts "Error when calling IdentityApi->disable_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | ID is the session&#39;s ID. |  |

### Return type

nil (empty response body)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

