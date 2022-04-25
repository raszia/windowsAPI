package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var FileDefaultPath = "./config.toml"

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
	if path != "" {
		main.configPath = path
		LoadMainConfig(path)
	}
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
	if filePath != "" {
		AAA.PolicyModelPath = filePath
	}
	return AAA
}

func (AAA *AAAConfigStruct) AddFilePolicy(filePath string) *AAAConfigStruct {
	if filePath != "" {
		AAA.PolicyFilePath = filePath
	}
	return AAA
}

func (AAA *AAAConfigStruct) Bypass(bypass bool) *AAAConfigStruct {
	AAA.BypassBool = bypass
	return AAA
}

func (webConfig *WebServerConfigStruct) AddCertKey(keyFile, certFile string) *WebServerConfigStruct {
	if keyFile != "" {
		webConfig.Key = keyFile
	}
	if certFile != "" {
		webConfig.Cert = certFile
	}
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddCA(caFile string) *WebServerConfigStruct {
	if caFile != "" {
		webConfig.Ca = caFile
	}
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddTLSAddr(httpsAddr string) *WebServerConfigStruct {
	if httpsAddr != "" {
		webConfig.HttpsAddr = httpsAddr
	}
	return webConfig
}

func (webConfig *WebServerConfigStruct) AddAddr(httpAddr string) *WebServerConfigStruct {
	if httpAddr != "" {
		webConfig.HttpAddr = httpAddr
	}
	return webConfig
}

func LoadMainConfig(configFilePath string) {
	_, err := toml.DecodeFile(configFilePath, MainConfig)
	log.Println(err)

}
