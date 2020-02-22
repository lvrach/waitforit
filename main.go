package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func waitFor(ctx context.Context, endpoint string) {

	wait := 50 * time.Millisecond

	for {
		conn, err := net.DialTimeout("tcp", endpoint, wait)
		if err == nil {
			conn.Close()
			log.Printf("%s ready!", endpoint)

			return
		}
		if ctx.Err() != nil {
			log.Printf("timeout waiting for %s ...", endpoint)
			return
		}
		if wait > time.Second {
			log.Printf("waiting for %s ...", endpoint)
		}

		select {
		case <-time.After(wait):
			wait = wait * 2
		case <-ctx.Done():
		}
	}
}

func main() {
	env := os.Getenv("WAIT_FOR_IT")
	if env == "" {
		os.Exit(0)
	}

	timeout := 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg := sync.WaitGroup{}
	endpoints := strings.Fields(env)
	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string) {
			waitFor(ctx, endpoint)
			wg.Done()
		}(endpoint)
	}
	wg.Wait()

	if len(os.Args) <= 1 {
		os.Exit(0)
	}

	cmd := exec.Command(os.Args[1], ...os.Args[1:])
	err := cmd.Run();

	type exitCode interface {
		
	}
}
