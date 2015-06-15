package token

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

const DATABASE_FILE_NAME = "token.db"
const SYSTEM_TABLE_NAME = ".system.token"
const TOKEN_KEY_NAME = "token"

type TokenService struct {
	epoch uint64
	db    *bolt.DB
}

func New() *TokenService {
	return &TokenService{}
}

func (service *TokenService) Listen() error {

	// Open the storage engine.
	db, err := bolt.Open(DATABASE_FILE_NAME, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	db.MaxBatchSize = 100
	db.MaxBatchDelay = 50 * time.Millisecond

	service.db = db

	err = service.CreateSystemTable()
	if err != nil {
		return err
	}

	res := serialization.MsgPackResolver{
		"gettoken": reflect.ValueOf(service.GetToken),
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

func (service *TokenService) Close() error {
	service.db.Close()
	return nil
}

func (service *TokenService) CreateSystemTable() error {
	service.db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(SYSTEM_TABLE_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return nil
}

func (service *TokenService) GetToken() (uint64, fmt.Stringer) {
	var token uint64

	service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(SYSTEM_TABLE_NAME))
		key_bytes := []byte(TOKEN_KEY_NAME)

		// Get the token and increment the value.
		value := bucket.Get(key_bytes)
		token,_ = binary.Uvarint(value)
		return nil
	})
	return token, nil
}

func (service *TokenService) Increment() (uint64, fmt.Stringer) {
	var token uint64

	service.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(SYSTEM_TABLE_NAME))
		key_bytes := []byte(TOKEN_KEY_NAME)

		// Get the token and increment the value.
		value := bucket.Get(key_bytes)
		token,_ = binary.Uvarint(value)
		token++

		// Store the new token.
		token_bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(token_bytes, token)

		_ = bucket.Put(key_bytes, token_bytes)
		return nil
	})
	return token, nil
}
