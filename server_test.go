package goland

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

// decode your data in this func
func StringDecoder(b []byte) (interface{}, error) {
	return string(b), nil
}

// encode your data in this func
func StringEncoder(msg interface{}) []byte {
	return []byte(msg.(string))
}

type StringHandler struct {
	mutex sync.Mutex
	times int
}

// process decoded message
func (t *StringHandler) Handle(ctx ConnectionHandler, msg interface{}) {
	t.mutex.Lock()
	t.times++
	t.mutex.Unlock()

	ctx.Write("copy")
	fmt.Println(t.times)
	fmt.Println("read from client: ", msg)
}

func TestServer(t *testing.T) {
	addr := net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 3333,
	}

	tcpServer, e := NewTcpServer(&addr)
	if e != nil {
		panic(e)
	}

	tcpServer.AddEncoder(StringEncoder)
	tcpServer.AddDecoder(StringDecoder)
	tcpServer.AddHandler(new(StringHandler))
	tcpServer.Start()
}