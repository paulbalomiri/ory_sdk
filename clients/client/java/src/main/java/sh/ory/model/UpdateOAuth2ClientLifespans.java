/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.60
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import sh.ory.JSON;

/**
 * The OIDC Hybrid grant type inherits token lifespan configuration from the implicit grant.
 */
@ApiModel(description = "The OIDC Hybrid grant type inherits token lifespan configuration from the implicit grant.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-10-24T09:56:23.976880967Z[Etc/UTC]")
public class UpdateOAuth2ClientLifespans {
  public static final String SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_ACCESS_TOKEN_LIFESPAN = "authorization_code_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String authorizationCodeGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_ID_TOKEN_LIFESPAN = "authorization_code_grant_id_token_lifespan";
  @SerializedName(SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_ID_TOKEN_LIFESPAN)
  private String authorizationCodeGrantIdTokenLifespan;

  public static final String SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_REFRESH_TOKEN_LIFESPAN = "authorization_code_grant_refresh_token_lifespan";
  @SerializedName(SERIALIZED_NAME_AUTHORIZATION_CODE_GRANT_REFRESH_TOKEN_LIFESPAN)
  private String authorizationCodeGrantRefreshTokenLifespan;

  public static final String SERIALIZED_NAME_CLIENT_CREDENTIALS_GRANT_ACCESS_TOKEN_LIFESPAN = "client_credentials_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_CLIENT_CREDENTIALS_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String clientCredentialsGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_IMPLICIT_GRANT_ACCESS_TOKEN_LIFESPAN = "implicit_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_IMPLICIT_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String implicitGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_IMPLICIT_GRANT_ID_TOKEN_LIFESPAN = "implicit_grant_id_token_lifespan";
  @SerializedName(SERIALIZED_NAME_IMPLICIT_GRANT_ID_TOKEN_LIFESPAN)
  private String implicitGrantIdTokenLifespan;

  public static final String SERIALIZED_NAME_JWT_BEARER_GRANT_ACCESS_TOKEN_LIFESPAN = "jwt_bearer_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_JWT_BEARER_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String jwtBearerGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_PASSWORD_GRANT_ACCESS_TOKEN_LIFESPAN = "password_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_PASSWORD_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String passwordGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_PASSWORD_GRANT_REFRESH_TOKEN_LIFESPAN = "password_grant_refresh_token_lifespan";
  @SerializedName(SERIALIZED_NAME_PASSWORD_GRANT_REFRESH_TOKEN_LIFESPAN)
  private String passwordGrantRefreshTokenLifespan;

  public static final String SERIALIZED_NAME_REFRESH_TOKEN_GRANT_ACCESS_TOKEN_LIFESPAN = "refresh_token_grant_access_token_lifespan";
  @SerializedName(SERIALIZED_NAME_REFRESH_TOKEN_GRANT_ACCESS_TOKEN_LIFESPAN)
  private String refreshTokenGrantAccessTokenLifespan;

  public static final String SERIALIZED_NAME_REFRESH_TOKEN_GRANT_ID_TOKEN_LIFESPAN = "refresh_token_grant_id_token_lifespan";
  @SerializedName(SERIALIZED_NAME_REFRESH_TOKEN_GRANT_ID_TOKEN_LIFESPAN)
  private String refreshTokenGrantIdTokenLifespan;

  public static final String SERIALIZED_NAME_REFRESH_TOKEN_GRANT_REFRESH_TOKEN_LIFESPAN = "refresh_token_grant_refresh_token_lifespan";
  @SerializedName(SERIALIZED_NAME_REFRESH_TOKEN_GRANT_REFRESH_TOKEN_LIFESPAN)
  private String refreshTokenGrantRefreshTokenLifespan;

  public UpdateOAuth2ClientLifespans() { 
  }

