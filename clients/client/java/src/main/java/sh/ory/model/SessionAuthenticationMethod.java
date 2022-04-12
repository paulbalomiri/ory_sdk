/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.161
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
import java.time.OffsetDateTime;
import sh.ory.model.AuthenticatorAssuranceLevel;

/**
 * A singular authenticator used during authentication / login.
 */
@ApiModel(description = "A singular authenticator used during authentication / login.")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-04-12T09:46:31.378269840Z[Etc/UTC]")
public class SessionAuthenticationMethod {
  public static final String SERIALIZED_NAME_AAL = "aal";
  @SerializedName(SERIALIZED_NAME_AAL)
  private AuthenticatorAssuranceLevel aal;

  public static final String SERIALIZED_NAME_COMPLETED_AT = "completed_at";
  @SerializedName(SERIALIZED_NAME_COMPLETED_AT)
  private OffsetDateTime completedAt;

  /**
   * Gets or Sets method
   */
  @JsonAdapter(MethodEnum.Adapter.class)
  public enum MethodEnum {
    LINK_RECOVERY("link_recovery"),
    
    PASSWORD("password"),
    
    TOTP("totp"),
    
    OIDC("oidc"),
    
    WEBAUTHN("webauthn"),
    
    LOOKUP_SECRET("lookup_secret");

    private String value;

    MethodEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static MethodEnum fromValue(String value) {
      for (MethodEnum b : MethodEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<MethodEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final MethodEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public MethodEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return MethodEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private MethodEnum method;

  public SessionAuthenticationMethod() { 
  }

  public SessionAuthenticationMethod aal(AuthenticatorAssuranceLevel aal) {
    
    this.aal = aal;
    return this;
  }

   /**
   * Get aal
   * @return aal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public AuthenticatorAssuranceLevel getAal() {
    return aal;
  }


  public void setAal(AuthenticatorAssuranceLevel aal) {
    this.aal = aal;
  }


  public SessionAuthenticationMethod completedAt(OffsetDateTime completedAt) {
    
    this.completedAt = completedAt;
    return this;
  }

   /**
   * When the authentication challenge was completed.
   * @return completedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "When the authentication challenge was completed.")

  public OffsetDateTime getCompletedAt() {
    return completedAt;
  }


  public void setCompletedAt(OffsetDateTime completedAt) {
    this.completedAt = completedAt;
  }


  public SessionAuthenticationMethod method(MethodEnum method) {
    
    this.method = method;
    return this;
  }

   /**
   * Get method
   * @return method
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public MethodEnum getMethod() {
    return method;
  }


  public void setMethod(MethodEnum method) {
    this.method = method;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SessionAuthenticationMethod sessionAuthenticationMethod = (SessionAuthenticationMethod) o;
    return Objects.equals(this.aal, sessionAuthenticationMethod.aal) &&
        Objects.equals(this.completedAt, sessionAuthenticationMethod.completedAt) &&
        Objects.equals(this.method, sessionAuthenticationMethod.method);
  }

  @Override
  public int hashCode() {
    return Objects.hash(aal, completedAt, method);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SessionAuthenticationMethod {\n");
    sb.append("    aal: ").append(toIndentedString(aal)).append("\n");
    sb.append("    completedAt: ").append(toIndentedString(completedAt)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
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

}

