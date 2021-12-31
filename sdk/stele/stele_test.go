package stele

import (
	"fmt"
	"log"
	"testing"
)

func TestStele(t *testing.T) {
	local, err := NewLocal("./steledata")
	if err != nil {
		fmt.Println("aaaa...")
		log.Fatalln(err)
	}

	get, err := local.Get([]byte("9e76420c-02b2-4a31-a27e-ad6652883c34"))
	if err != nil {
		fmt.Println("bbbb...")
		panic(err)
	}
	fmt.Println("ccc...")
	fmt.Println(string(get))
}
