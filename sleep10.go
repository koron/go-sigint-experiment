package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	nosigint := flag.Bool("nosigint", false, "")
	flag.Parse()
	sig := make(chan os.Signal, 1)
	go func() {
		for {
			s := <-sig
			if !*nosigint {
				fmt.Println("sleep10: terminated")
				os.Exit(1)
			}
			fmt.Printf("sleep10: ignored: %+v\n", s)
		}
	}()
	signal.Notify(sig, os.Interrupt)
	fmt.Println("sleep10: sleeping")
	time.Sleep(10 * time.Second)
	signal.Stop(sig)
	fmt.Println("sleep10: completed")
	close(sig)
}
