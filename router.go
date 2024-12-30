package hengo

type HandlerFunc func(c *Context)

type Router struct {
	routers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(path string, handler HandlerFunc) {
	r.routers[path] = handler
}

func (h *HenEngine) Get(path string, handler HandlerFunc) {
	h.router.addRoute("GET"+"-"+path, handler)
}

func (h *HenEngine) Post(path string, handler HandlerFunc) {
	h.router.addRoute("POST"+"-"+path, handler)
}
