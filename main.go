package main

import (
	"log"
	"os"

	"github.com/MarcosViniicius/port-scanner-go/app"
	"github.com/MarcosViniicius/port-scanner-go/scanner"
)

func main() {
	application := app.CliGen()

	for i := range application.Commands {
		switch application.Commands[i].Name {
		case "scanport":
			application.Commands[i].Action = scanner.ScanPort
		case "scanports":
			application.Commands[i].Action = scanner.ScanAllPorts
		}
	}

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
