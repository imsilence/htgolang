package main

import (
	"fmt"
	htmlTpl "html/template"
	"os"
	textTpl "text/template"
)

func main() {

	// 使用text/template
	fmt.Println(textTpl.HTMLEscapeString("<div>kk</div><script>alert(1);</script>"))
	fmt.Println(textTpl.JSEscapeString("<div>kk</div><script>alert(1);</script>"))

	txt := `
{{ . }}
{{ .|html }}
{{ .|js }}
`
	ttpl := textTpl.Must(textTpl.New("tpl").Parse(txt))
	ttpl.Execute(os.Stdout, "<div>kk</div><script>alert(1);</script>")


	// 使用html/template
	fmt.Println(htmlTpl.HTMLEscapeString("<div>kk</div><script>alert(1);</script>"))
	fmt.Println(htmlTpl.JSEscapeString("<div>kk</div><script>alert(1);</script>"))

	htpl := htmlTpl.Must(htmlTpl.New("tpl").Parse(txt))
	htpl.Execute(os.Stdout, "<div>kk</div><script>alert(1);</script>")
}
