package common

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func DecodeFromPositionalJSON(data []byte, msg proto.Message) error {
	m := msg.ProtoReflect()
	fields := m.Descriptor().Fields()

	var raw []interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw) > fields.Len() {
		return fmt.Errorf("too many elements")
	}

	for i := 0; i < len(raw); i++ {
		field := fields.Get(i)
		val := raw[i]

		if val == nil {
			continue
		}

		switch field.Kind() {
		case protoreflect.Int32Kind:
			m.Set(field, protoreflect.ValueOf(int32(val.(float64))))
		case protoreflect.StringKind:
			m.Set(field, protoreflect.ValueOf(val.(string)))
		case protoreflect.BoolKind:
			m.Set(field, protoreflect.ValueOf(val.(bool)))
		case protoreflect.MessageKind:
			subMsg := m.NewField(field).Message()
			rawBytes, _ := json.Marshal(val)
			err := DecodeFromPositionalJSON(rawBytes, subMsg.Interface())
			if err != nil {
				return err
			}
			m.Set(field, protoreflect.ValueOfMessage(subMsg))
		case protoreflect.EnumKind:
			m.Set(field, protoreflect.ValueOf(protoreflect.EnumNumber(int32(val.(float64)))))
		default:
			// Handle more types as needed
			return fmt.Errorf("unsupported kind: %s", field.Kind())
		}
	}
	return nil
}
