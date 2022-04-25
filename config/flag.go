package config

import "gopkg.in/alecthomas/kingpin.v2"

func FlagParser() {

	secureAddr := kingpin.Flag("sa", "Secure Address (HTTPS)").Default(HTTPSAddrDefault).String()
	Addr := kingpin.Flag("a", "Address (HTTP)").Default(HTTPAddrDefault).String()
	configPath := kingpin.Flag("c", "Config file path").Default(FileDefaultPath).String()
	ca := kingpin.Flag("ca", "CA file path").Default(CADefaultPath).String()
	key := kingpin.Flag("key", "Private Key file path").Default(KeyDefaultPath).String()
	cert := kingpin.Flag("cert", "cert file path").Default(CertDefaultPath).String()

	pModel := kingpin.Flag("pmodel", "PolicyModel file path").Default(AuthModelDefaultPath).String()
	pFile := kingpin.Flag("pfile", "PolicyFile file path").Default(AuthPolicyDefaultPath).String()
	bypass := kingpin.Flag("authBypass", "No authorization").Default("false").Bool()
	kingpin.Parse()

	Main().
		ConfigPathSet(*configPath).
		WebServer().AddTLSAddr(*secureAddr).
		AddAddr(*Addr).AddCertKey(*key, *cert).
		AddCA(*ca)

	Main().
		AAA().AddFilePolicy(*pFile).
		AddModelPolicy(*pModel).Bypass(*bypass)

	Main().Complete <- true
}
