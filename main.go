package main

import (
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := hazelcast.NewConfig()
	config.StatsConfig.Enabled = true
	client, err := hazelcast.StartNewClientWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client is running...")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	client.Shutdown()
}
