package golang_embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*
var files embed.FS

func TestMultipleEmbed(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T){
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir(){
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/"+entry.Name())
			fmt.Println(string(file))
		}
	}
}