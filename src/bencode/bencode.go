package bencode

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(msg map[string]string) string {
	result := "d"
	for k, v := range msg {
		result += strconv.Itoa(len(k)) + ":" + k + strconv.Itoa(len(v)) + ":" + v
	}
	result += "e"
	return result
}

// decoder should work in a streaming fashion.
// one byte at a time.

func Decode(emsg string) (map[string]interface{}, bool) {
	result := map[string]interface{}{}

	ptr := 0

	if emsg[ptr] == 'd' { // we are decoding a dictionary
		ptr++
		key := "" // current key being decoded
		inList := false
		list := []string{}

		for ptr <= len(emsg) {
			nextRune := emsg[ptr : ptr+1]

			if nextRune == "l" {
				inList = true
				ptr++
			} else if nextRune == "e" {
				if inList {
					inList = false
					result[key] = list
					ptr++
				} else {
					break
				}
			} else {
				cpos := strings.Index(emsg[ptr:], ":") + ptr

				if ptr > cpos { // incomplete msg
					return nil, false
				}

				length, err := strconv.Atoi(emsg[ptr:cpos])

				if ptr+length > len(emsg) { // length exceeds remaining message
					return nil, false
				}

				if err != nil {
					fmt.Errorf("problem converting %v to int\n", err)
				}

				ptr = cpos + 1

				if len(key) == 0 { // decoding key
					key = emsg[ptr : ptr+length]
				} else { // decoding value
					if inList {
						list = append(list, emsg[ptr:ptr+length])
					} else {
						result[key] = emsg[ptr : ptr+length+1]
						key = ""
					}
				}

				ptr += length
			}
		}
	}

	return result, true
}
