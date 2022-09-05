package main

import (
	"fmt"
)

func main() {
	db.conex()
	db.closeDB()
	fmt.Println("estaCorriendo")
}
