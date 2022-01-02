package main

import (
	"fmt"
	"net"
	"time"
	"github.com/atdykema/blockchain_golang/connection"
)

func main(){

	bootstrapIP := net.IPv4(108, 26, 172, 97)
	
	peers := []net.IP{bootstrapIP}

	go conn.StartServer("localhost", "8000")

	for i := 0; i < len(peers); i++{

		address := peers[i].String()
		fmt.Println(address)

	}
	

	for {

		fmt.Print("[",time.Now(),"]: ")

		var currCmd string
		fmt.Scanln(&currCmd)

		if handleCmd(currCmd) == -1{
			break
		}

	}


}

func handleCmd(command string) int {

	fmt.Print("[",time.Now(),"]: ")

	switch command{

	case "find peers":
		//go conn.FindPeers()
		return 0

	case "exit":
		fmt.Println("Exiting...")
		return -1
		
	default:
		fmt.Println("Unrecognized command")
		return 0
	}


}