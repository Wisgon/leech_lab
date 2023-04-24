package main

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	leech "graph_robot/simulate_leech"
	"math/rand"
	"time"
)

func cleanup() {
	fmt.Println("closing db~~~")
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

	leechObj := leech.Leech{}
	leechObj.InitLeech()
	cleanup()
}
