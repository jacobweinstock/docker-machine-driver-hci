package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	hci "github.com/jacobweinstock/docker-machine-driver-hci/pkg/drivers/hci"
)

func main() {
	plugin.RegisterDriver(new(hci.Driver))
}
