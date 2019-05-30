package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	fmt.Println([]byte("abcdef"))
	fmt.Println(string([]byte("abcdef")))

	fmt.Println(bytes.Compare([]byte("abcdefabc"), []byte("abcdefg")))
	fmt.Println(bytes.Contains([]byte("abcdefabc"), []byte("cde")))
	fmt.Println(bytes.Count([]byte("abcdefabc"), []byte("cde")))
	fmt.Println(bytes.Count([]byte("abcdefabc"), []byte("cde")))
	fmt.Printf("%#v\n", bytes.Fields([]byte("ab c\nde\tfa\fb\r\nc")))
	fmt.Println(bytes.HasPrefix([]byte("abcdefabc"), []byte("abc")))
	fmt.Println(bytes.HasSuffix([]byte("abcdefabc"), []byte("abc")))
	fmt.Println(bytes.Index([]byte("abcdefabc"), []byte("abc")))
	fmt.Println(bytes.LastIndex([]byte("abcdefabc"), []byte("abc")))
	fmt.Printf("%#v\n", bytes.Split([]byte("abcdefabc"), []byte("b")))
	fmt.Printf("%#v\n", bytes.SplitN([]byte("abcdefabc"), []byte("b"), 2))
	fmt.Println(bytes.Join([][]byte{[]byte("abc"), []byte("bcd"), []byte("cde")}, []byte("-")))
	fmt.Println(bytes.Map(func(r rune) rune { return r - 32 }, []byte("abc")))
	fmt.Println(bytes.Repeat([]byte("abc"), 3))
	fmt.Println(bytes.Replace([]byte("abcabcabcdefabc"), []byte("abc"), []byte("x"), 2))
	fmt.Println(bytes.ReplaceAll([]byte("abcabcabcdefabc"), []byte("abc"), []byte("x")))
	fmt.Println(bytes.Title([]byte("abc defc c")))
	fmt.Println(bytes.ToLower([]byte("abc defc c")))
	fmt.Println(bytes.ToUpper([]byte("abc defc c")))
	fmt.Println(bytes.Trim([]byte("abc defc c"), "a c"))
	fmt.Println(bytes.TrimSpace([]byte(" \t \f abc \r\n ")))
	fmt.Println(bytes.TrimLeft([]byte("abc defc c"), "a c"))
	fmt.Println(bytes.TrimRight([]byte("abc defc c"), "a c"))
	fmt.Println(bytes.TrimLeft([]byte("abc defc c"), "a c"))
	fmt.Println(bytes.TrimRight([]byte("abc defc c"), "a c"))
	fmt.Println(bytes.TrimPrefix([]byte("abc defc c"), []byte("a c")))
	fmt.Println(bytes.TrimSuffix([]byte("abc defc c"), []byte(" c")))

	builder := bytes.NewBufferString("abcdef")

	fmt.Println(builder.Bytes())
	fmt.Println(builder.String())

	cxt := make([]byte, 3)

	builder.Read(cxt)
	fmt.Println(cxt)

	fmt.Println(builder.ReadString('\n'))

	builder.WriteString("xyz")
	fmt.Println(builder.ReadString('\n'))

	builder.WriteString("xyz")
	builder.Reset()
	fmt.Println(builder.ReadString('\n'))

	builder.WriteString("xyz")
	builder.WriteTo(os.Stdout)

	var reader *bytes.Reader = bytes.NewReader([]byte("abcdef"))

	var cxt1 []byte = make([]byte, 10)
	fmt.Println(reader.Read(cxt1))
	fmt.Println(string(cxt1))

	var cxt2 []byte = make([]byte, 10)
	reader.Reset([]byte("xyz"))
	fmt.Println(reader.Read(cxt2))
	fmt.Println(string(cxt2))

	var cxt3 []byte = make([]byte, 10)
	reader.Seek(1, 0)
	fmt.Println(reader.Read(cxt3))
	fmt.Println(string(cxt3))

	var cxt4 []byte = make([]byte, 10)
	reader.Seek(10, 0)
	fmt.Println(reader.Read(cxt4))
	fmt.Println(string(cxt4))

	reader.Seek(0, 0)
	reader.WriteTo(os.Stdout)

}
