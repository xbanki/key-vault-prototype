package main

import "fmt"

const FILE_NAME = "/home/banki/test-db.json"

func main() {
	db := CreateDatabase(WithFilepath(FILE_NAME), WithHydrate(true))
}
