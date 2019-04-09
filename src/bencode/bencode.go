package bencode

import (
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
	inList := false
	inDict := false
	key := ""          // current key being decoded
	list := []string{} // store values when building a list

	for ptr < len(emsg) {

		switch emsg[ptr : ptr+1] {
		case "d":
			inDict = true
			ptr++
		case "l":
			inList = true
			ptr++
		case "e":
			if inList {
				inList = false
				result[key] = list
				ptr++
			} else {
				inDict = false
				ptr++
				break
			}
		default:
			// what happens if there is no colon left?
			nextColonIndex := strings.Index(emsg[ptr:], ":")

			if nextColonIndex == -1 { // no more elements but dict is unfinished
				return nil, false
			}

			cpos := nextColonIndex + ptr              // colon position
			length, _ := strconv.Atoi(emsg[ptr:cpos]) // next element length

			if ptr+length > len(emsg) { // length exceeds remaining message
				return nil, false
			}

			ptr = cpos + 1

			if ptr+length >= len(emsg) { // incomplete key
				return nil, false
			}

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

	if inDict {
		return nil, false // dict is unfinished
	}

	return result, true
}
