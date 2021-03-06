package node

import (
	"github.com/go-redis/redis/v8"
	"github.com/weblazy/socket-cluster/discovery"
	"github.com/weblazy/socket-cluster/protocol"
	"github.com/weblazy/socket-cluster/session_storage"
)

type (
	// RedisNode the consistent hash redis node config
	RedisNode struct {
		RedisConf *redis.Options
		Position  uint32 //the position of hash ring
	}
	// NodeConf node config
	NodeConf struct {
		Host                    string // the ip or domain of the node
		Key                     string // the transport path
		Port                    int64  // Node port
		InternalPort            int64  // Node port
		Password                string // Password for auth when connect to other node
		ClientPingInterval      int64
		NodePingInterval        int64                  // Heartbeat interval
		onMsg                   func(context *Context) // callback function when receive client message
		discoveryHandler        discovery.ServiceDiscovery
		protocolHandler         protocol.Protocol
		internalProtocolHandler protocol.Protocol
		sessionStorageHandler   session_storage.SessionStorage
	}
	// Params of onMsg
	Context struct {
		Conn     protocol.Connection
		ClientId string
		Msg      []byte
	}
)

// NewNodeConf creates a new NodeConf.
func NewNodeConf(host string, protocolHandler protocol.Protocol, sessionStorageHandler session_storage.SessionStorage, discoveryHandler discovery.ServiceDiscovery, onMsg func(context *Context)) *NodeConf {
	return &NodeConf{
		Host:                  host,
		Key:                   GetUUID(),
		Port:                  defaultPort,
		Password:              defaultPassword,
		ClientPingInterval:    defaultClientPingInterval,
		NodePingInterval:      defaultNodePingInterval,
		protocolHandler:       protocolHandler,
		sessionStorageHandler: sessionStorageHandler,
		discoveryHandler:      discoveryHandler,
		onMsg:                 onMsg,
	}

}

// WithPassword sets the password for transport node
func (conf *NodeConf) WithPassword(password string) *NodeConf {
	conf.Password = password
	return conf
}

// WithPort sets the port for websocket
func (conf *NodeConf) WithPort(port int64) *NodeConf {
	conf.Port = port
	return conf
}

// WithClientInterval sets the heartbeat interval
func (conf *NodeConf) WithClientInterval(pingInterval int64) *NodeConf {
	conf.ClientPingInterval = pingInterval
	return conf
}
