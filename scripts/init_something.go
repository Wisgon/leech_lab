package main

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	leech "graph_robot/simulate_leech"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
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

	// get control c signal and invole cleanup, because control c will not execute the defer function
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	leechObj := leech.Leech{}
	leechObj.InitLeech()
}
