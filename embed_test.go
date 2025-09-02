package golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
	"io/ioutil"
	"io/fs"
)

//go:embed version.txt
var version string

func TestString(t *testing.T){
	fmt.Println(version)
}

//go:embed zoro.jpeg
var zoro []byte

func TestByte(t *testing.T){
	err := ioutil.WriteFile("zoro_new.jpeg", zoro, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}