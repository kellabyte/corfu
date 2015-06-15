package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/paulbellamy/ratecounter"
	"github.com/kellabyte/corfu/token"
	"time"
	"sync"
)

func main() {
	log.Info("Creating token service")
	tokenService := token.New()
	defer tokenService.Close()

	log.Info("Starting token service")
	tokenService.Listen()

	counter := ratecounter.NewRateCounter(1 * time.Second)
	token, _ := tokenService.GetToken()

	log.Infof("Token: %d", token)

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := token; i < token + 10000; i++ {
		wg.Add(1)

		lock.Lock()
		//go func() {
			defer wg.Done()
		//	tokenService.Increment()
			counter.Incr(1)
		//}()
		lock.Unlock()
	}

	wg.Wait()
	token, _ = tokenService.GetToken()
	log.Infof("Rate: %d Token: %d", counter.Rate(), token)
}
