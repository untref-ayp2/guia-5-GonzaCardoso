package main

import (
	"fmt"

	"guia5/mapadebits"
)

func main() {
	m := mapadebits.NewMapaDeBits()

	m.Encender(3)
	fmt.Printf("%032b\n", *m)

	m.Encender(5)
	m.Encender(6)
	fmt.Printf("%032b\n", *m)

	m.Apagar(3)
	fmt.Printf("%032b\n", *m)

	var x int32 = 63
	var y int32 = 56345
	fmt.Printf("%032b = x, representa el conjunto {1,2,3,4,5,6}\n", x)
	fmt.Printf("%032b = y, representa el conjunto {1,4,5,11,12,13,15,16}\n", y)
	fmt.Printf("%032b = x&y, representa el conjunto {1,4,5}\n", x&y)                      // Interseccion
	fmt.Printf("%032b = x|y, representa el conjunto {1,2,3,4,5,6,11,12,13,15,16}\n", x|y) //Union
	fmt.Printf("%032b = x^y, representa el conjunto {2,3,6,11,12,13,15,16}\n", x^y)       //diferencia entre ambos
	fmt.Printf("%032b = x&^y,representa el conjunto {2,3,6}\n", x&^y)                     //diferencia referente a x
	fmt.Printf("%032b = y&^x,representa el conjunto {11,12,13,15,16}\n", y&^x)            //diferencia referencia a y
}
