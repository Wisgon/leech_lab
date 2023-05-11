package main

import (
	"context"
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

func cleanup(cancel context.CancelFunc) {
	// save neure map
	log.Println("closing db~~~")
	neure.NeureMap.Range(func(key, value any) bool {
		neureObj := value.(*neure.Neure)
		neureObj.UpdateNeure2DB()
		neureObj.NeureSleep() // todo:运行久了会关不掉，可能是这个sleep的原因，排查
		return true
	})
	database.CloseDb()

	log.Println("closing all go rouitines")
	cancel()
	// some other cleanup here ~~~

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cleanup(cancel)

	rand.Seed(time.Now().UnixNano()) // set rand seed
	database.InitDb(config.LeechDatasPath, config.SeqBandwidth)
	defer func() {
		if r := recover(); r != nil {
			cleanup(cancel)
		}
	}()

	// get control c signal and invole cleanup, because control c will not execute the defer function
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup(cancel)
		os.Exit(1)
	}()

	go neure.CheckNeureMap(ctx)

	// leech---------------------------------------------------------------------
	websocketRequest := make(chan map[string]interface{})
	websocketResponse := make(chan map[string]interface{})
	leech := leech.Leech{
		Body: &leech.LeechBody{
			Organ: &sync.Map{},
		},
		Brain: &leech.LeechBrain{
			Area: &sync.Map{},
		},
		EnvResponse:        websocketResponse,
		EnvRequest:         websocketRequest,
		SignalPassRecorder: neure.SignalPassRecorder,
	}
	leech.LoadLeech()

	go leech.WakeUpLeech(ctx)
	go leech.RecordSignalPass(ctx)
	go interact.StartInteract(ctx, websocketRequest, websocketResponse)
	// leech-------------------------------------------------------------------

	go utils.ServerStaticFile(config.ProjectRoot + "/visualization/")
	for {
		// run forever
		time.Sleep(60 * time.Minute)
	}
}
