package stmt

// Keywords by SQL.
type Keyword string

const (
	keySelect       Keyword = "select"
	keyDistinct     Keyword = "distinct"
	keyOffset       Keyword = "offset"
	keyRows         Keyword = "rows"
	keyFetchNext    Keyword = "fetch next"
	keyPercent      Keyword = "percent"
	keyRowsWithTies Keyword = "rows with ties"
	keyRowsOnly     Keyword = "rows only"
)

func (k Keyword) Accept(ctx *Context) {
	s := string(k)
	ctx.Sql(s)
}
