package demon

import "log/slog"

type ErrorHandler func(error, *Context) error

func defauleErrorHandler(err error, c *Context) {
	// c
	slog.Error("error", "err", err)
}
