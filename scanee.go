package main

import (
	"net"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout,"Usage: %s hostname",os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]

	res := make(chan []byte)
	go serve(address,res)

	result := <-res //receive from res

	fmt.Printf("%v ",string(result))
}

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

func Check(e error) {
	if e != nil{
		fmt.Println(e)
	}
}
