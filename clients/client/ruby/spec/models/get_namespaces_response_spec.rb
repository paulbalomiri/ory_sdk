=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v0.2.0-alpha.60
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::GetNamespacesResponse
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::GetNamespacesResponse do
  let(:instance) { OryClient::GetNamespacesResponse.new }

  describe 'test an instance of GetNamespacesResponse' do
    it 'should create an instance of GetNamespacesResponse' do
      expect(instance).to be_instance_of(OryClient::GetNamespacesResponse)
    end
  end
  describe 'test attribute "namespaces"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
