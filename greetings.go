package main

import (
	"fmt"
	"strconv"
	"strings"
)

var x int
var y string
var z bool

type ilustre int

var pedro ilustre

func main() {
	var key string
	fmt.Println("Digite qual função deseja executar:")
	fmt.Println("1 - Exercício 1")
	fmt.Println("2 - Exercício 2")
	fmt.Println("3 - Exercício 3")
	fmt.Println("4 - Exercício 4")
	fmt.Println("5 - Calculadora IMC")

	fmt.Scan(&key)

	if key == "1" {
		exercicio1()
	} else if key == "2" {
		exercicio2()
	} else if key == "3" {
		exercicio3()
	} else if key == "4" {
		exercicio4()
	} else {
		calculadoraIMC()
	}
}

func exercicio1() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println("### Somente uma declaração de print! ###")
	fmt.Println("O valor de x é:", x, "\n", "O valor de y é:", y, "\n", "O valor de z é:", z)
	fmt.Println(" ")
	fmt.Println("### Uma declaração de print por variavel! ###")
	fmt.Println("O valor de x é:", x)
	fmt.Println("O valor de y é:", y)
	fmt.Println("O valor de z é:", z)
}

func exercicio2() {
	fmt.Println(" ")
	fmt.Print("O valor zero de x é: ", x, "\n", "O valor zero de y é: ", y, "\n", "O valor zero de z é: ", z)
}

func exercicio3() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("Valor de x: %v \n Valor de y: %v \n Valor de z: %v", x, y, z)
	fmt.Println(s)
}

func exercicio4() {
	fmt.Print(pedro)
	fmt.Printf("\n %T", pedro)
	pedro = 42
	fmt.Print("\n", pedro)
}

func calculadoraIMC() {

	var result float64
	var alt, p string
  
	fmt.Println("*********************************************************")
	fmt.Println("----------------Calculadora de IMC-----------------------")
	fmt.Println("*********************************************************")

	fmt.Println("Digite sua altura em metros (exemplo: 1,75)")
	fmt.Scan(&alt)
	alt = strings.Replace(alt, ",", ".", -1)

	fmt.Println("Digite seu peso em quilos (exemplo: 95,1)")
	fmt.Scan(&p)
	p = strings.Replace(p, ",", ".", -1)

	altura, _ := strconv.ParseFloat(alt, 64)
	peso, _ := strconv.ParseFloat(p, 64)
	result = peso / (altura * altura)

	if result > 18.50 && result < 24.90 {
		fmt.Printf(" Seu IMC é %.2f. Você tem o peso normal.", result)
	} else if result > 25 && result < 29.99 {
		fmt.Printf(" Seu IMC é %.2f. Você está com sobrepeso.", result)
	} else if result > 30 && result < 34.90 {
		fmt.Printf(" Seu IMC é %.2f. Você tem Obesidade Grau I .", result)
	} else if result > 35 && result < 39.99 {
		fmt.Printf(" Seu IMC é %.2f. Você tem Obesidade Grau II .", result)
	} else if result > 40 {
		fmt.Printf(" Seu IMC é .2%f. Você tem Obesidade Grau III, CUIDADO!!!", result)
	} else {
		fmt.Printf(" Seu IMC é %.2f", result)
	}
}

// No "GO playground" deu certo, mas não deu pra ler os valores, testa no teu ambiente vê se funciona ;)
