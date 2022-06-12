package main

import (
	"flag"
	"imgConverter/pkg/app"
)

const ConfigPath = "configs"

var mode = flag.String("mode", "cli",
	"Modes are:\n"+
		"'web' - Run web convertor\n"+
		"'cli' - Run cli convertor\n")

func main() {
	flag.Parse()

	if *mode == "web" {
		app.RunWeb(ConfigPath)
	} else {
		app.RunCli()
	}
}
