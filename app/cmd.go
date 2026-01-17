package app

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/urfave/cli"
)

// Scan individual ports and ranges (e.g., 80, 10-20, 22,80,443)
func scanPort(c *cli.Context){
	timeout := 1 * time.Second
    address := c.String("host")
    port := c.String("port")
    connectionType := c.String("type")

	// Verify tcp or udp port
	if connectionType != "tcp"{
		connectionType = "udp"
	}

	// Check if entry contains "-" to detect port ranges vs single ports
	if strings.Contains(port, "-") {
		var ports []int 

		bounds := strings.Split(port, "-")
		if len(bounds) == 2 {
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			for p := start; p <= end; p++ {
				ports = append(ports, p)
			}
		}

		var wg sync.WaitGroup
			
			for i:=0; i < len(ports); i++{

				if ports[i] == 0 {
					continue 
				}
				
				wg.Add(1)
				go func(PortInfo PortInfo){
					defer wg.Done()
					formatedAddress := fmt.Sprintf("%s:%d", address, ports[i])
					conn,err := dialScan(connectionType, formatedAddress, timeout)
					if err != nil {
						fmt.Println("Port ",ports[i], "Is closed", ) // add var 'err' for debug
						return
					}
					fmt.Println("Port ",ports[i], "Is open", ) // add var 'err' for debug
					conn.Close()
				}(PortInfo{})
			}
			wg.Wait() 
	}

	// If no "-", treat as single port

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

	var timeout time.Duration= 1 * time.Second
    address := c.String("host")
    connectionType := c.String("type")
	
	fmt.Printf("\nNetwork Address: %s\nconnection type: %s\n",address, connectionType)

	for key, port := range CommonPorts {
		if connectionType == "tcp" && port.Type == "udp" {
			delete(CommonPorts, key)
			continue
		}
		if connectionType == "udp" && port.Type == "tcp" {
			delete(CommonPorts, key)
			continue
		}
		// fmt.Println(port.Port, port.Type)
	}

	var wg sync.WaitGroup
	
	for i:=0; i < len(CommonPorts); i++{

		if CommonPorts[i].Port == 0 {
			continue // ignora portas 0
		}
		
		wg.Add(1)
		go func(PortInfo PortInfo){
			defer wg.Done()
			formatedAddress := fmt.Sprintf("%s:%d", address, CommonPorts[i].Port)
			conn,err := dialScan(connectionType, formatedAddress, timeout)
			if err != nil {
				fmt.Println("Port ",CommonPorts[i].Port, "Is closed", ) // add var 'err' for debug
				return
			}
			fmt.Println("Port ",CommonPorts[i].Port, "Is open", ) // add var 'err' for debug
			conn.Close()
		}(PortInfo{})
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

