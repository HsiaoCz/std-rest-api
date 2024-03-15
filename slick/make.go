package slick

// handle the apifunc error 
// in the serveHTTP(w http.ResponseWriter,r *http.Request)
type apifunc func(c *Context)error

