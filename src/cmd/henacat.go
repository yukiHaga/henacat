package henacat

import (
	"fmt"
	"log"
	"os"
)

func Run() error {
	filePath := os.Args[1]

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("fail to read file: %v\n", err)
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
