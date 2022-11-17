//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/json_object.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'identity_schema.g.dart';

/// IdentitySchema
///
/// Properties:
/// * [id] - The ID of the Identity JSON Schema
/// * [schema] - The actual Identity JSON Schema
@BuiltValue()
abstract class IdentitySchema implements Built<IdentitySchema, IdentitySchemaBuilder> {
  /// The ID of the Identity JSON Schema
  @BuiltValueField(wireName: r'id')
  String? get id;

  /// The actual Identity JSON Schema
  @BuiltValueField(wireName: r'schema')
  JsonObject? get schema;

  IdentitySchema._();

  factory IdentitySchema([void updates(IdentitySchemaBuilder b)]) = _$IdentitySchema;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(IdentitySchemaBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<IdentitySchema> get serializer => _$IdentitySchemaSerializer();
}

class _$IdentitySchemaSerializer implements PrimitiveSerializer<IdentitySchema> {
  @override
  final Iterable<Type> types = const [IdentitySchema, _$IdentitySchema];

  @override
  final String wireName = r'IdentitySchema';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    IdentitySchema object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    if (object.id != null) {
      yield r'id';
      yield serializers.serialize(
        object.id,
        specifiedType: const FullType(String),
      );
    }
    if (object.schema != null) {
      yield r'schema';
      yield serializers.serialize(
        object.schema,
        specifiedType: const FullType(JsonObject),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    IdentitySchema object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required IdentitySchemaBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'id':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.id = valueDes;
          break;
        case r'schema':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(JsonObject),
          ) as JsonObject;
          result.schema = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  IdentitySchema deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = IdentitySchemaBuilder();
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

