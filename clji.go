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

	conn, err := net.Dial("tcp", "127.0.0.1:"+port)

	if err != nil {
		fmt.Errorf("could not connect: %v\n", err)
	}

	fmt.Fprintf(conn, emsg+"\n")

	r := bufio.NewReader(conn)
	var b []byte = make([]byte, 1)

	response := ""

	for {
		r.Read(b)
		response += string(b)
		fmt.Println(response)
		msg, ok := bencode.Decode(response)

		if ok {
			fmt.Println(msg)
			break
		}
	}
}
