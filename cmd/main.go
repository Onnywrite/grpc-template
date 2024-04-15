package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Onnywrite/grpc-template/gen"
	"github.com/Onnywrite/grpc-template/internal/grpc/tester"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func mustLoadTLSCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("/etc/ssl/server-cert.pem", "/etc/ssl/server-key.pem")
	if err != nil {
		panic(err)
	}
	return credentials.NewServerTLSFromCert(&cert)
}

func mustLoadClientTLSCredentials() credentials.TransportCredentials {
	ca, err := os.ReadFile("/etc/ssl/ca-cert.pem")
	if err != nil {
		panic(err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(ca) {
		panic("failed to add CA's certificate")
	}
	return credentials.NewClientTLSFromCert(pool, "")
}

func main() {
	s := grpc.NewServer(grpc.Creds(mustLoadTLSCredentials()))
	srv := &tester.GRPCServer{}
	gen.RegisterTesterServer(s, srv)

	l, err := net.Listen("tcp", ":5055")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	mainClient()

	shut := make(chan os.Signal, 1)
	signal.Notify(shut, syscall.SIGTERM, syscall.SIGINT)
	<-shut
	s.GracefulStop()
}

func mainClient() {
	conn, err := grpc.Dial("localhost:5055", grpc.WithTransportCredentials(mustLoadClientTLSCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := gen.NewTesterClient(conn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	add, err := c.Add(ctx, &gen.AddRequest{
		X: "I ",
		Y: "love you",
	})
	if err != nil {
		log.Print("error: ", err.Error())
	}

	fmt.Printf("addition result: %s", add.Result)

	resp, err := c.SayHello(ctx, &gen.Name{
		Name: "Victory",
	})
	if err != nil {
		log.Print("error: ", err.Error())
		return
	}

	fmt.Printf("response: %s", resp.Message)

	_, err = c.GetError(ctx, nil)
	if err != nil {
		log.Print("error: ", err.Error())
	}
}
