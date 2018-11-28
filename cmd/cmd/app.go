package cmd

import (
	"flag"
	"fmt"
	"os"
)

var testCmd string

func init() {
	flag.StringVar(&testCmd, "t", "222", "test cmd")
}

// Run 启动命令
func Run() {
	flag.Parse()
	if testCmd != "" {
		fmt.Println(testCmd)
		return
	}

	usage()
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage: -t

	options:
	`)
	flag.PrintDefaults()
}
