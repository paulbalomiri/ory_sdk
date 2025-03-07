=begin
#Ory Keto API

#Documentation for all of Ory Keto's REST APIs. gRPC is documented separately. 

The version of the OpenAPI document: v0.10.0-alpha.0
Contact: hi@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'cgi'

module OryKetoClient
  class WriteApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Create a Relation Tuple
    # Use this endpoint to create a relation tuple.
    # @param [Hash] opts the optional parameters
    # @option opts [RelationQuery] :relation_query 
    # @return [RelationQuery]
    def create_relation_tuple(opts = {})
      data, _status_code, _headers = create_relation_tuple_with_http_info(opts)
      data
    end

    # Create a Relation Tuple
    # Use this endpoint to create a relation tuple.
    # @param [Hash] opts the optional parameters
    # @option opts [RelationQuery] :relation_query 
    # @return [Array<(RelationQuery, Integer, Hash)>] RelationQuery data, response status code and response headers
    def create_relation_tuple_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WriteApi.create_relation_tuple ...'
      end
      # resource path
      local_var_path = '/admin/relation-tuples'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'relation_query'])

      # return_type
      return_type = opts[:debug_return_type] || 'RelationQuery'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"WriteApi.create_relation_tuple",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WriteApi#create_relation_tuple\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Delete Relation Tuples
    # Use this endpoint to delete relation tuples
    # @param [Hash] opts the optional parameters
    # @option opts [String] :namespace Namespace of the Relation Tuple
    # @option opts [String] :object Object of the Relation Tuple
    # @option opts [String] :relation Relation of the Relation Tuple
    # @option opts [String] :subject_id SubjectID of the Relation Tuple
    # @option opts [String] :subject_set_namespace Namespace of the Subject Set
    # @option opts [String] :subject_set_object Object of the Subject Set
    # @option opts [String] :subject_set_relation Relation of the Subject Set
    # @return [nil]
    def delete_relation_tuples(opts = {})
      delete_relation_tuples_with_http_info(opts)
      nil
    end

    # Delete Relation Tuples
    # Use this endpoint to delete relation tuples
    # @param [Hash] opts the optional parameters
    # @option opts [String] :namespace Namespace of the Relation Tuple
    # @option opts [String] :object Object of the Relation Tuple
    # @option opts [String] :relation Relation of the Relation Tuple
    # @option opts [String] :subject_id SubjectID of the Relation Tuple
    # @option opts [String] :subject_set_namespace Namespace of the Subject Set
    # @option opts [String] :subject_set_object Object of the Subject Set
    # @option opts [String] :subject_set_relation Relation of the Subject Set
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def delete_relation_tuples_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WriteApi.delete_relation_tuples ...'
      end
      # resource path
      local_var_path = '/admin/relation-tuples'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'namespace'] = opts[:'namespace'] if !opts[:'namespace'].nil?
      query_params[:'object'] = opts[:'object'] if !opts[:'object'].nil?
      query_params[:'relation'] = opts[:'relation'] if !opts[:'relation'].nil?
      query_params[:'subject_id'] = opts[:'subject_id'] if !opts[:'subject_id'].nil?
      query_params[:'subject_set.namespace'] = opts[:'subject_set_namespace'] if !opts[:'subject_set_namespace'].nil?
      query_params[:'subject_set.object'] = opts[:'subject_set_object'] if !opts[:'subject_set_object'].nil?
      query_params[:'subject_set.relation'] = opts[:'subject_set_relation'] if !opts[:'subject_set_relation'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"WriteApi.delete_relation_tuples",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WriteApi#delete_relation_tuples\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Patch Multiple Relation Tuples
    # Use this endpoint to patch one or more relation tuples.
    # @param [Hash] opts the optional parameters
    # @option opts [Array<PatchDelta>] :patch_delta 
    # @return [nil]
    def patch_relation_tuples(opts = {})
      patch_relation_tuples_with_http_info(opts)
      nil
    end

    # Patch Multiple Relation Tuples
    # Use this endpoint to patch one or more relation tuples.
    # @param [Hash] opts the optional parameters
    # @option opts [Array<PatchDelta>] :patch_delta 
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def patch_relation_tuples_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: WriteApi.patch_relation_tuples ...'
      end
      # resource path
      local_var_path = '/admin/relation-tuples'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'patch_delta'])

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"WriteApi.patch_relation_tuples",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: WriteApi#patch_relation_tuples\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
