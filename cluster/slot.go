package cluster

import "github.com/garyburd/redigo/redis"

func AddSlots(c redis.Conn, slots []int) (interface{}, error) {
	res, err := c.Do("CLUSTER", "ADDSLOTS", slots)
	if err != nil {
		return nil, err
	}
	return res, nil
}
