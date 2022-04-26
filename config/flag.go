package config

import "gopkg.in/alecthomas/kingpin.v2"

func FlagParser() {

	secureAddr := kingpin.Flag("sa", "Secure Address (HTTPS)").Default(HTTPSAddrDefault).String()
	Addr := kingpin.Flag("a", "Address (HTTP)").Default(HTTPAddrDefault).String()
	configPath := kingpin.Flag("c", "Config file path").Default(configFileDefaultPath).String()
	ca := kingpin.Flag("ca", "CA file path").Default(CADefaultPath).String()
	key := kingpin.Flag("key", "Private Key file path").Default(KeyDefaultPath).String()
	cert := kingpin.Flag("cert", "Cert file path").Default(CertDefaultPath).String()
	httpsDis := kingpin.Flag("disableHTTPS", "No HTTPS server").Default(FalseDefault).String()
	httpDis := kingpin.Flag("disableHTTP", "No HTTP server").Default(FalseDefault).String()

	pModel := kingpin.Flag("pmodel", "PolicyModel path for authorization").Default(AuthModelDefaultPath).String()
	pFile := kingpin.Flag("pfile", "PolicyFile path for authorization").Default(AuthPolicyDefaultPath).String()
	bypass := kingpin.Flag("authBypass", "No authorization").Default(FalseDefault).String()
	kingpin.Parse()

	Main().
		ConfigPathSet(*configPath).
		WebServer().
		AddTLSAddr(*secureAddr).
		AddAddr(*Addr).
		AddKey(*key).
		AddCert(*cert).
		AddCA(*ca).
		DisableHTTP(*httpDis).
		DisableHTTPS(*httpsDis)

	Main().
		AAA().
		AddFilePolicy(*pFile).
		AddModelPolicy(*pModel).
		Bypass(*bypass)

	Main().Complete <- true
}
