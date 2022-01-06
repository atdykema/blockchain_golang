package conn

import (
	"fmt"
	"net"
	"os"

)

func StartServer(address string, port string){

	listener, err := net.Listen("tcp", (address + ":" + port))
	if err != nil{
		fmt.Println("error listening: ", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Currently listening on " + address + ":" + port  + "...")

	for{

		connection, err := listener.Accept()
		if err != nil{
			fmt.Println("error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(connection)

	}

}


func handleRequest(connection net.Conn){

	buf := make([]byte, 1024)

	_, err := connection.Read(buf)
	if err != nil {
		fmt.Println("error reading: ", err.Error(), " ")
		os.Exit(1)
	}

	//TODO: actions for requests

	connection.Write([]byte("returning message"))

	connection.Close()
}

func FindPeers() {

	
}

