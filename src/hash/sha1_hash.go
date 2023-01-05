package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha1.New();
	h.Write([]byte("test"));
	bs := h.Sum([]byte{});
	fmt.Println(bs);
	he := hex.EncodeToString(bs);
	fmt.Println(he);
}