package stmt

import "bytes"

type Context struct {
	Errors []error

	buf *bytes.Buffer
}

func NewContext() *Context {
	return &Context{
		buf: &bytes.Buffer{},
	}
}

func (c *Context) Visit(e Expression) {
	e.Accept(c)
}

func (c *Context) Sql(s string) {
	_, err := c.buf.WriteString(s)
	if err != nil {
		c.Err(err)
		return
	}
}

func (c *Context) Err(e error) {
	c.Errors = append(c.Errors, e)
}

func (c *Context) String() (string, error) {
	if len(c.Errors) > 0 {
		return "", c.Errors[0]
	}

	return c.buf.String(), nil
}
