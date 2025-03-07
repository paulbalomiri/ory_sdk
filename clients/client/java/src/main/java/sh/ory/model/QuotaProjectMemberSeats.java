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
 * QuotaProjectMemberSeats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-10-24T09:56:23.976880967Z[Etc/UTC]")
public class QuotaProjectMemberSeats {
  public static final String SERIALIZED_NAME_PROJECT_ID = "project_id";
  @SerializedName(SERIALIZED_NAME_PROJECT_ID)
  private String projectId;

  public static final String SERIALIZED_NAME_REMAINING_SEATS = "remaining_seats";
  @SerializedName(SERIALIZED_NAME_REMAINING_SEATS)
  private Long remainingSeats;

  public static final String SERIALIZED_NAME_TOTAL_SEATS = "total_seats";
  @SerializedName(SERIALIZED_NAME_TOTAL_SEATS)
  private Long totalSeats;

  public QuotaProjectMemberSeats() { 
  }

  public QuotaProjectMemberSeats projectId(String projectId) {
    
    this.projectId = projectId;
    return this;
  }

   /**
   * Get projectId
   * @return projectId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProjectId() {
    return projectId;
  }


  public void setProjectId(String projectId) {
    this.projectId = projectId;
  }


  public QuotaProjectMemberSeats remainingSeats(Long remainingSeats) {
    
    this.remainingSeats = remainingSeats;
    return this;
  }

   /**
   * Get remainingSeats
   * @return remainingSeats
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRemainingSeats() {
    return remainingSeats;
  }


  public void setRemainingSeats(Long remainingSeats) {
    this.remainingSeats = remainingSeats;
  }


  public QuotaProjectMemberSeats totalSeats(Long totalSeats) {
    
    this.totalSeats = totalSeats;
    return this;
  }

   /**
   * Get totalSeats
   * @return totalSeats
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalSeats() {
    return totalSeats;
  }


  public void setTotalSeats(Long totalSeats) {
    this.totalSeats = totalSeats;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    QuotaProjectMemberSeats quotaProjectMemberSeats = (QuotaProjectMemberSeats) o;
    return Objects.equals(this.projectId, quotaProjectMemberSeats.projectId) &&
        Objects.equals(this.remainingSeats, quotaProjectMemberSeats.remainingSeats) &&
        Objects.equals(this.totalSeats, quotaProjectMemberSeats.totalSeats);
  }

  @Override
  public int hashCode() {
    return Objects.hash(projectId, remainingSeats, totalSeats);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class QuotaProjectMemberSeats {\n");
    sb.append("    projectId: ").append(toIndentedString(projectId)).append("\n");
    sb.append("    remainingSeats: ").append(toIndentedString(remainingSeats)).append("\n");
    sb.append("    totalSeats: ").append(toIndentedString(totalSeats)).append("\n");
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
    openapiFields.add("project_id");
    openapiFields.add("remaining_seats");
    openapiFields.add("total_seats");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to QuotaProjectMemberSeats
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (QuotaProjectMemberSeats.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in QuotaProjectMemberSeats is not found in the empty JSON string", QuotaProjectMemberSeats.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!QuotaProjectMemberSeats.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `QuotaProjectMemberSeats` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("project_id") != null && !jsonObj.get("project_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `project_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("project_id").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!QuotaProjectMemberSeats.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'QuotaProjectMemberSeats' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<QuotaProjectMemberSeats> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(QuotaProjectMemberSeats.class));

       return (TypeAdapter<T>) new TypeAdapter<QuotaProjectMemberSeats>() {
           @Override
           public void write(JsonWriter out, QuotaProjectMemberSeats value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public QuotaProjectMemberSeats read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of QuotaProjectMemberSeats given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of QuotaProjectMemberSeats
  * @throws IOException if the JSON string is invalid with respect to QuotaProjectMemberSeats
  */
  public static QuotaProjectMemberSeats fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, QuotaProjectMemberSeats.class);
  }

 /**
  * Convert an instance of QuotaProjectMemberSeats to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

