package main

import (
	"fmt"
	"time"
)

func main(){

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

	case "hello":
		fmt.Println("hi there")
		return 0

	case "exit":
		fmt.Println("Exiting...")
		return -1
		
	default:
		fmt.Println("Unrecognized command")
		return 0
	}


}