package plugincom

type TransportPlugin interface {
	ServerPlugin
	ClientPlugin
}

type ServerPlugin interface {
	Start() error
	Stop() error
	// Accept used by server to accept a new connection.
	// This function corresponds to the function Dail(addr).
	Accept() (ConnPlugin, error)
}

type ClientPlugin interface {
	// Dail used by client to dail the remote server address. The format of addr is host:port.
	// This function corresponds to the function Accept().
	Dail(addr string) (ConnPlugin, error)
}

type ConnPlugin interface {
	SendMsg([]byte) error
	RecvMsg() ([]byte, error)

	RemoteAddr() string
	Close() error
}
