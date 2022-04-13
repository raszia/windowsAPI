package utility

import "regexp"

var (
	validName      = regexp.MustCompile(`[^A-Za-z0-9\.-]+`)
	validIPAddress = regexp.MustCompile(`^(([1-9]?\d|1\d\d|25[0-5]|2[0-4]\d)\.){3}([1-9]?\d|1\d\d|25[0-5]|2[0-4]\d)$`)
)

func ZoneNameIsValid(name string) bool {
	return nameIsValid(name)
}

func NodeNameIsValid(name string) bool {
	return nameIsValid(name)
}
func IpIsValid(ip string) bool {
	return ipIsValid(ip)
}

func ipIsValid(ipaddress string) bool {
	return validIPAddress.MatchString(ipaddress)
}
func nameIsValid(name string) bool {
	return validName.MatchString(name)
}
