package stmt

import (
	"strconv"

	"github.com/drawmoon/ants/meta"
)

// The common base type for all objects that operand property.
type Expression interface {
	// Accept is the method that the visitor will call to visit the expression.
	Accept(ctx *Context)
}

// A QueryPart to be used exclusively in SELECT clauses.
type Asterisk struct{}

func (a Asterisk) Accept(ctx *Context) {
	ctx.Sql("*")
}

// A QueryPart to be used exclusively in SELECT clauses.
type SelectField struct {
	Field *meta.TableField
}

func (s SelectField) Accept(ctx *Context) {
	ctx.Sql("")
}

type Limit struct {
	Limit   int
	Offset  int
	Percent bool
	Ties    bool
}

func (l Limit) Accept(ctx *Context) {
	if l.Offset != 0 {
		ctx.Visit(keyOffset)
		ctx.Sql(" ")
		ctx.Sql(strconv.Itoa(l.Offset))
		ctx.Sql(" ")
		ctx.Visit(keyRows)
	}

	if l.Limit != 0 {
		ctx.Sql(" ")
		ctx.Visit(keyFetchNext)
		ctx.Sql(" ")
		ctx.Sql(strconv.Itoa(l.Limit))

		if l.Percent {
			ctx.Sql(" ")
			ctx.Visit(keyPercent)
		}
		if l.Ties {
			ctx.Sql(" ")
			ctx.Visit(keyRowsWithTies)
		} else {
			ctx.Sql(" ")
			ctx.Visit(keyRowsOnly)
		}
	}
}
