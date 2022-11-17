/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.10.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

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

import sh.ory.kratos.JSON;

/**
 * SubmitSelfServiceSettingsFlowWithTotpMethodBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-11-17T18:49:16.514253300Z[Etc/UTC]")
public class SubmitSelfServiceSettingsFlowWithTotpMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_TOTP_CODE = "totp_code";
  @SerializedName(SERIALIZED_NAME_TOTP_CODE)
  private String totpCode;

  public static final String SERIALIZED_NAME_TOTP_UNLINK = "totp_unlink";
  @SerializedName(SERIALIZED_NAME_TOTP_UNLINK)
  private Boolean totpUnlink;

  public SubmitSelfServiceSettingsFlowWithTotpMethodBody() { 
  }

  public SubmitSelfServiceSettingsFlowWithTotpMethodBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * CSRFToken is the anti-CSRF token
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "CSRFToken is the anti-CSRF token")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceSettingsFlowWithTotpMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to \&quot;totp\&quot; when trying to add, update, or remove a totp pairing.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method  Should be set to \"totp\" when trying to add, update, or remove a totp pairing.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceSettingsFlowWithTotpMethodBody totpCode(String totpCode) {
    
    this.totpCode = totpCode;
    return this;
  }

   /**
   * ValidationTOTP must contain a valid TOTP based on the
   * @return totpCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "ValidationTOTP must contain a valid TOTP based on the")

  public String getTotpCode() {
    return totpCode;
  }


  public void setTotpCode(String totpCode) {
    this.totpCode = totpCode;
  }


  public SubmitSelfServiceSettingsFlowWithTotpMethodBody totpUnlink(Boolean totpUnlink) {
    
    this.totpUnlink = totpUnlink;
    return this;
  }

   /**
   * UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.
   * @return totpUnlink
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.")

  public Boolean getTotpUnlink() {
    return totpUnlink;
  }


  public void setTotpUnlink(Boolean totpUnlink) {
    this.totpUnlink = totpUnlink;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceSettingsFlowWithTotpMethodBody submitSelfServiceSettingsFlowWithTotpMethodBody = (SubmitSelfServiceSettingsFlowWithTotpMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceSettingsFlowWithTotpMethodBody.csrfToken) &&
        Objects.equals(this.method, submitSelfServiceSettingsFlowWithTotpMethodBody.method) &&
        Objects.equals(this.totpCode, submitSelfServiceSettingsFlowWithTotpMethodBody.totpCode) &&
        Objects.equals(this.totpUnlink, submitSelfServiceSettingsFlowWithTotpMethodBody.totpUnlink);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, totpCode, totpUnlink);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceSettingsFlowWithTotpMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    totpCode: ").append(toIndentedString(totpCode)).append("\n");
    sb.append("    totpUnlink: ").append(toIndentedString(totpUnlink)).append("\n");
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
    openapiFields.add("csrf_token");
    openapiFields.add("method");
    openapiFields.add("totp_code");
    openapiFields.add("totp_unlink");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("method");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to SubmitSelfServiceSettingsFlowWithTotpMethodBody
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (SubmitSelfServiceSettingsFlowWithTotpMethodBody.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in SubmitSelfServiceSettingsFlowWithTotpMethodBody is not found in the empty JSON string", SubmitSelfServiceSettingsFlowWithTotpMethodBody.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!SubmitSelfServiceSettingsFlowWithTotpMethodBody.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `SubmitSelfServiceSettingsFlowWithTotpMethodBody` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : SubmitSelfServiceSettingsFlowWithTotpMethodBody.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (jsonObj.get("csrf_token") != null && !jsonObj.get("csrf_token").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `csrf_token` to be a primitive type in the JSON string but got `%s`", jsonObj.get("csrf_token").toString()));
      }
      if (jsonObj.get("method") != null && !jsonObj.get("method").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `method` to be a primitive type in the JSON string but got `%s`", jsonObj.get("method").toString()));
      }
      if (jsonObj.get("totp_code") != null && !jsonObj.get("totp_code").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `totp_code` to be a primitive type in the JSON string but got `%s`", jsonObj.get("totp_code").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!SubmitSelfServiceSettingsFlowWithTotpMethodBody.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'SubmitSelfServiceSettingsFlowWithTotpMethodBody' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<SubmitSelfServiceSettingsFlowWithTotpMethodBody> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(SubmitSelfServiceSettingsFlowWithTotpMethodBody.class));

       return (TypeAdapter<T>) new TypeAdapter<SubmitSelfServiceSettingsFlowWithTotpMethodBody>() {
           @Override
           public void write(JsonWriter out, SubmitSelfServiceSettingsFlowWithTotpMethodBody value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public SubmitSelfServiceSettingsFlowWithTotpMethodBody read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of SubmitSelfServiceSettingsFlowWithTotpMethodBody given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of SubmitSelfServiceSettingsFlowWithTotpMethodBody
  * @throws IOException if the JSON string is invalid with respect to SubmitSelfServiceSettingsFlowWithTotpMethodBody
  */
  public static SubmitSelfServiceSettingsFlowWithTotpMethodBody fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, SubmitSelfServiceSettingsFlowWithTotpMethodBody.class);
  }

 /**
  * Convert an instance of SubmitSelfServiceSettingsFlowWithTotpMethodBody to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

