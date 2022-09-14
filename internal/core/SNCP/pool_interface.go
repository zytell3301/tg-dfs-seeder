package SNCP

type Pool interface {
	AssignServerToOtherNode(destinationServer string) error
	GetServers() ([]server, error)
	GetAllClientDfsServersCount() (int, error)
	RequestServer(serverToRequest *server) error
}
