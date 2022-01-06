package main

import (
	"fmt"
	"net"
	"time"
	"os"
	//"github.com/atdykema/blockchain_golang/connection"
)

func main(){

	bootstrapIP := net.IPv4(108, 26, 172, 97)
	
	peers := []net.IP{bootstrapIP}

	//go conn.StartServer("localhost", "8080")

	for i := 0; i < len(peers); i++{


		address := peers[i].String() + ":8080"

		tcpAddr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			fmt.Println("resolve fail: ", err.Error())
			break
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil{
			fmt.Println("dial fail: ", err.Error())
			break
		}

		listOfPeers, err := conn.Write([]byte("listOfPeers"))
		if err != nil{
			fmt.Println("peer write to node fail: ", err.Error())
			conn.Close()
			break
		}

		fmt.Println(listOfPeers)
		//TODO: add peers to node list of peers

		returnMessage := make([]byte, 1024)

		_, err = conn.Read(returnMessage)
		if err != nil{
			fmt.Println("return message send fail: ", err.Error())
		}


		conn.Close()

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