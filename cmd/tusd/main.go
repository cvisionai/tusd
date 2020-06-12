package main

import (
	"github.com/tus/tusd/cmd/tusd/cli"
  "os/signal"
  "syscall"
)

func main() {
	signal.Ignore(syscall.SIGTERM)
	cli.ParseFlags()
	cli.PrepareGreeting()

	// Print version and other information and exit if the -version flag has been
	// passed else we will start the HTTP server
	if cli.Flags.ShowVersion {
		cli.ShowVersion()
	} else {
		cli.CreateComposer()
		cli.Serve()
	}
}
