package main

import (
	"os"

	"github.com/named-data/ndnd/dv/executor"
)

func main() {
	args := os.Args
	args[0] = "ndn-dv"
	executor.Main(args)
}
