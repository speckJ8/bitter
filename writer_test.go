package bitter

import (
        "bytes"
        "testing"
)

var writerTestOutput = []byte{ 0xff, 0xff, 0x18, 0xff, 0xff, 0xff }
var testWrites = []struct{ size int; value uint64; result []byte } {
    { size: 12, value: 0xfff,   result: []byte{0xff, 0x0f} },
    { size: 8,  value: 0x8f,    result: []byte{0xff, 0xff, 0x08} },
    { size: 8,  value: 0xf1,    result: []byte{0xff, 0xff, 0x18, 0x0f} },
    { size: 20, value: 0xfffff, result: []byte{0xff, 0xff, 0x18, 0xff, 0xff, 0xff} },
}

func TestWriter(t *testing.T) {
    var output bytes.Buffer
    var writer = NewBitWriter(&output)
    for _, w := range testWrites {
        bitsWritten, err := writer.Write(w.value, w.size)
        if err != nil || bitsWritten != w.size {
            t.Fatalf("writer err: {written=%d}; %v\n", bitsWritten, err)
        }
    }
    err := writer.Close()
    if err != nil {
        t.Fatalf("writer closing err: %s\n", err.Error())
    }
    if !bytes.Equal(output.Bytes(), writerTestOutput) {
        t.Fatalf("invalid answer: output=%v{%v}\n", output.Bytes(), writerTestOutput)
    }
}
