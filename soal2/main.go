package main

import (
	"fmt"
	"math"
)

func calculateChange(totalBelanja, uangDibayar int) string {
	if uangDibayar < totalBelanja {
		return "kurang bayar"
	}

	pecahanLembar := []int{1000, 2000, 5000, 10000, 20000, 50000, 100000}
	pecahanKoin := []int{100, 200, 500}

	change := uangDibayar - totalBelanja
	roundedChange := int(math.Floor(float64(change)/100) * 100)

	result := fmt.Sprintf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d\n", change, roundedChange)

	result += "Pecahan uang:\n"

	// Menghitung pecahan lembar terlebih dahulu
	for i := len(pecahanLembar) - 1; i >= 0; i-- {
		pecahan := pecahanLembar[i]
		count := roundedChange / pecahan
		if count > 0 {
			result += fmt.Sprintf("%d lembar %d\n", count, pecahan)
		}
		roundedChange %= pecahan
	}

	// Menghitung pecahan koin
	for i := len(pecahanKoin) - 1; i >= 0; i-- {
		pecahan := pecahanKoin[i]
		count := roundedChange / pecahan
		if count > 0 {
			result += fmt.Sprintf("%d koin %d\n", count, pecahan)
		}
		roundedChange %= pecahan
	}

	return result
}

func main() {
	totalBelanja := 700649
	uangDibayar := 800000
	result := calculateChange(totalBelanja, uangDibayar)
	fmt.Println(result)

}
