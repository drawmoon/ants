package meta

// A description of the type of a TableField.
type FieldType int32

const (
	// An ordinary field that is stored in the table.
	PhysicalField FieldType = 0
	// A virtual field, such as a derived field, a common field expression, etc.
	VirtualField FieldType = 1
)

// A table info in meta info.
type TableField struct {
	Id      string `json:"id"`      // The unique id of the table field.
	Name    string `json:"name"`    // The name of the table field.
	Exprstr string `json:"exprstr"` // The expression string of the table field.
	TableId string `json:"tableId"` // The id of the table.

	MetaId string `json:"-"`
	Table  *Table `json:"-"`
}

// The type of the table field.
//
// Possible values are:
//
//	0: physical field
//	1: virtual field
func (t *TableField) GetFieldType() FieldType {
	if t.Exprstr == "" {
		return PhysicalField
	}

	return VirtualField
}
