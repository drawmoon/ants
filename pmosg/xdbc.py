import base64
import json
from typing import Any, Dict

import jsonschema
from cryptography.fernet import Fernet

from pmosg.meta import SQLDialect

XDBC_PREFIX = "xdbc://"

DEFAULT_TEMPLATE = {
    "type": "object",
    "default": {},
    "required": ["host", "port", "auth"],
    "properties": {
        "host": {"type": "string"},
        "port": {"type": "integer"},
        "database": {"type": "string"},
        "auth": {
            "type": "object",
            "required": ["type"],
            "properties": {
                "type": {"type": "string", "enum": ["none", "default"]},
                "user": {"type": "string"},
                "password": {"type": "string"},
            },
            "examples": [{"type": "none"}],
        },
    },
    "examples": [
        {
            "host": "127.0.0.1",
            "port": 5432,
            "database": "postgres",
            "auth": {"type": "default", "user": "postgres", "password": "postgres"},
        }
    ],
}


class XDBC:

    dialect: SQLDialect
    parameters: Dict[str, Any]

    _url: str

    def __init__(self, **parameters):
        if "url" in parameters:
            self._url = parameters.pop("url")
            arr = str.split(self._url, "://")
            try:
                k = arr[1][:44]
                b = Fernet(k).decrypt(base64.urlsafe_b64decode(arr[1][44:]))
                self.parameters = json.loads(b)

                dialect = arr[0].upper()
                if dialect not in SQLDialect.__members__:
                    raise ValueError(f"Unsupported dialect: {dialect}")
                self.dialect = SQLDialect[dialect]
            except Exception:
                raise ValueError(f"Invalid URL: {self._url}")
        else:
            dialect = parameters.pop("dialect") if "dialect" in parameters else SQLDialect.DEFAULT

            jsonschema.validate(instance=parameters, schema=DEFAULT_TEMPLATE)
            s = json.dumps(parameters)
            k = Fernet.generate_key()
            p = base64.urlsafe_b64encode(Fernet(k).encrypt(s.encode())).decode()
            self._url = f"{dialect.name.lower()}://{k.decode()}{p}"
            self.dialect = dialect
            self.parameters = parameters

    def __str__(self) -> str:
        return self._url
