package config

import "github.com/BurntSushi/toml"

type Node struct {
	Server string `toml:"server"`
	Master string `toml:"master"`
}

type Config struct {
	Nodes []*Node `toml:"node"`
}

func Load(p string) (*Config, error) {
	c := new(Config)
	_, err := toml.DecodeFile(p, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
