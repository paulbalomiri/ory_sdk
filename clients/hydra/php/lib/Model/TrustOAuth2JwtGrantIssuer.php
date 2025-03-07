<?php
/**
 * TrustOAuth2JwtGrantIssuer
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Hydra\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory Hydra API
 *
 * Documentation for all of Ory Hydra's APIs.
 *
 * The version of the OpenAPI document: v2.0.2
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Hydra\Client\Model;

use \ArrayAccess;
use \Ory\Hydra\Client\ObjectSerializer;

/**
 * TrustOAuth2JwtGrantIssuer Class Doc Comment
 *
 * @category Class
 * @description Trust OAuth2 JWT Bearer Grant Type Issuer Request Body
 * @package  Ory\Hydra\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class TrustOAuth2JwtGrantIssuer implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'trustOAuth2JwtGrantIssuer';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'allowAnySubject' => 'bool',
        'expiresAt' => '\DateTime',
        'issuer' => 'string',
        'jwk' => '\Ory\Hydra\Client\Model\JsonWebKey',
        'scope' => 'string[]',
        'subject' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'allowAnySubject' => null,
        'expiresAt' => 'date-time',
        'issuer' => null,
        'jwk' => null,
        'scope' => null,
        'subject' => null
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
        'allowAnySubject' => 'allow_any_subject',
        'expiresAt' => 'expires_at',
        'issuer' => 'issuer',
        'jwk' => 'jwk',
        'scope' => 'scope',
        'subject' => 'subject'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'allowAnySubject' => 'setAllowAnySubject',
        'expiresAt' => 'setExpiresAt',
        'issuer' => 'setIssuer',
        'jwk' => 'setJwk',
        'scope' => 'setScope',
        'subject' => 'setSubject'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'allowAnySubject' => 'getAllowAnySubject',
        'expiresAt' => 'getExpiresAt',
        'issuer' => 'getIssuer',
        'jwk' => 'getJwk',
        'scope' => 'getScope',
        'subject' => 'getSubject'
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
        $this->container['allowAnySubject'] = $data['allowAnySubject'] ?? null;
        $this->container['expiresAt'] = $data['expiresAt'] ?? null;
        $this->container['issuer'] = $data['issuer'] ?? null;
        $this->container['jwk'] = $data['jwk'] ?? null;
        $this->container['scope'] = $data['scope'] ?? null;
        $this->container['subject'] = $data['subject'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['expiresAt'] === null) {
            $invalidProperties[] = "'expiresAt' can't be null";
        }
        if ($this->container['issuer'] === null) {
            $invalidProperties[] = "'issuer' can't be null";
        }
        if ($this->container['jwk'] === null) {
            $invalidProperties[] = "'jwk' can't be null";
        }
        if ($this->container['scope'] === null) {
            $invalidProperties[] = "'scope' can't be null";
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
     * Gets allowAnySubject
     *
     * @return bool|null
     */
    public function getAllowAnySubject()
    {
        return $this->container['allowAnySubject'];
    }

    /**
     * Sets allowAnySubject
     *
     * @param bool|null $allowAnySubject The \"allow_any_subject\" indicates that the issuer is allowed to have any principal as the subject of the JWT.
     *
     * @return self
     */
    public function setAllowAnySubject($allowAnySubject)
    {
        $this->container['allowAnySubject'] = $allowAnySubject;

        return $this;
    }

    /**
     * Gets expiresAt
     *
     * @return \DateTime
     */
    public function getExpiresAt()
    {
        return $this->container['expiresAt'];
    }

    /**
     * Sets expiresAt
     *
     * @param \DateTime $expiresAt The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".
     *
     * @return self
     */
    public function setExpiresAt($expiresAt)
    {
        $this->container['expiresAt'] = $expiresAt;

        return $this;
    }

    /**
     * Gets issuer
     *
     * @return string
     */
    public function getIssuer()
    {
        return $this->container['issuer'];
    }

    /**
     * Sets issuer
     *
     * @param string $issuer The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).
     *
     * @return self
     */
    public function setIssuer($issuer)
    {
        $this->container['issuer'] = $issuer;

        return $this;
    }

    /**
     * Gets jwk
     *
     * @return \Ory\Hydra\Client\Model\JsonWebKey
     */
    public function getJwk()
    {
        return $this->container['jwk'];
    }

    /**
     * Sets jwk
     *
     * @param \Ory\Hydra\Client\Model\JsonWebKey $jwk jwk
     *
     * @return self
     */
    public function setJwk($jwk)
    {
        $this->container['jwk'] = $jwk;

        return $this;
    }

    /**
     * Gets scope
     *
     * @return string[]
     */
    public function getScope()
    {
        return $this->container['scope'];
    }

    /**
     * Sets scope
     *
     * @param string[] $scope The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
     *
     * @return self
     */
    public function setScope($scope)
    {
        $this->container['scope'] = $scope;

        return $this;
    }

    /**
     * Gets subject
     *
     * @return string|null
     */
    public function getSubject()
    {
        return $this->container['subject'];
    }

    /**
     * Sets subject
     *
     * @param string|null $subject The \"subject\" identifies the principal that is the subject of the JWT.
     *
     * @return self
     */
    public function setSubject($subject)
    {
        $this->container['subject'] = $subject;

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


