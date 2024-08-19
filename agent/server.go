package agent

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"os"
	"sync"
	"time"
)

var bus *models.Bus

func Server(cn *gochannel.GoChannel) {

	port := ":3001"
	// Resolve the string address to a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start listening for TCP connections on the given address
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		// Accept new connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// Handle new connections in a Goroutine for concurrency

		go handleConnection(conn, cn)
	}

}

func handleConnection(conn net.Conn, cn *gochannel.GoChannel) {
	defer conn.Close()
	fmt.Println("agent listening on ", conn.RemoteAddr())
	AppendToAgentsPool(WorkerAgent{IP: conn.RemoteAddr().String()})
	ticker := time.NewTicker(1 * time.Second)
	timeoutTicker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
		timeoutTicker.Stop()
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ticker.C:
				go SendMessage("", nil, MessageTypePing, conn)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-timeoutTicker.C:
				conn.Close()
				DeleteFromAgentsPool(conn.RemoteAddr().String())
			}
		}
	}()

	go func() {
		bus = models.GetBus()
		for msg := range bus.Bus {
			go SendMessage(msg, models.GetBus(), MessageTypeData, conn)
		}
	}()

	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err == nil {

				// Print the data read from the connection to the terminal
				var t TransportModel

				err = json.Unmarshal([]byte(data), &t)
				if err != nil {
					zap.L().Error("error in packet unmarshal = " + err.Error() + " data: " + data)

				}

				if t.Type == MessageTypePong {
					err := UpdatePingTimeWithIP(conn.RemoteAddr().String())
					if err != nil {
						zap.L().Error("error updating ping time: " + err.Error())
					}
				}
				if t.Type == MessageTypeAck {
					var ack AckMessage
					e := json.Unmarshal(t.Payload, &ack)
					if e != nil {
						zap.L().Error("error in packet unmarshal = " + e.Error())
					}
					zap.L().Info("ack from client: " + conn.RemoteAddr().String() + " => " + ack.ID.String())
				}
				if t.Type == MessageTypeData {
					zap.L().Info("received data from client: " + conn.RemoteAddr().String() + " => " + fmt.Sprintf("%+v", t))
				}
			}
		}
	}()

	wg.Wait()
}

func SendMessage(msg interface{}, payload interface{}, messageType MessageType, conn net.Conn) (*TransportModel, error) {

	p, e := json.Marshal(payload)
	if e != nil {
		zap.L().Error("error in packet marshal = " + e.Error())
		return nil, errors.New("error in packet marshal = " + e.Error())
	}
	transportModel := TransportModel{
		Type:    messageType,
		ID:      uuid.New(),
		Message: msg,
		Payload: p,
	}
	b, err := json.Marshal(transportModel)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 100)

	_, err = conn.Write([]byte(string(b) + "\n"))

	zap.L().Info("message sent to " + conn.RemoteAddr().String() + " : " + fmt.Sprintf("%v", transportModel))

	return &transportModel, nil

}
