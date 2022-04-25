package aaa

import "github.com/casbin/casbin/v2"

const (
	policyConf = `
	[request_definition]
	r = sub, obj, act
	[policy_definition]
	p = sub, obj, act
	[policy_effect]
	e = some(where (p.eft == allow))
	[matchers]
	m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

var enforcer, _ = casbin.NewEnforcer(policyConf, "path/to/policy.csv")

func isAuthorize(user, pass string) bool {

	return true
}
