package main

import (
	"fmt"
	"graph_robot/database"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("closing db~~~")
	database.CloseDb()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			cleanup()
		}
	}()

	// get control c signal and invole cleanup, because control c will not exec defer function
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("thinking...")
		time.Sleep(2 * time.Second) // or runtime.Gosched() or similar per @misterbee
		panic("dfdfdf")
	}
}
