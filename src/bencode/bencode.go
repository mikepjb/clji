package bencode

import "strconv"

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

func Decode(emsg string) map[string]interface{} {
	var result map[string]interface{}

	return result
}
