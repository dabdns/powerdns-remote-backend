package connector

type Connector interface {
	Config() (err error)
	Start() (err error)
}
