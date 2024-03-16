package demon

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (d *Demon) makeHTTPRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &Context{
			w:   w,
			r:   r,
			ctx: context.Background(),
		}
		if err := h(ctx); err != nil {
			// handle the error from the handler
			d.ErrorHandler(err, ctx)
		}
	}
}
