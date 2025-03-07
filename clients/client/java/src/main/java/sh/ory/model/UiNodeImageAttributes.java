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
 * UiNodeImageAttributes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-10-24T09:56:23.976880967Z[Etc/UTC]")
public class UiNodeImageAttributes {
  public static final String SERIALIZED_NAME_HEIGHT = "height";
  @SerializedName(SERIALIZED_NAME_HEIGHT)
  private Long height;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_NODE_TYPE = "node_type";
  @SerializedName(SERIALIZED_NAME_NODE_TYPE)
  private String nodeType;

  public static final String SERIALIZED_NAME_SRC = "src";
  @SerializedName(SERIALIZED_NAME_SRC)
  private String src;

  public static final String SERIALIZED_NAME_WIDTH = "width";
  @SerializedName(SERIALIZED_NAME_WIDTH)
  private Long width;

  public UiNodeImageAttributes() { 
  }

  public UiNodeImageAttributes height(Long height) {
    
    this.height = height;
    return this;
  }

   /**
   * Height of the image
   * @return height
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Height of the image")

  public Long getHeight() {
    return height;
  }


  public void setHeight(Long height) {
    this.height = height;
  }


  public UiNodeImageAttributes id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * A unique identifier
   * @return id
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "A unique identifier")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public UiNodeImageAttributes nodeType(String nodeType) {
    
    this.nodeType = nodeType;
    return this;
  }

   /**
   * NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;.
   * @return nodeType
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"img\".")

  public String getNodeType() {
    return nodeType;
  }


  public void setNodeType(String nodeType) {
    this.nodeType = nodeType;
  }


  public UiNodeImageAttributes src(String src) {
    
    this.src = src;
    return this;
  }

   /**
   * The image&#39;s source URL.  format: uri
   * @return src
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "The image's source URL.  format: uri")

  public String getSrc() {
    return src;
  }


  public void setSrc(String src) {
    this.src = src;
  }


  public UiNodeImageAttributes width(Long width) {
    
    this.width = width;
    return this;
  }

   /**
   * Width of the image
   * @return width
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Width of the image")

  public Long getWidth() {
    return width;
  }


  public void setWidth(Long width) {
    this.width = width;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UiNodeImageAttributes uiNodeImageAttributes = (UiNodeImageAttributes) o;
    return Objects.equals(this.height, uiNodeImageAttributes.height) &&
        Objects.equals(this.id, uiNodeImageAttributes.id) &&
        Objects.equals(this.nodeType, uiNodeImageAttributes.nodeType) &&
        Objects.equals(this.src, uiNodeImageAttributes.src) &&
        Objects.equals(this.width, uiNodeImageAttributes.width);
  }

  @Override
  public int hashCode() {
    return Objects.hash(height, id, nodeType, src, width);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UiNodeImageAttributes {\n");
    sb.append("    height: ").append(toIndentedString(height)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    nodeType: ").append(toIndentedString(nodeType)).append("\n");
    sb.append("    src: ").append(toIndentedString(src)).append("\n");
    sb.append("    width: ").append(toIndentedString(width)).append("\n");
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
    openapiFields.add("height");
    openapiFields.add("id");
    openapiFields.add("node_type");
    openapiFields.add("src");
    openapiFields.add("width");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("height");
    openapiRequiredFields.add("id");
    openapiRequiredFields.add("node_type");
    openapiRequiredFields.add("src");
    openapiRequiredFields.add("width");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to UiNodeImageAttributes
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (UiNodeImageAttributes.openapiRequiredFields.isEmpty()) {
          return;
        } else { // has required fields
          throw new IllegalArgumentException(String.format("The required field(s) %s in UiNodeImageAttributes is not found in the empty JSON string", UiNodeImageAttributes.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!UiNodeImageAttributes.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UiNodeImageAttributes` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UiNodeImageAttributes.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (jsonObj.get("id") != null && !jsonObj.get("id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("id").toString()));
      }
      if (jsonObj.get("node_type") != null && !jsonObj.get("node_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `node_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("node_type").toString()));
      }
      if (jsonObj.get("src") != null && !jsonObj.get("src").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `src` to be a primitive type in the JSON string but got `%s`", jsonObj.get("src").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UiNodeImageAttributes.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UiNodeImageAttributes' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UiNodeImageAttributes> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UiNodeImageAttributes.class));

       return (TypeAdapter<T>) new TypeAdapter<UiNodeImageAttributes>() {
           @Override
           public void write(JsonWriter out, UiNodeImageAttributes value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UiNodeImageAttributes read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UiNodeImageAttributes given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UiNodeImageAttributes
  * @throws IOException if the JSON string is invalid with respect to UiNodeImageAttributes
  */
  public static UiNodeImageAttributes fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UiNodeImageAttributes.class);
  }

 /**
  * Convert an instance of UiNodeImageAttributes to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

