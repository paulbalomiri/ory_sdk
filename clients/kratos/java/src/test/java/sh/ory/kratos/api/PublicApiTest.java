/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.17
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.api;

import sh.ory.kratos.ApiException;
import sh.ory.kratos.model.ErrorContainer;
import sh.ory.kratos.model.GenericError;
import sh.ory.kratos.model.LoginFlow;
import sh.ory.kratos.model.LoginViaApiResponse;
import sh.ory.kratos.model.RecoveryFlow;
import sh.ory.kratos.model.RegistrationFlow;
import sh.ory.kratos.model.RegistrationViaApiResponse;
import sh.ory.kratos.model.RevokeSession;
import sh.ory.kratos.model.Session;
import sh.ory.kratos.model.SettingsFlow;
import sh.ory.kratos.model.SettingsViaApiResponse;
import sh.ory.kratos.model.SubmitSelfServiceLoginFlow;
import sh.ory.kratos.model.SubmitSelfServiceRecoveryFlowWithLinkMethod;
import sh.ory.kratos.model.SubmitSelfServiceRegistrationFlow;
import sh.ory.kratos.model.SubmitSelfServiceSettingsFlow;
import sh.ory.kratos.model.VerificationFlow;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for PublicApi
 */
@Ignore
public class PublicApiTest {

