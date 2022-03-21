package zsw

import (
	"fmt"
	"log"
	"strings"
)

var ZSWAttributeVariant = NewVariantDefinition([]VariantType{
	{Name: "int8", Type: int8(0)},
	{Name: "int16", Type: int16(0)},
	{Name: "int32", Type: int32(0)},
	{Name: "int64", Type: int64(0)},
	{Name: "uint8", Type: uint8(0)},
	{Name: "uint16", Type: uint16(0)},
	{Name: "uint32", Type: uint32(0)},
	{Name: "uint64", Type: uint64(0)},
	{Name: "float", Type: float32(0)},
	{Name: "double", Type: float64(0)},
	{Name: "string", Type: ""},
	{Name: "INT8_VEC", Type: []int8{}},
	{Name: "INT16_VEC", Type: []int16{}},
	{Name: "INT32_VEC", Type: []int32{}},
	{Name: "INT64_VEC", Type: []int64{}},
	{Name: "UINT8_VEC", Type: []uint8{}},
	{Name: "UINT16_VEC", Type: []uint16{}},
	{Name: "UINT32_VEC", Type: []uint32{}},
	{Name: "UINT64_VEC", Type: []uint64{}},
	{Name: "FLOAT_VEC", Type: []float32{}},
	{Name: "DOUBLE_VEC", Type: []float64{}},
	{Name: "STRING_VEC", Type: []string{}},
})

type InvalidTypeError struct {
	Label        string
	ExpectedType string
	Attribute    *ZSWAttribute
}

func (c *InvalidTypeError) Error() string {
	return fmt.Sprintf("received an unexpected type %T for metadata variant %T", c.ExpectedType, c.Attribute)
}

type AttributeMap map[string]*ZSWAttribute

type ZSWAttribute struct {
	*BaseVariant
}

func NewZSWAttribute(typeId string, value interface{}) *ZSWAttribute {
	return &ZSWAttribute{
		&BaseVariant{
			TypeID: ZSWAttributeVariant.TypeID(typeId),
			Impl:   value,
		},
	}
}

func ToZSWAttribute(value interface{}) *ZSWAttribute {
	switch v := value.(type) {
	case float32:
		return NewZSWAttribute("float", v)
	case float64:
		return NewZSWAttribute("double", v)
	case []int8, []int16, []int32, []int64, []uint8, []uint16, []uint32, []uint64, []string:
		typeId := fmt.Sprintf("%v_VEC", strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%T", v), "[]", "")))
		return NewZSWAttribute(typeId, v)
	case []float32:
		return NewZSWAttribute("FLOAT_VEC", v)
	case []float64:
		return NewZSWAttribute("DOUBLE_VEC", v)
	default:
		return NewZSWAttribute(fmt.Sprintf("%T", v), v)
	}
}

func (m *ZSWAttribute) InvalidTypeError(expectedType string) *InvalidTypeError {
	return &InvalidTypeError{
		Label:        fmt.Sprintf("received an unexpected type %T for variant %T", m.Impl, m),
		ExpectedType: "int8",
		Attribute:    m,
	}
}

func (m *ZSWAttribute) String() string {
	return fmt.Sprint(m.Impl)
}

func (m *ZSWAttribute) Int8() (int8, error) {
	switch v := m.Impl.(type) {
	case int8:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int8")
	}
}

func (m *ZSWAttribute) Int16() (int16, error) {
	switch v := m.Impl.(type) {
	case int16:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int16")
	}
}

func (m *ZSWAttribute) Int32() (int32, error) {
	switch v := m.Impl.(type) {
	case int32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int32")
	}
}

func (m *ZSWAttribute) Int64() (int64, error) {
	switch v := m.Impl.(type) {
	case int64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("int64")
	}
}

func (m *ZSWAttribute) UInt8() (uint8, error) {
	switch v := m.Impl.(type) {
	case uint8:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint8")
	}
}

func (m *ZSWAttribute) UInt16() (uint16, error) {
	switch v := m.Impl.(type) {
	case uint16:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint16")
	}
}

func (m *ZSWAttribute) UInt32() (uint32, error) {
	switch v := m.Impl.(type) {
	case uint32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint32")
	}
}

func (m *ZSWAttribute) UInt64() (uint64, error) {
	switch v := m.Impl.(type) {
	case uint64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("uint64")
	}
}

func (m *ZSWAttribute) Float32() (float32, error) {
	switch v := m.Impl.(type) {
	case float32:
		return v, nil
	default:
		return 0, m.InvalidTypeError("float32")
	}
}

func (m *ZSWAttribute) Float64() (float64, error) {
	switch v := m.Impl.(type) {
	case float64:
		return v, nil
	default:
		return 0, m.InvalidTypeError("float64")
	}
}

func (m *ZSWAttribute) Int8Slice() ([]int8, error) {
	switch v := m.Impl.(type) {
	case []int8:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int8")
	}
}

func (m *ZSWAttribute) Int16Slice() ([]int16, error) {
	switch v := m.Impl.(type) {
	case []int16:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int16")
	}
}

func (m *ZSWAttribute) Int32Slice() ([]int32, error) {
	switch v := m.Impl.(type) {
	case []int32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int32")
	}
}

func (m *ZSWAttribute) Int64Slice() ([]int64, error) {
	switch v := m.Impl.(type) {
	case []int64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]int64")
	}
}

