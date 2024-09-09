package p2p

// Peer is an interface that represent the remote node.
type Peer interface{}

// Transport is anything that handles communication between the nodes in the network.

type Transport interface {
	ListenAndAccept() error
}
