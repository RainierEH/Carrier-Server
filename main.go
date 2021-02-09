package main

import (
	db "./src/database"
	"fmt"
)

func main() {
	fmt.Println("Test?")
	teststr := db.Test("Test!")
	fmt.Println(teststr)
}
