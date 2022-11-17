"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v0.3.0
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.o_auth2_client import OAuth2Client
from ory_client.model.open_id_connect_context import OpenIDConnectContext
globals()['OAuth2Client'] = OAuth2Client
globals()['OpenIDConnectContext'] = OpenIDConnectContext
from ory_client.model.login_request import LoginRequest


class TestLoginRequest(unittest.TestCase):
    """LoginRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testLoginRequest(self):
        """Test LoginRequest"""
        # FIXME: construct object with mandatory attributes with example values
        # model = LoginRequest()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
