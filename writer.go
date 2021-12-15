package bitter

import (
        "errors"
        "io"
)

type BitWriter struct {
    writer io.Writer
    currByte []byte
    currByteOffset int
}

func NewBitWriter(writer io.Writer) BitWriter {
    return BitWriter{ writer: writer, currByte: []byte{ 0x00 } }
}

// Write `bits` bits to the underlying writer.
// It writes the least significant `bits` bits from `from` to the writer.
// Returns the number of bits written and an error in case one occurs
func (w *BitWriter) Write(from uint64, bits int) (int, error) {
    var bitsWritten = 0
    for bits > 0 {
        var availableInByte = 8 - w.currByteOffset
        if bits < availableInByte {
            w.currByte[0] |=
                byte(writerKeepBits(from >> bitsWritten, bits) << w.currByteOffset)
            w.currByteOffset += bits
            bitsWritten += bits
            bits -= bits
        } else {
            w.currByte[0] |= 
                byte(writerKeepBits(from >> bitsWritten, availableInByte) << w.currByteOffset)
            written, err := w.writer.Write(w.currByte)
            if err != nil || written == 0 {
                return bitsWritten, err
            }
            w.currByte = []byte { 0x00 }
            w.currByteOffset = 0
            bitsWritten += availableInByte
            bits -= availableInByte
        }
    }
    return bitsWritten, nil
}

func (w *BitWriter) Close() error {
    if w.currByteOffset > 0 {
        _, err := w.writer.Write(w.currByte)
        if err != nil {
            return errors.New("failed to flush last byte: " + err.Error())
        }
    }
    return nil
}

// remove the higher 64 - `bits` bits of u
func writerKeepBits(u uint64, bits int) uint64 {
    return (u << (64 - bits)) >> (64 - bits)
}
