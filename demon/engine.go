package demon

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Demon struct {
	router *httprouter.Router
	ErrorHandler
}

func New() *Demon {
	return &Demon{
		router: httprouter.New(),
	}
}

func (d *Demon) Get(path string, h Handler, plugs ...Handler) {
	d.router.GET(path, d.makeHTTPRouterHandler(h))
}

func (d *Demon) Start(port string) error {
	return http.ListenAndServe(port, d.router)
}
