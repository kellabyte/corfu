package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kellabyte/corfu/store"
)

func main() {
	log.Info("Creating storage service")
	storeService := store.New()

	log.Info("Starting storage service")
	storeService.Listen()

	log.Info("Creating storage client")
	storeClient := store.NewStoreClient()

	log.Info("Starting storage client")
	storeClient.Listen()

	storeClient.Append([]byte("hello, world!"))
	storeClient.Read(0)
}
