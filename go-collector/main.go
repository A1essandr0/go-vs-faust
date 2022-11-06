package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	// brokers = []string{"localhost:9093"}
	brokers = []string{"kafka-go-vs-faust:9092"}
	port = "0.0.0.0:5005"
	topic goka.Stream = "test-events-from"

	tmc *goka.TopicManagerConfig
	cfg *sarama.Config
	Emitter *goka.Emitter
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
	var err error
	Emitter, err = goka.NewEmitter(brokers, topic, new(codec.String))
	if err != nil { log.Fatalf("error creating emitter: %v", err) }

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	server := &http.Server{
		Addr: port,
		Handler: mux,
	}

	log.Printf("Starting server on %s...", port)
	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("...server closed")
	} else if err != nil {
		log.Printf("Error starting server: %s", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	event := RawEvent{
		Source: r.URL.Query().Get("source"),
		EventName: r.URL.Query().Get("event_name"),
		EventStatus: r.URL.Query().Get("event_status"),
		Created: r.URL.Query().Get("created"),
		Payout: r.URL.Query().Get("payout"),
	}
	marshalled, err := json.Marshal(event)
	serialized := string(marshalled)

	if err == nil {
		io.WriteString(w, "OK")
		log.Printf("received OK: %s", serialized)
		e := Emitter.EmitSync("key", serialized)
		if e == nil { log.Println("sent OK") }
	}	
}