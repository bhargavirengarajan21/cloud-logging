package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/nats-io/nats.go"
)

var (
	natsConn *nats.Conn
	once     sync.Once
)

// Event structure for logging
type Event struct {
	Message string `json:"message"`
}

// Initialize NATS connection and subscription
func init() {
	once.Do(func() {
		var err error
		natsConn, err = nats.Connect("nats://nats.openfaas.svc.cluster.local:4222")
		if err != nil {
			log.Fatalf("Failed to connect to NATS: %v", err)
			return
		}
		log.Println("Connected to NATS")

		// Subscribe to "log-events" topic for real-time logging
		_, err = natsConn.Subscribe("log-events", func(msg *nats.Msg) {
			log.Printf("Received message: %s", string(msg.Data))
		})

		if err != nil {
			log.Fatalf("Failed to subscribe to NATS topic: %v", err)
			return
		}
		log.Println("Subscribed to log-events topic")
	})
}

// HTTP handler to publish events to NATS
func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if event.Message == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
	}

	// Publish event to "log-events" topic
	err = natsConn.Publish("log-events", []byte(event.Message))
	if err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		log.Printf("Failed to publish event: %v", err)
		return
	}

	log.Printf("Published event: %s", event.Message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Event published successfully: %s", event.Message)))
}
