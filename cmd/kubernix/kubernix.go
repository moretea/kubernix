package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	systemdDaemon "github.com/coreos/go-systemd/daemon"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"

	"github.com/moretea/kubernix/server"
)

func main() {
	const (
		DEFAULT_SOCKET_PATH = "/var/run/kubernix.sock"
	)
	var defaultOrEnvPath string
	var enableDebug bool
	var socketPath string

	defaultOrEnvPath, ok := os.LookupEnv("CRI_RUNTIME_ENDPOINT")
	if !ok {
		defaultOrEnvPath = DEFAULT_SOCKET_PATH
	}

	flag.StringVar(&socketPath, "socket", defaultOrEnvPath, fmt.Sprintf("path of the kubernix CRI socket. Defaults to %s, can also be set with the CRI_RUNTIME_ENDPOINT env var", DEFAULT_SOCKET_PATH))
	flag.BoolVar(&enableDebug, "debug", false, "print out debug messages")

	flag.Parse()

	if enableDebug {
		log.SetLevel(log.DebugLevel)
	}

	grpcServer := grpc.NewServer()

	kubernix_server, err := server.New()

	if err != nil {
		log.Fatalf("Could not create kubernix server")
	}

	cri.RegisterRuntimeServiceServer(grpcServer, kubernix_server)
	cri.RegisterImageServiceServer(grpcServer, kubernix_server)

	log.Infof("Listening on %s", socketPath)
	unix_socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Could not bind kubernix socket '%s', because %s\n", socketPath, err)
	} else {
		defer func() {
			unix_socket.Close()
		}()
	}

	// Start server
	go func() {
		log.Infof("Starting kubernix")
		err = grpcServer.Serve(unix_socket)
		log.Fatalf("Server stopped %v", err)
	}()

	// Notify systemd that the daemon is ready
	go func() {
		systemdDaemon.SdNotify(true, "READY=1")
	}()

	// Handle interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c
	fmt.Println("Got signal", s)
}
