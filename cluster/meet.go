package cluster

import "github.com/garyburd/redigo/redis"

func Meet(c redis.Conn, h string, p int) (interface{}, error) {
	res, err := c.Do("CLUSTER", "MEET", h, p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
