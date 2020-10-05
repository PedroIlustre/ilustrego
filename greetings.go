package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
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
