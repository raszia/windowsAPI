package aaa

import (
	"net/http"
	"windows/aaa/authentication"
	"windows/aaa/authorization"
	"windows/config"
	"windows/utility"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.MainConfig.AAA().BypassBool == "false" {
			user, pass, ok := authentication.Check(r)
			res := utility.ResStruct{
				Msg:    "unAuthenticated",
				Status: "failed",
			}
			if !ok {
				utility.HttpConnectionClose(w, r, http.StatusUnauthorized, res)
				return
			}
			if !authorization.IsAuthorize(user, pass, r) {
				res.Msg = "unAuthorized"
				utility.HttpConnectionClose(w, r, http.StatusUnauthorized, res)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
