package cluster

import (
	"github.com/garyburd/redigo/redis"
	"github.com/hirokazumiyaji/redis-cluster/config"
)

type Client struct {
	Conn   redis.Conn
	Master redis.Conn
	Node   *config.Node
	NodeId string
}

func NewClients(nodes []*config.Node) (map[string]*Client, error) {
	clients := make(map[string]*Client)
	for _, n := range nodes {
		c, ok := clients[n.Server]
		if !ok {
			c = &Client{}
			conn, err := redis.Dial("tcp", n.Server)
			if err != nil {
				return nil, err
			}
			c.Conn = conn
		}
		c.Node = n
		if c.Master != nil {
			clients[n.Server] = c
			continue
		}
		if n.Master != "" {
			if m, ok := clients[n.Master]; ok {
				c.Master = m.Conn
			} else {
				mc, err := redis.Dial("tcp", n.Master)
				if err != nil {
					return nil, err
				}
				c.Master = mc
				clients[n.Master] = &Client{
					Conn: mc,
				}
			}
		}
		clients[n.Server] = c
	}
	return clients, nil
}

func CloseClientAll(clients map[string]*Client) {
	for _, c := range clients {
		c.Conn.Close()
	}
}
