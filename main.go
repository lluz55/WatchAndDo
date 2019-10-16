package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/inancgumus/screen"
)

var (
	port                int  // Server port
	opensDefaultBrowser bool // Opens browser
	showQr              bool // Shows QR Code in console to open browser
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

	localAddr := fmt.Sprintf("http://%s:%d", server.addr, server.port)

	fmt.Printf("Listening on %s ...\n", localAddr)

	openBrowser(fmt.Sprintf("%s/wad/", localAddr))

	qr := qrcodeTerminal.New()
	qr.Get(localAddr).Print()

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
	screen.Clear()
	screen.MoveTopLeft()

	// WebSocket mode
	var wsMode bool

	flag.BoolVar(&wsMode, "ws", false, "For configure and listen on webpage")
	flag.BoolVar(&opensDefaultBrowser, "o", false, "Opens default browser")
	flag.BoolVar(&showQr, "qr", false, "Shows QR CODE in console to open default browser")

	flag.Parse()

	// Only work in ws mode for now
	if !wsMode {
		fmt.Println("[!] Only web mode implemented yet\n\tExiting...")
		os.Exit(1)
	}

	// Start server to connect with listeners
	startServer()

	for {
		time.Sleep(24 * time.Hour)
	}

}
