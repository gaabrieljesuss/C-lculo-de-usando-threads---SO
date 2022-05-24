package main

import (
	"fmt"
	"math"
	"time"
)

const numeroExec = 5

var numeroThreads = 2
var numeroTermosTotal = 1000000 * 50
var valoresPi [numeroExec]float64

var somaTempo float64
var tempoPorExecucao [numeroExec]time.Duration

func calcularSomaParcial(termoInicio int, numeroTermosParcial int, c chan float64) {
	var soma float64

	for n := termoInicio; n < termoInicio+numeroTermosParcial; n++ {
		soma += math.Pow(-1, float64(n)) / float64(2*n+1)
	}

	c <- soma
}

func calcularPi(indice int, canal []chan float64) {
	for _, c := range canal {
		valoresPi[indice] += <-c
	}

	valoresPi[indice] = valoresPi[indice] * 4
}

func main() {
	inicioTempo := time.Now()
	var termoInicio int

	for i := 0; i < numeroExec; i++ {
		inicioThread := time.Now()
		var canal []chan float64
		numeroTermosParcial := numeroTermosTotal / numeroThreads
		termoInicio = 0

		for j := 0; j < numeroThreads; j++ {
			canal = append(canal, make(chan float64))
			go calcularSomaParcial(termoInicio, numeroTermosParcial, canal[j])

			termoInicio += numeroTermosParcial
		}
		calcularPi(i, canal)

		tempoPorExecucao[i] = time.Since(inicioThread)
		somaTempo = float64(tempoPorExecucao[i])
	}

	tempoDecorrido := time.Since(inicioTempo)

	mediaTempo := somaTempo / numeroExec
	var valorDados float64

	for _, tempo := range tempoPorExecucao {
		valor := float64(tempo) - mediaTempo
		valorDados += math.Pow(valor, 2) / numeroExec
	}

	desvioPadrao := math.Sqrt(valorDados)

	fmt.Println("Valor do pi:", valoresPi)
	fmt.Println("Tempo de execução tota: ", tempoDecorrido)
	fmt.Println("Tempo de execução por thread: ", tempoDecorrido/numeroExec)
	fmt.Println("Desvio padrão: ", desvioPadrao)
	fmt.Println("Tempo de cada execucao: ", tempoPorExecucao)
}
