package util

import (
	"os"
	"strings"
)

// func ReadFile(filename string, capacity uint) ([]string, error) {

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	data := make([]string, 0, capacity)
// 	read := bufio.NewScanner(file)
// 	for read.Scan() {
// 		data = append(data, read.Text())
// 	}

// 	if err := read.Err(); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func ReadFileUnix(filename string) ([]string, error) {
	return ReadFile(filename, "\n")
}

func ReadFile(filename string, lineSeparator string) ([]string, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(raw), lineSeparator), nil
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !fi.Mode().IsDir() {
		return false
	}
	return true
}
