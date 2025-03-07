<?php
/**
 * Session
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v0.2.0-alpha.60
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * Session Class Doc Comment
 *
 * @category Class
 * @description A Session
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class Session implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'session';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'active' => 'bool',
        'authenticatedAt' => '\DateTime',
        'authenticationMethods' => '\Ory\Client\Model\SessionAuthenticationMethod[]',
        'authenticatorAssuranceLevel' => '\Ory\Client\Model\AuthenticatorAssuranceLevel',
        'devices' => '\Ory\Client\Model\SessionDevice[]',
        'expiresAt' => '\DateTime',
        'id' => 'string',
        'identity' => '\Ory\Client\Model\Identity',
        'issuedAt' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'active' => null,
        'authenticatedAt' => 'date-time',
        'authenticationMethods' => null,
        'authenticatorAssuranceLevel' => null,
        'devices' => null,
        'expiresAt' => 'date-time',
        'id' => 'uuid',
        'identity' => null,
        'issuedAt' => 'date-time'
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'active' => 'active',
        'authenticatedAt' => 'authenticated_at',
        'authenticationMethods' => 'authentication_methods',
        'authenticatorAssuranceLevel' => 'authenticator_assurance_level',
        'devices' => 'devices',
        'expiresAt' => 'expires_at',
        'id' => 'id',
        'identity' => 'identity',
        'issuedAt' => 'issued_at'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'active' => 'setActive',
        'authenticatedAt' => 'setAuthenticatedAt',
        'authenticationMethods' => 'setAuthenticationMethods',
        'authenticatorAssuranceLevel' => 'setAuthenticatorAssuranceLevel',
        'devices' => 'setDevices',
        'expiresAt' => 'setExpiresAt',
        'id' => 'setId',
        'identity' => 'setIdentity',
        'issuedAt' => 'setIssuedAt'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'active' => 'getActive',
        'authenticatedAt' => 'getAuthenticatedAt',
        'authenticationMethods' => 'getAuthenticationMethods',
        'authenticatorAssuranceLevel' => 'getAuthenticatorAssuranceLevel',
        'devices' => 'getDevices',
        'expiresAt' => 'getExpiresAt',
        'id' => 'getId',
        'identity' => 'getIdentity',
        'issuedAt' => 'getIssuedAt'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['active'] = $data['active'] ?? null;
        $this->container['authenticatedAt'] = $data['authenticatedAt'] ?? null;
        $this->container['authenticationMethods'] = $data['authenticationMethods'] ?? null;
        $this->container['authenticatorAssuranceLevel'] = $data['authenticatorAssuranceLevel'] ?? null;
        $this->container['devices'] = $data['devices'] ?? null;
        $this->container['expiresAt'] = $data['expiresAt'] ?? null;
        $this->container['id'] = $data['id'] ?? null;
        $this->container['identity'] = $data['identity'] ?? null;
        $this->container['issuedAt'] = $data['issuedAt'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['id'] === null) {
            $invalidProperties[] = "'id' can't be null";
        }
        if ($this->container['identity'] === null) {
            $invalidProperties[] = "'identity' can't be null";
        }
        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active Active state. If false the session is no longer active.
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets authenticatedAt
     *
     * @return \DateTime|null
     */
    public function getAuthenticatedAt()
    {
        return $this->container['authenticatedAt'];
    }

    /**
     * Sets authenticatedAt
     *
     * @param \DateTime|null $authenticatedAt The Session Authentication Timestamp  When this session was authenticated at. If multi-factor authentication was used this is the time when the last factor was authenticated (e.g. the TOTP code challenge was completed).
     *
     * @return self
     */
    public function setAuthenticatedAt($authenticatedAt)
    {
        $this->container['authenticatedAt'] = $authenticatedAt;

        return $this;
    }

    /**
     * Gets authenticationMethods
     *
     * @return \Ory\Client\Model\SessionAuthenticationMethod[]|null
     */
    public function getAuthenticationMethods()
    {
        return $this->container['authenticationMethods'];
    }

    /**
     * Sets authenticationMethods
     *
     * @param \Ory\Client\Model\SessionAuthenticationMethod[]|null $authenticationMethods A list of authenticators which were used to authenticate the session.
     *
     * @return self
     */
    public function setAuthenticationMethods($authenticationMethods)
    {
        $this->container['authenticationMethods'] = $authenticationMethods;

        return $this;
    }

    /**
     * Gets authenticatorAssuranceLevel
     *
     * @return \Ory\Client\Model\AuthenticatorAssuranceLevel|null
     */
    public function getAuthenticatorAssuranceLevel()
    {
        return $this->container['authenticatorAssuranceLevel'];
    }

    /**
     * Sets authenticatorAssuranceLevel
     *
     * @param \Ory\Client\Model\AuthenticatorAssuranceLevel|null $authenticatorAssuranceLevel authenticatorAssuranceLevel
     *
     * @return self
     */
    public function setAuthenticatorAssuranceLevel($authenticatorAssuranceLevel)
    {
        $this->container['authenticatorAssuranceLevel'] = $authenticatorAssuranceLevel;

        return $this;
    }

    /**
     * Gets devices
     *
     * @return \Ory\Client\Model\SessionDevice[]|null
     */
    public function getDevices()
    {
        return $this->container['devices'];
    }

    /**
     * Sets devices
     *
     * @param \Ory\Client\Model\SessionDevice[]|null $devices Devices has history of all endpoints where the session was used
     *
     * @return self
     */
    public function setDevices($devices)
    {
        $this->container['devices'] = $devices;

        return $this;
    }

    /**
     * Gets expiresAt
     *
     * @return \DateTime|null
     */
    public function getExpiresAt()
    {
        return $this->container['expiresAt'];
    }

    /**
     * Sets expiresAt
     *
     * @param \DateTime|null $expiresAt The Session Expiry  When this session expires at.
     *
     * @return self
     */
    public function setExpiresAt($expiresAt)
    {
        $this->container['expiresAt'] = $expiresAt;

        return $this;
    }

    /**
     * Gets id
     *
     * @return string
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param string $id Session ID
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets identity
     *
     * @return \Ory\Client\Model\Identity
     */
    public function getIdentity()
    {
        return $this->container['identity'];
    }

    /**
     * Sets identity
     *
     * @param \Ory\Client\Model\Identity $identity identity
     *
     * @return self
     */
    public function setIdentity($identity)
    {
        $this->container['identity'] = $identity;

        return $this;
    }

    /**
     * Gets issuedAt
     *
     * @return \DateTime|null
     */
    public function getIssuedAt()
    {
        return $this->container['issuedAt'];
    }

    /**
     * Sets issuedAt
     *
     * @param \DateTime|null $issuedAt The Session Issuance Timestamp  When this session was issued at. Usually equal or close to `authenticated_at`.
     *
     * @return self
     */
    public function setIssuedAt($issuedAt)
    {
        $this->container['issuedAt'] = $issuedAt;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


