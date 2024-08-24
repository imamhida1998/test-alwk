package main

import (
	"fmt"
	"math"
	"time"
)

const cutiKantorPerTahun = 14
const masaTungguHari = 180
const maxCutiPribadiBerturut = 3

func HitungCuti(jumlahCutiBersama int, tanggalJoin, tanggalCuti time.Time, durasiCuti int) (bool, string) {
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
	var jumlahCutiBersama int
	var tanggalJoinStr, tanggalCutiStr string
	var durasiCuti int

	tanggalJoinStr = "2021-05-01"
	tanggalCutiStr = "2021-07-05"
	durasiCuti = 1

	tanggalJoin, _ := time.Parse("2006-01-02", tanggalJoinStr)
	tanggalCuti, _ := time.Parse("2006-01-02", tanggalCutiStr)

	_, alasan := HitungCuti(jumlahCutiBersama, tanggalJoin, tanggalCuti, durasiCuti)
	fmt.Println("Output:")
	fmt.Println("False")
	fmt.Println("Alasan:", alasan)
	fmt.Println()

	jumlahCutiBersama = 7
	tanggalJoinStr = "2021-05-01"
	tanggalCutiStr = "2021-11-05"
	durasiCuti = 3

	tanggalJoin, _ = time.Parse("2006-01-02", tanggalJoinStr)
	tanggalCuti, _ = time.Parse("2006-01-02", tanggalCutiStr)

	boleh, alasan := HitungCuti(jumlahCutiBersama, tanggalJoin, tanggalCuti, durasiCuti)
	fmt.Println("Output:")
	fmt.Println(boleh)
	fmt.Println("Alasan:", alasan)
	fmt.Println()

}
