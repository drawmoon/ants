package stmt

type Count struct {
	Expr Expression
}

func (c Count) Accept(ctx *Context) {
}
