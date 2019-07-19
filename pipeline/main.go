package main

import "fmt"

func generador(runs ...int) <-chan int {
	salida := make(chan int)
	go func() {
		for _, n := range runs {
			salida <- n
		}
		close(salida) //cierro el channel para que no se pueda agregar más nada y saber el tamaño exacto
	}()
	return salida
}

func cuadrado(in <-chan int) <-chan int {
	salida := make(chan int)
	go func() {
		for n := range in {
			salida <- (n * n)
		}
		close(salida)
	}()
	return salida
}

func main() {
	c1 := generador(1, 2, 3, 4, 5)
	c2 := cuadrado(c1)
	for n := range c2 {
		fmt.Println(n)
	}
}
