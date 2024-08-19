package agent

import (
	"encoding/json"
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"go.uber.org/zap"
	"net"
	"time"
)

func connectToServer(objectPool *models.ObjectPool) *net.TCPConn {
	server := "127.0.0.1:3001"
	zap.L().Info("agentd is connecting to : " + server)

	isConnected = false
	objectPool = models.NewObjectPool()

	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		zap.L().Error("connection error: " + err.Error())
		time.Sleep(time.Second * 1)
		return connectToServer(objectPool)
	}

	// Connect to the address with tcp
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		time.Sleep(time.Second * 1)
		return connectToServer(objectPool)
	}

	pingPacket, _ := json.Marshal(agent.TransportModel{Type: agent.MessageTypePing})
	// Send a message to the server
	_, err = conn.Write([]byte(string(pingPacket) + "\n"))
	if err != nil {
		time.Sleep(time.Second * 1)
		return connectToServer(nil)
	}

	timer = time.Now()

	zap.L().Info("connected to server: " + server)

	isConnected = true

	return conn
}
