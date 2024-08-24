package main

import (
	"fmt"
	"strings"
)

func findFirstMatch(N int, stringsList []string) interface{} { // interface untuk membeaskan balikan type data
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if strings.EqualFold(stringsList[i], stringsList[j]) {
				// Mengembalikan index (1-based)
				return fmt.Sprintf("%d %d", i+1, j+1)
			}
		}
	}
	// misalkan data tidak menemukan yang cocok bakal ngembaliin false
	return false
}

func main() {
	N := 11
	stringsList := []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"}
	result := findFirstMatch(N, stringsList)
	fmt.Println(result)

}
