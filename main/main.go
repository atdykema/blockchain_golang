package main

import (
	"fmt" 
	/* "net" */
	/* "os" */
	"github.com/atdykema/blockchain_golang/connection"
	
)

func main(){
	status := connection.Bootstrap()
	fmt.Println(status)
}