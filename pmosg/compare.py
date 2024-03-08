from pmosg.expr import Expression, ExpressionVisitContext
from pmosg.meta import Field
from pmosg.statement import Condition


class Eq(Expression, Condition):
    """The equals statement.
    """

    _arg1: Field
    _arg2: Field

    def __init__(self, arg1: Field, arg2: Field) -> None:
        super().__init__()
        self._arg1 = arg1
        self._arg2 = arg2

    def accept(self, ctx: ExpressionVisitContext):
        pass


class Ne(Expression, Condition):
    """The not-equals statement.
    """

    _arg1: Field
    _arg2: Field

    def __init__(self, arg1: Field, arg2: Field) -> None:
        super().__init__()
        self._arg1 = arg1
        self._arg2 = arg2

    def accept(self, ctx: ExpressionVisitContext):
        pass


class Gt(Expression, Condition):
    """The greater statement.
    """
    pass


class Ge(Expression, Condition):
    """The greater-equal statement.
    """
    pass


class Lt(Expression, Condition):
    """The less statement.
    """
    pass


class Le(Expression, Condition):
    """The less-equal statement.
    """
    pass
