package main

import (
	"os"

	"github.com/rarimo/verify-proxy/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
