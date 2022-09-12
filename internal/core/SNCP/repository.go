package SNCP

import "github.com/zytell3301/tg-dfs-seeder/internal/domain"

// This interface is only responsible for repository
// related actions.
type Repository interface {

	// This method's main usage is when CP is requested
	// to deliver a server. This method only gives a list
	// of all seed node servers. No logic applied to servers
	// sorting or anything in here.
	GetServers() []server

	// When a dfs client joins the network, it introduces
	// itself to SN ring. This method is used for this to
	// assign newly added client to a seed node.
	PutServer(id string)

	// This method transfers the supplied server's
	// responsibility to a target seed node. Mostly used when a
	// fresh seed node wants to take responsibility
	// of some dfs clients.
	TransferServerResponsibility(serverId string, targetSNId string)

	// Gets a list of all client dfs servers.
	GetAllClientDfsServers() []domain.Server
}
