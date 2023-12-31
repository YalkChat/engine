package server

import (
	"encoding/json"
	"log"

	"github.com/AleRosmo/engine/client"
	"github.com/AleRosmo/engine/models/events"
)

func (s *serverImpl) StartReceiver(client client.Client, quit chan struct{}) {
	for {
		select {
		case <-quit:
			// Handle cleanup if needed
			return
		default:
			messageType, receivedEvent, err := client.ReadMessage()
			if err != nil {
				// Handle the error, possibly by logging it and breaking the loop to stop the receiver
				log.Printf("Error reading event from client %d: %v", client.ID(), err)
				break
			}
			// TODO: Missing EOF check but unsure if needed
			if messageType.String() == "MessageText" {

				baseEvent := &events.BaseEvent{} // Assuming BaseEvent is your new RawEvent
				err := json.Unmarshal(receivedEvent, baseEvent)
				if err != nil {
					log.Println("Error unmarshalling:", err)
					continue
				}
				// Create EventWithMetadata and add UserID
				eventWithMetadata := &events.BaseEventWithMetadata{
					Event:  baseEvent,
					UserID: client.ID(), // Assuming you have UserID in your Client struct
				}

				// Forward the event to the server for handling
				if err := s.HandleEvent(eventWithMetadata); err != nil {
					// Handle the error, possibly by logging it
					log.Printf("Error handling event from client %d: %v", client.ID(), err)
				}
			}
		}
	}
}
