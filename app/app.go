package app

import (
	"github.com/urfave/cli"
)


type PortInfo struct {
	Port 	int
    Type    string
    Service string
}

func CliGen() *cli.App {
	app := cli.NewApp()
	app.Name = "Port Scanner"
	app.Usage = "scans for open and closed ports on the connection."

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "marcostech.com.br",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "443",
			Usage: "Port ranges (ex: 80 ou 20-100)",
    	},
		cli.StringFlag{
			Name:  "type",
			Value: "tcp",
			Usage: "tcp or udp",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "scanport",
			Usage:  "Scan ports with Address",
			Flags:  flags,
			Action: scanPort,
		},
		{
			Name:   "scanports",
			Usage:  "Scan all ports with Address",
			Flags:  flags,
			Action: scanAllPorts,
		},
	}
	return app
}
