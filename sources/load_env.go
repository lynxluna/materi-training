package main

import (
	"fmt"
	"os"
)

func main() {
	e, ok := os.LookupEnv("ARTICLE_HOST") // <1>

	if !ok {
		fmt.Println("Tidak ada ARTICLE_HOST")
		return
	}

	fmt.Println("Host : ", e)

}
