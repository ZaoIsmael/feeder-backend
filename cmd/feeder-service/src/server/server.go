package server

import (
	"bufio"
	"deporvillage-feeder-backend/cmd/feeder-service/src/util"
	"deporvillage-feeder-backend/internal/inventory/application"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type server struct {
	listener              net.Listener
	handler               application.AddProductApplicationService
	totalConnectedClients int
}

type Server interface {
	Run()
	Shutdown()
}

var limitConnectedClients = 5

func CreateServer(h application.AddProductApplicationService) (server, error) {
	l, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4000,
	})

	if err != nil {
		fmt.Println(err)
		return server{}, err
	}

	rand.Seed(time.Now().Unix())
	err = l.SetDeadline(time.Now().Add(time.Minute * 1))

	if err != nil {
		return server{}, err
	}

	return server{
		l, h, 0,
	}, nil
}

func handleConnection(c net.Conn, s *server) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	defer func(c net.Conn, s *server) {
		s.totalConnectedClients--
		err := c.Close()

		if err != nil {
			fmt.Println(err)
		}
	}(c, s)

	for {
		input, err := s.parseInput(c)

		if err != nil {
			return
		}

		if input == "terminate" {
			s.terminate(c)
			break
		}

		err = s.handler.Execute(input)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *server) Run() {
	s.listenSignals()

	for {
		c, err := s.listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		if s.totalConnectedClients < limitConnectedClients {
			s.totalConnectedClients++
			go handleConnection(c, s)
		} else {
			err := c.Close()

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (s *server) Shutdown() {
	err := s.listener.Close()

	if err != nil {
		fmt.Println(err)
	}
}

func (s server) listenSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func(s server) {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			s.Shutdown()
		case syscall.SIGTERM:
			s.Shutdown()
		}
	}(s)
}

func (s server) parseInput(c net.Conn) (string, error) {
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return strings.TrimSpace(netData), nil
}

func (s server) terminate(c net.Conn) {
	err := c.Close()

	if err != nil {
		fmt.Println(err)
	}

	util.KillSystem()
}
