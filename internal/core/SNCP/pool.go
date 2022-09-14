package SNCP

import (
	"errors"
	"math"
	"sort"
)

// This package is used to give us the ability
// to communicate with other instances of dfs seed node.
// For example consider a situation that current node
// went down. So all dfs clients that were using this
// seed node will use another seed node to send their
// reports. After that this service again becomes online,
// there is no client that is sending its report to this
// instance. So it will easily just get server from
// other instances and joins the flow again.
type pool struct {
	repository        Repository
	transport         Transport
	currentInstanceId string
}

func NewSNCP(repository Repository, transport Transport, currentInstanceId string) *pool {
	return &pool{
		repository:        repository,
		transport:         transport,
		currentInstanceId: currentInstanceId,
	}
}

// This method is called when a seed node server
// is getting online. When a seed node is getting
// online (either for first time or after being down)
// it has no dfs client server to receive reports from.
// So we request other seed nodes to assign this node
// some client dfs nodes so the pressure will be balanced
// between SN servers.
func (p pool) Bootstrap() {
	servers := p.repository.GetServers()
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].ServersCount > servers[j].ServersCount
	})
	clientServers := p.repository.GetAllClientDfsServers()

	// This is the number of servers that
	// this server must have (or maybe one less)
	requiredServers := int(math.Ceil(float64(len(clientServers))/float64(len(servers)))) + 1
	for i := 1; i <= requiredServers; {
		serverToRequest := &servers[0]

		// This condition indicates that current node is now
		// has the most responsibility between all servers.
		// However, we can't actually say that because there
		// is always inconsistency in the network.
		// But at least we can say this one has reached
		// its limits.
		if serverToRequest.InstanceId == p.currentInstanceId {
			break
		}
		err := p.transport.RequestServer(serverToRequest)
		if err != nil {
			// Remove the server that encountered the error
			servers = append(servers[:0], servers[1:]...)
			continue
		}
		serverToRequest.ServersCount--
		i++
	}
}

// This method is used when another SN instance
// requests for dfs client server to get reports from.
// So we must give one of current node's client dfs
// servers, to requesting node.
func (p pool) AssignServerToOtherNode(destinationServer string) error {
	// So at first step we fetch a list of all servers
	// assigned to current node.
	clientServers := p.repository.GetNodeClientDfsServers(p.currentInstanceId)
	if len(clientServers) < 1 {
		return errors.New("an error encountered while assigning a client dfs. No node assigned")
	}

	err := p.repository.TransferServerResponsibility(clientServers[0].InstanceId, destinationServer)
	if err != nil {
		return errors.New("an error encountered while transferring server responsibility")
	}

	return nil
}
