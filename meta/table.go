package meta

// A description of the type of a Table.
type TableType int32

const (
	// An ordinary table that is stored in the schema.
	PhysicalTable TableType = 0
	// A virtual table, such as a derived table, a joined table, a common table expression, etc.
	VirtualTable TableType = 1
)

// A table info in meta info.
type Table struct {
	Id      string `json:"id"`      // The unique id of the table.
	Name    string `json:"name"`    // The physical table name. For virtual table, it is not useful.
	Schema  string `json:"schema"`  // The schema name.
	Exprstr string `json:"exprstr"` // The expression string of the table. Indicates that it is a virtual table.

	MetaId             string          `json:"-"`
	TableFields        []*TableField   `json:"-"`
	Relationships      []*Relationship `json:"-"`
	LeftRelationships  []*Relationship `json:"-"`
	RightRelationships []*Relationship `json:"-"`
}

// The type of the table.
//
// Possible values are:
//
//	0: physical table
//	1: virtual table
func (t *Table) GetTableType() TableType {
	if t.Exprstr == "" {
		return PhysicalTable
	}

	return VirtualTable
}
