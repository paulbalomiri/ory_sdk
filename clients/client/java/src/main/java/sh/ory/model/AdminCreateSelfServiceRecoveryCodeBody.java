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
 * AdminCreateSelfServiceRecoveryCodeBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-10-24T09:56:23.976880967Z[Etc/UTC]")
public class AdminCreateSelfServiceRecoveryCodeBody {
  public static final String SERIALIZED_NAME_EXPIRES_IN = "expires_in";
  @SerializedName(SERIALIZED_NAME_EXPIRES_IN)
  private String expiresIn;

  public static final String SERIALIZED_NAME_IDENTITY_ID = "identity_id";
  @SerializedName(SERIALIZED_NAME_IDENTITY_ID)
  private String identityId;

  public AdminCreateSelfServiceRecoveryCodeBody() { 
  }

  public AdminCreateSelfServiceRecoveryCodeBody expiresIn(String expiresIn) {
    
    this.expiresIn = expiresIn;
    return this;
  }

   /**
   * Code Expires In  The recovery code will expire at that point in time. Defaults to the configuration value of &#x60;selfservice.flows.recovery.request_lifespan&#x60;.
   * @return expiresIn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Code Expires In  The recovery code will expire at that point in time. Defaults to the configuration value of `selfservice.flows.recovery.request_lifespan`.")

  public String getExpiresIn() {
    return expiresIn;
  }


  public void setExpiresIn(String expiresIn) {
    this.expiresIn = expiresIn;
  }


  public AdminCreateSelfServiceRecoveryCodeBody identityId(String identityId) {
    
    this.identityId = identityId;
    return this;
  }

   /**
   * Identity to Recover  The identity&#39;s ID you wish to recover.
   * @return identityId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Identity to Recover  The identity's ID you wish to recover.")

  public String getIdentityId() {
    return identityId;
  }


  public void setIdentityId(String identityId) {
    this.identityId = identityId;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AdminCreateSelfServiceRecoveryCodeBody adminCreateSelfServiceRecoveryCodeBody = (AdminCreateSelfServiceRecoveryCodeBody) o;
    return Objects.equals(this.expiresIn, adminCreateSelfServiceRecoveryCodeBody.expiresIn) &&
        Objects.equals(this.identityId, adminCreateSelfServiceRecoveryCodeBody.identityId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(expiresIn, identityId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AdminCreateSelfServiceRecoveryCodeBody {\n");
    sb.append("    expiresIn: ").append(toIndentedString(expiresIn)).append("\n");
    sb.append("    identityId: ").append(toIndentedString(identityId)).append("\n");
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
    openapiFields.add("expires_in");
    openapiFields.add("identity_id");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("identity_id");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to AdminCreateSelfServiceRecoveryCodeBody
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (AdminCreateSelfServiceRecoveryCodeBody.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in AdminCreateSelfServiceRecoveryCodeBody is not found in the empty JSON string", AdminCreateSelfServiceRecoveryCodeBody.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AdminCreateSelfServiceRecoveryCodeBody.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AdminCreateSelfServiceRecoveryCodeBody` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : AdminCreateSelfServiceRecoveryCodeBody.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (jsonObj.get("expires_in") != null && !jsonObj.get("expires_in").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `expires_in` to be a primitive type in the JSON string but got `%s`", jsonObj.get("expires_in").toString()));
      }
      if (jsonObj.get("identity_id") != null && !jsonObj.get("identity_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `identity_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("identity_id").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AdminCreateSelfServiceRecoveryCodeBody.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AdminCreateSelfServiceRecoveryCodeBody' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AdminCreateSelfServiceRecoveryCodeBody> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AdminCreateSelfServiceRecoveryCodeBody.class));

       return (TypeAdapter<T>) new TypeAdapter<AdminCreateSelfServiceRecoveryCodeBody>() {
           @Override
           public void write(JsonWriter out, AdminCreateSelfServiceRecoveryCodeBody value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AdminCreateSelfServiceRecoveryCodeBody read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AdminCreateSelfServiceRecoveryCodeBody given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AdminCreateSelfServiceRecoveryCodeBody
  * @throws IOException if the JSON string is invalid with respect to AdminCreateSelfServiceRecoveryCodeBody
  */
  public static AdminCreateSelfServiceRecoveryCodeBody fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AdminCreateSelfServiceRecoveryCodeBody.class);
  }

 /**
  * Convert an instance of AdminCreateSelfServiceRecoveryCodeBody to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

