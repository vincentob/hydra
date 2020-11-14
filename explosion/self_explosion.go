package explosion

import (
	"time"

	"github.com/sirupsen/logrus"
)

func CountDownAndExplosion(s int) {
	logrus.Infof("Server will exit in %v seconds.", s)
	for i := s; i >= 0; i-- {
		logrus.Info(i)
		time.Sleep(time.Second * 1)
	}
	logrus.Info("Bye Bye :)")
}
