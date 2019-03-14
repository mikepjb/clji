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

// d11:new-session36:7db0cecd-6f2a-4b57-a29b-c01c18eb7c897:session36:9340a182-e4b5-4bda-a0a9-74671af021486:statusl4:doneee
