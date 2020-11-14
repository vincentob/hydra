package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dantin-s/hydra/utils/whois"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	FlagDomainFile = "file"
)

var (
	whoisChan chan string = make(chan string, 100000)
)

func Whois(c *cli.Context) error {
	file, err := os.Open(c.String(FlagDomainFile))
	if err != nil {
		return err
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		whoisChan <- scanner.Text()
		wg.Add(1)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		go func() {
			for true {
				domain := <-whoisChan
				if result, err := whois.GetWhoisTimeout(domain, time.Second*10); err != nil {
					logrus.Error(domain, "    ", err)
				} else {
					// get registrar substring
					firstIndex := hasRegistrar(result)
					if firstIndex != -1 {
						subString := result[firstIndex:]

						// for common domain suffix, get first line
						begin := 0
						end := strings.Index(subString, "\n")

						// for .uk suffix, get second line
						if strings.HasSuffix(domain, ".uk") || strings.HasSuffix(domain, ".be") {
							begin = end + 1
							end = strings.Index(result[begin:], "\n") + begin + 1
						}

						if end != -1 {
							registrar := subString[begin:end]
							fmt.Println(domain, "    ", registrar)
							//logrus.Info(domain, "    ", registrar)
						} else {
							logrus.Error(domain, "    registrar not found 1")
						}
					} else {
						logrus.Error(domain, "    registrar not found 2")
					}
				}
				time.Sleep(100 * time.Millisecond)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	return nil
}

func hasRegistrar(str string) int {
	if index := strings.Index(str, "Registrar:"); index == -1 {
		return strings.Index(str, "registrar:")
	} else {
		return index
	}
}
