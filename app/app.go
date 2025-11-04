package app

import (
	"fmt"
	"net"
	"sync"
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

    conn, err := dialScan(connectionType, formatedAddres, timeout)
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
        0:  {Port: 20, Type: "tcp", Service: "ftp-data"},
        1:  {Port: 21, Type: "tcp", Service: "ftp"},
        2:  {Port: 22, Type: "tcp", Service: "ssh"},
        3:  {Port: 23, Type: "tcp", Service: "telnet"},
        4:  {Port: 25, Type: "tcp", Service: "smtp"},
        5:  {Port: 53, Type: "tcp", Service: "domain"},
        6:  {Port: 67, Type: "udp", Service: "dhcp-server"}, // incluí udp comum
        7:  {Port: 68, Type: "udp", Service: "dhcp-client"},
        8:  {Port: 69, Type: "udp", Service: "tftp"},
        9:  {Port: 80, Type: "tcp", Service: "http"},
        10: {Port: 110, Type: "tcp", Service: "pop3"},
        11: {Port: 111, Type: "tcp", Service: "rpcbind"},
        12: {Port: 123, Type: "udp", Service: "ntp"},
        13: {Port: 135, Type: "tcp", Service: "msrpc"},
        14: {Port: 137, Type: "udp", Service: "netbios-ns"},
        15: {Port: 138, Type: "udp", Service: "netbios-dgm"},
        16: {Port: 139, Type: "tcp", Service: "netbios-ssn"},
        17: {Port: 143, Type: "tcp", Service: "imap"},
        18: {Port: 161, Type: "udp", Service: "snmp"},
        19: {Port: 389, Type: "tcp", Service: "ldap"},
        20: {Port: 443, Type: "tcp", Service: "https"},
        21: {Port: 445, Type: "tcp", Service: "microsoft-ds"},
        22: {Port: 465, Type: "tcp", Service: "smtps"},
        23: {Port: 587, Type: "tcp", Service: "submission"},
        24: {Port: 636, Type: "tcp", Service: "ldaps"},
        25: {Port: 993, Type: "tcp", Service: "imaps"},
        26: {Port: 995, Type: "tcp", Service: "pop3s"},
        27: {Port: 1025, Type: "tcp", Service: "NFS-or-IIS"},
        28: {Port: 1080, Type: "tcp", Service: "socks"},
        29: {Port: 1433, Type: "tcp", Service: "ms-sql-s"},
        30: {Port: 1521, Type: "tcp", Service: "oracle"},
        31: {Port: 1723, Type: "tcp", Service: "pptp"},
        32: {Port: 2049, Type: "tcp", Service: "nfs"},
        33: {Port: 3306, Type: "tcp", Service: "mysql"},
        34: {Port: 3389, Type: "tcp", Service: "ms-wbt-server"},
        35: {Port: 5432, Type: "tcp", Service: "postgresql"},
        36: {Port: 5900, Type: "tcp", Service: "vnc"},
        37: {Port: 6001, Type: "tcp", Service: "X11:1"},
        38: {Port: 6379, Type: "tcp", Service: "redis"},
        39: {Port: 8080, Type: "tcp", Service: "http-proxy"},
        40: {Port: 8443, Type: "tcp", Service: "https-alt"},
        41: {Port: 9001, Type: "tcp", Service: "tor-orport"},
        // Pode adicionar mais conforme necessidade
    }

	var timeout time.Duration= 1 * time.Second
    address := c.String("host")
    connectionType := c.String("type")
	
	fmt.Printf("\nNetwork Address: %s\nconnection type: %s\n",address, connectionType)

	if connectionType == "tcp"{
		for i := 0; i <= len(ports); i++ {
			if ports[i].Type == "udp"{
				delete(ports,i)
			}
			fmt.Println(ports[i].Port,ports[i].Type)
		}
	} else {
		for i := 0; i <= len(ports); i++ {
			if ports[i].Type == "tcp"{
				delete(ports,i)
			}
			fmt.Println(ports[i].Port,ports[i].Type)
		}
	}

	
	var wg sync.WaitGroup
	for i:=0; i < len(ports); i++{
		wg.Add(1)
		go func(PortInfo PortInfo){
			defer wg.Done()
			formatedAddress := fmt.Sprintf("%s:%d", address, ports[i].Port)
			conn,err := dialScan(connectionType, formatedAddress, timeout)
			if err != nil {
				fmt.Println("Port ",ports[i].Port, "Is closed", err)
				return
			} else {
				fmt.Println("Port ",ports[i].Port, "Is open", err)
			}
			conn.Close()
		}(ports[i])
	}
	wg.Wait() 
}


func dialScan(typeConn string, address string, duration time.Duration) (net.Conn, error) {

	// var waitGroup sync.WaitGroup

	// waitGroup.Add(2)
	
	conn, err := net.DialTimeout(typeConn,address,duration)
		if err != nil {
        return nil, err
    }
    return conn, err
}

