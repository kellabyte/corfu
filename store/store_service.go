package store
import (
	"reflect"
	"fmt"
	"net"
	"github.com/kellabyte/corfu/serialization"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"github.com/boltdb/bolt"
	"time"
	"encoding/binary"
)

const DATABASE_FILE_NAME = "corfu.db"
const SYSTEM_TABLE_NAME = ".system.store"

type StoreService struct {
	epoch uint64
	db *bolt.DB
}

func New() *StoreService {
	return &StoreService{}
}

func (service *StoreService) Listen() error {

	// Open the storage engine.
	db, err := bolt.Open(DATABASE_FILE_NAME, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	//defer db.Close()
	service.db = db

	err = service.CreateSystemTable()
	if err != nil {
		return err
	}

	res := serialization.MsgPackResolver {
		"read": reflect.ValueOf(service.Read),
		"write": reflect.ValueOf(service.Write),
		"delete": reflect.ValueOf(service.Delete),
		"seal": reflect.ValueOf(service.Seal),
	}

	serv := rpc.NewServer(res, true, nil)
	listener, error := net.Listen("tcp", "127.0.0.1:50000")
	if (error != nil) {
		return error
	}

	serv.Listen(listener)
	go (func() { serv.Run() })()

	return nil
}

func (service *StoreService) CreateSystemTable() error {
	service.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(SYSTEM_TABLE_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return nil
}

func (service *StoreService) Read(epoch uint64, address uint64) ([]byte, fmt.Stringer) {
	if (epoch != service.epoch) {
		//return nil, SealedError { message: "Sealed", epoch: service.epoch, address: address }
		return []byte("error"), nil
	}

	var value []byte

	service.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(SYSTEM_TABLE_NAME))

		key_bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(key_bytes, address)

		value = bucket.Get(key_bytes)
		return nil
	})
	return value, nil
}

func (service *StoreService) Write(epoch uint64, address uint64, data []byte) (bool, fmt.Stringer) {
	if (epoch != service.epoch) {
		return false, SealedError { message: "Sealed", epoch: service.epoch, address: address }
	}

	err := service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(SYSTEM_TABLE_NAME))

		key_bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(key_bytes, address)

		err := bucket.Put(key_bytes, data)
		return err
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (service *StoreService) Delete(address uint64) (bool, fmt.Stringer) {
	return true, nil
}

func (service *StoreService) Seal(epoch uint64) (bool, uint64, fmt.Stringer) {
	return true, 0, nil
}
