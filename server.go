package main

import (
	// "errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// MessageSend through ws listerners
type MessageSend struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// WatchingNewFile comunicate new file through ws
type WatchingNewFile struct {
	Filename string
}

// Makes sure ws has at least one listener to send message to
var oneListener bool = false

// Upgrade connection to ws
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Server for communicate with clients
type Server struct {
	port       int
	addr       string
	listenErrs func(err error)
	_sendMsgWS func(msg []byte)
}

// Start server with random port
func (s *Server) Start() error {
	s.port = getRandomPort()
	addr, err := getIPAddress()
	if err != nil {
		return err
	}
	s.addr = addr
	s.listen()
	return nil
}

// StartOnPort server On difined port
func (s *Server) StartOnPort(port int) error {
	s.port = port
	addr, err := getIPAddress()
	if err != nil {
		return err
	}
	s.addr = addr
	s.listen()
	return nil
}

// SendMsg through ws
func (s *Server) SendMsg(msg string) {
	if !oneListener {
		return
	}
	s._sendMsgWS([]byte(msg))
}

// SendMsgBytes through ws
func (s *Server) SendMsgBytes(msg []byte) {
	if !oneListener {
		return
	}
	s._sendMsgWS(msg)
}

// handles ws connection
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	s._sendMsgWS = func(msg []byte) {
		if socket != nil {
			socket.WriteMessage(1, msg)
		}
	}

	oneListener = true
	for {
		// Vamos ler a mensagem recebida via Websocket
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			oneListener = false
			return
		}

		// Logando no console do Webserver
		fmt.Println("Message received: ", string(msg))

		// Devolvendo a mensagem recebida de volta para o cliente
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			oneListener = false
			return
		}
	}
}

// Listen on choosen port
func (s *Server) listen() {
	go func() {
		http.HandleFunc("/wad/", index)
		http.HandleFunc("/ws", s.handler)
		err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
		if err != nil {
			s.listenErrs(fmt.Errorf("Listen and serve error: %s", err))
		}
	}()
}

// get local ip address
func getIPAddress() (string, error) {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", nil
}

// get random port
func getRandomPort() int {
	const MAXPORT = 65535
	const MINPORT = 5001

	freePort := false
	var rnd int

	for !freePort {
		rand.Seed(time.Now().Unix())
		rnd = rand.Intn(MAXPORT-MINPORT) + MINPORT
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", rnd))
		if err == nil {
			freePort = true
			fmt.Printf("Listening on port %d", rnd)
		}
		ln.Close()
	}
	return rnd
}
