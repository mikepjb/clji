package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mikepjb/clji/src/bencode"
)

func main() {
	port := "9999"

	if len(os.Args) != 1 {
		port = os.Args[1]
	}

	clone := map[string]string{"op": "clone"}
	emsg := bencode.Encode(clone)
	fmt.Println(emsg)

	conn, _ := net.Dial("tcp", "127.0.0.1:"+port)
	fmt.Fprintf(conn, emsg+"\n")

	// b, _ := bufio.NewReader(conn).ReadLine()
	// fmt.Print("response: " + string(b))
}
