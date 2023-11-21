// O segundo código sempre que exibir múltiplo de 3 o participante deve dizer "Pin" em vez do número e múltiplo de 5 deve dizer "Pan". Então o programa deve imprimir numeros de 1 a 100, mas quando for múltiplo de 3 deve imprimir "Pin" e quando for múltiplo de 5 deve imprimir "Pan". Nos casos citados, não deve imprimir o número.
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		if math.Mod(float64(i), 3) == 0 {
			fmt.Println("Pin")
		} else if math.Mod(float64(i), 5) == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
