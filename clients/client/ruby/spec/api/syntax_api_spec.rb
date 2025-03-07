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

# Unit tests for OryClient::SyntaxApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'SyntaxApi' do
  before do
    # run before each test
    @api_instance = OryClient::SyntaxApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of SyntaxApi' do
    it 'should create an instance of SyntaxApi' do
      expect(@api_instance).to be_instance_of(OryClient::SyntaxApi)
    end
  end

  # unit tests for check_opl_syntax
  # Check the syntax of an OPL file
  # The OPL file is expected in the body of the request.
  # @param body the OPL content to check
  # @param [Hash] opts the optional parameters
  # @return [PostCheckOplSyntaxResponse]
  describe 'check_opl_syntax test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
