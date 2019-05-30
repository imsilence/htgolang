package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abcdefabc", "abcdefg"))
	fmt.Println(strings.Contains("abcdefabc", "cde"))
	fmt.Println(strings.Count("abcdefabc", "cde"))
	fmt.Println(strings.Count("abcdefabc", "cde"))
	fmt.Printf("%#v\n", strings.Fields("ab c\nde\tfa\fb\r\nc"))
	fmt.Println(strings.HasPrefix("abcdefabc", "abc"))
	fmt.Println(strings.HasSuffix("abcdefabc", "abc"))
	fmt.Println(strings.Index("abcdefabc", "abc"))
	fmt.Println(strings.LastIndex("abcdefabc", "abc"))
	fmt.Printf("%#v\n", strings.Split("abcdefabc", "b"))
	fmt.Printf("%#v\n", strings.SplitN("abcdefabc", "b", 2))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println(strings.Map(func(r rune) rune { return r - 32 }, "abc"))
	fmt.Println(strings.Repeat("abc", 3))
	fmt.Println(strings.Replace("abcabcabcdefabc", "abc", "x", 2))
	fmt.Println(strings.ReplaceAll("abcabcabcdefabc", "abc", "x"))
	fmt.Println(strings.Title("abc defc c"))
	fmt.Println(strings.ToLower("abc defc c"))
	fmt.Println(strings.ToUpper("abc defc c"))
	fmt.Println(strings.Trim("abc defc c", "a c"))
	fmt.Println(strings.TrimSpace(" \t \f abc \r\n "))
	fmt.Println(strings.TrimLeft("abc defc c", "a c"))
	fmt.Println(strings.TrimRight("abc defc c", "a c"))
	fmt.Println(strings.TrimLeft("abc defc c", "a c"))
	fmt.Println(strings.TrimRight("abc defc c", "a c"))
	fmt.Println(strings.TrimPrefix("abc defc c", "a c"))
	fmt.Println(strings.TrimSuffix("abc defc c", " c"))

	var builder strings.Builder

	builder.Write([]byte{'a', 'b', 'c'})
	builder.WriteByte('m')
	builder.WriteRune('棒')
	builder.WriteString("我是谁?")

	fmt.Println(builder.String())
	builder.Reset()
	fmt.Println(builder.String())

	var reader *strings.Reader = strings.NewReader("abcdef")

	var cxt []byte = make([]byte, 10)
	fmt.Println(reader.Read(cxt))
	fmt.Println(string(cxt))

	var cxt2 []byte = make([]byte, 10)
	reader.Reset("xyz")
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
