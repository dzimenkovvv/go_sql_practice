package main

import (
	"Study/feature_postgres/simple_connection"
	"Study/features"
	"fmt"
)

func main() {
	fmt.Println("Hello, It's main!")
	features.Feature1()
	features.Feature2()
	simple_connection.CheckConnection()
}
