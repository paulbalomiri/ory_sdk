# OryClient::AdminCreateSelfServiceRecoveryLinkBody

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **expires_in** | **String** | Link Expires In  The recovery link will expire at that point in time. Defaults to the configuration value of &#x60;selfservice.flows.recovery.request_lifespan&#x60;. | [optional] |
| **identity_id** | **String** | Identity to Recover  The identity&#39;s ID you wish to recover. |  |

## Example

```ruby
require 'ory-client'

instance = OryClient::AdminCreateSelfServiceRecoveryLinkBody.new(
  expires_in: null,
  identity_id: null
)
```

