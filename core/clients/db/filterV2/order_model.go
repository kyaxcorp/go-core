package filter

type Order struct {
	FieldName string  `json:"field_name"`
	Direction *string `json:"direction"`
}
