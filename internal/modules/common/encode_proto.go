package common

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// EncodeToPositionalJSON marshals a protobuf message to a JSON positional array
func EncodeToPositionalJSON(msg proto.Message) ([]byte, error) {
	encoded, err := encodeMessageToArray(msg.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return json.Marshal(encoded)
}

// encodeMessageToArray recursively converts a protobuf message to a []interface{}
func encodeMessageToArray(m protoreflect.Message) ([]interface{}, error) {
	var result []interface{}
	desc := m.Descriptor()
	fields := desc.Fields()

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		// Handle oneof (non-synthetic only)
		if oneof := field.ContainingOneof(); oneof != nil && !oneof.IsSynthetic() {
			// Only handle first field of oneof group
			if field.Index() != oneof.Fields().Get(0).Index() {
				continue
			}
			activeField := m.WhichOneof(oneof)
			if activeField == nil {
				result = append(result, nil)
				continue
			}
			val := m.Get(activeField)
			kind := activeField.Kind()

			if kind == protoreflect.MessageKind {
				nested, err := encodeMessageToArray(val.Message())
				if err != nil {
					return nil, err
				}
				result = append(result, nested)
			} else {
				result = append(result, val.Interface())
			}
			continue
		}

		// Skip already handled oneof field
		if field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic() {
			continue
		}

		val := m.Get(field)

		if !m.Has(field) {
			result = append(result, nil)
			continue
		}

		// Map fields
		if field.IsMap() {
			mapVal := val.Map()
			var mapResult []interface{}

			mapKeyKind := field.MapKey().Kind()
			mapValueKind := field.MapValue().Kind()

			mapVal.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				var key interface{}
				switch mapKeyKind {
				case protoreflect.StringKind:
					key = k.String()
				default:
					key = k.Interface()
				}

				var value interface{}
				if mapValueKind == protoreflect.MessageKind {
					nested, err := encodeMessageToArray(v.Message())
					if err != nil {
						return false
					}
					value = nested
				} else {
					value = v.Interface()
				}

				mapResult = append(mapResult, []interface{}{key, value})
				return true
			})

			result = append(result, mapResult)
			continue
		}

		// Repeated fields
		if field.IsList() {
			list := val.List()
			var arr []interface{}
			for j := 0; j < list.Len(); j++ {
				item := list.Get(j)
				if field.Kind() == protoreflect.MessageKind {
					nested, err := encodeMessageToArray(item.Message())
					if err != nil {
						return nil, err
					}
					arr = append(arr, nested)
				} else {
					arr = append(arr, item.Interface())
				}
			}
			result = append(result, arr)
			continue
		}

		// Nested message
		if field.Kind() == protoreflect.MessageKind {
			nested, err := encodeMessageToArray(val.Message())
			if err != nil {
				return nil, err
			}
			result = append(result, nested)
			continue
		}

		// Scalar field
		result = append(result, val.Interface())
	}

	return result, nil
}
