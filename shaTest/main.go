package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

// func GetSha256(a []byte) [32]byte {
func GetSha256(a []byte) [32]byte {
	c1 := sha256.Sum256(a)
	return c1
}

func GetSha384(a []byte) [48]byte {
	c1 := sha512.Sum384(a)
	return c1
}

func GetSha512(a []byte) [64]byte {
	c1 := sha512.Sum512(a)
	return c1
}

func main() {
	// var a []byte := os.Args[2]
	// var a = os.Args[2]
	// var b
	if os.Args[1] == "256" {
		fmt.Print("using sha256\n")
		b := GetSha256([]byte(os.Args[2]))
		fmt.Printf("%s %x\n", os.Args[2], b)
	} else if os.Args[1] == "384" {
		fmt.Print("using sha384\n")
		b := GetSha384([]byte(os.Args[2]))
		fmt.Printf("%s %x\n", os.Args[2], b)

	} else if os.Args[1] == "512" {
		fmt.Print("using sha512\n")
		b := GetSha512([]byte(os.Args[2]))
		fmt.Printf("%s %x\n", os.Args[2], b)

	} else {
		fmt.Print("method err 256 | 384 | 512\n")
		return
	}

}
