package main

import (
	"encoding/json"
	"fmt"
	"github.com/inancgumus/screen"
	"log"
	"os"
	"time"
)

func startServer() {
	// Create and start a server
	server := Server{}
	// TODO: [IMPLEMENT] Usage by command line parsed by flag
	// TODO: port setter just for debug
	err := server.StartOnPort(5555)
	// handles server errors on start up
	if err != nil {
		log.Println(err)
	}

	// handles asynchronous server errors
	server.listenErrs = func(err error) {
		log.Println(err)
	}

	time.Sleep(5 * time.Second)

	b, _ := json.Marshal(&MessageSend{Type: "watchingnewfile",
		Payload: &WatchingNewFile{Filename: "Arquivo.txt"},
	})

	server.SendMsgBytes(b)

}

func main() {
	// TODO: REMOVE
	screen.Clear()
	screen.MoveTopLeft()

	// WebSocket mode
	wsMode := false

	// TODO: Use flag to parse os args
	// checks if has ws arg
	if len(os.Args) > 1 {
		for i := range os.Args {
			if os.Args[i] == "--ws" {
				wsMode = true
			}
		}
	}

	// Only work in ws mode for now
	if !wsMode {
		fmt.Println("<!> Only web mode implemented yet\n>>> Exiting...")
		os.Exit(1)
	}

	startServer()

	for {
		time.Sleep(24 * time.Hour)
	}

}
