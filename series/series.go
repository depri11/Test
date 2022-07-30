package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := deretAngka("789101112")
    fmt.Println(res)
}

func deretAngka(input string) int {
	var panjang_angka, hasilKali, output int
	var stat = false

	for panjang_angka < len(input)/4 && !stat {
		panjang_angka++

		a1, _ := strconv.Atoi(input[0:panjang_angka])
		a2, _ := strconv.Atoi(input[panjang_angka : panjang_angka*2])
		a3, _ := strconv.Atoi(input[panjang_angka*2 : panjang_angka*3])
		a4, _ := strconv.Atoi(input[panjang_angka*3 : panjang_angka*4])

		hasilKali = a2 - a1
		hasilKali = hasilKali * (a3 - a2)
		hasilKali = hasilKali * (a4 - a3)

		if hasilKali == 1 || hasilKali == 2 {
			stat = true
		}

	}

	if stat {
		deret := len(input) / panjang_angka

		for i := 0; i < deret-1; i++ {
			a1, _ := strconv.Atoi(input[i*panjang_angka : panjang_angka*(i+1)])
			a2, _ := strconv.Atoi(input[(i+1)*panjang_angka : panjang_angka*(i+2)])

			if a2-a1 > 1 {
				output = a1 + 1
			}
		}
	}

	return output
}
