# bitter
Perform bit level I/O from any `io.Reader` or `io.Writer`.

Example
-------

Reading
```go
var input = []byte{ 0x35, 0x02 }
var reader = bitter.NewBitReader(bytes.NewReader(input))
bits, _, _ := reader.Read(10)
fmt.Println("%b\n", bits) // 1000110101
```

Writing
```go
var output bytes.Buffer
var writer = bitter.NewBitReader(&output)
writer.Write(0b1, 1)
writer.Write(0b10, 2)
writer.Write(0b10, 2)
writer.Write(0b001, 3)
writer.Write(0b010, 3)
writer.Close()
fmt.Println("%v\n", output.Bytes()) // { 0x35, 0x02 } 
```
