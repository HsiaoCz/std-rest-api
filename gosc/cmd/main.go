package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Context struct {
	w   http.ResponseWriter
	req *http.Request
	t   *Template
}

func (c *Context) WriteJSON(code int, value any) error {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	return json.NewEncoder(c.w).Encode(value)
}
func (c *Context) render(name string, data any) error {
	return c.t.template.ExecuteTemplate(c.w, name, data)
}

type apifunc func(c *Context) error

func makeHTTPHandler(fn apifunc, t *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			t:   t,
			req: r,
			w:   w,
		}
		if err := fn(ctx); err != nil {
			ctx.WriteJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

var t *Template

type User struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	Age      int    `json:"age"`
}

type Template struct {
	template *template.Template
}

func main() {
	t = &Template{
		template: template.Must(template.ParseGlob("www/*.html")),
	}
	http.HandleFunc("/user", makeHTTPHandler(handleUser, t))
	http.HandleFunc("/", makeHTTPHandler(handleHome, t))
	http.HandleFunc("/cat", makeHTTPHandler(showCatFacts, t))
	http.ListenAndServe(":3001", nil)
}

func handleUser(c *Context) error {
	return c.WriteJSON(http.StatusOK, map[string]any{
		"message": "hello some user",
	})
}

func handleHome(c *Context) error {
	user := User{
		Username: "hason",
		IsAdmin:  true,
		Age:      31,
	}
	return c.render("index.html", user)
}

type CatFact struct {
	Fact string `json:"fact"`
}

func showCatFacts(c *Context) error {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fact := &CatFact{}

	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return err
	}
	return c.render("index.html", fact)
}
