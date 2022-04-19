package ldap

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"windows/ldap/auth"
)

const (
	serverAddr = "192.168.32.50"
	serverPort = 389
)

type ConnectionStruct struct {
	conn     *auth.Conn
	upn      string
	userName string
	password string
	limit    int
}

type userInfoResStruct struct {
	GroupList  []string `json:"groupList"`
	MobileList []int64  `json:"mobileList"`
	IPPhone    string   `json:"ipPhone"`
	ADFullName string   `json:"adFullName"`
	Status     string   `json:"status"`
	Msg        string   `json:"msg"`
}

func (ldapConn *ConnectionStruct) close() {
	ldapConn.conn.Conn.Close()
}

//baseDN "DC=myco,DC=local"
func (ldapReq *ReqStruct) bindConnection() (*ConnectionStruct, error) {
	config := &auth.Config{
		Server:   serverAddr,
		Port:     serverPort,
		BaseDN:   ldapReq.BaseDN,
		Security: auth.SecurityNone,
	}

	upn, err := config.UPN(ldapReq.UserName)
	if err != nil {
		return nil, err
	}

	conn, err := config.Connect()
	if err != nil {
		return nil, err
	}

	ldapConn := &ConnectionStruct{upn: upn,
		userName: ldapReq.UserName,
		password: ldapReq.Password}

	status, err := conn.Bind(upn, ldapReq.Password)
	if !status {
		return nil, errors.New(ldapConnIsNotTrueErr)
	}
	if ldapReq.LimitSize < 1 {
		ldapReq.LimitSize = 10
	}
	ldapConn.limit = ldapReq.LimitSize
	ldapConn.conn = conn
	return ldapConn, err

}

func (ldapConn *ConnectionStruct) getUserInfoLdap() (*userInfoResStruct, error) {

	defer ldapConn.close()
	entry, err := ldapConn.conn.GetAttributes(userPrincipalNameAttr, ldapConn.upn, []string{cnAttr})
	if err != nil {
		return nil, err
	}

	foundGroups, err := ldapConn.conn.Search(fmt.Sprintf("(member:%s:=%s)", auth.LDAPMatchingRuleInChain, entry.DN), []string{""}, ldapConn.limit)
	if err != nil {
		return nil, err
	}

	userInfo := &userInfoResStruct{}
	for _, userGroup := range foundGroups {
		userInfo.GroupList = append(userInfo.GroupList, userGroup.DN)
	}

	_, entry, _, err = auth.AuthenticateExtended(ldapConn.conn.Config, ldapConn.userName, ldapConn.password, []string{}, []string{})
	if err != nil {
		return userInfo, err
	}

	userInfo.ADFullName = entry.GetAttributeValue(cnAttr)
	userInfo.IPPhone = entry.GetAttributeValue(ipPhoneAttr)

	mobileNumberList := strings.Split(entry.GetAttributeValue(telephoneNumberAttr), ",")
	for _, mobileNumber := range mobileNumberList {
		mobileNumberint, err := strconv.ParseInt(mobileNumber, 10, 64)
		if err != nil {
			continue
		}
		userInfo.MobileList = append(userInfo.MobileList, mobileNumberint)
	}

	return userInfo, nil
}
