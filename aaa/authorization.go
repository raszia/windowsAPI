package aaa

import (
	"fmt"
	"net/http"

	"windows/config"

	"github.com/casbin/casbin/v2"
)

var (
	enforcer    *casbin.Enforcer
	enforcerErr error
)

func isAuthorize(user, pass string, r *http.Request) bool {

	if enforcerErr != nil {
		//TODO: log
		return false
	}
	res, err := enforcer.Enforce(user, pass, r.URL.Path, r.Method)
	if err != nil {
		//TODO: log
		return false
	}
	return res
}

func SetEnforcer() {
	fmt.Printf("config.MainConfig.AAA().PolicyFilePath: %v\n", config.MainConfig.AAA().PolicyFilePath)
	fmt.Printf("config.MainConfig.AAA().PolicyModelPath: %v\n", config.MainConfig.AAA().PolicyModelPath)
	enforcer, enforcerErr = casbin.NewEnforcer(config.MainConfig.AAA().PolicyFilePath, config.MainConfig.AAA().PolicyModelPath)
}
