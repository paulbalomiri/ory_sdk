/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.60
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using Xunit;

using Ory.Client.Client;
using Ory.Client.Api;
// uncomment below to import models
//using Ory.Client.Model;

namespace Ory.Client.Test.Api
{
    /// <summary>
    ///  Class for testing NamespacesApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class NamespacesApiTests : IDisposable
    {
        private NamespacesApi instance;

        public NamespacesApiTests()
        {
            instance = new NamespacesApi();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of NamespacesApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsType' NamespacesApi
            //Assert.IsType<NamespacesApi>(instance);
        }

        /// <summary>
        /// Test GetNamespaces
        /// </summary>
        [Fact]
        public void GetNamespacesTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //var response = instance.GetNamespaces();
            //Assert.IsType<ClientGetNamespacesResponse>(response);
        }
    }
}
