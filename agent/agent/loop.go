package agent

import (
	"bufio"
	"encoding/json"
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"go.uber.org/zap"
	"net"
	"time"
)

func loop(reader *bufio.Reader, conn *net.TCPConn, objectPool *models.ObjectPool) {
	for {
		// Read from the connection until a new line is sent
		data, err := reader.ReadString('\n')
		if err != nil {
			zap.L().Error("connection error: " + err.Error())
			zap.L().Info("trying to reconnect")
			time.Sleep(time.Second * 1)
			conn = connectToServer(nil)
			reader = bufio.NewReader(conn)
		}

		// Unmarshal the data and handle errors
		var model agent.TransportModel
		if err := json.Unmarshal([]byte(data), &model); err != nil {
			zap.L().Error("failed to parse data: " + err.Error())
			continue
		}

		// Process the received model based on its type
		switch model.Type {
		case agent.MessageTypePing:
			packetsCount["ping"]++
		case agent.MessageTypePong:
			packetsCount["pong"]++
		case agent.MessageTypeData:
			packetsCount["data"]++
			objectPool.Append("data", model)
		default:
			zap.L().Error("Unhandled model type: " + model.Type.String())
		}

		go agent.SendMessage(nil, agent.AckMessage{ID: model.ID}, agent.MessageTypeAck, conn)

	}
}
