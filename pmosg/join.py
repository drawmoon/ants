from __future__ import annotations

from pmosg.meta import Table
from pmosg.statement import Joinable


class Join(Joinable):

    _lhs: Joinable
    _rhs: Table
    _alias: str

    def __init__(self, lhs: Joinable, rhs: Table) -> None:
        super().__init__()
        self._lhs = lhs
        self._rhs = rhs
        self._alias = None

    def qua(self, s: str):
        self._alias = s
        return self

    def join(self, to: Table):
        base = self
        if self._rhs is not None:
            base = base.join()
        return Join(base, to)

    def to_str(self):
        pass
