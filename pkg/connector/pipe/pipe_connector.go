package pipe

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

type ConnectorPipe struct {
	Backend backend.Backend
}

func NewConnectorPipe(backend backend.Backend) *ConnectorPipe {
	return &ConnectorPipe{
		Backend: backend,
	}
}

func (*ConnectorPipe) Config() (err error) {
	return
}

func (*ConnectorPipe) Start() (err error) {
	return
}
