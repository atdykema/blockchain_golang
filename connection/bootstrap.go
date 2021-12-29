package conn

import (
	"fmt"
	"net"
	/* "os" */
	)


func BootstrapConnection() string {
	fmt.Println("entered func")
	status := "connect"
	return status
}

func InitListen() error{

	listen, err := net.Listen("tcp", ":10001")

	if err != nil{

		fmt.Println(err)
		return err
	}

	for {

		connection, err := listen.Accept()
		if err != nil{

			fmt.Println(err)
			return err

		}

		//handle connection
	}

}