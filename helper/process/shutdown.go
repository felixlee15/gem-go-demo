package process

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// ShutdownCallback executes callback function on SIGTERM, SIGINT and SIGSTOP signals
func ShutdownCallback(fn func()) chan struct{} {
	done := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGQUIT)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error(err)
			}
		}()
		s := <-signals
		signal.Stop(signals)
		logrus.Infof("Received %v signal, shutting down gracefully", s)

		fn()
		logrus.Info("Shutdown callback completed")
		close(done)
	}()
	return done
}
