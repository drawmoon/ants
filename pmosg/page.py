from pmosg.expr import (
    FETCH_NEXT,
    OFFSET,
    PERCENT,
    ROWS,
    ROWS_ONLY,
    ROWS_WITH_TIES,
    Expression,
    ExpressionVisitContext,
)


class Limit(Expression):
    _limit: int
    _offset: int
    _percent: bool
    _with_ties: bool

    def __init__(self, limit: int, offset=0, percent=False, with_ties=False) -> None:
        super().__init__()
        self._limit = limit
        self._offset = offset
        self._percent = percent
        self._with_ties = with_ties

    def accept(self, ctx: ExpressionVisitContext):
        if self._offset != 0:
            ctx.visit(OFFSET)
            ctx.sql(" ")
            ctx.sql(str(self._offset))
            ctx.sql(" ")
            ctx.visit(ROWS)

        if self._limit != 0:
            ctx.sql(" ")
            ctx.visit(FETCH_NEXT)
            ctx.sql(" ")
            ctx.sql(str(self._limit))

            if self._percent:
                ctx.sql(" ")
                ctx.visit(PERCENT)

            ctx.sql(" ")
            ctx.visit(ROWS_WITH_TIES if self._with_ties else ROWS_ONLY)
