package app

import (
	"fmt"
	"net"
	"time"

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

func scanPort(c *cli.Context){
	timeout := 1 * time.Second
    address := c.String("host")
    port := c.String("port")
    connectionType := c.String("type")

	if connectionType != "tcp"{
		connectionType = "udp"
	}
	
	formatedAddres := address + ":"+port

    conn, err := net.DialTimeout(connectionType, formatedAddres, timeout)
    if err != nil {
		fmt.Println("=====Porta is Closed=====\n", err)
		fmt.Printf("\nNetwork Address: %s\nPort:%s\nconnection type: %s\n",address, port, connectionType)
        return
    }
    defer conn.Close()
    fmt.Println("=====Port is Open=====\n", err)
	fmt.Printf("\nNetwork Address: %s\nPort:%s\nconnection type: %s\n",address, port, connectionType)
}

func scanAllPorts (c *cli.Context){

	ports := map[int]PortInfo{
			0:  {Port: 21, Type: "TCP", Service: "ftp"},
			1:  {Port: 22, Type: "TCP", Service: "ssh"},
			2:  {Port: 23, Type: "TCP", Service: "telnet"},
			3:  {Port: 25, Type: "TCP", Service: "smtp"},
			4:  {Port: 53, Type: "TCP", Service: "domain"},
			5:  {Port: 80, Type: "TCP", Service: "http"},
			6:  {Port: 110, Type: "TCP", Service: "pop3"},
			7:  {Port: 111, Type: "TCP", Service: "rpcbind"},
			8:  {Port: 135, Type: "TCP", Service: "msrpc"},
			9:  {Port: 139, Type: "TCP", Service: "netbios-ssn"},
			10: {Port: 143, Type: "TCP", Service: "imap"},
			11: {Port: 389, Type: "TCP", Service: "ldap"},
			12: {Port: 443, Type: "TCP", Service: "https"},
			13: {Port: 445, Type: "TCP", Service: "microsoft-ds"},
			14: {Port: 587, Type: "TCP", Service: "submission"},
			15: {Port: 1025, Type: "TCP", Service: "NFS-or-IIS"},
			16: {Port: 1080, Type: "TCP", Service: "socks"},
			17: {Port: 1433, Type: "TCP", Service: "ms-sql-s"},
			18: {Port: 2049, Type: "TCP", Service: "nfs"},
			19: {Port: 3306, Type: "TCP", Service: "mysql"},
			20: {Port: 3389, Type: "TCP", Service: "ms-wbt-server"},
			21: {Port: 5900, Type: "TCP", Service: "vnc"},
			22: {Port: 6001, Type: "TCP", Service: "X11:1"},
			23: {Port: 6379, Type: "TCP", Service: "redis"},
			24: {Port: 8080, Type: "TCP", Service: "http-proxy"},
			25: {Port: 9001, Type: "TCP", Service: "tor-orport"},
		}

	
	timeout := 1 * time.Second
    address := c.String("host")
    connectionType := c.String("type")

	fmt.Printf("\nNetwork Address: %s\nconnection type: %s\n",address, connectionType)
	for i:=0; i < len(ports); i++{
		formatedAddres := fmt.Sprintf("%s:%d", address, ports[i].Port)
		conn, err := net.DialTimeout(connectionType, formatedAddres, timeout)
		
		if err != nil {
			fmt.Println("Port ",ports[i].Port, "Is closed", err)
			continue
		} else {
			fmt.Println("Port ",ports[i].Port, "Is open", err)
		}
		conn.Close()
	}
}
