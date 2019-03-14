package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/mikepjb/clji/src/bencode"
)

func main() {
	port := "9999"

	fileb, err := ioutil.ReadFile(".nrepl-port")

	if err != nil {
		fmt.Println(".nrepl-port not found")
	} else {
		fmt.Printf("setting port to %v\n", string(fileb))
		port = string(fileb)
	}

	if len(os.Args) != 1 {
		port = os.Args[1]
	}

	clone := map[string]string{"op": "clone"}
	emsg := bencode.Encode(clone)
	fmt.Println(emsg)

	conn, _ := net.Dial("tcp", "127.0.0.1:"+port)
	fmt.Fprintf(conn, emsg+"\n")

	r := bufio.NewReader(conn)
	var b []byte = make([]byte, 1)

	for {
		r.Read(b)
		fmt.Printf("%v", string(b))
	}
}
