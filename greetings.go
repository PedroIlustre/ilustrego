package main

import (
	"fmt"
	"strconv"
	"strings"
)

var x2 int
var y2 string
var z2 bool

type ilustre int

var pedro ilustre
var pedro5 int

func main() {
	var key string
	fmt.Println("Digite qual função deseja executar:")
	fmt.Println("1 - Exercício 1")
	fmt.Println("2 - Exercício 2")
	fmt.Println("3 - Exercício 3")
	fmt.Println("4 - Exercício 4")
	fmt.Println("5 - Exercício 5")
	fmt.Println("outro - Calculadora IMC")

	fmt.Scan(&key)

	if key == "1" {
		exercicio1()
	} else if key == "2" {
		exercicio2()
	} else if key == "3" {
		exercicio3()
	} else if key == "4" {
		exercicio4()
	} else if key == "5" {
		exercicio5()
	} else {
		aulaBooleano()
		//calculadoraIMC()
	}
}

func exercicio1() {

	x1 := 42
	y1 := "James Bond"
	z1 := true

	fmt.Println("### Somente uma declaração de print! ###")
	fmt.Println("O valor de x é:", x1, "\n", "O valor de y é:", y1, "\n", "O valor de z é:", z1)
	fmt.Println(" ")
	fmt.Println("### Uma declaração de print por variavel! ###")
	fmt.Println("O valor de x é:", x1)
	fmt.Println("O valor de y é:", y1)
	fmt.Println("O valor de z é:", z1)
}

func exercicio2() {
	fmt.Println(" ")
	fmt.Print("O valor zero de x é: ", x2, "\n", "O valor zero de y é: ", y2, "\n", "O valor zero de z é: ", z2)
}

func exercicio3() {
	x2 = 42
	y2 = "James Bond"
	z2 = true
	s := fmt.Sprintf("Valor de x: %v \n Valor de y: %v \n Valor de z: %v", x2, y2, z2)
	fmt.Println(s)
}

func exercicio4() {
	fmt.Print(pedro)
	fmt.Printf("\n %T", pedro)
	pedro = 42
	fmt.Print("\n", pedro)
}

func exercicio5() {
	exercicio4()
	pedro5 = int(pedro)
	fmt.Printf("\n %v", pedro)
	fmt.Printf("\n %T", pedro5)
}

func aulaBooleano() {
	varBoleana := true
	fmt.Println(varBoleana)
	varBoleana = 10 > 100
	fmt.Println(varBoleana)

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
