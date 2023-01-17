package socket

import (
	"encoding/json"
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
	"io"
	"net"
	"os"
	"time"
)

type ConnectorSocket struct {
	Backend  backend.Backend
	Network  string
	Address  string
	Timeout  int64
	listener net.Listener
}

func NewConnectorSocket(backend backend.Backend, network string, address string, timeout int64) *ConnectorSocket {
	return &ConnectorSocket{
		Backend: backend,
		Network: network,
		Address: address,
		Timeout: timeout,
	}
}

func (connectorSocket *ConnectorSocket) Config() (err error) {
	if connectorSocket.Network == "unix" {
		// Create file if needed
		_, _ = os.OpenFile(connectorSocket.Address, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModeSocket|0660)
	}
	var listener net.Listener
	listener, err = net.Listen(connectorSocket.Network, connectorSocket.Address)
	if err == nil {
		defer listener.Close()
		connectorSocket.listener = listener
	} else {
		panic(err)
	}
	return
}

func (connectorSocket *ConnectorSocket) Start() (err error) {
	go connectorSocket.run()
	return
}

func (connectorSocket *ConnectorSocket) run() {
	for {
		conn, err := connectorSocket.listener.Accept()
		if err != nil {
			go connectorSocket.handleConnection(conn)
		}
	}
}

func (connectorSocket *ConnectorSocket) handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		// Handle timeout
		if connectorSocket.Timeout > 0 {
			_ = conn.SetDeadline(time.UnixMilli(connectorSocket.Timeout))
		} else {
			_ = conn.SetDeadline(time.UnixMilli(0))
		}
		n, err := conn.Read(buffer)
		if err == nil {
			if n > 0 {
				var req backend.Request
				var resp backend.Response
				err = json.Unmarshal(buffer, &req)
				if err == nil {
					fmt.Printf("%s\n", buffer[:n])
					serviceErr := connectorSocket.Backend.Service(&req, &resp)
					if serviceErr != nil {
						resp.Result = false
					}
					respBytes, marshallErr := json.Marshal(resp)
					if marshallErr != nil {
						_, _ = conn.Write(respBytes)
					}
				}
			}
		} else if err == io.EOF {
			break
		}
	}
}
