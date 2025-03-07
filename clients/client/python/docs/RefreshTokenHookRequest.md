# RefreshTokenHookRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**client_id** | **str** | ClientID is the identifier of the OAuth 2.0 client. | [optional] 
**granted_audience** | **[str]** | GrantedAudience is the list of audiences granted to the OAuth 2.0 client. | [optional] 
**granted_scopes** | **[str]** | GrantedScopes is the list of scopes granted to the OAuth 2.0 client. | [optional] 
**requester** | [**OAuth2AccessRequest**](OAuth2AccessRequest.md) |  | [optional] 
**session** | [**OAuth2ConsentSession**](OAuth2ConsentSession.md) |  | [optional] 
**subject** | **str** | Subject is the identifier of the authenticated end-user. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


