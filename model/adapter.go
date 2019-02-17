package model

import "net/http"

func BuildFunc(
	h func(http.ResponseWriter, *http.Request),
	stk ...func(http.ResponseWriter, *http.Request,
		func(w http.ResponseWriter, req *http.Request))) func(http.ResponseWriter, *http.Request) {

	if len(stk) == 0 {
		return h
	}
	return BuildFunc(func(w http.ResponseWriter, req *http.Request) {
		stk[0](w, req, h)
	}, stk[1:]...)
}

func Build(
	h http.Handler,
	stk ...func(http.ResponseWriter, *http.Request,
		func(w http.ResponseWriter, req *http.Request))) http.Handler {

	return http.HandlerFunc(BuildFunc(h.ServeHTTP, stk...))
}
