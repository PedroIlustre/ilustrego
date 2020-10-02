package main

import ( 
	"fmt" 
)

func main() {
	var alt, peso, result float64
	
	fmt.Println("*********************************************************")
	fmt.Println("----------------Calculadora de IMC-----------------------")
	fmt.Println("*********************************************************")

	fmt.Println("Digite sua altura em centímetros (exemplo: 1,75m = 175cm)")
	fmt.Scan(&alt)
	fmt.Println("Digite seu peso em gramas (exemplo: 95,80kg = 9580g)")
	fmt.Scan(&peso)
	result = (alt * alt) / peso

	if result > 18.50 && result < 24.90 {
		fmt.Printf(" Seu IMC é %.2f. Você tem o peso normal.",result)
	} else if result > 25 && result < 29.99 {
		fmt.Printf(" Seu IMC é %.2f. Você está com sobrepeso.",result)
	} else if result > 30 && result < 34.90 {
		fmt.Printf(" Seu IMC é %.2f. Você tem Obesidade Grau I .",result)
	} else if result > 35 && result < 39.99 {
		fmt.Printf(" Seu IMC é %.2f. Você tem Obesidade Grau II .",result)
	} else if result > 40 {
		fmt.Printf(" Seu IMC é %.2f. Você tem Obesidade Grau III, CUIDADO!!!",result)
	}

}

// No "GO playground" deu certo, mas não deu pra ler os valores, testa no teu ambiente vê se funciona ;)