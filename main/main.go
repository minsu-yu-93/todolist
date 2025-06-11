package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd_line := flag.Bool("c", false, "CLI mode")
	server := flag.Bool("s", false, "Server mode")

	flag.Parse()

	if *cmd_line {
		fmt.Println("Start Command line mode")
		cmd_line_run()
	} else if *server {
		fmt.Println("Start Server mode")
	}
}
