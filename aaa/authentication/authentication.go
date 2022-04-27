package authentication

import (
	"net/http"
	"strings"
	"sync"
	"windows/aaa/audit"
)

var (
	UserManager = &userManagerStruct{
		users: make(map[string]*userStruct),
	}
)

type userStruct struct {
	mu       sync.Mutex
	userName string
	password string
}

type userManagerStruct struct {
	mu    sync.Mutex
	users map[string]*userStruct
}

func CreateUser(userName, password string) *userStruct {
	return &userStruct{
		userName: userName,
		password: password,
	}
}
func (userMan *userManagerStruct) Insert(user *userStruct) {
	userMan.mu.Lock()
	defer userMan.mu.Unlock()
	userMan.users[user.userName] = user
}

func (userMan *userManagerStruct) DeleteUserByUserName(userName string) {
	userMan.mu.Lock()
	defer userMan.mu.Unlock()
	delete(userMan.users, userName)
}

func (userMan *userManagerStruct) GetUserByUserName(userName string) (*userStruct, bool) {
	userMan.mu.Lock()
	defer userMan.mu.Unlock()
	user, ok := userMan.users[userName]
	return user, ok
}

func (userMan *userManagerStruct) IsAuthenticateByUserPass(userName, password string) bool {
	user, ok := userMan.GetUserByUserName(userName)
	if !ok {
		return false
	}
	user.mu.Lock()
	defer user.mu.Unlock()
	return password == user.password
}

func Check(req *http.Request) (user, password string, ok bool) {

	if strings.HasPrefix(req.Header.Get("authorization"), "Basic") {
		return basic(req)
	}

	return "", "", false
}

func basic(req *http.Request) (user, password string, ok bool) {

	auditLog := audit.Logger().Info().Str("authenticationType", "basic")
	defer auditLog.Send()
	if user, password, ok = req.BasicAuth(); !ok {
		auditLog.Bool("isAuthenticate", false).Str("reason", "basic decode err")
		return "", "", false
	}
	auditLog.Str("user", user)
	if ok := UserManager.IsAuthenticateByUserPass(user, password); ok {
		auditLog.Bool("isAuthenticate", true)
		return user, password, true
	}
	auditLog.Bool("isAuthenticate", false).Str("reason", "bad user or pass")
	return "", "", false
}
