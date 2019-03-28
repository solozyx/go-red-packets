package main

import (
	"go-red-packets/infra/algorithm"
	"log"
)

func main() {
	log.Println(algorithm.AfterShuffle(int64(10), int64(100)*100))
}
