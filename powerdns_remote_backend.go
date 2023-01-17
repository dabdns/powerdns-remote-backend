package main

import (
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
	"github.com/dabdns/powerdns-remote-backend/pkg/config"
	connectorBase "github.com/dabdns/powerdns-remote-backend/pkg/connector"
	connectorHTTP "github.com/dabdns/powerdns-remote-backend/pkg/connector/http"
	connectorPipe "github.com/dabdns/powerdns-remote-backend/pkg/connector/pipe"
	connectorSocket "github.com/dabdns/powerdns-remote-backend/pkg/connector/socket"
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
	defaultsMap := config.GetDefaultConfigMap()
	err = v.MergeConfigMap(defaultsMap)
	err = v.MergeInConfig()
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
	switch *conf.Connector.Type {
	case "http":
		connector = connectorHTTP.NewConnectorHTTP(delegate, *conf.Connector.Host, *conf.Connector.Port)
	case "pipe":
		connector = connectorPipe.NewConnectorPipe(delegate)
	case "socket":
		connector = connectorSocket.NewConnectorSocket(delegate, *conf.Connector.Network, *conf.Connector.Address, *conf.Connector.Timeout)
	default:
	}

	err = connector.Config()
	if err == nil {
		err = connector.Start()
	}
	if err != nil {
		os.Exit(1)
	}
}
