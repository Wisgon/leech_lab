package main

import (
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/interact"
	"graph_robot/neure"
	leech "graph_robot/simulate_leech"
	"graph_robot/utils"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func cleanup() {
	log.Println("closing db~~~")
	// save neure map
	// neure.NeureMap.Range(func(key, value any) bool {
	// 	neureObj := value.(*neure.Neure)
	// 	neureObj.UpdateNeure2DB()
	// 	return true
	// })
	database.CloseDb()
	// some other cleanup here ~~~

}

func main() {
	defer cleanup()
	stopCheckNeureMapSignal := make(chan bool, 1)
	rand.Seed(time.Now().UnixNano()) // set rand seed
	database.InitDb(config.LeechDatasPath, config.SeqBandwidth)
	defer func() {
		if r := recover(); r != nil {
			stopCheckNeureMapSignal <- true
			cleanup()
		}
	}()

	go neure.CheckNeureMap(stopCheckNeureMapSignal)

	// get control c signal and invole cleanup, because control c will not execute the defer function
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	done := make(chan int, 1)
	websocketRequest := make(chan map[string]interface{})
	websocketResponse := make(chan map[string]interface{})
	leech := leech.Leech{
		Body: &leech.LeechBody{
			Organ: &sync.Map{},
		},
		Brain: &leech.LeechBrain{
			Area: &sync.Map{},
		},
		EnvResponse: websocketResponse,
		EnvRequest:  websocketRequest,
	}
	leech.LoadLeech()

	go leech.WakeUp()
	go interact.StartInteract(done, websocketRequest, websocketResponse)
	go utils.ServerStaticFile(config.ProjectRoot + "/visualization/")

	for {
		log.Println("thinking...")
		time.Sleep(10 * time.Minute)
	}
}
