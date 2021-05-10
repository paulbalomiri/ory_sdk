/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.0-alpha.4
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
import java.util.ArrayList;
import java.util.List;
import sh.ory.kratos.model.PluginDevice;
import sh.ory.kratos.model.PluginMount;

/**
 * PluginSettings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-05-10T10:42:51.879955097Z[Etc/UTC]")
public class PluginSettings {
  public static final String SERIALIZED_NAME_ARGS = "Args";
  @SerializedName(SERIALIZED_NAME_ARGS)
  private List<String> args = new ArrayList<>();

  public static final String SERIALIZED_NAME_DEVICES = "Devices";
  @SerializedName(SERIALIZED_NAME_DEVICES)
  private List<PluginDevice> devices = new ArrayList<>();

  public static final String SERIALIZED_NAME_ENV = "Env";
  @SerializedName(SERIALIZED_NAME_ENV)
  private List<String> env = new ArrayList<>();

  public static final String SERIALIZED_NAME_MOUNTS = "Mounts";
  @SerializedName(SERIALIZED_NAME_MOUNTS)
  private List<PluginMount> mounts = new ArrayList<>();


  public PluginSettings args(List<String> args) {
    
    this.args = args;
    return this;
  }

  public PluginSettings addArgsItem(String argsItem) {
    this.args.add(argsItem);
    return this;
  }

   /**
   * args
   * @return args
  **/
  @ApiModelProperty(required = true, value = "args")

  public List<String> getArgs() {
    return args;
  }


  public void setArgs(List<String> args) {
    this.args = args;
  }


  public PluginSettings devices(List<PluginDevice> devices) {
    
    this.devices = devices;
    return this;
  }

  public PluginSettings addDevicesItem(PluginDevice devicesItem) {
    this.devices.add(devicesItem);
    return this;
  }

   /**
   * devices
   * @return devices
  **/
  @ApiModelProperty(required = true, value = "devices")

  public List<PluginDevice> getDevices() {
    return devices;
  }


  public void setDevices(List<PluginDevice> devices) {
    this.devices = devices;
  }


  public PluginSettings env(List<String> env) {
    
    this.env = env;
    return this;
  }

  public PluginSettings addEnvItem(String envItem) {
    this.env.add(envItem);
    return this;
  }

   /**
   * env
   * @return env
  **/
  @ApiModelProperty(required = true, value = "env")

  public List<String> getEnv() {
    return env;
  }


  public void setEnv(List<String> env) {
    this.env = env;
  }


  public PluginSettings mounts(List<PluginMount> mounts) {
    
    this.mounts = mounts;
    return this;
  }

  public PluginSettings addMountsItem(PluginMount mountsItem) {
    this.mounts.add(mountsItem);
    return this;
  }

   /**
   * mounts
   * @return mounts
  **/
  @ApiModelProperty(required = true, value = "mounts")

  public List<PluginMount> getMounts() {
    return mounts;
  }


  public void setMounts(List<PluginMount> mounts) {
    this.mounts = mounts;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PluginSettings pluginSettings = (PluginSettings) o;
    return Objects.equals(this.args, pluginSettings.args) &&
        Objects.equals(this.devices, pluginSettings.devices) &&
        Objects.equals(this.env, pluginSettings.env) &&
        Objects.equals(this.mounts, pluginSettings.mounts);
  }

  @Override
  public int hashCode() {
    return Objects.hash(args, devices, env, mounts);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PluginSettings {\n");
    sb.append("    args: ").append(toIndentedString(args)).append("\n");
    sb.append("    devices: ").append(toIndentedString(devices)).append("\n");
    sb.append("    env: ").append(toIndentedString(env)).append("\n");
    sb.append("    mounts: ").append(toIndentedString(mounts)).append("\n");
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

