package aaa

import (
	"windows/aaa/authentication"
	"windows/aaa/authorization"
	"windows/config"

	"github.com/casbin/casbin/v2"
)

func SetEnforcer() error {
	Enforcer, err := casbin.NewEnforcer(config.MainConfig.AAA().PolicyModelPath, config.MainConfig.AAA().PolicyFilePath)
	if err != nil {
		return err
	}
	authorization.Enforcer = Enforcer
	pList := Enforcer.GetPolicy()
	for _, policy := range pList {
		if len(policy) < 2 {
			continue
		}
		authentication.UserManager.Insert(
			authentication.CreateUser(policy[0], policy[1]),
		)
	}

	return nil
}
