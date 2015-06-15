package store

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"strconv"
)

type StoreClient struct {
	rpc *rpc.Session
}

func NewStoreClient() *StoreClient {
	return &StoreClient {}
}

func (client *StoreClient) Listen() error {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		log.Info("fail to connect to server.")
		return nil
	}
	client.rpc = rpc.NewSession(conn, true)

//	retval, xerr := client.rpc.Send("read", 0, 1)
//	if xerr != nil {
//		log.Error(xerr)
//		return nil
//	}
//	log.Info(retval.String())

	return nil
}

func (client *StoreClient) Append(entry []byte) error {
	retval, xerr := client.rpc.Send("write", 0, 0, entry)
	if xerr != nil {
		log.Error(xerr)
		return nil
	}
	//log.Info("WRITE: " + retval.Kind().String())
	log.Info("WRITE: " + strconv.FormatBool(retval.Bool()))
	return nil
}

func (client *StoreClient) Read(address uint64) error {
	retval, xerr := client.rpc.Send("read", 0, 0)
	if xerr != nil {
		log.Error(xerr)
		return nil
	}
	log.Info("READ: " + retval.String())
	return nil
}

func (client *StoreClient) Fill(address uint64) error {
	return nil
}

func (client *StoreClient) Trim(address uint64) error {
	return nil
}

func (client *StoreClient) Reconfigure() error {
	return nil
}
