package main

import (
	"app/infrastructure"
	"log"
	"net"
	"net/http/fcgi"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const SOCK = "./socket/server.sock"
	time.Local = time.FixedZone("JST", 9*60*60)

	db := infrastructure.NewDB()
	awsS3 := infrastructure.NewAwsS3()
	r := infrastructure.NewRouting(db, awsS3)
	r.SetMiddleware()

	if os.Getenv("GIN_MODE") == "release" {
		r.SetReleaseMode()
		production(r)
	} else {
		develop(r)
	}
}

func develop(r *infrastructure.Routing) {
	r.Run()
}

func production(r *infrastructure.Routing) {
	const SOCK = "./socket/server.sock"

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, syscall.SIGINT)

	l, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	if err := os.Chmod(SOCK, 0777); err != nil {
		log.Fatal("error:", err)
	}

	go func() {
		fcgi.Serve(l, r.Gin)
	}()

	<-sig

	if err := os.Remove(SOCK); err != nil {
		log.Fatal("socket file remove error:", err)
	}
}
