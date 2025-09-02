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