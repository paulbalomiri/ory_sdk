# IdentityApi

All URIs are relative to *https://playground.projects.oryapis.com*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**disableSession**](IdentityApi.md#disableSession) | **DELETE** /admin/sessions/{id} | Deactivate a Session |


<a name="disableSession"></a>
# **disableSession**
> disableSession(id)

Deactivate a Session

Calling this endpoint deactivates the specified session. Session data is not deleted.

### Example
```java
// Import classes:
import sh.ory.ApiClient;
import sh.ory.ApiException;
import sh.ory.Configuration;
import sh.ory.auth.*;
import sh.ory.models.*;
import sh.ory.api.IdentityApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://playground.projects.oryapis.com");
    
    // Configure HTTP bearer authorization: oryAccessToken
    HttpBearerAuth oryAccessToken = (HttpBearerAuth) defaultClient.getAuthentication("oryAccessToken");
    oryAccessToken.setBearerToken("BEARER TOKEN");

    IdentityApi apiInstance = new IdentityApi(defaultClient);
    String id = "id_example"; // String | ID is the session's ID.
    try {
      apiInstance.disableSession(id);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentityApi#disableSession");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| ID is the session&#39;s ID. | |

### Return type

null (empty response body)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 201. |  -  |
| **400** | jsonError |  -  |
| **401** | jsonError |  -  |
| **0** | jsonError |  -  |

