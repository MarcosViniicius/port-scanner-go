package main

import (
	"log"
	"os"

	"github.com/MarcosViniicius/port-scanner-go/app"
)

func main() {
	aplication := app.CliGen()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}