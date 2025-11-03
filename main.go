package main

import (
	"log"
	"os"
	"port-scanner/app"
)

func main() {
	aplication := app.CliGen()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}