func (m *ZSWAttribute) UInt8Slice() ([]uint8, error) {
	switch v := m.Impl.(type) {
	case []uint8:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint8")
	}
}

func (m *ZSWAttribute) UInt16Slice() ([]uint16, error) {
	switch v := m.Impl.(type) {
	case []uint16:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint16")
	}
}

func (m *ZSWAttribute) UInt32Slice() ([]uint32, error) {
	switch v := m.Impl.(type) {
	case []uint32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint32")
	}
}

func (m *ZSWAttribute) UInt64Slice() ([]uint64, error) {
	switch v := m.Impl.(type) {
	case []uint64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]uint64")
	}
}

func (m *ZSWAttribute) Float32Slice() ([]float32, error) {
	switch v := m.Impl.(type) {
	case []float32:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]float32")
	}
}

func (m *ZSWAttribute) Float64Slice() ([]float64, error) {
	switch v := m.Impl.(type) {
	case []float64:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]float64")
	}
}

func (m *ZSWAttribute) StringSlice() ([]string, error) {
	switch v := m.Impl.(type) {
	case []string:
		return v, nil
	default:
		return nil, m.InvalidTypeError("[]string")
	}
}

// IsEqual evaluates if the two ZSWAttributes have the same types and values (deep compare)
func (m *ZSWAttribute) IsEqual(m2 *ZSWAttribute) bool {

	if m.TypeID != m2.TypeID {
		log.Println("ZSWAttribute types inequal: ", m.TypeID, " vs ", m2.TypeID)
		return false
	}

	if m.String() != m2.String() {
		log.Println("ZSWAttribute Values.String() inequal: ", m.String(), " vs ", m2.String())
		return false
	}

	return true
}

// MarshalJSON translates to []byte
func (m *ZSWAttribute) MarshalJSON() ([]byte, error) {
	return m.BaseVariant.MarshalJSON(ZSWAttributeVariant)
}

// UnmarshalJSON translates ZSWAttributeVariant
func (m *ZSWAttribute) UnmarshalJSON(data []byte) error {
	return m.BaseVariant.UnmarshalJSON(data, ZSWAttributeVariant)
}

// UnmarshalBinary ...
func (m *ZSWAttribute) UnmarshalBinary(decoder *Decoder) error {
	return m.BaseVariant.UnmarshalBinaryVariant(decoder, ZSWAttributeVariant)
}
