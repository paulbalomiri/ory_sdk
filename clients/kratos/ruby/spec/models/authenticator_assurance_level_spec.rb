=begin
#Ory Kratos API

#Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 

The version of the OpenAPI document: v0.10.1
Contact: hi@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryKratosClient::AuthenticatorAssuranceLevel
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryKratosClient::AuthenticatorAssuranceLevel do
  let(:instance) { OryKratosClient::AuthenticatorAssuranceLevel.new }

  describe 'test an instance of AuthenticatorAssuranceLevel' do
    it 'should create an instance of AuthenticatorAssuranceLevel' do
      expect(instance).to be_instance_of(OryKratosClient::AuthenticatorAssuranceLevel)
    end
  end
end
