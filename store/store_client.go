package store

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
)

type StoreClient struct {

}

func NewStoreClient() *StoreClient {
	return &StoreClient {}
}

func (storeClient *StoreClient) Listen() error {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		log.Info("fail to connect to server.")
		return nil
	}
	client := rpc.NewSession(conn, true)

	retval, xerr := client.Send("read", 0, 1)
	if xerr != nil {
		log.Error(xerr)
		return nil
	}
	log.Info(retval.String())

	return nil
}
