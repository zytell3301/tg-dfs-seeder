package domain

// The reason that this type has created separately is
// that same instance id may be assigned to another node.
// This basically happens when because of any technical
// reason, we have to copy a whole server to another physical
// location. The concept of instance id, makes physical
// details and application view of network independent.
type ServerCommunicationDetails struct {
	Ip string

	// There may be several ports to a service.
	// But for now just pinging port is considered.
	PingingPort int
}
