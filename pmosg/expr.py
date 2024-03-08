from __future__ import annotations

import io
from enum import Enum, auto
from typing import List


class ExpressionVisitContext:
    """The context object that the visitor and the current state.
    """

    buf: io.BytesIO
    errors: List[Exception]
    render_style: RenderStyle

    def __init__(self) -> None:
        self.buf = io.BytesIO()
        self.errors = []

    def visit(self, expr: Expression):
        expr.accept(self)

    def sql(self, s: str):
        self.buf.write(s)

    def err(self, e: Exception):
        self.errors.append(e)

    def __str__(self) -> str:
        if len(self.errors) > 0:
            return '\n'.join(map(lambda e: str(e), self.errors))
        return self.buf.getvalue().decode()


class Expression:
    """The common base type for all objects that operand property.
    """

    def accept(self, ctx: ExpressionVisitContext):
        """Accept is the method that the visitor will call to visit the expression.

        Args:
            ctx (Context): The context object that contains the visitor and the current state.
        """
        pass


class RenderKeywordCase(Enum):
    """SQL keyword render style.
    """

    AS_IS = auto()
    LOWER = auto()
    UPPER = auto()


class RenderStyle:
    """Render style for the generated SQL.
    """

    format: RenderKeywordCase
    delimiter_required: bool
    delimiter: str

    def __init__(self, style=RenderKeywordCase.AS_IS, delimiter_required=False, delimiter="") -> None:
        self.format = style
        self.delimiter_required = delimiter_required
        self.delimiter = delimiter


class Keyword(Expression):
    """A SQL keyword.
    """

    as_is: str

    def __init__(self, s: str) -> None:
        super().__init__()
        self.as_is = s

    def accept(self, ctx: ExpressionVisitContext):
        ctx.sql(self.render(ctx))

    def render(self, ctx: ExpressionVisitContext) -> str:
        style = ctx.render_style.format

        if style == RenderKeywordCase.AS_IS:
            return self.as_is
        if style == RenderKeywordCase.LOWER:
            return self.as_is.lower()
        if style == RenderKeywordCase.UPPER:
            return self.as_is.upper()
        raise Exception(f"Unknown render style: {style}")


SELECT = Keyword("SELECT")
DISTINCT = Keyword("DISTINCT")
FROM = Keyword("FROM")
WHERE = Keyword("WHERE")
GROUP_BY = Keyword("GROUP BY")
ORDER_BY = Keyword("ORDER BY")
OFFSET = Keyword("OFFSET")
ROWS = Keyword("ROWS")
FETCH_NEXT = Keyword("FETCH NEXT")
PERCENT = Keyword("PERCENT")
ROWS_WITH_TIES = Keyword("ROWS WITH TIES")
ROWS_ONLY = Keyword("ROWS ONLY")
