package bitter

import (
	"errors"
	"io"
)

type BitReader struct {
    reader io.Reader
    currByte []byte
    currByteOffset int
}

func NewBitReader(reader io.Reader) BitReader {
    return BitReader{ reader: reader, currByte: []byte{ 0x00 } }
}

// Read `bits` bits from the underlying reader.
// Returns the bits read (in unsigned64 variable), the number
// of bits read and an error in case one occurs
// At most 64 bits can be read at a time.
func (r *BitReader) Read(bits int) (uint64, int, error) {
    if bits > 64 {
        return 0, 0, errors.New("can read at most 64 bits at a time")
    }

    if r.currByteOffset == 0 {
        read, err := r.reader.Read(r.currByte)
        if err != nil || read == 0 {
            return 0, 0, err
        }
    }

    var result uint64
    var bitsRead int = 0
    for bits > 0 {
        var availableInByte = 8 - r.currByteOffset
        var b = r.currByte[0] >> r.currByteOffset
        if bits <= availableInByte {
            result |= (uint64(readerKeepBits(b, bits)) << bitsRead) 
            r.currByteOffset = (r.currByteOffset + bits) % 8
            bitsRead += bits
            bits -= bits
        } else {
            result |= (uint64(readerKeepBits(b, availableInByte)) << bitsRead) 
            r.currByteOffset = 0
            bitsRead += availableInByte
            bits -= availableInByte
            read, err := r.reader.Read(r.currByte)
            if err != nil || read == 0 {
                return result, bitsRead, err
            }
        }
    }

    return result, bitsRead, nil
}

// remove the higher 8 - `bits` bits of b
func readerKeepBits(b byte, bits int) byte {
    return (b << (8 - bits)) >> (8 - bits)
}
