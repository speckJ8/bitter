package bitter

import (
	"bytes"
	"testing"
)

var readerTestInput = []byte{ 0xff, 0xff, 0x18, 0xff, 0xff, 0xff }
var testReads = []struct{ size int; result uint64 } {
    { size: 12, result: 0xfff },
    { size: 8,  result: 0x8f },
    { size: 8,  result: 0xf1 },
    { size: 20, result: 0xfffff },
}

func TestRead(t *testing.T) {
    var reader = NewBitReader(bytes.NewReader(readerTestInput))
    for _, r := range testReads {
        var bits, bitsRead, err = reader.Read(r.size)
        if err != nil {
            t.Fatalf("reader err: %s\n", err.Error())
        }
        if bits != r.result || bitsRead != r.size {
            t.Fatalf("invalid answer: bits=%b{%b}, bitsRead=%d{%d}\n",
                bits, r.result, bitsRead, r.size)
        }
    }
}
