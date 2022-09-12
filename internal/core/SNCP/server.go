package SNCP

type server struct {
	// This is the value of instance_id
	// in configs.yaml file.
	InstanceId string

	// This value determines that corresponding server
	// is taking report from how many servers. This is
	// used for example when bootstrapping the server.
	// SNCP chooses most heavy server to take server from.
	ServersCount int32
}
