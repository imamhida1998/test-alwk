package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func findFirstMatch(N int, stringsList []string) interface{} {
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if strings.EqualFold(stringsList[i], stringsList[j]) {
				return fmt.Sprintf("%d %d", i+1, j+1)
			}
		}
	}
	return false
}

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

func HitungCuti(jumlahCutiBersama int, tanggalJoin, tanggalCuti time.Time, durasiCuti int) (bool, string) {
	cutiKantorPerTahun := 14
	masaTungguHari := 180
	maxCutiPribadiBerturut := 3

	jumlahCutiPribadi := cutiKantorPerTahun - jumlahCutiBersama

	tanggalBisaCuti := tanggalJoin.AddDate(0, 0, masaTungguHari)
	if tanggalCuti.Before(tanggalBisaCuti) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	tanggalAkhirTahun := time.Date(tanggalCuti.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	jumlahHariTersisa := tanggalAkhirTahun.Sub(tanggalCuti).Hours() / 24
	kuotaCutiPribadi := int(math.Floor((jumlahHariTersisa / 365) * float64(jumlahCutiPribadi)))

	if durasiCuti > kuotaCutiPribadi {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", kuotaCutiPribadi)
	}

	if durasiCuti > maxCutiPribadiBerturut {
		return false, fmt.Sprintf("Karena durasi cuti melebihi maksimum cuti pribadi berturut-turut (%d hari)", maxCutiPribadiBerturut)
	}

	return true, ""
}

func main() {
	var choice int
	fmt.Println("Pilih Soal:")
	fmt.Println("1. Mencari String yang sesuai")
	fmt.Println("2. Menghitung Kembalian Kasir")
	fmt.Println("3. Penentuan Cuti Pribadi")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var N int
		fmt.Println("Masukkan jumlah string:")
		fmt.Scan(&N)
		stringsList := make([]string, N)
		fmt.Println("Masukkan string satu per satu:")
		for i := 0; i < N; i++ {
			fmt.Scan(&stringsList[i])
		}
		result := findFirstMatch(N, stringsList)
		fmt.Println("Hasil:", result)

	case 2:
		var totalBelanja, uangDibayar int
		fmt.Println("Masukkan total belanja:")
		fmt.Scan(&totalBelanja)
		fmt.Println("Masukkan jumlah uang yang dibayarkan:")
		fmt.Scan(&uangDibayar)
		result := calculateChange(totalBelanja, uangDibayar)
		fmt.Println("Hasil:", result)

	case 3:
		var jumlahCutiBersama int
		var tanggalJoinStr, tanggalCutiStr string
		var durasiCuti int

		fmt.Println("Masukkan jumlah cuti bersama:")
		fmt.Scan(&jumlahCutiBersama)

		fmt.Println("Masukkan tanggal join karyawan (format: yyyy-mm-dd):")
		fmt.Scan(&tanggalJoinStr)

		fmt.Println("Masukkan tanggal rencana cuti (format: yyyy-mm-dd):")
		fmt.Scan(&tanggalCutiStr)

		fmt.Println("Masukkan durasi cuti (hari):")
		fmt.Scan(&durasiCuti)

		tanggalJoin, _ := time.Parse("2006-01-02", tanggalJoinStr)
		tanggalCuti, _ := time.Parse("2006-01-02", tanggalCutiStr)
		boleh, alasan := HitungCuti(jumlahCutiBersama, tanggalJoin, tanggalCuti, durasiCuti)

		if boleh {
			fmt.Println("Hasil: ", boleh)
		} else {
			fmt.Println("Hasil: ", boleh)
			fmt.Println("Alasan:", alasan)
		}

	default:
		fmt.Println("Pilihan Tidak ada")
	}
}
