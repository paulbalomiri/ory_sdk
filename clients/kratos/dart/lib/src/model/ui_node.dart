//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_collection/built_collection.dart';
import 'package:ory_kratos_client/src/model/ui_text.dart';
import 'package:ory_kratos_client/src/model/ui_node_meta.dart';
import 'package:ory_kratos_client/src/model/ui_node_attributes.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'ui_node.g.dart';

/// Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.
///
/// Properties:
/// * [attributes] 
/// * [group] - Group specifies which group (e.g. password authenticator) this node belongs to.
/// * [messages] 
/// * [meta] 
/// * [type] - The node's type
@BuiltValue()
abstract class UiNode implements Built<UiNode, UiNodeBuilder> {
  @BuiltValueField(wireName: r'attributes')
  UiNodeAttributes get attributes;

  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueField(wireName: r'group')
  UiNodeGroupEnum get group;
  // enum groupEnum {  default,  password,  oidc,  profile,  link,  totp,  lookup_secret,  webauthn,  };

  @BuiltValueField(wireName: r'messages')
  BuiltList<UiText> get messages;

  @BuiltValueField(wireName: r'meta')
  UiNodeMeta get meta;

  /// The node's type
  @BuiltValueField(wireName: r'type')
  UiNodeTypeEnum get type;
  // enum typeEnum {  text,  input,  img,  a,  script,  };

  UiNode._();

  factory UiNode([void updates(UiNodeBuilder b)]) = _$UiNode;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(UiNodeBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<UiNode> get serializer => _$UiNodeSerializer();
}

class _$UiNodeSerializer implements PrimitiveSerializer<UiNode> {
  @override
  final Iterable<Type> types = const [UiNode, _$UiNode];

  @override
  final String wireName = r'UiNode';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    UiNode object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'attributes';
    yield serializers.serialize(
      object.attributes,
      specifiedType: const FullType(UiNodeAttributes),
    );
    yield r'group';
    yield serializers.serialize(
      object.group,
      specifiedType: const FullType(UiNodeGroupEnum),
    );
    yield r'messages';
    yield serializers.serialize(
      object.messages,
      specifiedType: const FullType(BuiltList, [FullType(UiText)]),
    );
    yield r'meta';
    yield serializers.serialize(
      object.meta,
      specifiedType: const FullType(UiNodeMeta),
    );
    yield r'type';
    yield serializers.serialize(
      object.type,
      specifiedType: const FullType(UiNodeTypeEnum),
    );
  }

  @override
  Object serialize(
    Serializers serializers,
    UiNode object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required UiNodeBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'attributes':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(UiNodeAttributes),
          ) as UiNodeAttributes;
          result.attributes.replace(valueDes);
          break;
        case r'group':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(UiNodeGroupEnum),
          ) as UiNodeGroupEnum;
          result.group = valueDes;
          break;
        case r'messages':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(UiText)]),
          ) as BuiltList<UiText>;
          result.messages.replace(valueDes);
          break;
        case r'meta':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(UiNodeMeta),
          ) as UiNodeMeta;
          result.meta.replace(valueDes);
          break;
        case r'type':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(UiNodeTypeEnum),
          ) as UiNodeTypeEnum;
          result.type = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  UiNode deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = UiNodeBuilder();
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

class UiNodeGroupEnum extends EnumClass {

  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'default')
  static const UiNodeGroupEnum default_ = _$uiNodeGroupEnum_default_;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'password')
  static const UiNodeGroupEnum password = _$uiNodeGroupEnum_password;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'oidc')
  static const UiNodeGroupEnum oidc = _$uiNodeGroupEnum_oidc;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'profile')
  static const UiNodeGroupEnum profile = _$uiNodeGroupEnum_profile;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'link')
  static const UiNodeGroupEnum link = _$uiNodeGroupEnum_link;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'totp')
  static const UiNodeGroupEnum totp = _$uiNodeGroupEnum_totp;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'lookup_secret')
  static const UiNodeGroupEnum lookupSecret = _$uiNodeGroupEnum_lookupSecret;
  /// Group specifies which group (e.g. password authenticator) this node belongs to.
  @BuiltValueEnumConst(wireName: r'webauthn')
  static const UiNodeGroupEnum webauthn = _$uiNodeGroupEnum_webauthn;

  static Serializer<UiNodeGroupEnum> get serializer => _$uiNodeGroupEnumSerializer;

  const UiNodeGroupEnum._(String name): super(name);

  static BuiltSet<UiNodeGroupEnum> get values => _$uiNodeGroupEnumValues;
  static UiNodeGroupEnum valueOf(String name) => _$uiNodeGroupEnumValueOf(name);
}

class UiNodeTypeEnum extends EnumClass {

  /// The node's type
  @BuiltValueEnumConst(wireName: r'text')
  static const UiNodeTypeEnum text = _$uiNodeTypeEnum_text;
  /// The node's type
  @BuiltValueEnumConst(wireName: r'input')
  static const UiNodeTypeEnum input = _$uiNodeTypeEnum_input;
  /// The node's type
  @BuiltValueEnumConst(wireName: r'img')
  static const UiNodeTypeEnum img = _$uiNodeTypeEnum_img;
  /// The node's type
  @BuiltValueEnumConst(wireName: r'a')
  static const UiNodeTypeEnum a = _$uiNodeTypeEnum_a;
  /// The node's type
  @BuiltValueEnumConst(wireName: r'script')
  static const UiNodeTypeEnum script = _$uiNodeTypeEnum_script;

  static Serializer<UiNodeTypeEnum> get serializer => _$uiNodeTypeEnumSerializer;

  const UiNodeTypeEnum._(String name): super(name);

  static BuiltSet<UiNodeTypeEnum> get values => _$uiNodeTypeEnumValues;
  static UiNodeTypeEnum valueOf(String name) => _$uiNodeTypeEnumValueOf(name);
}

