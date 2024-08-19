package agent

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MessageType int

const (
	MessageTypeAck          = 0
	MessageTypePing         = 1
	MessageTypePong         = 2
	MessageTypeData         = 3
	MessageTypeCreateServer = 4
)

func (t MessageType) String() string {
	switch t {
	case MessageTypeAck:
		return "Ack"
	case MessageTypePing:
		return "Ping"
	case MessageTypePong:
		return "Pong"
	case MessageTypeData:
		return "Data"
	case MessageTypeCreateServer:
		return "CreateServer"
	}
	return "invalid!"
}

type AckMessage struct {
	ID uuid.UUID `json:"id"`
}

type TransportModel struct {
	ID      uuid.UUID   `json:"id"`
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Data    string      `json:"data"`
	Type    MessageType `json:"type"`
	Payload []byte      `json:"paylod"`
}

type WorkerAgent struct {
	IP           string
	LastPingTime time.Time
}

func (w *WorkerAgent) Status() string {
	diff := time.Now().Sub(w.LastPingTime)
	if diff > time.Second*3 && diff < 10 {
		return "waiting to reconnect!"
	}
	if diff > time.Second*10 {
		return "seems disconnected!"
	}

	return "connected"
}

// UpdatePingTimeWithIP updates the LastPingTime of the agent with the specified IP.
func UpdatePingTimeWithIP(ip string) error {
	// Get the current time
	now := time.Now()
	pool := GetAgentsPool()

	// Iterate through the AgentsPool to find the agent with the matching IP
	for i := range *pool {
		if (*pool)[i].IP == ip {
			// Update the LastPingTime for the matching agent
			(*pool)[i].LastPingTime = now
			return nil
		}
	}

	// If IP address is not found in the pool, return an error
	return fmt.Errorf("agent with IP %s not found", ip)
}
