//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:ory_kratos_client/src/model/ui_text.dart';
import 'package:built_value/json_object.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'ui_node_input_attributes.g.dart';

/// InputAttributes represents the attributes of an input node
///
/// Properties:
/// * [disabled] - Sets the input's disabled field to true or false.
/// * [label] 
/// * [name] - The input's element name.
/// * [nodeType] - NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"input\".
/// * [onclick] - OnClick may contain javascript which should be executed on click. This is primarily used for WebAuthn.
/// * [pattern] - The input's pattern.
/// * [required_] - Mark this input field as required.
/// * [type] 
/// * [value] - The input's value.
@BuiltValue()
abstract class UiNodeInputAttributes implements Built<UiNodeInputAttributes, UiNodeInputAttributesBuilder> {
  /// Sets the input's disabled field to true or false.
  @BuiltValueField(wireName: r'disabled')
  bool get disabled;

  @BuiltValueField(wireName: r'label')
  UiText? get label;

  /// The input's element name.
  @BuiltValueField(wireName: r'name')
  String get name;

  /// NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"input\".
  @BuiltValueField(wireName: r'node_type')
  String get nodeType;

  /// OnClick may contain javascript which should be executed on click. This is primarily used for WebAuthn.
  @BuiltValueField(wireName: r'onclick')
  String? get onclick;

  /// The input's pattern.
  @BuiltValueField(wireName: r'pattern')
  String? get pattern;

  /// Mark this input field as required.
  @BuiltValueField(wireName: r'required')
  bool? get required_;

  @BuiltValueField(wireName: r'type')
  String get type;

  /// The input's value.
  @BuiltValueField(wireName: r'value')
  JsonObject? get value;

  UiNodeInputAttributes._();

  factory UiNodeInputAttributes([void updates(UiNodeInputAttributesBuilder b)]) = _$UiNodeInputAttributes;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(UiNodeInputAttributesBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<UiNodeInputAttributes> get serializer => _$UiNodeInputAttributesSerializer();
}

class _$UiNodeInputAttributesSerializer implements PrimitiveSerializer<UiNodeInputAttributes> {
  @override
  final Iterable<Type> types = const [UiNodeInputAttributes, _$UiNodeInputAttributes];

  @override
  final String wireName = r'UiNodeInputAttributes';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    UiNodeInputAttributes object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'disabled';
    yield serializers.serialize(
      object.disabled,
      specifiedType: const FullType(bool),
    );
    if (object.label != null) {
      yield r'label';
      yield serializers.serialize(
        object.label,
        specifiedType: const FullType(UiText),
      );
    }
    yield r'name';
    yield serializers.serialize(
      object.name,
      specifiedType: const FullType(String),
    );
    yield r'node_type';
    yield serializers.serialize(
      object.nodeType,
      specifiedType: const FullType(String),
    );
    if (object.onclick != null) {
      yield r'onclick';
      yield serializers.serialize(
        object.onclick,
        specifiedType: const FullType(String),
      );
    }
    if (object.pattern != null) {
      yield r'pattern';
      yield serializers.serialize(
        object.pattern,
        specifiedType: const FullType(String),
      );
    }
    if (object.required_ != null) {
      yield r'required';
      yield serializers.serialize(
        object.required_,
        specifiedType: const FullType(bool),
      );
    }
    yield r'type';
    yield serializers.serialize(
      object.type,
      specifiedType: const FullType(String),
    );
    if (object.value != null) {
      yield r'value';
      yield serializers.serialize(
        object.value,
        specifiedType: const FullType.nullable(JsonObject),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    UiNodeInputAttributes object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required UiNodeInputAttributesBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'disabled':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(bool),
          ) as bool;
          result.disabled = valueDes;
          break;
        case r'label':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(UiText),
          ) as UiText;
          result.label.replace(valueDes);
          break;
        case r'name':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.name = valueDes;
          break;
        case r'node_type':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.nodeType = valueDes;
          break;
        case r'onclick':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.onclick = valueDes;
          break;
        case r'pattern':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.pattern = valueDes;
          break;
        case r'required':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(bool),
          ) as bool;
          result.required_ = valueDes;
          break;
        case r'type':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.type = valueDes;
          break;
        case r'value':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType.nullable(JsonObject),
          ) as JsonObject?;
          if (valueDes == null) continue;
          result.value = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  UiNodeInputAttributes deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = UiNodeInputAttributesBuilder();
    final serializedList = (serialized as Iterable<Object?>).toList();
    final unhandled = <Object?>[];
    _deserializeProperties(
      serializers,
      serialized,
      specifiedType: specifiedType,
      serializedList: serializedList,
      unhandled: unhandled,
      result: result,
    );
    return result.build();
  }
}

