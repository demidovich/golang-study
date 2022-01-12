package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/demidovich/go-learn/util"
)

func main() {

	file := file()
	data := data(file)

	for _, row := range data {
		fmt.Println(row)
	}
}

func data(filename string) []string {
	data, err := util.ReadFileUnix(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return data
}

func file() string {
	path, err := filepath.Abs("01/urls.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(path)
	return path
}
