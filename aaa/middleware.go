package aaa

import (
	"net/http"
	"windows/config"
	"windows/utility"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !config.MainConfig.AAA().BypassBool {
			user, pass, ok := r.BasicAuth()
			res := utility.ResStruct{
				Msg:    "unauthorized",
				Status: "failed",
			}
			if !ok {
				utility.HttpConnectionClose(w, r, http.StatusUnauthorized, res)
				return
			}
			if !isAuthorize(user, pass, r) {
				utility.HttpConnectionClose(w, r, http.StatusUnauthorized, res)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}