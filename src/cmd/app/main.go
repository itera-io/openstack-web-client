package main

import (
	"github.com/itera-io/openstack-web-client/api"
	"github.com/itera-io/openstack-web-client/config"
)

func main() {
	cfg := config.GetConfig()

	api.InitServer(cfg)
}
