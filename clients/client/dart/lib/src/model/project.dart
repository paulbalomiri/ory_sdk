//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_collection/built_collection.dart';
import 'package:ory_client/src/model/project_services.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'project.g.dart';

/// Project
///
/// Properties:
/// * [id] - The project's ID.
/// * [name] - The name of the project.
/// * [revisionId] - The configuration revision ID.
/// * [services] 
/// * [slug] - The project's slug
/// * [state] - The state of the project. running Running halted Halted
@BuiltValue()
abstract class Project implements Built<Project, ProjectBuilder> {
  /// The project's ID.
  @BuiltValueField(wireName: r'id')
  String get id;

  /// The name of the project.
  @BuiltValueField(wireName: r'name')
  String get name;

  /// The configuration revision ID.
  @BuiltValueField(wireName: r'revision_id')
  String get revisionId;

  @BuiltValueField(wireName: r'services')
  ProjectServices get services;

  /// The project's slug
  @BuiltValueField(wireName: r'slug')
  String get slug;

  /// The state of the project. running Running halted Halted
  @BuiltValueField(wireName: r'state')
  ProjectStateEnum get state;
  // enum stateEnum {  running,  halted,  };

  Project._();

  factory Project([void updates(ProjectBuilder b)]) = _$Project;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ProjectBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<Project> get serializer => _$ProjectSerializer();
}

class _$ProjectSerializer implements PrimitiveSerializer<Project> {
  @override
  final Iterable<Type> types = const [Project, _$Project];

  @override
  final String wireName = r'Project';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    Project object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'id';
    yield serializers.serialize(
      object.id,
      specifiedType: const FullType(String),
    );
    yield r'name';
    yield serializers.serialize(
      object.name,
      specifiedType: const FullType(String),
    );
    yield r'revision_id';
    yield serializers.serialize(
      object.revisionId,
      specifiedType: const FullType(String),
    );
    yield r'services';
    yield serializers.serialize(
      object.services,
      specifiedType: const FullType(ProjectServices),
    );
    yield r'slug';
    yield serializers.serialize(
      object.slug,
      specifiedType: const FullType(String),
    );
    yield r'state';
    yield serializers.serialize(
      object.state,
      specifiedType: const FullType(ProjectStateEnum),
    );
  }

  @override
  Object serialize(
    Serializers serializers,
    Project object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ProjectBuilder result,
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
        case r'name':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.name = valueDes;
          break;
        case r'revision_id':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.revisionId = valueDes;
          break;
        case r'services':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(ProjectServices),
          ) as ProjectServices;
          result.services.replace(valueDes);
          break;
        case r'slug':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.slug = valueDes;
          break;
        case r'state':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(ProjectStateEnum),
          ) as ProjectStateEnum;
          result.state = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  Project deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ProjectBuilder();
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

class ProjectStateEnum extends EnumClass {

  /// The state of the project. running Running halted Halted
  @BuiltValueEnumConst(wireName: r'running')
  static const ProjectStateEnum running = _$projectStateEnum_running;
  /// The state of the project. running Running halted Halted
  @BuiltValueEnumConst(wireName: r'halted')
  static const ProjectStateEnum halted = _$projectStateEnum_halted;

  static Serializer<ProjectStateEnum> get serializer => _$projectStateEnumSerializer;

  const ProjectStateEnum._(String name): super(name);

  static BuiltSet<ProjectStateEnum> get values => _$projectStateEnumValues;
  static ProjectStateEnum valueOf(String name) => _$projectStateEnumValueOf(name);
}

