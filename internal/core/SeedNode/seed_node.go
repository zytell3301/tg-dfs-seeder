package SeedNode

import (
	"github.com/zytell3301/tg-dfs-seeder/internal/core/SNCP"
	"math"
	"sort"
)

type SeedNode struct {
	sNCP              SNCP.Pool
	currentInstanceId string
}

// This method is called when a seed node server
// is getting online. When a seed node is getting
// online (either for first time or after being down)
// it has no dfs client server to receive reports from.
// So we request other seed nodes to assign this node
// some client dfs nodes so the pressure will be balanced
// between SN servers.
func (sn SeedNode) Bootstrap() error {
	servers, err := sn.sNCP.GetServers()
	if err != nil {
		return getSeedNodeServersError
	}
	clientServersCount, err := sn.sNCP.GetAllClientDfsServersCount()
	if err != nil {
		return getClientServersCountError
	}

	// This is the number of servers that
	// this server must have (or maybe one less)
	requiredServers := int(math.Ceil(float64(clientServersCount)/float64(len(servers)))) + 1
	for i := 1; i <= requiredServers; {
		sort.Slice(servers, func(i, j int) bool {
			return servers[i].ServersCount > servers[j].ServersCount
		})
		serverToRequest := &servers[0]

		// This condition indicates that current node is now
		// has the most responsibility between all servers.
		// However, we can't actually say that because there
		// is always inconsistency in the network.
		// But at least we can say this one has reached
		// its limits.
		if serverToRequest.InstanceId == sn.currentInstanceId {
			break
		}
		err := sn.sNCP.RequestServer(serverToRequest)
		if err != nil {
			// Remove the server that encountered the error
			servers = append(servers[:0], servers[1:]...)
			continue
		}
		serverToRequest.ServersCount--
		i++
	}

	// Everything done as planned
	return nil
}
