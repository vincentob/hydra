package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vincentob/hydra/explosion"
)

var (
	SigTerminated chan os.Signal
	SigCompleted  chan struct{}
)

func init() {
	// register the given channel to receive notifications
	// of the specified signals.
	// SIGHUP : terminal connection loss
	// SIGINT : interrupt
	// SIGTERM: program shutdown
	SigTerminated = make(chan os.Signal, 1)
	SigCompleted = make(chan struct{})

	signal.Notify(SigTerminated, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
}

// WaitForExit block and wait for complete or terminated signal.
func WaitForExit() {
	select {
	case sig := <-SigTerminated:
		logrus.Infof("Receive %v signal.", sig)
		explosion.CountDownAndExplosion(3)
	case <-SigCompleted:
		logrus.Info("Completed.")
	}
}
