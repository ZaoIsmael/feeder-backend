package server

import (
	"bufio"
	"deporvillage-feeder-backend/cmd/feeder-service/src/controller"
	"deporvillage-feeder-backend/cmd/feeder-service/src/util"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Server struct {
	listener              net.Listener
	handler               controller.Controller
	totalConnectedClients int
}

const limitConnectedClients = 5

func CreateServer(h controller.Controller) (Server, error) {
	l, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4000,
	})

	if err != nil {
		return Server{}, err
	}

	rand.Seed(time.Now().Unix())
	err = l.SetDeadline(time.Now().Add(time.Minute * 1))

	if err != nil {
		return Server{}, err
	}

	return Server{
		l, h, 0,
	}, nil
}

func handleConnection(c net.Conn, s *Server) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	defer func(c net.Conn, s *Server) {
		s.totalConnectedClients--
		err := c.Close()

		if err != nil {
			log.Println(err)
		}
	}(c, s)

	for {
		input, err := s.parseInput(c)

		if err != nil {
			log.Println(err)
			return
		}

		if input == "terminate" {
			s.terminate(c)
			break
		} else {
			s.handler.Run(input)
		}
	}
}

func (s *Server) Run() {
	fmt.Printf("Server run on %s\n", s.listener.Addr().String())
	s.listenSignals()

	for {
		c, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		if s.totalConnectedClients < limitConnectedClients {
			s.totalConnectedClients++
			go handleConnection(c, s)
		} else {
			err := c.Close()

			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (s *Server) Shutdown() {
	err := s.listener.Close()

	if err != nil {
		log.Println(err)
	}
}

func (s Server) listenSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func(s Server) {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			s.Shutdown()
		case syscall.SIGTERM:
			s.Shutdown()
		}
	}(s)
}

func (s Server) parseInput(c net.Conn) (string, error) {
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(netData), nil
}

func (s Server) terminate(c net.Conn) {
	err := c.Close()

	if err != nil {
		log.Println(err)
	}

	util.KillSystem()
}
