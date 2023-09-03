package main

import (
	"log"

	henacat "github.com/yukiHaga/henacat/src/cmd"
)

func main() {
	if err := henacat.Run(); err != nil {
		log.Printf("fail to run henacat: %v\n", err)
		return
	}
}
