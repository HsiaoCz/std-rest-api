package views

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func compontentToString(component templ.Component) (string, error) {
	r, w := io.Pipe()

	go func() {
		component.Render(context.Background(), w)
		w.Close()
	}()
	data, err := io.ReadAll(r)
	return string(data), err
}
