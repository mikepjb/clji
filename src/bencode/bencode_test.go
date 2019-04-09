package bencode_test

import (
	"reflect"
	"testing"

	"github.com/mikepjb/clji/src/bencode"
)

func TestEncode(t *testing.T) {
	emsg := bencode.Encode(map[string]string{"op": "clone"})

	if emsg != "d2:op5:clonee" {
		t.Errorf("wrong encoded message: %v\n", emsg)
	}
}

func TestDecode(t *testing.T) {
	emsg := "d11:new-session36:7db0cecd-6f2a-4b57-a29b-c01c18eb7c897:session36:9340a182-e4b5-4bda-a0a9-74671af021486:statusl4:doneee"

	msg := map[string]interface{}{
		"new-session": "7db0cecd-6f2a-4b57-a29b-c01c18eb7c897",
		"session":     "9340a182-e4b5-4bda-a0a9-74671af021486",
		"status":      []string{"done"},
	}

	dmsg, ok := bencode.Decode(emsg)

	if ok && !reflect.DeepEqual(dmsg, msg) {
		t.Errorf("incorrect decoding, got: %v\n", dmsg)
	}
}

// because we are reading bencoded info in a stream we'd like to know if we have
// a complete data structure.
func TestFlagIncompleteBencode(t *testing.T) {
	// open dictionary only
	_, ok := bencode.Decode("d")

	if ok {
		t.Errorf("open dictionary is not a complete bencode message")
	}

	// missing dictionary close with complete keypair
	_, ok = bencode.Decode("d11:new-session")
	if ok {
		t.Errorf("got ok for missing dictionary close with complete keypair")
	}

	// with complete value
	partialWithCompleteValue := "d11:new-session36:7db0cecd-6f2a-4b57-a29b-c01c18eb7c897"

	_, ok = bencode.Decode(partialWithCompleteValue)

	if ok {
		t.Errorf("message should not be complete for partialWithCompleteValue")
	}

	// with incomplete value (asking for it would exceed the range of the string)
	partialWithIncompleteValue := "d11:new-session36:7db0cecd-6f2a-4b57-a29b-c01c1"

	_, ok = bencode.Decode(partialWithIncompleteValue)

	if ok {
		t.Errorf("message should not be complete for partialWithIncompleteValue")
	}
}
