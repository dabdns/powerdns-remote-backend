package zeromq

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

type ConnectorZeroMQ struct {
	Backend backend.Backend
}

func NewConnectorZeroMQ() *ConnectorZeroMQ {
	return &ConnectorZeroMQ{}
}

func (*ConnectorZeroMQ) Config() (err error) {
	return
}

func (*ConnectorZeroMQ) Start() (err error) {
	return
}