    private final PublicApi api = new PublicApi();

    
    /**
     * 
     *
     * Get a Traits Schema Definition
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSchemaTest() throws ApiException {
        String id = null;
        Object response = api.getSchema(id);

        // TODO: test validations
    }
    
    /**
     * Get User-Facing Self-Service Errors
     *
     * This endpoint returns the error associated with a user-facing self service errors.  This endpoint supports stub values to help you implement the error UI:  &#x60;?error&#x3D;stub:500&#x60; - returns a stub 500 (Internal Server Error) error.  More information can be found at [Ory Kratos User User Facing Error Documentation](https://www.ory.sh/docs/kratos/self-service/flows/user-facing-errors).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceErrorTest() throws ApiException {
        String error = null;
        ErrorContainer response = api.getSelfServiceError(error);

        // TODO: test validations
    }
    
    /**
     * Get Login Flow
     *
     * This endpoint returns a login flow&#39;s context with, for example, error details and other information.  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceLoginFlowTest() throws ApiException {
        String id = null;
        LoginFlow response = api.getSelfServiceLoginFlow(id);

        // TODO: test validations
    }
    
    /**
     * Get information about a recovery flow
     *
     * This endpoint returns a recovery flow&#39;s context with, for example, error details and other information.  More information can be found at [Ory Kratos Account Recovery Documentation](../self-service/flows/account-recovery.mdx).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceRecoveryFlowTest() throws ApiException {
        String id = null;
        RecoveryFlow response = api.getSelfServiceRecoveryFlow(id);

        // TODO: test validations
    }
    
    /**
     * Get Registration Flow
     *
     * This endpoint returns a registration flow&#39;s context with, for example, error details and other information.  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceRegistrationFlowTest() throws ApiException {
        String id = null;
        RegistrationFlow response = api.getSelfServiceRegistrationFlow(id);

        // TODO: test validations
    }
    
    /**
     * Get Settings Flow
     *
     * When accessing this endpoint through Ory Kratos&#39; Public API you must ensure that either the Ory Kratos Session Cookie or the Ory Kratos Session Token are set. The public endpoint does not return 404 status codes but instead 403 or 500 to improve data privacy.  You can access this endpoint without credentials when using Ory Kratos&#39; Admin API.  More information can be found at [Ory Kratos User Settings &amp; Profile Management Documentation](../self-service/flows/user-settings).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceSettingsFlowTest() throws ApiException {
        String id = null;
        SettingsFlow response = api.getSelfServiceSettingsFlow(id);

        // TODO: test validations
    }
    
    /**
     * Get Verification Flow
     *
     * This endpoint returns a verification flow&#39;s context with, for example, error details and other information.  More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSelfServiceVerificationFlowTest() throws ApiException {
        String id = null;
        VerificationFlow response = api.getSelfServiceVerificationFlow(id);

        // TODO: test validations
    }
    
    /**
     * Initialize Browser-Based Logout User Flow
     *
     * This endpoint initializes a logout flow.  &gt; This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  On successful logout, the browser will be redirected (HTTP 302 Found) to the &#x60;return_to&#x60; parameter of the initial request or fall back to &#x60;urls.default_return_to&#x60;.  More information can be found at [Ory Kratos User Logout Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-logout).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceBrowserLogoutFlowTest() throws ApiException {
        api.initializeSelfServiceBrowserLogoutFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Login Flow for API clients
     *
     * This endpoint initiates a login flow for API clients such as mobile devices, smart TVs, and so on.  If a valid provided session cookie or session token is provided, a 400 Bad Request error will be returned unless the URL query parameter &#x60;?refresh&#x3D;true&#x60; is set.  To fetch an existing login flow call &#x60;/self-service/login/flows?flow&#x3D;&lt;flow_id&gt;&#x60;.  :::warning  You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make you vulnerable to a variety of CSRF attacks, including CSRF login attacks.  This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).  :::  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceLoginViaAPIFlowTest() throws ApiException {
        Boolean refresh = null;
        LoginFlow response = api.initializeSelfServiceLoginViaAPIFlow(refresh);

        // TODO: test validations
    }
    
    /**
     * Initialize Login Flow for browsers
     *
     * This endpoint initializes a browser-based user login flow. Once initialized, the browser will be redirected to &#x60;selfservice.flows.login.ui_url&#x60; with the flow ID set as the query parameter &#x60;?flow&#x3D;&#x60;. If a valid user session exists already, the browser will be redirected to &#x60;urls.default_redirect_url&#x60; unless the query parameter &#x60;?refresh&#x3D;true&#x60; was set.  This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceLoginViaBrowserFlowTest() throws ApiException {
        api.initializeSelfServiceLoginViaBrowserFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Recovery Flow for API Clients
     *
     * This endpoint initiates a recovery flow for API clients such as mobile devices, smart TVs, and so on.  If a valid provided session cookie or session token is provided, a 400 Bad Request error.  To fetch an existing recovery flow call &#x60;/self-service/recovery/flows?flow&#x3D;&lt;flow_id&gt;&#x60;.  :::warning  You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make you vulnerable to a variety of CSRF attacks.  This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).  :::  More information can be found at [Ory Kratos Account Recovery Documentation](../self-service/flows/account-recovery.mdx).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceRecoveryViaAPIFlowTest() throws ApiException {
        RecoveryFlow response = api.initializeSelfServiceRecoveryViaAPIFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Recovery Flow for Browser Clients
     *
     * This endpoint initializes a browser-based account recovery flow. Once initialized, the browser will be redirected to &#x60;selfservice.flows.recovery.ui_url&#x60; with the flow ID set as the query parameter &#x60;?flow&#x3D;&#x60;. If a valid user session exists, the browser is returned to the configured return URL.  This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  More information can be found at [Ory Kratos Account Recovery Documentation](../self-service/flows/account-recovery.mdx).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceRecoveryViaBrowserFlowTest() throws ApiException {
        api.initializeSelfServiceRecoveryViaBrowserFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Registration Flow for API clients
     *
     * This endpoint initiates a registration flow for API clients such as mobile devices, smart TVs, and so on.  If a valid provided session cookie or session token is provided, a 400 Bad Request error will be returned unless the URL query parameter &#x60;?refresh&#x3D;true&#x60; is set.  To fetch an existing registration flow call &#x60;/self-service/registration/flows?flow&#x3D;&lt;flow_id&gt;&#x60;.  :::warning  You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make you vulnerable to a variety of CSRF attacks.  This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).  :::  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceRegistrationViaAPIFlowTest() throws ApiException {
        RegistrationFlow response = api.initializeSelfServiceRegistrationViaAPIFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Registration Flow for browsers
     *
     * This endpoint initializes a browser-based user registration flow. Once initialized, the browser will be redirected to &#x60;selfservice.flows.registration.ui_url&#x60; with the flow ID set as the query parameter &#x60;?flow&#x3D;&#x60;. If a valid user session exists already, the browser will be redirected to &#x60;urls.default_redirect_url&#x60; unless the query parameter &#x60;?refresh&#x3D;true&#x60; was set.  :::note  This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  :::  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceRegistrationViaBrowserFlowTest() throws ApiException {
        api.initializeSelfServiceRegistrationViaBrowserFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Settings Flow for API Clients
     *
     * This endpoint initiates a settings flow for API clients such as mobile devices, smart TVs, and so on. You must provide a valid Ory Kratos Session Token for this endpoint to respond with HTTP 200 OK.  To fetch an existing settings flow call &#x60;/self-service/settings/flows?flow&#x3D;&lt;flow_id&gt;&#x60;.  :::warning  You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make you vulnerable to a variety of CSRF attacks.  This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).  :::  More information can be found at [Ory Kratos User Settings &amp; Profile Management Documentation](../self-service/flows/user-settings).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceSettingsViaAPIFlowTest() throws ApiException {
        SettingsFlow response = api.initializeSelfServiceSettingsViaAPIFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Settings Flow for Browsers
     *
     * This endpoint initializes a browser-based user settings flow. Once initialized, the browser will be redirected to &#x60;selfservice.flows.settings.ui_url&#x60; with the flow ID set as the query parameter &#x60;?flow&#x3D;&#x60;. If no valid Ory Kratos Session Cookie is included in the request, a login flow will be initialized.  :::note  This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  :::  More information can be found at [Ory Kratos User Settings &amp; Profile Management Documentation](../self-service/flows/user-settings).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceSettingsViaBrowserFlowTest() throws ApiException {
        api.initializeSelfServiceSettingsViaBrowserFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Verification Flow for API Clients
     *
     * This endpoint initiates a verification flow for API clients such as mobile devices, smart TVs, and so on.  To fetch an existing verification flow call &#x60;/self-service/verification/flows?flow&#x3D;&lt;flow_id&gt;&#x60;.  :::warning  You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make you vulnerable to a variety of CSRF attacks.  This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).  :::  More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceVerificationViaAPIFlowTest() throws ApiException {
        VerificationFlow response = api.initializeSelfServiceVerificationViaAPIFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Verification Flow for Browser Clients
     *
     * This endpoint initializes a browser-based account verification flow. Once initialized, the browser will be redirected to &#x60;selfservice.flows.verification.ui_url&#x60; with the flow ID set as the query parameter &#x60;?flow&#x3D;&#x60;.  This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).  More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void initializeSelfServiceVerificationViaBrowserFlowTest() throws ApiException {
        api.initializeSelfServiceVerificationViaBrowserFlow();

        // TODO: test validations
    }
    
    /**
     * Initialize Logout Flow for API Clients - Revoke a Session
     *
     * Use this endpoint to revoke a session using its token. This endpoint is particularly useful for API clients such as mobile apps to log the user out of the system and invalidate the session.  This endpoint does not remove any HTTP Cookies - use the Browser-Based Self-Service Logout Flow instead.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void revokeSessionTest() throws ApiException {
        RevokeSession revokeSession = null;
        api.revokeSession(revokeSession);

        // TODO: test validations
    }
    
    /**
     * Submit a Login Flow
     *
     * Use this endpoint to complete a login flow. This endpoint behaves differently for API and browser flows.  API flows expect &#x60;application/json&#x60; to be sent in the body and responds with HTTP 200 and a application/json body with the session token on success; HTTP 302 redirect to a fresh login flow if the original flow expired with the appropriate error messages set; HTTP 400 on form validation errors.  Browser flows expect &#x60;application/x-www-form-urlencoded&#x60; to be sent in the body and responds with a HTTP 302 redirect to the post/after login URL or the &#x60;return_to&#x60; value if it was set and if the login succeeded; a HTTP 302 redirect to the login UI URL with the flow ID containing the validation errors otherwise.  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceLoginFlowTest() throws ApiException {
        String flow = null;
        SubmitSelfServiceLoginFlow submitSelfServiceLoginFlow = null;
        LoginViaApiResponse response = api.submitSelfServiceLoginFlow(flow, submitSelfServiceLoginFlow);

        // TODO: test validations
    }
    
    /**
     * Complete Recovery Flow
     *
     * Use this endpoint to complete a recovery flow. This endpoint behaves differently for API and browser flows and has several states:  &#x60;choose_method&#x60; expects &#x60;flow&#x60; (in the URL query) and &#x60;email&#x60; (in the body) to be sent and works with API- and Browser-initiated flows. For API clients it either returns a HTTP 200 OK when the form is valid and HTTP 400 OK when the form is invalid and a HTTP 302 Found redirect with a fresh recovery flow if the flow was otherwise invalid (e.g. expired). For Browser clients it returns a HTTP 302 Found redirect to the Recovery UI URL with the Recovery Flow ID appended. &#x60;sent_email&#x60; is the success state after &#x60;choose_method&#x60; for the &#x60;link&#x60; method and allows the user to request another recovery email. It works for both API and Browser-initiated flows and returns the same responses as the flow in &#x60;choose_method&#x60; state. &#x60;passed_challenge&#x60; expects a &#x60;token&#x60; to be sent in the URL query and given the nature of the flow (\&quot;sending a recovery link\&quot;) does not have any API capabilities. The server responds with a HTTP 302 Found redirect either to the Settings UI URL (if the link was valid) and instructs the user to update their password, or a redirect to the Recover UI URL with a new Recovery Flow ID which contains an error message that the recovery link was invalid.  More information can be found at [Ory Kratos Account Recovery Documentation](../self-service/flows/account-recovery.mdx).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceRecoveryFlowTest() throws ApiException {
        String flow = null;
        Object body = null;
        api.submitSelfServiceRecoveryFlow(flow, body);

        // TODO: test validations
    }
    
    /**
     * Complete Recovery Flow with Link Method
     *
     * Use this endpoint to complete a recovery flow using the link method. This endpoint behaves differently for API and browser flows and has several states:  &#x60;choose_method&#x60; expects &#x60;flow&#x60; (in the URL query) and &#x60;email&#x60; (in the body) to be sent and works with API- and Browser-initiated flows. For API clients it either returns a HTTP 200 OK when the form is valid and HTTP 400 OK when the form is invalid and a HTTP 302 Found redirect with a fresh recovery flow if the flow was otherwise invalid (e.g. expired). For Browser clients it returns a HTTP 302 Found redirect to the Recovery UI URL with the Recovery Flow ID appended. &#x60;sent_email&#x60; is the success state after &#x60;choose_method&#x60; and allows the user to request another recovery email. It works for both API and Browser-initiated flows and returns the same responses as the flow in &#x60;choose_method&#x60; state. &#x60;passed_challenge&#x60; expects a &#x60;token&#x60; to be sent in the URL query and given the nature of the flow (\&quot;sending a recovery link\&quot;) does not have any API capabilities. The server responds with a HTTP 302 Found redirect either to the Settings UI URL (if the link was valid) and instructs the user to update their password, or a redirect to the Recover UI URL with a new Recovery Flow ID which contains an error message that the recovery link was invalid.  More information can be found at [Ory Kratos Account Recovery Documentation](../self-service/flows/account-recovery.mdx).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceRecoveryFlowWithLinkMethodTest() throws ApiException {
        String token = null;
        String flow = null;
        SubmitSelfServiceRecoveryFlowWithLinkMethod submitSelfServiceRecoveryFlowWithLinkMethod = null;
        api.submitSelfServiceRecoveryFlowWithLinkMethod(token, flow, submitSelfServiceRecoveryFlowWithLinkMethod);

        // TODO: test validations
    }
    
    /**
     * Submit a Registration Flow
     *
     * Use this endpoint to complete a registration flow by sending an identity&#39;s traits and password. This endpoint behaves differently for API and browser flows.  API flows expect &#x60;application/json&#x60; to be sent in the body and respond with HTTP 200 and a application/json body with the created identity success - if the session hook is configured the &#x60;session&#x60; and &#x60;session_token&#x60; will also be included; HTTP 302 redirect to a fresh registration flow if the original flow expired with the appropriate error messages set; HTTP 400 on form validation errors.  Browser flows expect &#x60;application/x-www-form-urlencoded&#x60; to be sent in the body and responds with a HTTP 302 redirect to the post/after registration URL or the &#x60;return_to&#x60; value if it was set and if the registration succeeded; a HTTP 302 redirect to the registration UI URL with the flow ID containing the validation errors otherwise.  More information can be found at [Ory Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceRegistrationFlowTest() throws ApiException {
        String flow = null;
        SubmitSelfServiceRegistrationFlow submitSelfServiceRegistrationFlow = null;
        RegistrationViaApiResponse response = api.submitSelfServiceRegistrationFlow(flow, submitSelfServiceRegistrationFlow);

        // TODO: test validations
    }
    
    /**
     * Complete Settings Flow
     *
     * Use this endpoint to complete a settings flow by sending an identity&#39;s updated password. This endpoint behaves differently for API and browser flows.  API-initiated flows expect &#x60;application/json&#x60; to be sent in the body and respond with HTTP 200 and an application/json body with the session token on success; HTTP 302 redirect to a fresh settings flow if the original flow expired with the appropriate error messages set; HTTP 400 on form validation errors. HTTP 401 when the endpoint is called without a valid session token. HTTP 403 when &#x60;selfservice.flows.settings.privileged_session_max_age&#x60; was reached. Implies that the user needs to re-authenticate.  Browser flows expect &#x60;application/x-www-form-urlencoded&#x60; to be sent in the body and responds with a HTTP 302 redirect to the post/after settings URL or the &#x60;return_to&#x60; value if it was set and if the flow succeeded; a HTTP 302 redirect to the Settings UI URL with the flow ID containing the validation errors otherwise. a HTTP 302 redirect to the login endpoint when &#x60;selfservice.flows.settings.privileged_session_max_age&#x60; was reached.  More information can be found at [Ory Kratos User Settings &amp; Profile Management Documentation](../self-service/flows/user-settings).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceSettingsFlowTest() throws ApiException {
        String flow = null;
        SubmitSelfServiceSettingsFlow submitSelfServiceSettingsFlow = null;
        SettingsViaApiResponse response = api.submitSelfServiceSettingsFlow(flow, submitSelfServiceSettingsFlow);

        // TODO: test validations
    }
    
    /**
     * Complete Verification Flow
     *
     * Use this endpoint to complete a verification flow. This endpoint behaves differently for API and browser flows and has several states:  &#x60;choose_method&#x60; expects &#x60;flow&#x60; (in the URL query) and &#x60;email&#x60; (in the body) to be sent and works with API- and Browser-initiated flows. For API clients it either returns a HTTP 200 OK when the form is valid and HTTP 400 OK when the form is invalid and a HTTP 302 Found redirect with a fresh verification flow if the flow was otherwise invalid (e.g. expired). For Browser clients it returns a HTTP 302 Found redirect to the Verification UI URL with the Verification Flow ID appended. &#x60;sent_email&#x60; is the success state after &#x60;choose_method&#x60; when using the &#x60;link&#x60; method and allows the user to request another verification email. It works for both API and Browser-initiated flows and returns the same responses as the flow in &#x60;choose_method&#x60; state. &#x60;passed_challenge&#x60; expects a &#x60;token&#x60; to be sent in the URL query and given the nature of the flow (\&quot;sending a verification link\&quot;) does not have any API capabilities. The server responds with a HTTP 302 Found redirect either to the Settings UI URL (if the link was valid) and instructs the user to update their password, or a redirect to the Verification UI URL with a new Verification Flow ID which contains an error message that the verification link was invalid.  More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void submitSelfServiceVerificationFlowTest() throws ApiException {
        String flow = null;
        Object body = null;
        api.submitSelfServiceVerificationFlow(flow, body);

        // TODO: test validations
    }
    
    /**
     * Check Who the Current HTTP Session Belongs To
     *
     * Uses the HTTP Headers in the GET request to determine (e.g. by using checking the cookies) who is authenticated. Returns a session object in the body or 401 if the credentials are invalid or no credentials were sent. Additionally when the request it successful it adds the user ID to the &#39;X-Kratos-Authenticated-Identity-Id&#39; header in the response.  This endpoint is useful for reverse proxies and API Gateways.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void whoamiTest() throws ApiException {
        String cookie = null;
        String authorization = null;
        Session response = api.whoami(cookie, authorization);

        // TODO: test validations
    }
    
}
