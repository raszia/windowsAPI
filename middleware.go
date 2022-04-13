package main

import "net/http"

func middlewareDispatch(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler { //a dispatch for middlewares
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
