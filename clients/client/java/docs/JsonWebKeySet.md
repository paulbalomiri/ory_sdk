

# JsonWebKeySet

It is important that this model object is named JSONWebKeySet for \"swagger generate spec\" to generate only on definition of a JSONWebKeySet. Since one with the same name is previously defined as client.Client.JSONWebKeys and this one is last, this one will be effectively written in the swagger spec.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**keys** | [**List&lt;JsonWebKey&gt;**](JsonWebKey.md) | The value of the \&quot;keys\&quot; parameter is an array of JSON Web Key (JWK) values. By default, the order of the JWK values within the array does not imply an order of preference among them, although applications of JWK Sets can choose to assign a meaning to the order for their purposes, if desired. |  [optional] |



