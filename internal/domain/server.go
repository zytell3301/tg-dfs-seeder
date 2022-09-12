package domain

type Server struct {
	// This value will be used to uniquely identify
	// every server.
	InstanceId string

	// This is the total data that a server holds at
	// pinging moment. It is used for calculating
	// selection probability.
	VolumeSize int64

	// This is a value at maximum of 1 and minimum of 0.
	// It is measured based on volume size of this server
	// and neighbor servers. By neighbor servers I mean
	// the first server that has higher volume size and the
	// first server that has lower volume size in comparison
	// to current server
	SelectionProbability float64

	ServerCommunicationDetails ServerCommunicationDetails
}
