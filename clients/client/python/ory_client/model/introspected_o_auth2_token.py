"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v0.2.0-alpha.60
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from ory_client.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
    OpenApiModel
)
from ory_client.exceptions import ApiAttributeError



class IntrospectedOAuth2Token(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    validations = {
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        return {
            'active': (bool,),  # noqa: E501
            'aud': ([str],),  # noqa: E501
            'client_id': (str,),  # noqa: E501
            'exp': (int,),  # noqa: E501
            'ext': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)},),  # noqa: E501
            'iat': (int,),  # noqa: E501
            'iss': (str,),  # noqa: E501
            'nbf': (int,),  # noqa: E501
            'obfuscated_subject': (str,),  # noqa: E501
            'scope': (str,),  # noqa: E501
            'sub': (str,),  # noqa: E501
            'token_type': (str,),  # noqa: E501
            'token_use': (str,),  # noqa: E501
            'username': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'active': 'active',  # noqa: E501
        'aud': 'aud',  # noqa: E501
        'client_id': 'client_id',  # noqa: E501
        'exp': 'exp',  # noqa: E501
        'ext': 'ext',  # noqa: E501
        'iat': 'iat',  # noqa: E501
        'iss': 'iss',  # noqa: E501
        'nbf': 'nbf',  # noqa: E501
        'obfuscated_subject': 'obfuscated_subject',  # noqa: E501
        'scope': 'scope',  # noqa: E501
        'sub': 'sub',  # noqa: E501
        'token_type': 'token_type',  # noqa: E501
        'token_use': 'token_use',  # noqa: E501
        'username': 'username',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, active, *args, **kwargs):  # noqa: E501
        """IntrospectedOAuth2Token - a model defined in OpenAPI

        Args:
            active (bool): Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            aud ([str]): Audience contains a list of the token's intended audiences.. [optional]  # noqa: E501
            client_id (str): ID is aclient identifier for the OAuth 2.0 client that requested this token.. [optional]  # noqa: E501
            exp (int): Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.. [optional]  # noqa: E501
            ext ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): Extra is arbitrary data set by the session.. [optional]  # noqa: E501
            iat (int): Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.. [optional]  # noqa: E501
            iss (str): IssuerURL is a string representing the issuer of this token. [optional]  # noqa: E501
            nbf (int): NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.. [optional]  # noqa: E501
            obfuscated_subject (str): ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.. [optional]  # noqa: E501
            scope (str): Scope is a JSON string containing a space-separated list of scopes associated with this token.. [optional]  # noqa: E501
            sub (str): Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.. [optional]  # noqa: E501
            token_type (str): TokenType is the introspected token's type, typically `Bearer`.. [optional]  # noqa: E501
            token_use (str): TokenUse is the introspected token's use, for example `access_token` or `refresh_token`.. [optional]  # noqa: E501
            username (str): Username is a human-readable identifier for the resource owner who authorized this token.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', True)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.active = active
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, active, *args, **kwargs):  # noqa: E501
        """IntrospectedOAuth2Token - a model defined in OpenAPI

        Args:
            active (bool): Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            aud ([str]): Audience contains a list of the token's intended audiences.. [optional]  # noqa: E501
            client_id (str): ID is aclient identifier for the OAuth 2.0 client that requested this token.. [optional]  # noqa: E501
            exp (int): Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.. [optional]  # noqa: E501
            ext ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): Extra is arbitrary data set by the session.. [optional]  # noqa: E501
            iat (int): Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.. [optional]  # noqa: E501
            iss (str): IssuerURL is a string representing the issuer of this token. [optional]  # noqa: E501
            nbf (int): NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.. [optional]  # noqa: E501
            obfuscated_subject (str): ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.. [optional]  # noqa: E501
            scope (str): Scope is a JSON string containing a space-separated list of scopes associated with this token.. [optional]  # noqa: E501
            sub (str): Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.. [optional]  # noqa: E501
            token_type (str): TokenType is the introspected token's type, typically `Bearer`.. [optional]  # noqa: E501
            token_use (str): TokenUse is the introspected token's use, for example `access_token` or `refresh_token`.. [optional]  # noqa: E501
            username (str): Username is a human-readable identifier for the resource owner who authorized this token.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.active = active
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")
