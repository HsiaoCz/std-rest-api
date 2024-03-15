package slick

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Context struct {
	r *http.Request
	w http.ResponseWriter
	t *template.Template
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
