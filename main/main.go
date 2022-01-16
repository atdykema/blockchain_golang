package main

import (
	"fmt"
	"time"
	"github.com/atdykema/blockchain_golang/connection"
)

func main(){

	bootstrapIPs := conn.GetBootstrapIPs()

	go conn.StartServer("localhost", "8080")

	
	for i := 0; i < len(bootstrapIPs); i++{

		
		

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