  public UpdateOAuth2ClientLifespans authorizationCodeGrantAccessTokenLifespan(String authorizationCodeGrantAccessTokenLifespan) {
    
    this.authorizationCodeGrantAccessTokenLifespan = authorizationCodeGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get authorizationCodeGrantAccessTokenLifespan
   * @return authorizationCodeGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAuthorizationCodeGrantAccessTokenLifespan() {
    return authorizationCodeGrantAccessTokenLifespan;
  }


  public void setAuthorizationCodeGrantAccessTokenLifespan(String authorizationCodeGrantAccessTokenLifespan) {
    this.authorizationCodeGrantAccessTokenLifespan = authorizationCodeGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans authorizationCodeGrantIdTokenLifespan(String authorizationCodeGrantIdTokenLifespan) {
    
    this.authorizationCodeGrantIdTokenLifespan = authorizationCodeGrantIdTokenLifespan;
    return this;
  }

   /**
   * Get authorizationCodeGrantIdTokenLifespan
   * @return authorizationCodeGrantIdTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAuthorizationCodeGrantIdTokenLifespan() {
    return authorizationCodeGrantIdTokenLifespan;
  }


  public void setAuthorizationCodeGrantIdTokenLifespan(String authorizationCodeGrantIdTokenLifespan) {
    this.authorizationCodeGrantIdTokenLifespan = authorizationCodeGrantIdTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans authorizationCodeGrantRefreshTokenLifespan(String authorizationCodeGrantRefreshTokenLifespan) {
    
    this.authorizationCodeGrantRefreshTokenLifespan = authorizationCodeGrantRefreshTokenLifespan;
    return this;
  }

   /**
   * Get authorizationCodeGrantRefreshTokenLifespan
   * @return authorizationCodeGrantRefreshTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAuthorizationCodeGrantRefreshTokenLifespan() {
    return authorizationCodeGrantRefreshTokenLifespan;
  }


  public void setAuthorizationCodeGrantRefreshTokenLifespan(String authorizationCodeGrantRefreshTokenLifespan) {
    this.authorizationCodeGrantRefreshTokenLifespan = authorizationCodeGrantRefreshTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans clientCredentialsGrantAccessTokenLifespan(String clientCredentialsGrantAccessTokenLifespan) {
    
    this.clientCredentialsGrantAccessTokenLifespan = clientCredentialsGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get clientCredentialsGrantAccessTokenLifespan
   * @return clientCredentialsGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getClientCredentialsGrantAccessTokenLifespan() {
    return clientCredentialsGrantAccessTokenLifespan;
  }


  public void setClientCredentialsGrantAccessTokenLifespan(String clientCredentialsGrantAccessTokenLifespan) {
    this.clientCredentialsGrantAccessTokenLifespan = clientCredentialsGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans implicitGrantAccessTokenLifespan(String implicitGrantAccessTokenLifespan) {
    
    this.implicitGrantAccessTokenLifespan = implicitGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get implicitGrantAccessTokenLifespan
   * @return implicitGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImplicitGrantAccessTokenLifespan() {
    return implicitGrantAccessTokenLifespan;
  }


  public void setImplicitGrantAccessTokenLifespan(String implicitGrantAccessTokenLifespan) {
    this.implicitGrantAccessTokenLifespan = implicitGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans implicitGrantIdTokenLifespan(String implicitGrantIdTokenLifespan) {
    
    this.implicitGrantIdTokenLifespan = implicitGrantIdTokenLifespan;
    return this;
  }

   /**
   * Get implicitGrantIdTokenLifespan
   * @return implicitGrantIdTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getImplicitGrantIdTokenLifespan() {
    return implicitGrantIdTokenLifespan;
  }


  public void setImplicitGrantIdTokenLifespan(String implicitGrantIdTokenLifespan) {
    this.implicitGrantIdTokenLifespan = implicitGrantIdTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans jwtBearerGrantAccessTokenLifespan(String jwtBearerGrantAccessTokenLifespan) {
    
    this.jwtBearerGrantAccessTokenLifespan = jwtBearerGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get jwtBearerGrantAccessTokenLifespan
   * @return jwtBearerGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getJwtBearerGrantAccessTokenLifespan() {
    return jwtBearerGrantAccessTokenLifespan;
  }


  public void setJwtBearerGrantAccessTokenLifespan(String jwtBearerGrantAccessTokenLifespan) {
    this.jwtBearerGrantAccessTokenLifespan = jwtBearerGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans passwordGrantAccessTokenLifespan(String passwordGrantAccessTokenLifespan) {
    
    this.passwordGrantAccessTokenLifespan = passwordGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get passwordGrantAccessTokenLifespan
   * @return passwordGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPasswordGrantAccessTokenLifespan() {
    return passwordGrantAccessTokenLifespan;
  }


  public void setPasswordGrantAccessTokenLifespan(String passwordGrantAccessTokenLifespan) {
    this.passwordGrantAccessTokenLifespan = passwordGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans passwordGrantRefreshTokenLifespan(String passwordGrantRefreshTokenLifespan) {
    
    this.passwordGrantRefreshTokenLifespan = passwordGrantRefreshTokenLifespan;
    return this;
  }

   /**
   * Get passwordGrantRefreshTokenLifespan
   * @return passwordGrantRefreshTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPasswordGrantRefreshTokenLifespan() {
    return passwordGrantRefreshTokenLifespan;
  }


  public void setPasswordGrantRefreshTokenLifespan(String passwordGrantRefreshTokenLifespan) {
    this.passwordGrantRefreshTokenLifespan = passwordGrantRefreshTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans refreshTokenGrantAccessTokenLifespan(String refreshTokenGrantAccessTokenLifespan) {
    
    this.refreshTokenGrantAccessTokenLifespan = refreshTokenGrantAccessTokenLifespan;
    return this;
  }

   /**
   * Get refreshTokenGrantAccessTokenLifespan
   * @return refreshTokenGrantAccessTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefreshTokenGrantAccessTokenLifespan() {
    return refreshTokenGrantAccessTokenLifespan;
  }


  public void setRefreshTokenGrantAccessTokenLifespan(String refreshTokenGrantAccessTokenLifespan) {
    this.refreshTokenGrantAccessTokenLifespan = refreshTokenGrantAccessTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans refreshTokenGrantIdTokenLifespan(String refreshTokenGrantIdTokenLifespan) {
    
    this.refreshTokenGrantIdTokenLifespan = refreshTokenGrantIdTokenLifespan;
    return this;
  }

   /**
   * Get refreshTokenGrantIdTokenLifespan
   * @return refreshTokenGrantIdTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefreshTokenGrantIdTokenLifespan() {
    return refreshTokenGrantIdTokenLifespan;
  }


  public void setRefreshTokenGrantIdTokenLifespan(String refreshTokenGrantIdTokenLifespan) {
    this.refreshTokenGrantIdTokenLifespan = refreshTokenGrantIdTokenLifespan;
  }


  public UpdateOAuth2ClientLifespans refreshTokenGrantRefreshTokenLifespan(String refreshTokenGrantRefreshTokenLifespan) {
    
    this.refreshTokenGrantRefreshTokenLifespan = refreshTokenGrantRefreshTokenLifespan;
    return this;
  }

   /**
   * Get refreshTokenGrantRefreshTokenLifespan
   * @return refreshTokenGrantRefreshTokenLifespan
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefreshTokenGrantRefreshTokenLifespan() {
    return refreshTokenGrantRefreshTokenLifespan;
  }


  public void setRefreshTokenGrantRefreshTokenLifespan(String refreshTokenGrantRefreshTokenLifespan) {
    this.refreshTokenGrantRefreshTokenLifespan = refreshTokenGrantRefreshTokenLifespan;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdateOAuth2ClientLifespans updateOAuth2ClientLifespans = (UpdateOAuth2ClientLifespans) o;
    return Objects.equals(this.authorizationCodeGrantAccessTokenLifespan, updateOAuth2ClientLifespans.authorizationCodeGrantAccessTokenLifespan) &&
        Objects.equals(this.authorizationCodeGrantIdTokenLifespan, updateOAuth2ClientLifespans.authorizationCodeGrantIdTokenLifespan) &&
        Objects.equals(this.authorizationCodeGrantRefreshTokenLifespan, updateOAuth2ClientLifespans.authorizationCodeGrantRefreshTokenLifespan) &&
        Objects.equals(this.clientCredentialsGrantAccessTokenLifespan, updateOAuth2ClientLifespans.clientCredentialsGrantAccessTokenLifespan) &&
        Objects.equals(this.implicitGrantAccessTokenLifespan, updateOAuth2ClientLifespans.implicitGrantAccessTokenLifespan) &&
        Objects.equals(this.implicitGrantIdTokenLifespan, updateOAuth2ClientLifespans.implicitGrantIdTokenLifespan) &&
        Objects.equals(this.jwtBearerGrantAccessTokenLifespan, updateOAuth2ClientLifespans.jwtBearerGrantAccessTokenLifespan) &&
        Objects.equals(this.passwordGrantAccessTokenLifespan, updateOAuth2ClientLifespans.passwordGrantAccessTokenLifespan) &&
        Objects.equals(this.passwordGrantRefreshTokenLifespan, updateOAuth2ClientLifespans.passwordGrantRefreshTokenLifespan) &&
        Objects.equals(this.refreshTokenGrantAccessTokenLifespan, updateOAuth2ClientLifespans.refreshTokenGrantAccessTokenLifespan) &&
        Objects.equals(this.refreshTokenGrantIdTokenLifespan, updateOAuth2ClientLifespans.refreshTokenGrantIdTokenLifespan) &&
        Objects.equals(this.refreshTokenGrantRefreshTokenLifespan, updateOAuth2ClientLifespans.refreshTokenGrantRefreshTokenLifespan);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(authorizationCodeGrantAccessTokenLifespan, authorizationCodeGrantIdTokenLifespan, authorizationCodeGrantRefreshTokenLifespan, clientCredentialsGrantAccessTokenLifespan, implicitGrantAccessTokenLifespan, implicitGrantIdTokenLifespan, jwtBearerGrantAccessTokenLifespan, passwordGrantAccessTokenLifespan, passwordGrantRefreshTokenLifespan, refreshTokenGrantAccessTokenLifespan, refreshTokenGrantIdTokenLifespan, refreshTokenGrantRefreshTokenLifespan);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateOAuth2ClientLifespans {\n");
    sb.append("    authorizationCodeGrantAccessTokenLifespan: ").append(toIndentedString(authorizationCodeGrantAccessTokenLifespan)).append("\n");
    sb.append("    authorizationCodeGrantIdTokenLifespan: ").append(toIndentedString(authorizationCodeGrantIdTokenLifespan)).append("\n");
    sb.append("    authorizationCodeGrantRefreshTokenLifespan: ").append(toIndentedString(authorizationCodeGrantRefreshTokenLifespan)).append("\n");
    sb.append("    clientCredentialsGrantAccessTokenLifespan: ").append(toIndentedString(clientCredentialsGrantAccessTokenLifespan)).append("\n");
    sb.append("    implicitGrantAccessTokenLifespan: ").append(toIndentedString(implicitGrantAccessTokenLifespan)).append("\n");
    sb.append("    implicitGrantIdTokenLifespan: ").append(toIndentedString(implicitGrantIdTokenLifespan)).append("\n");
    sb.append("    jwtBearerGrantAccessTokenLifespan: ").append(toIndentedString(jwtBearerGrantAccessTokenLifespan)).append("\n");
    sb.append("    passwordGrantAccessTokenLifespan: ").append(toIndentedString(passwordGrantAccessTokenLifespan)).append("\n");
    sb.append("    passwordGrantRefreshTokenLifespan: ").append(toIndentedString(passwordGrantRefreshTokenLifespan)).append("\n");
    sb.append("    refreshTokenGrantAccessTokenLifespan: ").append(toIndentedString(refreshTokenGrantAccessTokenLifespan)).append("\n");
    sb.append("    refreshTokenGrantIdTokenLifespan: ").append(toIndentedString(refreshTokenGrantIdTokenLifespan)).append("\n");
    sb.append("    refreshTokenGrantRefreshTokenLifespan: ").append(toIndentedString(refreshTokenGrantRefreshTokenLifespan)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("authorization_code_grant_access_token_lifespan");
    openapiFields.add("authorization_code_grant_id_token_lifespan");
    openapiFields.add("authorization_code_grant_refresh_token_lifespan");
    openapiFields.add("client_credentials_grant_access_token_lifespan");
    openapiFields.add("implicit_grant_access_token_lifespan");
    openapiFields.add("implicit_grant_id_token_lifespan");
    openapiFields.add("jwt_bearer_grant_access_token_lifespan");
    openapiFields.add("password_grant_access_token_lifespan");
    openapiFields.add("password_grant_refresh_token_lifespan");
    openapiFields.add("refresh_token_grant_access_token_lifespan");
    openapiFields.add("refresh_token_grant_id_token_lifespan");
    openapiFields.add("refresh_token_grant_refresh_token_lifespan");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to UpdateOAuth2ClientLifespans
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (UpdateOAuth2ClientLifespans.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in UpdateOAuth2ClientLifespans is not found in the empty JSON string", UpdateOAuth2ClientLifespans.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!UpdateOAuth2ClientLifespans.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UpdateOAuth2ClientLifespans` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("authorization_code_grant_access_token_lifespan") != null && !jsonObj.get("authorization_code_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `authorization_code_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("authorization_code_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("authorization_code_grant_id_token_lifespan") != null && !jsonObj.get("authorization_code_grant_id_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `authorization_code_grant_id_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("authorization_code_grant_id_token_lifespan").toString()));
      }
      if (jsonObj.get("authorization_code_grant_refresh_token_lifespan") != null && !jsonObj.get("authorization_code_grant_refresh_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `authorization_code_grant_refresh_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("authorization_code_grant_refresh_token_lifespan").toString()));
      }
      if (jsonObj.get("client_credentials_grant_access_token_lifespan") != null && !jsonObj.get("client_credentials_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `client_credentials_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("client_credentials_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("implicit_grant_access_token_lifespan") != null && !jsonObj.get("implicit_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `implicit_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("implicit_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("implicit_grant_id_token_lifespan") != null && !jsonObj.get("implicit_grant_id_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `implicit_grant_id_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("implicit_grant_id_token_lifespan").toString()));
      }
      if (jsonObj.get("jwt_bearer_grant_access_token_lifespan") != null && !jsonObj.get("jwt_bearer_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `jwt_bearer_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("jwt_bearer_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("password_grant_access_token_lifespan") != null && !jsonObj.get("password_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `password_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("password_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("password_grant_refresh_token_lifespan") != null && !jsonObj.get("password_grant_refresh_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `password_grant_refresh_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("password_grant_refresh_token_lifespan").toString()));
      }
      if (jsonObj.get("refresh_token_grant_access_token_lifespan") != null && !jsonObj.get("refresh_token_grant_access_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `refresh_token_grant_access_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("refresh_token_grant_access_token_lifespan").toString()));
      }
      if (jsonObj.get("refresh_token_grant_id_token_lifespan") != null && !jsonObj.get("refresh_token_grant_id_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `refresh_token_grant_id_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("refresh_token_grant_id_token_lifespan").toString()));
      }
      if (jsonObj.get("refresh_token_grant_refresh_token_lifespan") != null && !jsonObj.get("refresh_token_grant_refresh_token_lifespan").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `refresh_token_grant_refresh_token_lifespan` to be a primitive type in the JSON string but got `%s`", jsonObj.get("refresh_token_grant_refresh_token_lifespan").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UpdateOAuth2ClientLifespans.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UpdateOAuth2ClientLifespans' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UpdateOAuth2ClientLifespans> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UpdateOAuth2ClientLifespans.class));

       return (TypeAdapter<T>) new TypeAdapter<UpdateOAuth2ClientLifespans>() {
           @Override
           public void write(JsonWriter out, UpdateOAuth2ClientLifespans value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UpdateOAuth2ClientLifespans read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UpdateOAuth2ClientLifespans given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UpdateOAuth2ClientLifespans
  * @throws IOException if the JSON string is invalid with respect to UpdateOAuth2ClientLifespans
  */
  public static UpdateOAuth2ClientLifespans fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UpdateOAuth2ClientLifespans.class);
  }

 /**
  * Convert an instance of UpdateOAuth2ClientLifespans to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

