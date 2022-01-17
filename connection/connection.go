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

func PeerConnect(ip net.IP) *net.TCPConn{

	address := ip.String() + ":8080"

		tcpAddr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			fmt.Println("resolve fail: ", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil{
			fmt.Println("dial fail: ", err.Error())
			os.Exit(1)
		}

		return conn
}

func WriteToPeer(conn *net.TCPConn, msg string, id string){

	_, err := conn.Write([]byte(id))
		if err != nil{
			fmt.Println("peer write id to node fail: ", err.Error())
			conn.Close()
		}
	
	

	_, err = conn.Write([]byte(msg))
		if err != nil{
			fmt.Println("peer write msg to node fail: ", err.Error())
			conn.Close()
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

