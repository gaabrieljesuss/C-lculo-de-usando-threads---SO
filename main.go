package main

import (
	"fmt"
	"math"
)


var valorPi float64
var numThreads = 1
var numTermos = 1000000


func calcularPi(termos int, c chan float64) {
	var pi float64

	for n := 0; n < termos; n++ {
		pi += math.Pow(-1, float64(n)) / float64(2 * n + 1)
	}

	pi = pi *  4
	c <- pi
}

func main() {
		var canal []chan float64
		for j := 0; j < numThreads; j++ {
			canal = append(canal, make(chan float64))
			go calcularPi(numTermos ,canal[j])
		}
		
		valorPi = <- canal[0]
	fmt.Println("Valor do pi ", valorPi)
}