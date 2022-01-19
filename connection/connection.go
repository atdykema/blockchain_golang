package conn

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
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

func WriteToPeer(conn *net.TCPConn, msg string){
	
	_, err := conn.Write([]byte(msg))
		if err != nil{
			fmt.Println("peer write msg to node fail: ", err.Error())
			conn.Close()
		}
}

func handleRequest(connection net.Conn){

	for{

		unformattedData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("error reading: ", err.Error(), " ")
			os.Exit(1)
		}

		data := strings.TrimSpace(string(unformattedData))

		//TODO: actions for requests
		
		if data == "end"{
			break
		}else{
			if data == "REQ_PEERS"{

				wd, err:= os.Getwd()
				if err != nil{
					fmt.Println("error getting wd: ", err.Error())
				}

				parentDir := filepath.Dir(wd)
				fmt.Println(parentDir)//test

				peers, err := os.Open(filepath.Join(parentDir, "peers.txt"))
				if err != nil {
					fmt.Printf("Unable to open file: %v\n", err)
					return
				}

				defer peers.Close()

				pscanner := bufio.NewScanner(peers)

				for pscanner.Scan(){
					_, err := connection.Write([]byte(pscanner.Text()))
					if err != nil{
						fmt.Println("unable to write file:", err.Error())
						return
					}
				}
			}
		}
	}

	connection.Write([]byte("end"))

	connection.Close()
}

func FindPeers() {

	
}

