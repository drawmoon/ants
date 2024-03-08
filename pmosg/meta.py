from enum import Enum, auto
from typing import Dict, List

from pmosg.statement import JoinHint, JoinType


class SQLDialect(Enum):
    """Dialects and dialect families as supported by pyants.
    """

    # Default is the default dialect.
    #
	# This dialect was chosen in the absence of a more definitive dialect. It is
	# not intended to be used with any actual database, and it is ultimately interpreted as
	# a serializable string format.
    DEFAULT = auto()

    # The MySQL dialect family.
    MYSQL = auto()

    # The PostgreSQL dialect family.
    POSTGRES = auto()

    # The Oracle dialect family.
    ORACLE = auto()

    # The SqlServer dialect family.
    SQL_SERVER = auto()


class Schema:
    """A schema info in meta info.
    """

    id: str
    meta_id: str
    name: str
    title: str
    description: str


class TableType(Enum):
    """A description of the type of a Table.
    """

    # An ordinary table that is stored in the schema.
    PHYSICAL_TABLE = auto()

    # A virtual table, such as a derived table, a joined table, a common table expression, etc.
    VIRTUAL_TABLE = auto()


class Table:
    """A table info in meta info.
    """

    id: str
    meta_id: str
    schema_id: str
    name: str
    title: str
    description: str
    type: TableType
    expr_str: str


class FieldType(Enum):
    """A description of the type of a TableField.
    """

    # An ordinary field that is stored in the table.
    PHYSICAL_FIELD = auto()

    # A virtual field, such as a derived field, a common field expression, etc.
    VIRTUAL_FIELD = auto()


class Field:
    """A field info in meta info.
    """

    id: str
    meta_id: str
    table_id: str
    name: str
    title: str
    description: str
    type: FieldType
    expr_str: str


class Relationship:
    """A table relationship.
    """

    id: str
    meta_id: str
    join_type: JoinType
    join_hint: JoinHint
    left_field_id: str
    right_field_id: str


class MetaSet:
    """A meta set, represents an ensemble of associative relationships.
    """

    id: str
    url: str
    version: str
    dialect: SQLDialect

    _tables: Dict[str, Table]
    _fields: Dict[str, Field]
    _relationships: Dict[str, Relationship]

    def __init__(self,
                 id: str,
                 url: str,
                 version: str,
                 dialect: SQLDialect,
                 tables: List[Table],
                 fields: List[Field],
                 relationships: List[Relationship]) -> None:
        """Initialize a MetaSet.
        """

        self.id = id
        self.url = url
        self.version = version
        self.dialect = dialect

        self._tables = {x.id: x for x in tables}
        self._fields = {x.id: x for x in fields}
        self._relationships = {x.id: x for x in relationships}
