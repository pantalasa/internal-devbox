// devbox is an internal CLI used by Pantalasa engineers to bootstrap local
// development environments. It is not customer-facing and is not deployed to
// production — it runs on laptops only.
package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "env":
		printEnv()
	case "version", "-v", "--version":
		fmt.Println(version)
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `devbox — internal dev-env helper for Pantalasa engineers.

Usage:
  devbox env                Print recommended local env vars
  devbox version            Print the tool version
  devbox help               Show this message`)
}

func printEnv() {
	fmt.Println("PANTALASA_ENV=local")
	fmt.Println("PANTALASA_LOG_LEVEL=debug")
	fmt.Println("PANTALASA_TRACE_SAMPLE=1.0")
}
