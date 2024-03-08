import unittest

from pmosg.meta import SQLDialect
from pmosg.xdbc import XDBC


class XDBCTest(unittest.TestCase):
    def test_generate(self):
        x = XDBC(dialect=SQLDialect.POSTGRES,
                 database="pg",
                 host="127.0.0.1",
                 port=5432,
                 auth={"type": "default", "user": "postgres", "password": "postgres"})

        s = str(x)
        assert s is not None
        assert len(s) != 0
        assert s.startswith("postgres://")

    def test_parse(self):
        url = "postgres://By4U5eSDZRSYmjhqzYM17mvQ9012YrTVdRSSYQGgbRY=Z0FBQUFBQmwweWI5MjlKOWh2am9iaXNJTTRaT2laU1BTYnlud0MwcW9fSlRNWGd6dzYzbldSdFVLWmVZbmhteDFQRWJkN3pLWXJ6SDNFajZZUTNrb0FpaHc4Z0NWamFhS042Sl9tcUx1di1GSWtIVXVJM05PRFFWZlU4RXdmZ3VfdVBMN1ZGcVROck1uWWdmX01oT0VoRXJHVTJVY2JwTDBPYjJQYmhSOEwyUmJmNkQyU1EtSGlmQnRYN2Q1a3E2RzRJblJRU1VJdkdoOGFtNWwxRENiakxhRXdhQ2ltb1BrbmhVQ0cyNnd1ZmNoY1p6cUl0ajlYaz0="
        x = XDBC(url=url)

        s = str(x)
        assert s == url
        assert x.dialect == SQLDialect.POSTGRES
        assert x.parameters is not None

        p = x.parameters
        assert p["database"] == "pg"
        assert p["host"] == "127.0.0.1"
        assert p["port"] == 5432
        assert p["auth"]["type"] == "default"
        assert p["auth"]["user"] == "postgres"
        assert p["auth"]["password"] == "postgres"


if __name__ == "__main__":
    unittest.main()
