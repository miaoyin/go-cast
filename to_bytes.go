package cast

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func UToB(v any) []byte {
	v2, _ := UToBE(v)
	return v2
}

func UToBE(v any) ([]byte, error) {
	switch converted := v.(type) {
	case uint8:
		return []byte{converted}, nil
	case uint16:
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, converted)
		return b, nil
	case uint32:
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, converted)
		return b, nil
	case uint64:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, converted)
		return b, nil
	default:
		return nil, fmt.Errorf("unknown type: %T", converted)
	}
}

func IToB(v any) []byte {
	v2, _ := IToBE(v)
	return v2
}

func IToBE(v any) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	switch converted := v.(type) {
	case int:
		if err := binary.Write(buff, binary.BigEndian, int32(converted)); err != nil {
			return nil, err
		}
		return buff.Bytes(), nil
	case int8, int16, int32, int64:
		if err := binary.Write(buff, binary.BigEndian, converted); err != nil {
			return nil, err
		}
		return buff.Bytes(), nil
	default:
		return nil, fmt.Errorf("unknown type: %T", converted)
	}
}

func NumberToB(v any) []byte {
	v2, _ := NumberToBE(v)
	return v2
}

func NumberToBE(v any) ([]byte, error) {
	switch converted := v.(type) {
	case int, int8, int16, int32, int64:
		return IToBE(converted)
	case uint, uint8, uint16, uint32, uint64:
		return UToBE(converted)
	default:
		return nil, fmt.Errorf("unknown type: %T", converted)
	}
}
