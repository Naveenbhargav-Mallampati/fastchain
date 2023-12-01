package main

import (
	"os"

	"fastchain.com/corechain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()

	// w := wallet.MakeWallet()
	// w.Address()
}
