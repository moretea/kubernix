package main

import (
	"flag"
	"fmt"
	"net"

	systemdDaemon "github.com/coreos/go-systemd/daemon"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	cri "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"

	"github.com/moretea/kubernix/server"
)

func main() {
	const (
		defaultSocketPath = "/var/run/kubernix.sock"
	)
	var enableDebug bool
	var socketPath string
	flag.StringVar(&socketPath, "socket", defaultSocketPath, fmt.Sprintf("path of the kubernix CRI socket. Defaults to %s", defaultSocketPath))
	flag.BoolVar(&enableDebug, "debug", false, "print out debug messages")

	flag.Parse()

	if enableDebug {
		log.SetLevel(log.DebugLevel)
	}

	grpcServer := grpc.NewServer()

	kubernix_service, err := server.New()

	if err != nil {
		log.Fatalf("Could not create kubernix server")
	}

	cri.RegisterRuntimeServiceServer(grpcServer, kubernix_service)
	//  cri.RegisterImageServiceServer(grpcServer, kubernix_service)

	unix_socket, err := net.Listen("unix", socketPath)
	log.Infof("Listening on %s", socketPath)

	if err != nil {
		log.Fatalf("Could not bind kubernix socket '%s', because %s\n", socketPath, err)
	}

	// Notify systemd that the daemon is ready
	go func() {
		systemdDaemon.SdNotify(true, "READY=1")
	}()

	log.Infof("Starting kubernix")
	err = grpcServer.Serve(unix_socket)
	log.Fatalf("Server stopped %v", err)
}
