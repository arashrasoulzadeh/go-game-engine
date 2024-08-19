package agent

import (
	"fmt"
	"github.com/arashrasoulzadeh/go-game-engine/agent"
	"github.com/arashrasoulzadeh/go-game-engine/models"
	"github.com/pterm/pterm"
	"math"
	"strconv"
	"time"
	"unsafe"
)

func refresh(objectPool *models.ObjectPool) {
	print("\033[H\033[2J")

	// Print an informational message.
	if isConnected {
		pterm.Success.Println("connected to server")
	} else {
		pterm.Error.Println("disconnected")
	}

	// Create three panels with text, some of them with titles.
	// The panels are created using the DefaultBox style.

	panel1Text := ""
	for key := range packetsCount {
		panel1Text += key + " : " + strconv.Itoa(packetsCount[key]) + "\n"
	}
	panel1 := pterm.DefaultBox.Sprint(panel1Text)
	elapsed := time.Since(timer)
	elapsedSeconds := int(math.Round(elapsed.Seconds()))
	panel2 := pterm.DefaultBox.WithTitle("time connected").Sprint(" seconds:", elapsedSeconds, "\n since:"+timer.Format("2006-01-02 15:04:05"))

	history := ""
	data := objectPool.Get("data") // Assuming this returns a slice

	// Iterate over the slice in reverse order
	for i := len(data) - 1; i >= 0; i-- {
		if transport, ok := data[i].(agent.TransportModel); ok {
			history += fmt.Sprintf("%s with %d bytes of payload\n", transport.ID, unsafe.Sizeof(transport.Payload))
		}
	}
	log := pterm.DefaultBox.WithTitle("latest data packets:\n this history does not include ping/pong packets.").Sprintf(history)

	// Combine the panels into a layout using the DefaultPanel style.
	// The layout is a 2D grid, with each row being an array of panels.
	// In this case, the first row contains panel1 and panel2, and the second row contains only panel3.
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: panel1}, {Data: panel2}},
		{{log}},
	}).Srender()

	// Print the panels layout inside a box with a title.
	// The box is created using the DefaultBox style, with the title positioned at the bottom right.
	pterm.DefaultBox.WithTitle("Agentd").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)

}
