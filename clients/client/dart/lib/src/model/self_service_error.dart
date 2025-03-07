//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/json_object.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'self_service_error.g.dart';

/// SelfServiceError
///
/// Properties:
/// * [createdAt] - CreatedAt is a helper struct field for gobuffalo.pop.
/// * [error] 
/// * [id] - ID of the error container.
/// * [updatedAt] - UpdatedAt is a helper struct field for gobuffalo.pop.
@BuiltValue()
abstract class SelfServiceError implements Built<SelfServiceError, SelfServiceErrorBuilder> {
  /// CreatedAt is a helper struct field for gobuffalo.pop.
  @BuiltValueField(wireName: r'created_at')
  DateTime? get createdAt;

  @BuiltValueField(wireName: r'error')
  JsonObject? get error;

  /// ID of the error container.
  @BuiltValueField(wireName: r'id')
  String get id;

  /// UpdatedAt is a helper struct field for gobuffalo.pop.
  @BuiltValueField(wireName: r'updated_at')
  DateTime? get updatedAt;

  SelfServiceError._();

  factory SelfServiceError([void updates(SelfServiceErrorBuilder b)]) = _$SelfServiceError;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(SelfServiceErrorBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<SelfServiceError> get serializer => _$SelfServiceErrorSerializer();
}

class _$SelfServiceErrorSerializer implements PrimitiveSerializer<SelfServiceError> {
  @override
  final Iterable<Type> types = const [SelfServiceError, _$SelfServiceError];

  @override
  final String wireName = r'SelfServiceError';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    SelfServiceError object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    if (object.createdAt != null) {
      yield r'created_at';
      yield serializers.serialize(
        object.createdAt,
        specifiedType: const FullType(DateTime),
      );
    }
    if (object.error != null) {
      yield r'error';
      yield serializers.serialize(
        object.error,
        specifiedType: const FullType(JsonObject),
      );
    }
    yield r'id';
    yield serializers.serialize(
      object.id,
      specifiedType: const FullType(String),
    );
    if (object.updatedAt != null) {
      yield r'updated_at';
      yield serializers.serialize(
        object.updatedAt,
        specifiedType: const FullType(DateTime),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    SelfServiceError object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required SelfServiceErrorBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'created_at':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(DateTime),
          ) as DateTime;
          result.createdAt = valueDes;
          break;
        case r'error':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(JsonObject),
          ) as JsonObject;
          result.error = valueDes;
          break;
        case r'id':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.id = valueDes;
          break;
        case r'updated_at':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(DateTime),
          ) as DateTime;
          result.updatedAt = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  SelfServiceError deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = SelfServiceErrorBuilder();
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

