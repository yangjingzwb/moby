package networkallocator

import (
	"github.com/docker/libnetwork/drivers/bridge/brmanager"
	"github.com/docker/libnetwork/drivers/host"
	"github.com/docker/libnetwork/drivers/ipvlan/ivmanager"
	"github.com/docker/libnetwork/drivers/macvlan/mvmanager"
	"github.com/docker/libnetwork/drivers/overlay/ovmanager"
	"github.com/docker/libnetwork/drivers/remote"
)

var initializers = []initializer{
	{remote.Init, "remote"},
	{ovmanager.Init, "overlay"},
	{mvmanager.Init, "macvlan"},
	{brmanager.Init, "bridge"},
	{ivmanager.Init, "ipvlan"},
	{host.Init, "host"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []PredefinedNetworkData {
	return []PredefinedNetworkData{
		{"bridge", "bridge"},
		{"host", "host"},
	}
}
