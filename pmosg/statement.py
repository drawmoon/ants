from __future__ import annotations

from enum import Enum, auto
from typing import List

from pmosg.expr import Expression
from pmosg.page import Limit


class QueryPart:
    """The common base type for all objects that can be used for query composition.
    """

    def qua(self, s: str):
        """Set the alias of this QueryPart.

        Args:
            s (str): The alias to set for this QueryPart.
        """
        pass

    def __str__(self) -> str:
        """Render a SQL string representation of this QueryPart.
        """
        pass


class Condition:
    """A condition or predicate.

    Conditions can be used in a variety of SQL clauses. They're mainly used in a
    Select statement's WHERE clause.
    """

    def __and__(self, c):
        """Performs a logical AND operation between this Condition and the specified expression.

        The expression to be logically ANDed with this Condition.

        Args:
            c (Condition): The expression to logically AND with this Condition.

        Returns:
            Condition: return a new Condition representing the logical AND operation.
        """
        pass

    def __or__(self, c):
        """Performs a logical OR operation between this Condition and the specified expression.

        The expression to be logically ORed with this Condition.

        Args:
            c (Condition): The expression to logically OR with this Condition.

        Returns:
            Condition: return a new Condition representing the logical OR operation.
        """
        pass


class JoinType(Enum):
    """The type of join.
    """

    # A self join is a regular join, but the table is joined with itself.
    SELF_JOIN = auto()

    # The INNER JOIN keyword selects records that have matching values in both tables.
    INNER_JOIN = auto()

    # The LEFT JOIN keyword returns all records from the left table (table1), and the
    # matching records from the right table (table2). The result is 0 records from
    # the right side, if there is no match.
    LEFT_OUTER_JOIN = auto()

    # The RIGHT JOIN keyword returns all records from the right table (table2), and the
    # matching records from the left table (table1). The result is 0 records from the
    # left side, if there is no match.
    RIGHT_OUTER_JOIN = auto()

    # The FULL OUTER JOIN keyword returns all records when there is a match in left (table1)
    # or right (table2) table records.
    FULL_OUTER_JOIN = auto()


class JoinHint(Enum):
    """A hint for join algorithms.
    """

    LOOP = auto()
    HASH = auto()
    MERGE = auto()
    REMOTE = auto()


class Joinable:
    """A joinable object.

    Example:
    ```python
    table1.join(table2)
    ```
    """

    def join(self, to: Joinable, t=JoinType.INNER_JOIN, h: JoinHint=None):
        """Join this Joinable with another Joinable.

        Args:
            to (Joinable): The Joinable to join with.
            t (JoinType): The type of join.
            h (JoinHint): A hint for join algorithms.
        """
        pass

    def qua(self, s: str):
        """Set the alias of this Joinable.

        Args:
            s (str): The alias to set for this Joinable.
        """
        pass

    def __str__(self) -> str:
        """Render a SQL string representation of this Joinable.
        """
        pass


class Statement(QueryPart):
    """A statement is a SELECT statement that can be combined with other SELECT statement in UNION and similar operations.
    """

    _select: List[QueryPart]
    # _distinct: bool
    _from: Joinable
    _predicate: Condition
    _group_by: List[None]
    _group_by_distinct: bool
    _having: List[None]
    _order_by: List[None]
    _limit: Limit
    _alias: str

    def __init__(self, source: Joinable) -> None:
        super().__init__()
        self._from = source

    def qua(self, s: str):
        self._alias = s
        return self

    def select(self, p: SelectPart):
        pass

    def where(self, p: Condition):
        pass

    def limit(self, *p):
        """To limit the number of rows returned by a statement.

        Example:
        ```python
        # Specify skip and limit parameters.
        st.limit(0, 10)

        # Or use a Limit with rich parameters.
        st.limit(Limit(10, with_ties=True))
        ```
        """
        pass

    def __str__(self) -> str:
        pass


class SelectPart(QueryPart):
    """A QueryPart to be used exclusively in SELECT clauses.

    This method is useful for things like:
    - t.abc abc
    - count(t.abc) abc_count
    - (select a from x) a
    """

    # The select element expression.
    _expr: Expression

    # The name of the select part to be displayed.
    _alias: str

    def __init__(self, expr: Expression) -> None:
        super().__init__()
        self._expr = expr

    def qua(self, s: str):
        self._alias = s
        return self

    def __str__(self) -> str:
        pass


# def select(*p: List[QueryPart]) -> Statement:
#     """Create a new statement.

#     Example:
#     ```python
#     select(field1, field2, ...)
#         .from(table1.join(table2, Eq(field1, field2)))
#         .where(field1.eq(1))
#         .limit(0, 10)
#     ```
#     """
#     pass
def query(j: Joinable) -> Statement:
    """Create a new statement.

    Example:
    ```python
    # Specify the data source.
    source = table1.join(table2)

    # Define the query expression.
    query(source)
        .where(field1.eq(1))
        .select(field1, field2, ...)
        .limit(0, 10)
    ```

    Args:
        j (Joinable): The data source.
    """
    pass
