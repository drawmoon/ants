package stmt

import "github.com/drawmoon/ants/meta"

// The common base type for all objects that can be used for query composition.
type QueryPart interface {
	// Set the alias of this QueryPart.
	As(alias string)

	// Render a SQL string representation of this QueryPart.
	String() (string, error)
}

// A joinable object.
type Joinable interface {
	// Specify the join between Joinable and Joinable.
	Join(to Joinable, t meta.JoinType)

	// Set the join condition between this Joinable and the specified Joinable.
	On(expr Condition)

	// Set the alias of this QueryPart.
	As(alias string)
}

// A condition or predicate.
// Conditions can be used in a variety of SQL clauses. They're mainly used in a
// Select statement's WHERE clause.
type Condition interface {
	Expression

	// Performs a logical AND operation between this Condition and the specified expression.
	//
	// The expression to be logically ANDed with this Condition.
	//
	// return a new Condition representing the logical AND operation.
	And(expr Expression) Condition

	// Performs a logical OR operation between this Condition and the specified expression.
	//
	// The expression to be logically ORed with this Condition.
	//
	// return a new Condition representing the logical OR operation.
	Or(expr Expression) Condition
}

// A sub-select is a SELECT statement that can be combined with
// other SELECT statement in UNION and similar operations.
type Statement struct {
	Select          []QueryPart    // The select statement elms.
	Distinct        bool           // Gets or sets whether the select statement is distinct.
	From            []Joinable     //
	Predicate       Condition      //
	GroupBy         []*interface{} //
	GroupByDistinct bool           //
	Having          []*interface{} //
	OrderBy         []*interface{} //
	Limit           *Limit         //
	Alias           string         // The name of the select part to be displayed.
}

func (s Statement) As(alias string) {
	x := &s
	x.Alias = alias
}

func (s Statement) String() (string, error) {
	c := NewContext()

	c.Visit(keySelect)
	c.Sql(" ")

	// generate select clause.
	if s.Select == nil || len(s.Select) == 0 {
		c.Visit(&Asterisk{})
	} else {
		if s.Distinct {
			c.Visit(keyDistinct)
			c.Sql(" ")
		}
		for _, s := range s.Select {
			str, err := s.String()
			if err != nil {
				return "", err
			}
			c.Sql(str)
		}
	}

	return c.String()
}

// A QueryPart to be used exclusively in SELECT clauses.
// This method is useful for things like
//
//	t.abc abc
//	count(t.abc) abc_count
//	(select a from x) a
type SelectPart struct {
	Expr  Expression // The select element expression.
	Alias string     // The name of the select part to be displayed.
}

func (p SelectPart) As(alias string) {
	x := &p
	x.Alias = alias
}

func (p SelectPart) String() (string, error) {
	return "", nil
}

// Create a new select statement.
//
// Example:
//
//	Select(field1, field2, ...)
//		.From(table1)
//		.Join(table2, field1.Eq(field2))
//		.Where(field1.Eq(1))
//		.Limit(1, 10)
//		.String()
func Select(p ...QueryPart) QueryPart {
	return &Statement{Select: p}
}
