package SNCP

// This interface is the implementation for
//connection pool's  outgoing requests to other seed nodes.
type Transport interface {
	// This request sends a server request to
	// supplied seed node. The seed node itself
	// assigns a server to this node and no
	// extra operation is required from this side.
	// This gives us the ability to apply a
	// selection algorithm to choose the best
	// server to assign.
	RequestServer(targetSN *server) error
}
