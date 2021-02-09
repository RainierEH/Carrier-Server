package main

import (
	cmsg "./src/crypto_msg"
	db "./src/database"
	"fmt"
)

func main() {
	fmt.Println("Test?")
	teststr := db.Test("Test!")
	teststr2 := cmsg.Test2("Test2!")
	fmt.Println(teststr)
	fmt.Println(teststr2)
}
