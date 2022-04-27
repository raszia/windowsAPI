package authorization

import (
	"net/http"

	"windows/aaa/audit"

	"github.com/casbin/casbin/v2"
)

var (
	Enforcer *casbin.Enforcer
)

func IsAuthorize(user, pass string, r *http.Request) bool {

	res, err := Enforcer.Enforce(user, pass, r.URL.Path, r.Method)
	auditLog := audit.Logger().Info().Str("user", user).Str("uri", r.URL.Path).Str("method", r.Method).Bool("isAuthorize", res)
	defer auditLog.Send()
	if err != nil {
		auditLog.Err(err)
		return false
	}
	return res
}
