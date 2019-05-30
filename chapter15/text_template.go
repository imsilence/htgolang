package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	fmt.Println(template.HTMLEscapeString("<div>kk</div><script>alert(1);</script>"))
	fmt.Println(template.JSEscapeString("<div>kk</div><script>alert(1);</script>"))

	// 显示信息
	txt := `show var:
{{ . }}`
	tpl, _ := template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, "kk")

	fmt.Println()

	// 显示切片元素
	txt = `show slice:
{{ index . 0 }}
{{ index . 1 }}
`
	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, []string{"silence", "kk"})

	fmt.Println()

	// 显示映射元素
	txt = `show map:
{{ .name }}
{{ .addr }}
`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, map[string]string{"name": "kk", "addr": "西安市"})

	fmt.Println()

	// if else表达式
	txt = `if-else:
{{ if . }}
True
{{ else }}
False
{{ end }}`
	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, true)
	tpl.Execute(os.Stdout, false)

	fmt.Println()

	// if else-if else表达式
	txt = `if-elseif-else:
{{ if eq . "white" }}
WHITE
{{ else if eq . "black" }}
BLACK
{{ else }}
unknow
{{ end }}`
	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, "white")
	tpl.Execute(os.Stdout, "black")
	tpl.Execute(os.Stdout, "red")

	fmt.Println()

	// range表达式
	txt = `range:
{{ range . }}
{{ . }}
{{ end }}`
	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, []string{"white", "black", "red"})

	fmt.Println()

	// 属性显示、if-else表达式、range-else表达式
	txt = `
Title: {{ .Title }}
Students:
{{ range .Students }}
	学号: {{ .ID }}
	性名: {{ .Name }}
	性别: {{ if .Gender }} 男 {{ else }} 女 {{ end }}
{{ else }}
	无学生信息
{{ end }}
`
	type Student struct {
		ID     int
		Name   string
		Gender bool
	}

	type Class struct {
		Title    string
		Students []Student
	}

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, Class{
		Title: "三年级二班",
		Students: []Student{
			{ID: 1, Name: "kk", Gender: true},
			{ID: 2, Name: "ada", Gender: false},
			{ID: 3, Name: "woniu", Gender: true},
		},
	})

	tpl.Execute(os.Stdout, Class{
		Title:    "三年级二班",
		Students: []Student{},
	})

	// range表达式
	txt = `
{{ range . }}
	Title: {{ .Title }}
	Students:
	{{ range .Students }}
		学号: {{ .ID }}
		性名: {{ .Name }}
		性别: {{ if .Gender }} 男 {{ else }} 女 {{ end }}
	{{ end }}
{{ end }}`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, []Class{
		{
			Title: "三年级二班",
			Students: []Student{
				{ID: 1, Name: "kk", Gender: true},
				{ID: 2, Name: "ada", Gender: false},
				{ID: 3, Name: "woniu", Gender: true},
			},
		}, {
			Title: "三年级三班",
			Students: []Student{
				{ID: 1, Name: "wd", Gender: true},
				{ID: 2, Name: "xiaolin", Gender: false},
				{ID: 3, Name: "xuegao", Gender: true},
			},
		},
	})

	// 定义变量
	txt = `
{{ $boy := false }}

{{ if $boy }}
男
{{ else }}
女
{{ end }}`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, nil)

	// with表达式指定上下文
	txt = `
	{{ . }}
{{ with "output" }}
	{{ . }}
{{ else }}
	{{ . }}
{{ end }}`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, "txt")

	// 使用内置函数
	txt = `
{{ . }}
{{ printf "%T" . }}
{{ .|printf "%T" }}
{{ len . }}
{{ .|len }}
{{ html . }}
{{ .|html }}
{{ js . }}
{{ .|js }}
`
	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, "<div>kk</div><script>alert(1);</script>")

	// 自定义函数
	txt = `
{{ . }}
{{ upper . }}
{{ .|lower }}

`
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}

	tpl, _ = template.New("tpl").Funcs(funcs).Parse(txt)
	tpl.Execute(os.Stdout, "abcdefgABCDEFG")

	// 使用块指定默认值
	txt = `
content: {{ block "content" . }} {{ . }} {{ end }}
`
	overlay := `{{ define "content" }} {{ . | len }}:{{ . }} {{ end }}`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, "content")

	tpl, _ = tpl.Parse(overlay)
	tpl.Execute(os.Stdout, "content")

	// 使用模板
	txt = `
{{ define "head" }}
<head><title>{{ .Title }}</title></head>
{{ end }}

{{ define "body" }}
<body>{{ .Body }}</body>
{{ end }}

{{ define "html" }}
<html>
{{ template "head" . }}
{{ template "body" . }}
</html>
{{ end }}

{{ template "html" . }}
`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.Execute(os.Stdout, struct {
		Title string
		Body  string
	}{
		"网页Title",
		"网页Body",
	})

	// 指定模板
	txt = `
{{ define "head" }}
<head><title>{{ .Title }}</title></head>
{{ end }}

{{ define "body" }}
<body>{{ .Body }}</body>
{{ end }}

{{ define "html" }}
<html>
{{ template "head" . }}
{{ template "body" . }}
</html>
{{ end }}
`

	tpl, _ = template.New("tpl").Parse(txt)
	tpl.ExecuteTemplate(os.Stdout, "html", struct {
		Title string
		Body  string
	}{
		"网页Title",
		"网页Body",
	})

	// 从文件中读取模板
	tpl, _ = template.New("tpl").ParseFiles("tpls/index.html", "tpls/layout.html")

	for _, template := range tpl.Templates() {
		fmt.Println(template.Name())
	}

	tpl.ExecuteTemplate(os.Stdout, "layout.html", struct {
		Title string
		Body  string
	}{
		"网页Title",
		"网页Body",
	})

	// 使用匹配格式的文件中读取模板
	tpl = template.Must(template.New("tpl").ParseGlob("tpls/*.html"))

	for _, template := range tpl.Templates() {
		fmt.Println(template.Name())
	}

	// 覆盖block块
	tpl.Parse(`{{ define "block" }} tpl block content {{ end }}`)

	tpl.ExecuteTemplate(os.Stdout, "layout.html", struct {
		Title string
		Body  string
	}{
		"网页Title",
		"网页Body",
	})
}
