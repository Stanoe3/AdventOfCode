package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	buf := make([]byte, 102400)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			return string(buf[:n])
		}
	}
	return ""
}
