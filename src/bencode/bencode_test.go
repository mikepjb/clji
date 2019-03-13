package bencode_test

import (
	"testing"

	"github.com/mikepjb/clji/src/bencode"
)

func TestEncode(t *testing.T) {
	emsg := bencode.Encode(map[string]string{"op": "clone"})

	if emsg != "d2:op5:clonee" {
		t.Errorf("wrong encoded message: %v\n", emsg)
	}
}
