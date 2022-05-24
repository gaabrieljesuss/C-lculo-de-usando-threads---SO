package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

const numeroExec = 5

var valoresPi, tempoPorExecucao [numeroExec]float64

var somaTempo float64

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

func calcularDesvioPadrao(mediaTempo float64) float64 {
	var soma float64
	var desvioPadrao float64

	for _, tempo := range tempoPorExecucao {
		dispersao := float64(tempo) - mediaTempo
		soma += math.Pow(dispersao, 2) / numeroExec
	}

	desvioPadrao = math.Sqrt(soma)
	return desvioPadrao
}

func getNumeroThreads() int {
	var numeroThreads, erro = strconv.Atoi(os.Getenv("NUMERO_THREADS"))

	if erro != nil {
		panic("Valor da variável NUMERO_THREADS inválido ou não definido!")
	}

	return numeroThreads
}

func getNumeroTermos() int {
	var numeroTermos, erro = strconv.Atoi(os.Getenv("NUMERO_TERMOS"))

	if erro != nil {
		panic("Valor da variável NUMERO_TERMOS inválido ou não definido!")
	}

	return numeroTermos
}

func main() {
	var numeroThreads = getNumeroThreads()
	var numeroTermosTotal = getNumeroTermos()

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

		tempoPorExecucao[i] = time.Since(inicioThread).Seconds()
		somaTempo += float64(tempoPorExecucao[i])
	}

	mediaTempo := somaTempo / numeroExec

	desvioPadrao := calcularDesvioPadrao(mediaTempo)

	fmt.Println("Número de Threads: ", numeroThreads)
	fmt.Println("Número de Termos: ", numeroTermosTotal)
	fmt.Println("Valores de pi:", valoresPi)
	fmt.Println("Tempo médio de execução: ", mediaTempo)
	fmt.Println("Desvio padrão entre execuções: ", desvioPadrao)
}
