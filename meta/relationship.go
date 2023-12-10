package meta

// A description of the Relationship table join type.
type JoinType int32

const (
	Inner JoinType = 0 // Inner join two tables.
	Left  JoinType = 1 // Left outer join two tables.
	Right JoinType = 2 // Right outer join two tables.
	Full  JoinType = 3 // Full outer join two tables.
	Cross JoinType = 4 // Cross join two tables.
)

// A table relationship.
type Relationship struct {
	Id             string   `json:"id"`             // The unique id of the relationship.
	JoinType       JoinType `json:"joinType"`       // The type of join, possible values are: 0: Inner, 1: Left, 2: Right, 3: Full, 4: Cross.
	ConditionLeft  string   `json:"conditionLeft"`  // Field id to the left of the association condition.
	ConditionRight string   `json:"conditionRight"` // Field id to the right of the association condition.

	MetaId     string      `json:"-"`
	LeftTable  *Table      `json:"-"`
	RightTable *Table      `json:"-"`
	LeftField  *TableField `json:"-"`
	RightField *TableField `json:"-"`
}
