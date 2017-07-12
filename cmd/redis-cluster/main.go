package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hirokazumiyaji/redis-cluster/cluster"
	"github.com/hirokazumiyaji/redis-cluster/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var path string
	cmd := &cobra.Command{}

	createCommand := &cobra.Command{
		Use:   "create",
		Short: "cluster create",
		Long: `cluster create
ex)
  redis-cluster create -c <config path>`,
		RunE: func(c *cobra.Command, args []string) error {
			conf, err := config.Load(path)
			if err != nil {
				return err
			}

			if len(conf.Nodes) == 0 {
				return errors.New("node is empty")
			}

			clients, err := cluster.NewClients(conf.Nodes)
			if err != nil {
				return err
			}

			defer cluster.CloseClientAll(clients)

			if len(clients) < 2 {
				return errors.New("one redis process can't create cluster")
			}

			servers := make([]string, 0, len(clients))
			for server, _ := range clients {
				servers = append(servers, server)
			}
			mainClient, _ := clients[servers[0]]
			for _, server := range servers[1:] {
				hostAndPort := strings.Split(server, ":")
				port, err := strconv.Atoi(hostAndPort[1])
				if err != nil {
					return errors.Wrap(
						err,
						fmt.Sprintf("connection server: %s, target server: %s", servers[0], server),
					)
				}
				result, err := cluster.Meet(mainClient.Conn, hostAndPort[0], port)
				if err != nil {
					return errors.Wrap(
						err,
						fmt.Sprintf("connection server: %s, target server: %s", servers[0], server),
					)
				}
				fmt.Printf("%v\n", result)
			}

			for _, client := range clients {
				if client.Master == nil {
					continue
				}
				result, err := cluster.Replicate(client.Conn, client.Node.Master)
				if err != nil {
					return errors.Wrap(
						err,
						fmt.Sprintf("master: %s, slave: %s", client.Node.Master, client.Node.Server),
					)
				}
				fmt.Printf("%v\n", result)
			}

			return nil
		},
	}
	createCommand.Flags().StringVarP(&path, "config", "c", "", "config path")
	cmd.AddCommand(createCommand)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
