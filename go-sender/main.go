package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	groupName goka.Group = "gosender"

	// brokers = []string{"localhost:9093"}
	brokers = []string{"kafka-go-vs-faust:9092"}
	sourceTopic goka.Stream = "test-events-from"
	processedTopic goka.Stream = "test-events-to"

	tmc *goka.TopicManagerConfig
	cfg *sarama.Config
)


func init() {
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	cfg = goka.DefaultConfig()
	cfg.Version = sarama.V2_4_0_0
	// cfg.Net.SASL.Enable = true
	// cfg.Net.SASL.User = "username"
	// cfg.Net.SASL.Password = "password"
	goka.ReplaceGlobalConfig(cfg)
}

func main() {
	runApplication()
}

// main logic and processing happens here
func OnEvent(ctx goka.Context, rawmsg interface{}) {
	msg := RawEvent{}
	json.Unmarshal([]byte(fmt.Sprint(rawmsg)), &msg)

	log.Printf("Received event: %+v", msg)

	// emitting result
	marshalled, err := json.Marshal(msg)
	if err != nil {
		log.Printf("error sending data %v; %v", msg, err)
	} else {
		log.Println("Sending serialized further: ", string(marshalled))
		ctx.Emit(processedTopic, "key", string(marshalled))		
	}
}


func runApplication() {
	group := goka.DefineGroup(
		groupName,
		goka.Input(sourceTopic, new(codec.String), OnEvent),
		goka.Output(processedTopic, new(codec.String)),
		goka.Persist(new(codec.Int64)),
	)
	application, err := goka.NewProcessor(
		brokers, 
		group,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
	)
	if err != nil { log.Fatalf("Couldn't create application: %v", err)}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)

	go func() {
		defer close(done)
		err = application.Run(ctx)
		if err != nil {
			log.Fatalf("Error running application: %v", err)
		} else {
			log.Printf("Application shutdown cleanly")
		}
	}()
	
	// mechanics to gracefully stop application
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Gracefully stopping...")
	<- wait
	cancel()
	<- done
}