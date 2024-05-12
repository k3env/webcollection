package handlers

import "net/http"

type FuncWrapperHandler struct {
	f http.HandlerFunc
}

func (h *FuncWrapperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(w, r)
}

func NewFuncWrapperHandler(f http.HandlerFunc) *FuncWrapperHandler {
	return &FuncWrapperHandler{f}
}
