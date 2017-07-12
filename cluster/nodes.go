package cluster

import "github.com/garyburd/redigo/redis"

func Nodes(c redis.Conn) (interface{}, error) {
	res, err := c.Do("CLUSTER", "NODES")
	if err != nil {
		return nil, err
	}
	return res, nil
}
