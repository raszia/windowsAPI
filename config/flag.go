package config

import "gopkg.in/alecthomas/kingpin.v2"

func FlagParser() {

	secureAddr := kingpin.Flag("sa", "Secure Address (HTTPS)").Default(HTTPSAddrDefault).String()
	Addr := kingpin.Flag("a", "Address (HTTP)").Default(HTTPAddrDefault).String()
	configPath := kingpin.Flag("c", "Config file path").Default(FileDefaultPath).String()
	ca := kingpin.Flag("ca", "CA file path").Default(CADefaultPath).String()
	key := kingpin.Flag("key", "Private Key file path").Default(KeyDefaultPath).String()
	cert := kingpin.Flag("cert", "cert file path").Default(CertDefaultPath).String()
	kingpin.Parse()

	Main().ConfigPathSet(*configPath).
		WebServer().AddTLSAddr(*secureAddr).
		AddAddr(*Addr).AddCertKey(*key, *cert).
		AddCA(*ca)

}
