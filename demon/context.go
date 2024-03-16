package demon

import (
	"context"
	"encoding/json"
	"net/http"
	"text/template"
)

type Context struct {
	r   *http.Request
	w   http.ResponseWriter
	t   *template.Template
	ctx context.Context
}

func (c *Context) JSON(v any) error {
	c.w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.w).Encode(v)
}

func (c *Context) Render(name string, v any) error {
	return c.t.ExecuteTemplate(c.w, name, v)
}

func (c *Context) String(v string) error {
	_, err := c.w.Write([]byte(v))
	return err
}

func (c *Context) BodyPrase(data any) error {
	return json.NewDecoder(c.r.Body).Decode(&data)
}

func (c *Context) Status(code int) *Context {
	c.w.WriteHeader(code)
	return c
}

func (c *Context) SetContextValue(key any, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Context) GetContextValue(key any) (value any) {
	return c.ctx.Value(key)
}
