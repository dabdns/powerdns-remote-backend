package socket

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

type ConnectorSocket struct {
	Backend backend.Backend
}

func NewConnectorSocket() *ConnectorSocket {
	return &ConnectorSocket{}
}

func (*ConnectorSocket) Config() (err error) {
	return
}

func (*ConnectorSocket) Start() (err error) {
	return
}
