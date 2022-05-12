package config

import (
	"path/filepath"
	"bufio"
	"strings"
	"strconv"
	"net"
	"os"
	"fmt"

)

func GetBootstrapIPs() []net.IP{

	wd, err:= os.Getwd()
	if err != nil{
		fmt.Println("error getting wd: ", err.Error())
	}

	par := filepath.Dir(wd)
	fmt.Println(par)//test

	bootstrapIPs := []net.IP{}

	bootstrapIPsf, err := os.Open(filepath.Join(par, "bootstrapIPs.txt"))
    if err != nil {
        fmt.Printf("Unable to open file: %v\n", err)
		os.Exit(1)
    }

	defer bootstrapIPsf.Close()

	bscanner := bufio.NewScanner(bootstrapIPsf)

	for bscanner.Scan(){
		ipstr := strings.Split(bscanner.Text(),".")
		ipint := []int{}
		for i := 0; i < len(ipstr); i++{
			temp, err := strconv.Atoi(ipstr[i])
			if err != nil{
				fmt.Println("ip str to int convers error: ", err.Error())
				os.Exit(1)
			}
			ipint = append(ipint, temp)
		}
		
		bootstrapIPs = append(bootstrapIPs, net.IPv4(byte(ipint[0]), byte(ipint[1]), byte(ipint[2]), byte(ipint[3])))
	}

	fmt.Println(bootstrapIPs) //test

	return bootstrapIPs
}