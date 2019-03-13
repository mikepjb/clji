package main

import (
	"fmt"

	"github.com/mikepjb/clji/src/bencode"
)

func main() {
	clone := map[string]string{"op": "clone"}
	emsg := bencode.Encode(clone)
	fmt.Println(emsg)
}
