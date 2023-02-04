package main

import (
	"fmt"

	"github.com/tv42/zbase32"
)

func main() {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	res := zbase32.EncodeToString(data)
	fmt.Println(res)
	// yrbygbyfyadoo

	res2 := zbase32.EncodeToString([]byte("hello"))
	fmt.Println(res2)
	// pb1sa5dx
}
