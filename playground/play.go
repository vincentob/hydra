package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/vincentob/hydra/utils/whois"
)

func main() {
	// ascii - rune
	////asciiValue := int('A')
	//logrus.Infof("Ascii Value of %c = %d", 'A', asciiValue)

	if result, err := whois.GetWhoisTimeout("1001store.online", 5*time.Second); err != nil {
		logrus.Info(err)
	} else {
		logrus.Info(result)
	}

}
