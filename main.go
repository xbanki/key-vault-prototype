package main

const FILE_NAME = "test-db.json"

func main() {
	CreateDatabase(WithFilepath(FILE_NAME), WithHydrate(true))
}
