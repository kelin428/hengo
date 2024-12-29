package hengo

import "net/http"

type HenClient struct {
	router *Router
}

func NewHen() *HenClient {
	return &HenClient{}
}

func (h *HenClient) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h *HenClient) Run() {
	http.ListenAndServe(":8000", h)
}
