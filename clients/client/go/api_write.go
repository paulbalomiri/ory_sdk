/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.60
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type WriteApi interface {

	/*
	CreateRelationTuple Create a Relation Tuple

	Use this endpoint to create a relation tuple.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WriteApiCreateRelationTupleRequest
	*/
	CreateRelationTuple(ctx context.Context) WriteApiCreateRelationTupleRequest

	// CreateRelationTupleExecute executes the request
	//  @return RelationQuery
	CreateRelationTupleExecute(r WriteApiCreateRelationTupleRequest) (*RelationQuery, *http.Response, error)

	/*
	DeleteRelationTuples Delete Relation Tuples

	Use this endpoint to delete relation tuples

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WriteApiDeleteRelationTuplesRequest
	*/
	DeleteRelationTuples(ctx context.Context) WriteApiDeleteRelationTuplesRequest

	// DeleteRelationTuplesExecute executes the request
	DeleteRelationTuplesExecute(r WriteApiDeleteRelationTuplesRequest) (*http.Response, error)

	/*
	PatchRelationTuples Patch Multiple Relation Tuples

	Use this endpoint to patch one or more relation tuples.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WriteApiPatchRelationTuplesRequest
	*/
	PatchRelationTuples(ctx context.Context) WriteApiPatchRelationTuplesRequest

	// PatchRelationTuplesExecute executes the request
	PatchRelationTuplesExecute(r WriteApiPatchRelationTuplesRequest) (*http.Response, error)
}

// WriteApiService WriteApi service
type WriteApiService service

type WriteApiCreateRelationTupleRequest struct {
	ctx context.Context
	ApiService WriteApi
	relationQuery *RelationQuery
}

func (r WriteApiCreateRelationTupleRequest) RelationQuery(relationQuery RelationQuery) WriteApiCreateRelationTupleRequest {
	r.relationQuery = &relationQuery
	return r
}

func (r WriteApiCreateRelationTupleRequest) Execute() (*RelationQuery, *http.Response, error) {
	return r.ApiService.CreateRelationTupleExecute(r)
}

/*
CreateRelationTuple Create a Relation Tuple

Use this endpoint to create a relation tuple.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WriteApiCreateRelationTupleRequest
*/
func (a *WriteApiService) CreateRelationTuple(ctx context.Context) WriteApiCreateRelationTupleRequest {
	return WriteApiCreateRelationTupleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RelationQuery
func (a *WriteApiService) CreateRelationTupleExecute(r WriteApiCreateRelationTupleRequest) (*RelationQuery, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationQuery
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.CreateRelationTuple")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.relationQuery
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WriteApiDeleteRelationTuplesRequest struct {
	ctx context.Context
	ApiService WriteApi
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

// Namespace of the Relation Tuple
func (r WriteApiDeleteRelationTuplesRequest) Namespace(namespace string) WriteApiDeleteRelationTuplesRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relation Tuple
func (r WriteApiDeleteRelationTuplesRequest) Object(object string) WriteApiDeleteRelationTuplesRequest {
	r.object = &object
	return r
}

// Relation of the Relation Tuple
func (r WriteApiDeleteRelationTuplesRequest) Relation(relation string) WriteApiDeleteRelationTuplesRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relation Tuple
func (r WriteApiDeleteRelationTuplesRequest) SubjectId(subjectId string) WriteApiDeleteRelationTuplesRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r WriteApiDeleteRelationTuplesRequest) SubjectSetNamespace(subjectSetNamespace string) WriteApiDeleteRelationTuplesRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r WriteApiDeleteRelationTuplesRequest) SubjectSetObject(subjectSetObject string) WriteApiDeleteRelationTuplesRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r WriteApiDeleteRelationTuplesRequest) SubjectSetRelation(subjectSetRelation string) WriteApiDeleteRelationTuplesRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r WriteApiDeleteRelationTuplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRelationTuplesExecute(r)
}

/*
DeleteRelationTuples Delete Relation Tuples

Use this endpoint to delete relation tuples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WriteApiDeleteRelationTuplesRequest
*/
func (a *WriteApiService) DeleteRelationTuples(ctx context.Context) WriteApiDeleteRelationTuplesRequest {
	return WriteApiDeleteRelationTuplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WriteApiService) DeleteRelationTuplesExecute(r WriteApiDeleteRelationTuplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.DeleteRelationTuples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type WriteApiPatchRelationTuplesRequest struct {
	ctx context.Context
	ApiService WriteApi
	patchDelta *[]PatchDelta
}

func (r WriteApiPatchRelationTuplesRequest) PatchDelta(patchDelta []PatchDelta) WriteApiPatchRelationTuplesRequest {
	r.patchDelta = &patchDelta
	return r
}

func (r WriteApiPatchRelationTuplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchRelationTuplesExecute(r)
}

/*
PatchRelationTuples Patch Multiple Relation Tuples

Use this endpoint to patch one or more relation tuples.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WriteApiPatchRelationTuplesRequest
*/
func (a *WriteApiService) PatchRelationTuples(ctx context.Context) WriteApiPatchRelationTuplesRequest {
	return WriteApiPatchRelationTuplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WriteApiService) PatchRelationTuplesExecute(r WriteApiPatchRelationTuplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.PatchRelationTuples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchDelta
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
