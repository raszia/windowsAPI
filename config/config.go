package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var configFileDefaultPath = "./config.toml"

type MainConfigStruct struct {
	configPath string
	WebServerS *WebServerConfigStruct `toml:"webServer"`
	AAAS       *AAAConfigStruct       `toml:"aaa"`
	Complete   chan bool
}

type WebServerConfigStruct struct {
	HttpsAddr string `toml:"HTTPSAddr"`
	HttpAddr  string `toml:"HTTPAddr"`
	Ca        string `toml:"CA"`
	Key       string `toml:"Key"`
	Cert      string `toml:"Cert"`
}
type AAAConfigStruct struct {
	PolicyModelPath string `toml:"PolicyModelPath"`
	PolicyFilePath  string `toml:"PolicyFilePath"`
	BypassBool      bool   `toml:"Bypass"`
}

var MainConfig = &MainConfigStruct{
	WebServerS: &WebServerConfigStruct{},
	AAAS:       &AAAConfigStruct{},
	Complete:   make(chan bool, 100),
}

func (main *MainConfigStruct) ConfigPathSet(path string) *MainConfigStruct {
	main.configPath = path
	LoadMainConfig(path)
	return main
}

func Main() *MainConfigStruct {
	return MainConfig
}

func (main *MainConfigStruct) WebServer() *WebServerConfigStruct {
	return MainConfig.WebServerS
}

func (main *MainConfigStruct) AAA() *AAAConfigStruct {
	return MainConfig.AAAS
}

func (AAA *AAAConfigStruct) AddModelPolicy(filePath string) *AAAConfigStruct {
	if filePath != "" && filePath == AuthModelDefaultPath {
		return AAA
	}
	AAA.PolicyModelPath = filePath
	return AAA
}

func (AAA *AAAConfigStruct) AddFilePolicy(filePath string) *AAAConfigStruct {
	if filePath != "" && filePath == AuthPolicyDefaultPath {
		return AAA
	}
	AAA.PolicyFilePath = filePath
	return AAA
}

func (AAA *AAAConfigStruct) Bypass(bypass bool) *AAAConfigStruct {
	AAA.BypassBool = bypass
	return AAA
}

func (webConfig *WebServerConfigStruct) AddKey(keyFile string) *WebServerConfigStruct {
	if webConfig.Key != "" && keyFile == KeyDefaultPath {
		return webConfig
	}
	webConfig.Key = keyFile
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddCert(certFile string) *WebServerConfigStruct {
	if webConfig.Cert != "" && certFile == CertDefaultPath {
		return webConfig
	}
	webConfig.Cert = certFile
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddCA(caFile string) *WebServerConfigStruct {
	if webConfig.Ca != "" && caFile == CADefaultPath {
		return webConfig
	}
	webConfig.Ca = caFile
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddTLSAddr(httpsAddr string) *WebServerConfigStruct {
	if webConfig.HttpsAddr != "" && httpsAddr == HTTPSAddrDefault {
		return webConfig
	}
	webConfig.HttpsAddr = httpsAddr

	return webConfig
}

func (webConfig *WebServerConfigStruct) AddAddr(httpAddr string) *WebServerConfigStruct {
	if webConfig.HttpAddr != "" && httpAddr == HTTPAddrDefault {
		return webConfig
	}
	webConfig.HttpAddr = httpAddr
	return webConfig
}

func LoadMainConfig(configFilePath string) {
	_, err := toml.DecodeFile(configFilePath, MainConfig)
	if err != nil {
		log.Println(err)
	}
}
