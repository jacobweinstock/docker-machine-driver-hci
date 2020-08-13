package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/jacobweinstock/docker-machine-driver-hci/pkg/drivers/netapp"
)

func main() {
	plugin.RegisterDriver(new(netapp.Driver))
}
