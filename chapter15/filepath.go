package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.IsAbs("."))

	exe, _ := os.Executable()

	fmt.Println(filepath.IsAbs(exe))
	fmt.Println(exe)
	fmt.Println(filepath.Dir(exe))
	fmt.Println(filepath.Base(exe))
	fmt.Println(filepath.Split(exe))

	fmt.Println(filepath.Ext(exe))

	fmt.Println(filepath.FromSlash(exe))
	fmt.Println(filepath.FromSlash("c:/test/exe"))
	fmt.Println(filepath.ToSlash(exe))

	fmt.Println(filepath.ToSlash("c:\\test\\exe"))

	fmt.Println(filepath.Glob("e:\\codes\\*.zip"))
	fmt.Println(filepath.Match("e:\\*\\codes\\*.go", "e:\\htgolang\\codes\\project.go"))

	fmt.Println(filepath.Join("e:\\", "codes", "project.go"))
	fmt.Println(filepath.SplitList(os.Getenv("PATH")))

	filepath.Walk("e:\\codes", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() && info.Name() != "codes" {
			return filepath.SkipDir
		}
		fmt.Println(path)
		return nil
	})
}
