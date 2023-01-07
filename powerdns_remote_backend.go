package main

import (
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
	"github.com/dabdns/powerdns-remote-backend/pkg/config"
	connectorBase "github.com/dabdns/powerdns-remote-backend/pkg/connector"
	connectorHTTP "github.com/dabdns/powerdns-remote-backend/pkg/connector/http"
	connectorPipe "github.com/dabdns/powerdns-remote-backend/pkg/connector/pipe"
	"github.com/spf13/viper"

	"os"
)

func loadConfig() (conf config.Config, err error) {
	conf = config.GetDefaultConfig()
	var vConf config.Config
	v := viper.New()
	// load configuration with viper
	v.AddConfigPath(".")
	v.SetConfigName("dabdns")
	v.SetEnvPrefix("dabdns")
	v.AutomaticEnv()
	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Error reading config file, %s", err))
		}
	}
	err = v.Unmarshal(&vConf)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct, %v", err))
	}
	conf.Merge(&vConf)
	return
}

func main() {
	conf, err := loadConfig()
	delegate := backend.NewDelegate()
	for _, delegateConfig := range conf.Delegates {
		var delegateBase *backend.DelegateBase
		delegateBase, err = backend.NewDelagateBase(delegateConfig)
		if err == nil {
			delegate.AddDelegate(delegateBase)
		} else {
			panic(err)
		}
	}

	var connector connectorBase.Connector
	o, _ := os.Stdout.Stat()
	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		//Display info to the terminal
		connector = connectorHTTP.NewConnectorHTTP(delegate, *conf.Connector.Host, *conf.Connector.Port)
	} else { //It is not the terminal
		// Display info to a pipe
		connector = connectorPipe.NewConnectorPipe(delegate)
	}

	err = connector.Config()
	if err == nil {
		err = connector.Start()
	}
	if err != nil {
		os.Exit(1)
	}
}
