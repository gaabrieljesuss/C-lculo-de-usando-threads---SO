package main

import (
	"fmt"
	"math"
)

const numeroExec = 5
var somaPi float64
var numeroThreads = 1
var numeroTermosTotal = 1000000
var valoresPi [numeroExec]float64


func calcularSomaParcial(termoInicio int, numeroTermosParcial int, c chan float64) {
	var soma float64

	for n := termoInicio; n < termoInicio + numeroTermosParcial; n++ {
		soma += math.Pow(-1, float64(n)) / float64(2 * n + 1)
	}
	
	c <- soma
}

func calcularPi(indice int, canal []chan float64){
	for _, c := range canal {
		valoresPi[indice] += <-c
	}

	valoresPi[indice] = valoresPi[indice] * 4
}

func main() {
	var termoInicio int

	for i := 0; i < numeroExec; i++ {
		var canal []chan float64
		numeroTermosParcial := numeroTermosTotal/numeroThreads
		termoInicio = 0

		for j := 0; j < numeroThreads; j++ {
			canal = append(canal, make(chan float64))
			go calcularSomaParcial(termoInicio, numeroTermosParcial ,canal[j])

			termoInicio += numeroTermosParcial
		}
		calcularPi(i, canal)	

	}
	
	fmt.Println("Valor do pi:", valoresPi)
}
