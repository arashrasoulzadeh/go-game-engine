package agent

import (
	"bufio"
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"sync"
	"time"
)

var packetsCount map[string]int = make(map[string]int)
var timer time.Time

var isConnected bool

func Agent(silent bool) {
	objectPool := models.NewObjectPool()
	conn := connectToServer(objectPool)

	// Create a ticker that triggers every 2 seconds
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			if !silent {
				refresh(objectPool)
			}
			go agent.SendMessage("", nil, agent.MessageTypePong, conn)

			time.Sleep(1 * time.Second)

		}
	}()

	var wg sync.WaitGroup
	reader := bufio.NewReader(conn)
	wg.Add(1)
	go loop(reader, conn, objectPool)

	wg.Wait()
}
