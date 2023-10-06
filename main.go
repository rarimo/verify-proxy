package main

import (
	"os"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
