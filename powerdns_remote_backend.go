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
	defaultHost   string = "0.0.0.0"
	defaultPort   int16  = 8080
	defaultDomain string = "example.com."
	defaultMNAME  string = "ns.icann.org."
	defaultRNAME  string = "noc.dns.icann.org."
)

func main() {
	backendDelegate := backend.NewBackendDelegate()
	soaConfig := ip.NewSOAConfig(defaultMNAME, defaultRNAME)
	backendDelegate.AddDelegate(ip.NewDelegateIP(defaultDomain, soaConfig))

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
