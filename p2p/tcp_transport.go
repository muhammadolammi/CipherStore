package p2p

import (
	"log"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAdrress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAdrress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAdrress)
	if err != nil {
		return err
	}

	return nil
}

func (t *TCPTransport) startAccept() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Printf("accept error. err:%s\n", err)
		}
		go t.handleConn(conn)

	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	log.Printf("new incomming conn. conn: %s\n", conn)
}
