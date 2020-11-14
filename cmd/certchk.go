package cmd

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli"
)

var (
	dialer = &net.Dialer{Timeout: 5 * time.Second}
)

func check(server string, width int) {
	conn, err := tls.DialWithDialer(dialer, "tcp", server+":443", nil)
	if err != nil {
		fmt.Printf("%*s | %v\n", width, server, err)
		return
	}
	defer conn.Close()
	valid := conn.VerifyHostname(server)

	for _, c := range conn.ConnectionState().PeerCertificates {
		if valid == nil {
			fmt.Printf("%*s | valid, expires on %s (%s)\n", width, server, c.NotAfter.Format("2006-01-02"), humanize.Time(c.NotAfter))
		} else {
			fmt.Printf("%*s | %v\n", width, server, valid)
		}
		return
	}
}

// CertChk
// Ref: https://www.opsdash.com/blog/check-ssl-certificate.html
func CertChk(c *cli.Context) error {
	domains, err := getDomains(c)
	if err != nil {
		return err
	}

	// for cosmetics
	width := 0
	for _, domain := range domains {
		if len(domain) > width {
			width = len(domain)
		}
	}

	fmt.Printf("%*s | Certificate status\n%s-+-%s\n", width, "Server",
		strings.Repeat("-", width), strings.Repeat("-", 80-width-2))

	for _, domain := range domains {
		check(domain, width)
	}

	return nil
}

func getDomains(c *cli.Context) (domains []string, err error) {
	// read names from the file
	if len(c.String(FlagDomainFile)) > 0 {
		f, err := os.Open(c.String(FlagDomainFile))
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) > 0 && line[0] != '#' {
				domains = append(domains, strings.Fields(line)[0])
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}

		_ = f.Close()
	}

	for _, i := range c.Args() {
		domains = append(domains, i)
	}

	return
}
