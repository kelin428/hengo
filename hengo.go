package hengo

import "net/http"

type HenEngine struct {
	router *Router
}

func New() *HenEngine {
	return &HenEngine{
		router: NewRouter(),
	}
}

func (h *HenEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	h.handleHTTPRequest(c)
	return
}

func (h *HenEngine) Run(addr string) {
	panic(http.ListenAndServe(addr, h))
}

func (h *HenEngine) handleHTTPRequest(c *Context) {
	path := c.method + "-" + c.path
	handler, ok := h.router.routers[path]
	if !ok {
		c.w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(c)
}
