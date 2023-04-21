package main

import (
	"fmt"
	"graph_robot/config"
	"graph_robot/database"
	"graph_robot/interact"
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

	done := make(chan int, 1)
	go interact.StartInteract(done)

	for {
		fmt.Println("thinking...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
		done <- 0
		// panic("dfdfdf")
	}
}
