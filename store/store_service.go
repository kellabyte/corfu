package store
import (
	"reflect"
	"fmt"
	"net"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
)

type StoreService struct {
	epoch uint64
}

func New() *StoreService {
	return &StoreService{}
}

func (service *StoreService) Listen() error {
	res := MsgPackResolver {
		"read": reflect.ValueOf(service.Read),
		"write": reflect.ValueOf(service.Write)}

	serv := rpc.NewServer(res, true, nil)
	listener, error := net.Listen("tcp", "127.0.0.1:50000")
	if (error != nil) {
		return error
	}

	serv.Listen(listener)
	go (func() { serv.Run() })()

	return nil
}

func (service *StoreService) Read(epoch uint64, address uint64) ([]byte, fmt.Stringer) {
	if (epoch != service.epoch) {
		//return nil, SealedError { message: "Sealed", epoch: service.epoch, address: address }
		return []byte("error"), nil
	}

	// TODO: Read bytes from storage and return them.
	return []byte("Hello, world"), nil
}

func (service *StoreService) Write(epoch uint64, address uint64, data []byte) fmt.Stringer {
	if (epoch != service.epoch) {
		return SealedError { message: "Sealed", epoch: service.epoch, address: address }
	}

	// TODO: Write bytes from storage and return them.
	return nil
}
