=begin
#ORY Oathkeeper

#ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.

The version of the OpenAPI document: v0.39.4
Contact: hi@ory.am
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

# Common files
require 'ory-oathkeeper-client/api_client'
require 'ory-oathkeeper-client/api_error'
require 'ory-oathkeeper-client/version'
require 'ory-oathkeeper-client/configuration'

# Models
require 'ory-oathkeeper-client/models/get_well_known_json_web_keys500_response'
require 'ory-oathkeeper-client/models/health_not_ready_status'
require 'ory-oathkeeper-client/models/health_status'
require 'ory-oathkeeper-client/models/json_web_key'
require 'ory-oathkeeper-client/models/json_web_key_set'
require 'ory-oathkeeper-client/models/rule'
require 'ory-oathkeeper-client/models/rule_handler'
require 'ory-oathkeeper-client/models/rule_match'
require 'ory-oathkeeper-client/models/upstream'
require 'ory-oathkeeper-client/models/version'

# APIs
require 'ory-oathkeeper-client/api/api_api'

module OryOathkeeperClient
  class << self
    # Customize default settings for the SDK using block.
    #   OryOathkeeperClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
