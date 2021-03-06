package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/mikepjb/clji/src/bencode"
)

func send(code string) string {
	port := "9999"

	fileb, err := ioutil.ReadFile(".nrepl-port")

	if err != nil {
		fileb, err = ioutil.ReadFile("~/.lein/repl-port")
		if err == nil {
			port = string(fileb)
		}
	} else {
		port = string(fileb)
	}

	clone := map[string]string{"op": "clone"}
	emsg := bencode.Encode(clone)

	conn, err := net.Dial("tcp", "127.0.0.1:"+port)

	if err != nil {
		fmt.Errorf("could not connect: %v\n", err)
	}

	fmt.Fprintf(conn, emsg)

	r := bufio.NewReader(conn)
	var b []byte = make([]byte, 1)

	response := ""
	value := ""

	for {
		r.Read(b)
		response += string(b)
		msg, ok := bencode.Decode(response)

		if ok {
			response := ""
			newSession := msg["new-session"].(string)

			defMsg := map[string]string{
				"session": msg["new-session"].(string),
				"op":      "eval",
				"ns":      "user",
				"code":    code,
			}

			emsg := bencode.Encode(defMsg)
			fmt.Fprintf(conn, emsg)

			for {
				r.Read(b)
				response += string(b)

				msg, ok = bencode.Decode(response)

				if ok {
					response = ""
					v, ok := msg["status"].([]string)

					if ok {
						if msg["session"].(string) == newSession &&
							v[0] == "done" {
							break
						}
					}

					val, ok := msg["value"].(string)

					if ok {
						value = val
					}

					errMsg, ok := msg["err"].(string)

					if ok {
						value = errMsg
					}
				}
			}

			break
		}
	}

	return value
}

func main() {
	input := "(+ 40 2)"

	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	fmt.Printf("%v", send(input))
}
