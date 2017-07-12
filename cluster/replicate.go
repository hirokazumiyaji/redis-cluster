package cluster

import "github.com/garyburd/redigo/redis"

func Replicate(c redis.Conn, n string) (interface{}, error) {
	res, err := c.Do("CLUSTER", "REPLICATE", n)
	if err != nil {
		return nil, err
	}
	return res, nil
}
