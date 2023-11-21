// Crie um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		if math.Mod(float64(i), 3) == 0 {
			fmt.Println(i)
		}
	}
}
