<?php
/**
 * IdentityState
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Kratos\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * The version of the OpenAPI document: v0.7.6-alpha.3
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.2.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Kratos\Client\Model;
use \Ory\Kratos\Client\ObjectSerializer;

/**
 * IdentityState Class Doc Comment
 *
 * @category Class
 * @description The state can either be &#x60;active&#x60; or &#x60;inactive&#x60;.
 * @package  Ory\Kratos\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class IdentityState
{
    /**
     * Possible values of this enum
     */
    const ACTIVE = 'active';
    const INACTIVE = 'inactive';
    
    /**
     * Gets allowable values of the enum
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::ACTIVE,
            self::INACTIVE,
        ];
    }
}


