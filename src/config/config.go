package config

import (
	"os"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/spf13/viper"
)

func defaults() {
	Config.SetDefault("name", petname.Generate(3, "-"))
	Config.SetDefault("log", "info")
	Config.SetDefault("rpc.ethereum", "https://eth.llamarpc.com")
	Config.SetDefault("broker", "wss://shinobi.brokers.kenshi.io")
}

var Config *viper.Viper
var Secrets *viper.Viper

func init() {
	Config = viper.New()
	Secrets = viper.New()
}

func LoadConfig(ConfigFileName string, SecretsFileName string) {
	defaults()

	Config.SetConfigFile(ConfigFileName)
	err := Config.ReadInConfig()

	if err != nil {
		if os.IsNotExist(err) {
			Config.WriteConfig()
		} else {
			panic(err)
		}
	}

	Secrets.SetConfigFile(SecretsFileName)
	err = Secrets.MergeInConfig()

	if err != nil && os.IsExist(err) {
		panic(err)
	}

}
