package main

import (
	"encoding/json"
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
	"github.com/dabdns/powerdns-remote-backend/pkg/config"
	connectorBase "github.com/dabdns/powerdns-remote-backend/pkg/connector"
	connectorHTTP "github.com/dabdns/powerdns-remote-backend/pkg/connector/http"
	connectorPipe "github.com/dabdns/powerdns-remote-backend/pkg/connector/pipe"

	"github.com/spf13/viper"

	"os"
)

func loadDefaultConfig(v *viper.Viper) {
	// Unmarshalling default config json
	var defaultConf config.Config
	err := json.Unmarshal([]byte(config.DefaultConfString), &defaultConf)
	if err != nil {
		panic(fmt.Errorf("Error unmarshalling default config json, %s", err))
	}

	v.SetDefault("connector", defaultConf.Connector)
	v.SetDefault("delegates", defaultConf.Delegates)
}
func loadConfig() (conf config.Config, err error) {
	v := viper.New()

	// load configuration with viper
	v.AddConfigPath(".")
	v.SetConfigName("dabdns")
	v.SetEnvPrefix("dabdns")
	v.AutomaticEnv()
	loadDefaultConfig(v)
	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Error reading config file, %s", err))
		}
	}
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct, %v", err))
	}
	return
}

func main() {
	conf, err := loadConfig()
	backendDelegate := backend.NewDelegate()
	for _, delegateConfig := range conf.Delegates {
		backendDelegate.AddDelegate(backend.NewDelagateBase(delegateConfig))
	}

	var connector connectorBase.Connector
	o, _ := os.Stdout.Stat()
	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		//Display info to the terminal
		connector = connectorHTTP.NewConnectorHTTP(backendDelegate, conf.Connector.Host, conf.Connector.Port)
	} else { //It is not the terminal
		// Display info to a pipe
		connector = connectorPipe.NewConnectorPipe(backendDelegate)
	}

	err = connector.Config()
	if err == nil {
		err = connector.Start()
	}
	if err != nil {
		os.Exit(1)
	}
}
