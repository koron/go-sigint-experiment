package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func setup(c *exec.Cmd) error {
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	return nil
}

func main() {
	nosigint := flag.Bool("nosigint", false, "")
	flag.Parse()
	args := make([]string, 0, 1)
	if *nosigint {
		args = append(args, "-nosigint")
	}
	c := exec.Command("./sleep10", args...)
	err := setup(c)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for {
			switch <-sig {
			case os.Interrupt:
				log.Printf("kick: terminated")
			}
		}
	}()
	defer func() {
		signal.Stop(sig)
		close(sig)
	}()

	c.Run()
	log.Printf("kick: completed")

}
