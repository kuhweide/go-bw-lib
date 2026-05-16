package bw

type FieldType int

const (
	FieldTypeText FieldType = iota
	FieldTypeHidden
	FieldTypeBoolean
)

type field struct {
	Name  string    `json:"name"`
	Value string    `json:"value"`
	Type  FieldType `json:"type"`
}

func NewField(name string, value string, fieldType FieldType) field {
	return field{Name: name, Value: value, Type: fieldType}
}
