package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

//Simple connection to whois servers with default timeout 5 sec
func GetWhois(domain string) (string, error) {
	return GetWhoisTimeout(domain, time.Second*5)
}

//Connection to whois servers with various time.Duration
func GetWhoisTimeout(domain string, timeout time.Duration) (result string, err error) {
	var (
		parts      []string
		zone       string
		buffer     []byte
		connection net.Conn
	)

	parts = strings.Split(domain, ".")
	if len(parts) < 2 {
		err = fmt.Errorf("domain(%s) name is wrong", domain)
		return
	}
	//last part of domain is zome
	zone = parts[len(parts)-1]

	server, ok := servers[zone]

	if !ok {
		//err = fmt.Errorf("no such server for zone %s. Domain %s", zone, domain)
		server = servers["com"]
	}

	connection, err = net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)

	if err != nil {
		return
	}

	defer connection.Close()

	connection.Write([]byte(domain + "\r\n"))

	buffer, err = ioutil.ReadAll(connection)

	if err != nil {
		return
	}

	result = string(buffer[:])

	return
}
