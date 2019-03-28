package main

import (
	"go-red-packets/infra/algorithm"
	"log"
)

func main() {
	count, amount := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		x := algorithm.DoubleAverage(count, amount*100)
		log.Printf("x = %.2f\n", float64(x)/float64(100))
	}
}
