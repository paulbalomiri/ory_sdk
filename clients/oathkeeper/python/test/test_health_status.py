"""
    ORY Oathkeeper

    ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.  # noqa: E501

    The version of the OpenAPI document: v0.39.4
    Contact: hi@ory.am
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_oathkeeper_client
from ory_oathkeeper_client.model.health_status import HealthStatus


class TestHealthStatus(unittest.TestCase):
    """HealthStatus unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testHealthStatus(self):
        """Test HealthStatus"""
        # FIXME: construct object with mandatory attributes with example values
        # model = HealthStatus()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
