package main

import (
	"graph_robot/config"
	"graph_robot/database"
	leech "graph_robot/simulate_leech"
	"log"
	"math/rand"
	"sync"
	"time"
)

func cleanup() {
	log.Println("closing db~~~")
	database.CloseDb()
	// some other cleanup here ~~~
}

func main() {
	rand.Seed(time.Now().UnixNano()) // set rand seed
	database.InitDb(config.LeechDatasPath, config.SeqBandwidth)
	defer func() {
		if r := recover(); r != nil {
			cleanup()
		}
	}()

	leechObj := leech.Leech{
		Brain: &leech.LeechBrain{
			Area: &sync.Map{},
		},
		Body: &leech.LeechBody{
			Organ: &sync.Map{},
		},
	}
	leechObj.InitLeech()
	cleanup()
}
