package main

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
	connectorBase "github.com/dabdns/powerdns-remote-backend/pkg/connector"
	connectorHTTP "github.com/dabdns/powerdns-remote-backend/pkg/connector/http"
	connectorPipe "github.com/dabdns/powerdns-remote-backend/pkg/connector/pipe"
	"github.com/dabdns/powerdns-remote-backend/pkg/ip"
	"os"
)

const (
	defaultHost   string = "localhost"
	defaultPort   int16  = 8080
	defaultDomain string = "example.com."
)

func main() {
	backendDelegate := backend.NewBackendDelegate()
	backendDelegate.AddDelegate(ip.NewDelegateIP(defaultDomain))

	var connector connectorBase.Connector
	o, _ := os.Stdout.Stat()
	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		//Display info to the terminal
		connector = connectorHTTP.NewConnectorHTTP(backendDelegate, defaultHost, defaultPort)
	} else { //It is not the terminal
		// Display info to a pipe
		connector = connectorPipe.NewConnectorPipe(backendDelegate)
	}

	err := connector.Config()
	if err == nil {
		err = connector.Start()
	}
	if err != nil {
		os.Exit(1)
	}
}
