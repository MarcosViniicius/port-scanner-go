package app

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/urfave/cli"
)

// Scan individual ports and ranges (e.g., 80, 10-20, 22,80,443)
func scanPort(c *cli.Context) {
    timeout := 2 * time.Second
    address := c.String("host")
    port := c.String("port")
    connectionType := c.String("type")
    
    // Verify tcp or udp
    if connectionType != "tcp" {
        connectionType = "udp"
    }
    
    ports := parsePortsFlexible(port)  // Single, range, multiple
    
    var wg sync.WaitGroup
    results := make([]ScanResult, len(ports))
    mu := sync.Mutex{}
    
    // Goroutines 
    for i, p := range ports {
        wg.Add(1)
        go func(idx int, portNum int) {
            defer wg.Done()
            formatedAddress := fmt.Sprintf("%s:%d", address, portNum)
            conn, err := dialScan(connectionType, formatedAddress, timeout)
            if conn != nil {
                conn.Close()
            }
            
            mu.Lock()
            results[idx] = ScanResult{Port: portNum, Open: err == nil}
            mu.Unlock()
        }(i, p)
    }
    wg.Wait()
    
    // Ordened output
	fmt.Printf("\nNetwork Address: %s\nPort:%s\nconnection type: %s\n====================================================\n",address, port, connectionType)

    for _, r := range results {
        if r.Open {
            fmt.Printf("Port %d Is open\n", r.Port)
        } else {
            fmt.Printf("Port %d Is closed\n", r.Port)
        }
    }
}


func scanAllPorts (c *cli.Context){

	var timeout time.Duration= 1 * time.Second
    address := c.String("host")
    connectionType := c.String("type")
	
	fmt.Printf("\nNetwork Address: %s\nconnection type: %s\n==================================================\n",address, connectionType)

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
			continue 
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


// Helper: parse single, range "-", multiple ","
func parsePortsFlexible(input string) []int {
    var ports []int
    parts := strings.Split(input, ",")
    
    for _, part := range parts {
        part = strings.TrimSpace(part)
        if strings.Contains(part, "-") {
            bounds := strings.Split(part, "-")
            if len(bounds) == 2 {
                start, _ := strconv.Atoi(bounds[0])
                end, _ := strconv.Atoi(bounds[1])
                for p := start; p <= end; p++ {
                    ports = append(ports, p)
                }
            }
        } else if p, err := strconv.Atoi(part); err == nil {
            ports = append(ports, p)
        }
    }
    sort.Ints(ports)
    return ports
}

type ScanResult struct {
    Port int
    Open bool
}