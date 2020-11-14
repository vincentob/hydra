package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

const (
	InputTypeBytes  = "bytes"
	InputTypeString = "string"

	FlagFrom = "from"
	FlagTo   = "to"
	FlagBase = "base"
)

// Convert can convert everything
// Online converter is: https://onlineutf8tools.com/convert-bytes-to-utf8
func Convert(cli *cli.Context) error {
	switch cli.String(FlagFrom) {
	case InputTypeBytes:
		return convertFromBytesToString(cli)
	default:
		return errors.New("only support bytes to string convert")
	}
}

func convertFromBytesToString(cli *cli.Context) error {
	if cli.NArg() != 1 {
		return errors.New("only support 1 parameter")
	}

	decode := ""
	encode := cli.Args().First()
	base := cli.Int("base")

	for i := 0; i <= len(encode)-2; i += 2 {
		ii, _ := strconv.ParseInt(encode[i:i+2], base, 8)
		decode += string(ii)
	}

	fmt.Println(decode)
	return nil
}
