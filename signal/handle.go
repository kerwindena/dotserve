/*
The signal package traps the signals TERM, INT and HUP and calls a callback
before exiting.
*/
package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleExit catches SIGTERM, SIGINT and SIGHUP and calls the callback before
// exiting  the program with the provided return value using os.Exit and thus
// ignoring all defers.
func HandleExit(callback func(), retval int) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP)
	go func() {
		<-c
		callback()
		os.Exit(retval)
	}()
}
