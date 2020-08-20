package main

import (
	"net"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout,"Usage: %s hostname\n",os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]

	res := make(chan []byte,2)
	go GetIpAddr(address,res)
	go serve(address,res)

	ipAddresses := <-res
	result := <-res //receive from res


	fmt.Printf("The IP Address of %s is:\n",os.Args[1])
/*	for _,ipAddress := range ipAddresses{
		fmt.Printf("%v",string(ipAddress))
		fmt.Println("\n\n")
	}
*/
	fmt.Println(string(ipAddresses))
	fmt.Println("Server Response:")
	fmt.Println("\n")
	fmt.Printf("%v ",string(result))
}

//Get server info
func serve(addr string,c chan []byte) {

	tcpAddr,err := net.ResolveTCPAddr("tcp4",addr+":80")
	Check(err)


	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	Check(err)


	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	Check(err)


	res,err := ioutil.ReadAll(conn)
	Check(err)

	c<-res //send to c channel
}

//Get ipaddress
func GetIpAddr(hostname string,c chan []byte){
	ipAddr,err := net.LookupHost(hostname)
	Check(err)

	var ip = []byte{}

	for i,_ := range ipAddr{
		b := []byte(ipAddr[i])
		for j,_ := range b{
			ip = append(ip,b[j])
		}
	}

	c <- ip //send ip to c channel

}

func Check(e error) {
	if e != nil{
		fmt.Println(e)
	}
}
