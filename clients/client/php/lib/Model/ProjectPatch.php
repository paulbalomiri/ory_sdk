<?php
/**
 * ProjectPatch
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
 * The version of the OpenAPI document: v0.0.1-alpha.37
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.2.1
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
 * ProjectPatch Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class ProjectPatch implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'projectPatch';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'apiUrl' => 'string',
        'applicationUrl' => 'string',
        'defaultIdentitySchemaUrl' => 'string',
        'errorUiUrl' => 'string',
        'kratosCustomSchemaId' => 'string',
        'loginUiUrl' => 'string',
        'lookupSecret' => '\Ory\Client\Model\ProjectLookupSecretConfig',
        'name' => 'string',
        'password' => '\Ory\Client\Model\ProjectPasswordConfig',
        'recovery' => '\Ory\Client\Model\ProjectRecoveryConfig',
        'recoveryUiUrl' => 'string',
        'redirectionConfig' => '\Ory\Client\Model\RedirectionConfig',
        'registrationUiUrl' => 'string',
        'sessionAfterSignUp' => 'bool',
        'sessionSoft2fa' => 'bool',
        'settingsPrivilegedSessionMaxAgeSeconds' => 'int',
        'settingsSoft2fa' => 'bool',
        'settingsUiUrl' => 'string',
        'totp' => '\Ory\Client\Model\ProjectTotpConfig',
        'verification' => '\Ory\Client\Model\ProjectVerificationConfig',
        'verificationUiUrl' => 'string',
        'webauthn' => '\Ory\Client\Model\ProjectWebAuthnConfig'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'apiUrl' => null,
        'applicationUrl' => null,
        'defaultIdentitySchemaUrl' => null,
        'errorUiUrl' => null,
        'kratosCustomSchemaId' => null,
        'loginUiUrl' => null,
        'lookupSecret' => null,
        'name' => null,
        'password' => null,
        'recovery' => null,
        'recoveryUiUrl' => null,
        'redirectionConfig' => null,
        'registrationUiUrl' => null,
        'sessionAfterSignUp' => null,
        'sessionSoft2fa' => null,
        'settingsPrivilegedSessionMaxAgeSeconds' => 'int64',
        'settingsSoft2fa' => null,
        'settingsUiUrl' => null,
        'totp' => null,
        'verification' => null,
        'verificationUiUrl' => null,
        'webauthn' => null
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
        'apiUrl' => 'api_url',
        'applicationUrl' => 'application_url',
        'defaultIdentitySchemaUrl' => 'default_identity_schema_url',
        'errorUiUrl' => 'error_ui_url',
        'kratosCustomSchemaId' => 'kratos_custom_schema_id',
        'loginUiUrl' => 'login_ui_url',
        'lookupSecret' => 'lookup_secret',
        'name' => 'name',
        'password' => 'password',
        'recovery' => 'recovery',
        'recoveryUiUrl' => 'recovery_ui_url',
        'redirectionConfig' => 'redirection_config',
        'registrationUiUrl' => 'registration_ui_url',
        'sessionAfterSignUp' => 'session_after_sign_up',
        'sessionSoft2fa' => 'session_soft_2fa',
        'settingsPrivilegedSessionMaxAgeSeconds' => 'settings_privileged_session_max_age_seconds',
        'settingsSoft2fa' => 'settings_soft_2fa',
        'settingsUiUrl' => 'settings_ui_url',
        'totp' => 'totp',
        'verification' => 'verification',
        'verificationUiUrl' => 'verification_ui_url',
        'webauthn' => 'webauthn'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'apiUrl' => 'setApiUrl',
        'applicationUrl' => 'setApplicationUrl',
        'defaultIdentitySchemaUrl' => 'setDefaultIdentitySchemaUrl',
        'errorUiUrl' => 'setErrorUiUrl',
        'kratosCustomSchemaId' => 'setKratosCustomSchemaId',
        'loginUiUrl' => 'setLoginUiUrl',
        'lookupSecret' => 'setLookupSecret',
        'name' => 'setName',
        'password' => 'setPassword',
        'recovery' => 'setRecovery',
        'recoveryUiUrl' => 'setRecoveryUiUrl',
        'redirectionConfig' => 'setRedirectionConfig',
        'registrationUiUrl' => 'setRegistrationUiUrl',
        'sessionAfterSignUp' => 'setSessionAfterSignUp',
        'sessionSoft2fa' => 'setSessionSoft2fa',
        'settingsPrivilegedSessionMaxAgeSeconds' => 'setSettingsPrivilegedSessionMaxAgeSeconds',
        'settingsSoft2fa' => 'setSettingsSoft2fa',
        'settingsUiUrl' => 'setSettingsUiUrl',
        'totp' => 'setTotp',
        'verification' => 'setVerification',
        'verificationUiUrl' => 'setVerificationUiUrl',
        'webauthn' => 'setWebauthn'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'apiUrl' => 'getApiUrl',
        'applicationUrl' => 'getApplicationUrl',
        'defaultIdentitySchemaUrl' => 'getDefaultIdentitySchemaUrl',
        'errorUiUrl' => 'getErrorUiUrl',
        'kratosCustomSchemaId' => 'getKratosCustomSchemaId',
        'loginUiUrl' => 'getLoginUiUrl',
        'lookupSecret' => 'getLookupSecret',
        'name' => 'getName',
        'password' => 'getPassword',
        'recovery' => 'getRecovery',
        'recoveryUiUrl' => 'getRecoveryUiUrl',
        'redirectionConfig' => 'getRedirectionConfig',
        'registrationUiUrl' => 'getRegistrationUiUrl',
        'sessionAfterSignUp' => 'getSessionAfterSignUp',
        'sessionSoft2fa' => 'getSessionSoft2fa',
        'settingsPrivilegedSessionMaxAgeSeconds' => 'getSettingsPrivilegedSessionMaxAgeSeconds',
        'settingsSoft2fa' => 'getSettingsSoft2fa',
        'settingsUiUrl' => 'getSettingsUiUrl',
        'totp' => 'getTotp',
        'verification' => 'getVerification',
        'verificationUiUrl' => 'getVerificationUiUrl',
        'webauthn' => 'getWebauthn'
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
        $this->container['apiUrl'] = $data['apiUrl'] ?? null;
        $this->container['applicationUrl'] = $data['applicationUrl'] ?? null;
        $this->container['defaultIdentitySchemaUrl'] = $data['defaultIdentitySchemaUrl'] ?? null;
        $this->container['errorUiUrl'] = $data['errorUiUrl'] ?? null;
        $this->container['kratosCustomSchemaId'] = $data['kratosCustomSchemaId'] ?? null;
        $this->container['loginUiUrl'] = $data['loginUiUrl'] ?? null;
        $this->container['lookupSecret'] = $data['lookupSecret'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['password'] = $data['password'] ?? null;
        $this->container['recovery'] = $data['recovery'] ?? null;
        $this->container['recoveryUiUrl'] = $data['recoveryUiUrl'] ?? null;
        $this->container['redirectionConfig'] = $data['redirectionConfig'] ?? null;
        $this->container['registrationUiUrl'] = $data['registrationUiUrl'] ?? null;
        $this->container['sessionAfterSignUp'] = $data['sessionAfterSignUp'] ?? null;
        $this->container['sessionSoft2fa'] = $data['sessionSoft2fa'] ?? null;
        $this->container['settingsPrivilegedSessionMaxAgeSeconds'] = $data['settingsPrivilegedSessionMaxAgeSeconds'] ?? null;
        $this->container['settingsSoft2fa'] = $data['settingsSoft2fa'] ?? null;
        $this->container['settingsUiUrl'] = $data['settingsUiUrl'] ?? null;
        $this->container['totp'] = $data['totp'] ?? null;
        $this->container['verification'] = $data['verification'] ?? null;
        $this->container['verificationUiUrl'] = $data['verificationUiUrl'] ?? null;
        $this->container['webauthn'] = $data['webauthn'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['errorUiUrl'] === null) {
            $invalidProperties[] = "'errorUiUrl' can't be null";
        }
        if ($this->container['loginUiUrl'] === null) {
            $invalidProperties[] = "'loginUiUrl' can't be null";
        }
        if ($this->container['lookupSecret'] === null) {
            $invalidProperties[] = "'lookupSecret' can't be null";
        }
        if ($this->container['password'] === null) {
            $invalidProperties[] = "'password' can't be null";
        }
        if ($this->container['recovery'] === null) {
            $invalidProperties[] = "'recovery' can't be null";
        }
        if ($this->container['recoveryUiUrl'] === null) {
            $invalidProperties[] = "'recoveryUiUrl' can't be null";
        }
        if ($this->container['registrationUiUrl'] === null) {
            $invalidProperties[] = "'registrationUiUrl' can't be null";
        }
        if ($this->container['sessionAfterSignUp'] === null) {
            $invalidProperties[] = "'sessionAfterSignUp' can't be null";
        }
        if ($this->container['sessionSoft2fa'] === null) {
            $invalidProperties[] = "'sessionSoft2fa' can't be null";
        }
        if ($this->container['settingsPrivilegedSessionMaxAgeSeconds'] === null) {
            $invalidProperties[] = "'settingsPrivilegedSessionMaxAgeSeconds' can't be null";
        }
        if ($this->container['settingsSoft2fa'] === null) {
            $invalidProperties[] = "'settingsSoft2fa' can't be null";
        }
        if ($this->container['settingsUiUrl'] === null) {
            $invalidProperties[] = "'settingsUiUrl' can't be null";
        }
        if ($this->container['totp'] === null) {
            $invalidProperties[] = "'totp' can't be null";
        }
        if ($this->container['verification'] === null) {
            $invalidProperties[] = "'verification' can't be null";
        }
        if ($this->container['verificationUiUrl'] === null) {
            $invalidProperties[] = "'verificationUiUrl' can't be null";
        }
        if ($this->container['webauthn'] === null) {
            $invalidProperties[] = "'webauthn' can't be null";
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
     * Gets apiUrl
     *
     * @return string|null
     */
    public function getApiUrl()
    {
        return $this->container['apiUrl'];
    }

    /**
     * Sets apiUrl
     *
     * @param string|null $apiUrl URL of Project API  In the future will allow to override Kratos project URL. If set, error is returned. If left empty our hosted URL is used
     *
     * @return self
     */
    public function setApiUrl($apiUrl)
    {
        $this->container['apiUrl'] = $apiUrl;

        return $this;
    }

    /**
     * Gets applicationUrl
     *
     * @return string|null
     */
    public function getApplicationUrl()
    {
        return $this->container['applicationUrl'];
    }

    /**
     * Sets applicationUrl
     *
     * @param string|null $applicationUrl URL where your application is available.  Your users will be redirected to this URL when they successfully complete a login, logout, verification, recovery, or registration flow. More fine-grained redirection patterns are available for the individual flows (login, registration, ...).  If left empty, our hosted URL will be used instead. It contains technical information on the identities and is useful when developing a new project or trying out Ory Cloud.
     *
     * @return self
     */
    public function setApplicationUrl($applicationUrl)
    {
        $this->container['applicationUrl'] = $applicationUrl;

        return $this;
    }

    /**
     * Gets defaultIdentitySchemaUrl
     *
     * @return string|null
     */
    public function getDefaultIdentitySchemaUrl()
    {
        return $this->container['defaultIdentitySchemaUrl'];
    }

    /**
     * Sets defaultIdentitySchemaUrl
     *
     * @param string|null $defaultIdentitySchemaUrl Default Identity Schema URL  This represents your Ory Kratos Default Identity Schema. It is your identity's default schema. This allows setting custom fields such as \"address\", specifying whether you want to log in using email or a username, and more. For more information on this topic, please check out the identity documentation.  If KratosCustomSchemaID is also added to the payload, it will be preferred and replace the value of DefaultIdentitySchemaURL set by the client.  The value of this field can be either any \"https://\" URL - for example a file hosted at GitHub, or a `preset://`-prefixed string.  Available presets are:  preset://email preset://username preset://basic preset://blank
     *
     * @return self
     */
    public function setDefaultIdentitySchemaUrl($defaultIdentitySchemaUrl)
    {
        $this->container['defaultIdentitySchemaUrl'] = $defaultIdentitySchemaUrl;

        return $this;
    }

    /**
     * Gets errorUiUrl
     *
     * @return string
     */
    public function getErrorUiUrl()
    {
        return $this->container['errorUiUrl'];
    }

    /**
     * Sets errorUiUrl
     *
     * @param string $errorUiUrl Self-Service Error UI URL  Sets the UI URL for the error UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setErrorUiUrl($errorUiUrl)
    {
        $this->container['errorUiUrl'] = $errorUiUrl;

        return $this;
    }

    /**
     * Gets kratosCustomSchemaId
     *
     * @return string|null
     */
    public function getKratosCustomSchemaId()
    {
        return $this->container['kratosCustomSchemaId'];
    }

    /**
     * Sets kratosCustomSchemaId
     *
     * @param string|null $kratosCustomSchemaId The Custom Schema ID  The custom schema ID is a preferred field over DefaultIdentitySchemaURL, meaning it will override the DefaultIdentitySchemaURL with the custom schema's URL instead of accepting it from the client-side. Once this field is set the given custom schema will be associated with this project.
     *
     * @return self
     */
    public function setKratosCustomSchemaId($kratosCustomSchemaId)
    {
        $this->container['kratosCustomSchemaId'] = $kratosCustomSchemaId;

        return $this;
    }

    /**
     * Gets loginUiUrl
     *
     * @return string
     */
    public function getLoginUiUrl()
    {
        return $this->container['loginUiUrl'];
    }

    /**
     * Sets loginUiUrl
     *
     * @param string $loginUiUrl Self-Service Login UI URL  Sets the UI URL for the login UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setLoginUiUrl($loginUiUrl)
    {
        $this->container['loginUiUrl'] = $loginUiUrl;

        return $this;
    }

    /**
     * Gets lookupSecret
     *
     * @return \Ory\Client\Model\ProjectLookupSecretConfig
     */
    public function getLookupSecret()
    {
        return $this->container['lookupSecret'];
    }

    /**
     * Sets lookupSecret
     *
     * @param \Ory\Client\Model\ProjectLookupSecretConfig $lookupSecret lookupSecret
     *
     * @return self
     */
    public function setLookupSecret($lookupSecret)
    {
        $this->container['lookupSecret'] = $lookupSecret;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name The project name.  Helps you identify your project.  in: body
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets password
     *
     * @return \Ory\Client\Model\ProjectPasswordConfig
     */
    public function getPassword()
    {
        return $this->container['password'];
    }

    /**
     * Sets password
     *
     * @param \Ory\Client\Model\ProjectPasswordConfig $password password
     *
     * @return self
     */
    public function setPassword($password)
    {
        $this->container['password'] = $password;

        return $this;
    }

    /**
     * Gets recovery
     *
     * @return \Ory\Client\Model\ProjectRecoveryConfig
     */
    public function getRecovery()
    {
        return $this->container['recovery'];
    }

    /**
     * Sets recovery
     *
     * @param \Ory\Client\Model\ProjectRecoveryConfig $recovery recovery
     *
     * @return self
     */
    public function setRecovery($recovery)
    {
        $this->container['recovery'] = $recovery;

        return $this;
    }

    /**
     * Gets recoveryUiUrl
     *
     * @return string
     */
    public function getRecoveryUiUrl()
    {
        return $this->container['recoveryUiUrl'];
    }

    /**
     * Sets recoveryUiUrl
     *
     * @param string $recoveryUiUrl Self-Service Login UI URL  Sets the UI URL for the recovery UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setRecoveryUiUrl($recoveryUiUrl)
    {
        $this->container['recoveryUiUrl'] = $recoveryUiUrl;

        return $this;
    }

    /**
     * Gets redirectionConfig
     *
     * @return \Ory\Client\Model\RedirectionConfig|null
     */
    public function getRedirectionConfig()
    {
        return $this->container['redirectionConfig'];
    }

    /**
     * Sets redirectionConfig
     *
     * @param \Ory\Client\Model\RedirectionConfig|null $redirectionConfig redirectionConfig
     *
     * @return self
     */
    public function setRedirectionConfig($redirectionConfig)
    {
        $this->container['redirectionConfig'] = $redirectionConfig;

        return $this;
    }

    /**
     * Gets registrationUiUrl
     *
     * @return string
     */
    public function getRegistrationUiUrl()
    {
        return $this->container['registrationUiUrl'];
    }

    /**
     * Sets registrationUiUrl
     *
     * @param string $registrationUiUrl Self-Service Login UI URL  Sets the UI URL for the registration UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setRegistrationUiUrl($registrationUiUrl)
    {
        $this->container['registrationUiUrl'] = $registrationUiUrl;

        return $this;
    }

    /**
     * Gets sessionAfterSignUp
     *
     * @return bool
     */
    public function getSessionAfterSignUp()
    {
        return $this->container['sessionAfterSignUp'];
    }

    /**
     * Sets sessionAfterSignUp
     *
     * @param bool $sessionAfterSignUp Issue Session after Sign Up  If set to true, users will receive a session after they successfully signed up. Enabling this option allows account enumeration during registration flows. Read more: https://www.ory.sh/kratos/docs/self-service/flows/user-registration#successful-registration
     *
     * @return self
     */
    public function setSessionAfterSignUp($sessionAfterSignUp)
    {
        $this->container['sessionAfterSignUp'] = $sessionAfterSignUp;

        return $this;
    }

    /**
     * Gets sessionSoft2fa
     *
     * @return bool
     */
    public function getSessionSoft2fa()
    {
        return $this->container['sessionSoft2fa'];
    }

    /**
     * Sets sessionSoft2fa
     *
     * @param bool $sessionSoft2fa Enable Soft 2FA for Login Sessions
     *
     * @return self
     */
    public function setSessionSoft2fa($sessionSoft2fa)
    {
        $this->container['sessionSoft2fa'] = $sessionSoft2fa;

        return $this;
    }

    /**
     * Gets settingsPrivilegedSessionMaxAgeSeconds
     *
     * @return int
     */
    public function getSettingsPrivilegedSessionMaxAgeSeconds()
    {
        return $this->container['settingsPrivilegedSessionMaxAgeSeconds'];
    }

    /**
     * Sets settingsPrivilegedSessionMaxAgeSeconds
     *
     * @param int $settingsPrivilegedSessionMaxAgeSeconds Duration in Seconds of how long a Session is Privileged  Defines how long a session is considered privileged in seconds. If the session's authenticated_at is older than the value specified here, the user needs to re-authenticate to perform certain actions (e.g. password change).
     *
     * @return self
     */
    public function setSettingsPrivilegedSessionMaxAgeSeconds($settingsPrivilegedSessionMaxAgeSeconds)
    {
        $this->container['settingsPrivilegedSessionMaxAgeSeconds'] = $settingsPrivilegedSessionMaxAgeSeconds;

        return $this;
    }

    /**
     * Gets settingsSoft2fa
     *
     * @return bool
     */
    public function getSettingsSoft2fa()
    {
        return $this->container['settingsSoft2fa'];
    }

    /**
     * Sets settingsSoft2fa
     *
     * @param bool $settingsSoft2fa Enable Soft 2FA for Self-Service Settings Flows
     *
     * @return self
     */
    public function setSettingsSoft2fa($settingsSoft2fa)
    {
        $this->container['settingsSoft2fa'] = $settingsSoft2fa;

        return $this;
    }

    /**
     * Gets settingsUiUrl
     *
     * @return string
     */
    public function getSettingsUiUrl()
    {
        return $this->container['settingsUiUrl'];
    }

    /**
     * Sets settingsUiUrl
     *
     * @param string $settingsUiUrl Self-Service Settings UI URL  Sets the UI URL for the settings UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setSettingsUiUrl($settingsUiUrl)
    {
        $this->container['settingsUiUrl'] = $settingsUiUrl;

        return $this;
    }

    /**
     * Gets totp
     *
     * @return \Ory\Client\Model\ProjectTotpConfig
     */
    public function getTotp()
    {
        return $this->container['totp'];
    }

    /**
     * Sets totp
     *
     * @param \Ory\Client\Model\ProjectTotpConfig $totp totp
     *
     * @return self
     */
    public function setTotp($totp)
    {
        $this->container['totp'] = $totp;

        return $this;
    }

    /**
     * Gets verification
     *
     * @return \Ory\Client\Model\ProjectVerificationConfig
     */
    public function getVerification()
    {
        return $this->container['verification'];
    }

    /**
     * Sets verification
     *
     * @param \Ory\Client\Model\ProjectVerificationConfig $verification verification
     *
     * @return self
     */
    public function setVerification($verification)
    {
        $this->container['verification'] = $verification;

        return $this;
    }

    /**
     * Gets verificationUiUrl
     *
     * @return string
     */
    public function getVerificationUiUrl()
    {
        return $this->container['verificationUiUrl'];
    }

    /**
     * Sets verificationUiUrl
     *
     * @param string $verificationUiUrl Self-Service Login UI URL  Sets the UI URL for the verification UI. If left empty, this will use Ory's hosted pages.
     *
     * @return self
     */
    public function setVerificationUiUrl($verificationUiUrl)
    {
        $this->container['verificationUiUrl'] = $verificationUiUrl;

        return $this;
    }

    /**
     * Gets webauthn
     *
     * @return \Ory\Client\Model\ProjectWebAuthnConfig
     */
    public function getWebauthn()
    {
        return $this->container['webauthn'];
    }

    /**
     * Sets webauthn
     *
     * @param \Ory\Client\Model\ProjectWebAuthnConfig $webauthn webauthn
     *
     * @return self
     */
    public function setWebauthn($webauthn)
    {
        $this->container['webauthn'] = $webauthn;

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


