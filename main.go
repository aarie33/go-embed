package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed version.txt
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")

	for _, d := range dir {
		if !d.IsDir() {
			fmt.Println(d.Name())
			content, _ := path.ReadFile("files/" + d.Name())
			fmt.Println(string(content))
		}
	}
}
